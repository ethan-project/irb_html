
package model

import (
	"github.com/gin-gonic/gin"
  "ipsap/common"
	"strings"
	"time"
	"fmt"
)

type DashBoard struct {
	App            Application
  DashBoard_type uint
}

func (board *DashBoard)GetDashBoardContent() (result gin.H) {
	moreSelct	:= fmt.Sprintf(` , FROM_UNIXTIME(app.submit_dttm, '%%Y-%%m-%%d') as tmp_submit_dttm,
															 IF(TIMESTAMPDIFF(DAY, FROM_UNIXTIME(app.submit_dttm, '%%Y-%%m-%%d'), NOW()) = 0,
																	'오늘', concat(TIMESTAMPDIFF(DAY, FROM_UNIXTIME(app.submit_dttm, '%%Y-%%m-%%d'), NOW()),'일 경과')
																 )AS time_diff`)
	switch board.DashBoard_type {
		case DEF_DASHBOARD_TYPE_ADMIN:
				 result = board.GetAdminDashBoard(moreSelct)
		case DEF_DASHBOARD_TYPE_RESEARCHER:
				 result = board.GetResearcherDashBoard(moreSelct)
		case DEF_DASHBOARD_TYPE_CHAIRPERSON:
				 result = board.GetChairpersonDashBoard(moreSelct)
	}
  return
}

func (board *DashBoard)GetAdminDashBoard(moreSelct string) (result gin.H) {
	check_condition	:= fmt.Sprintf( ` AND app.application_step = %d
																		AND app.application_result <> %d`,
																		DEF_APP_STEP_CHECKING,
																		DEF_APP_RESULT_DELETED)
	judge_condition	:= fmt.Sprintf( ` AND app.application_step IN(%d, %d)
																		AND app.application_result <> %d`,
																		DEF_APP_STEP_PRO,
																		DEF_APP_STEP_NORMAL,
																		DEF_APP_RESULT_DELETED)
	finial_condition := fmt.Sprintf(` AND NOT app.judge_type = IF(istt.ia_final_director = %d, %d, 0)
																		AND app.application_result <> %d`,
																		DEF_FINAL_DIRECTOR_CHAIRPERSON,
																		DEF_APP_JUDGE_TYPE_CODE_IACUC,
																		DEF_APP_RESULT_DELETED)
	check_rows := board.App.LoadList(moreSelct,"",check_condition,"").([]map[string]interface{})
	judge_rows := board.setDelayAppTimeDiff(board.App.LoadList(moreSelct,"",judge_condition,"").([]map[string]interface{}))
	final_rows := board.getFinalAppList(moreSelct, finial_condition)
	sup_condition := fmt.Sprintf( ` AND app.application_result = %d
																	AND app.application_result <> %d`,
																	DEF_APP_RESULT_SUPPLEMENT,
																	DEF_APP_RESULT_DELETED)
	sup_rows := board.App.LoadList(moreSelct,"",sup_condition,"").([]map[string]interface{})

	result = gin.H{ "app_cnt" : board.getAppCount(),
									"supplement_cnt" : len(sup_rows),
									"supplement" : sup_rows,
									"checking_cnt" : len(check_rows),
									"checking" : check_rows,
									"judge_ing_cnt" : len(judge_rows),
									"judge_ing" : judge_rows,
									"final_cnt" : len(final_rows),
									"final" : final_rows,
									"performance_app_list" : board.getPerformanceAppList(fmt.Sprintf(`AND app.application_result <> %d`,DEF_APP_RESULT_DELETED)),
									"approved_statistics" : board.getApprovedStatistics(),}
	return
}

// 내가 참여한 실험만 가져와야됨!
func (board *DashBoard)GetResearcherDashBoard(moreSelct string) (result gin.H) {
	moreCondition := fmt.Sprintf(`
										AND (
											app.application_seq = (
												SELECT
													application_seq
												FROM
													t_application_member member
												WHERE
													member.item_name IN (
														'general_director', 'general_expt'
													)
													AND app.application_seq = member.application_seq
													AND member.user_seq = %v
											)
											OR app.parent_app_seq = (
												SELECT
													application_seq
												FROM
													t_application_member member
												WHERE
													member.item_name IN (
														'general_director', 'general_expt'
													)
													AND app.parent_app_seq = member.application_seq
													AND member.user_seq = %v
											)
										)
										AND app.application_result <> %d`, board.App.LoginToken["user_seq"], board.App.LoginToken["user_seq"], DEF_APP_RESULT_DELETED)

	check_condition	:= fmt.Sprintf( ` AND app.application_step = %d`, DEF_APP_STEP_CHECKING) + moreCondition
	judge_condition	:= fmt.Sprintf( ` AND app.application_step IN(%d, %d)`, DEF_APP_STEP_PRO, DEF_APP_STEP_NORMAL) + moreCondition
	writing_condition := fmt.Sprintf( ` AND app.reg_user_seq = %v
																			AND app.application_step = %d
		                                  AND app.application_result IN (%d, %d)`, board.App.LoginToken["user_seq"],	DEF_APP_STEP_WRITE,	DEF_APP_RESULT_TEMP, DEF_APP_RESULT_SUPPLEMENT)
	writing_rows := board.App.LoadList(strings.Replace(moreSelct, "submit_dttm" , "reg_dttm", -1),"",writing_condition,"").([]map[string]interface{})
	check_rows := board.App.LoadList(moreSelct,"",check_condition,"").([]map[string]interface{})
	judge_rows := board.App.LoadList(moreSelct,"",judge_condition,"").([]map[string]interface{})
	final_rows := board.getFinalAppList(moreSelct, moreCondition)

	result = gin.H{ "writing_cnt" : len(writing_rows),
									"writing" : writing_rows,
									"checking_cnt" : len(check_rows),
									"checking" : check_rows,
									"judge_ing_cnt" : len(judge_rows),
									"judge_ing" : judge_rows,
									"final_cnt" : len(final_rows),
									"final" : final_rows,
									"performance_app_list" : board.getPerformanceAppList(moreCondition),}
	return
}

func (board *DashBoard)GetChairpersonDashBoard(moreSelct string) (result gin.H) {
	moreCondition := fmt.Sprintf(` AND application_result <> %d`, DEF_APP_RESULT_DELETED)
	performance_app_list := board.getPerformanceAppList(moreCondition)
	final_rows := board.getFinalAppList(moreSelct, moreCondition)
	result = gin.H{ "app_cnt" : board.getAppCount(),
									"final_cnt" : len(final_rows),
									"final" : final_rows,
									"performance_app_list_cnt" : len(performance_app_list),
									"performance_app_list" : performance_app_list,
									"approved_statistics" : board.getApprovedStatistics(),}
	return
}

func (board *DashBoard)getFinalAppList(moreSelct string, moreCondition string) (list []map[string]interface{}) {
	moreCondition	+= fmt.Sprintf(` AND app.application_step =  %d
																 AND app.application_result = %d`,
																DEF_APP_STEP_FINAL, DEF_APP_RESULT_DECISION_ING)
	list = board.App.LoadList(moreSelct,"",moreCondition,"").([]map[string]interface{})
	return
}

func (board *DashBoard)getPerformanceAppList(moreCondition string) (list []map[string]interface{}) {
	moreCondition	+= fmt.Sprintf( ` AND app.application_step = %d
																	AND app.application_result <> %d`, DEF_APP_STEP_PERFORMANCE, DEF_SUBMIT_TYPE_TASK_FINISH)

	sql2 := fmt.Sprintf(`
					SELECT etc2.contents
						FROM t_application app2, t_application_etc etc2
					 WHERE app2.application_seq = etc2.application_seq
						 AND app2.parent_app_seq = app.application_seq
						 AND app2.application_type = %d
						 AND app2.application_step = %d
						 AND app2.application_result IN (%d, %d)
						 AND etc2.item_name = 'general_end_date'
				ORDER BY app2.approved_dttm DESC
					 LIMIT 1`, DEF_APP_TYPE_CHANGE, DEF_APP_STEP_FINAL,
						DEF_APP_RESULT_APPROVED, DEF_APP_RESULT_APPROVED_C)

	moreSelect := fmt.Sprintf(` , IFNULL((%v),
																			(SELECT etc.contents
																				 FROM t_application_etc etc
																				WHERE app.application_seq = etc.application_seq
																					AND etc.item_name = 'general_end_date'))
																 			 AS general_end_date,
																			 (SELECT IF(etc.contents > date_format(now(), '%%Y-%%m-%%d'), false, true)
																			 	  FROM t_application_etc etc
																				 WHERE app.application_seq = etc.application_seq
																				 	 AND etc.item_name = 'general_end_date')
																				AS end_date_over`, sql2)
	list = board.App.LoadList(moreSelect,"",moreCondition,"").([]map[string]interface{})
	return
}

func (board *DashBoard)getApprovedStatistics() (result gin.H) {
 	sql := `SELECT count(case when app_select.select_ids = 1 then 1 end) as approved_cnt,
								 count(case when app_select.select_ids = 2 then 1 end) as conditional_approved_cnt,
								 count(case when app_select.select_ids = 3 then 1 end) as require_retry_cnt,
								 count(case when app_select.select_ids = 4 then 1 end) as reject_cnt,
								 count(app_select.application_seq) as total_cnt
						FROM t_application_select app_select, t_application app
					 WHERE app.application_seq = app_select.application_seq
						 AND app_select.item_name = 'final_judge_result'
						 AND app.judge_type = ?
						 AND app.institution_seq = ?
						 AND app.application_result <> 17`
	approved_iacuc	:= common.DB_fetch_one(sql, nil, DEF_APP_JUDGE_TYPE_CODE_IACUC, board.App.LoginToken["institution_seq"])
	approved_ibc		:= common.DB_fetch_one(sql, nil, DEF_APP_JUDGE_TYPE_CODE_IBC, board.App.LoginToken["institution_seq"])
	approved_irb		:= common.DB_fetch_one(sql, nil, DEF_APP_JUDGE_TYPE_CODE_IRB, board.App.LoginToken["institution_seq"])
	result = gin.H {
		"iacuc" : approved_iacuc,
		"ibc"	 : approved_ibc,
		"irb"	 : approved_irb,
	}
	return
}

func (board *DashBoard)getAppCount() (result gin.H) {
	sql := fmt.Sprintf(`
						 SELECT count(case when judge_type = %d then 1 end) as iacuc_cnt,
										count(case when judge_type = %d then 1 end) as ibc_cnt,
										count(case when judge_type = %d then 1 end) as irb_cnt
							 FROM t_application
							WHERE institution_seq = %v
								AND application_result <> %v
								 `,DEF_APP_JUDGE_TYPE_CODE_IACUC,
									 DEF_APP_JUDGE_TYPE_CODE_IBC,
									 DEF_APP_JUDGE_TYPE_CODE_IRB,
									 board.App.LoginToken["institution_seq"],
								 	 DEF_APP_RESULT_DELETED)

	row := common.DB_fetch_one(sql, nil)
	result = gin.H {
		"iacuc_cnt"	: row["iacuc_cnt"],
		"ibc_cnt"		: row["ibc_cnt"],
		"irb_cnt"		: row["irb_cnt"],
	}
	return
}

func (board *DashBoard)setDelayAppTimeDiff(judge_rows []map[string]interface{}) []map[string]interface{}{
	for _, judge_row := range judge_rows {
		application := Application{
			Application_seq : common.ToUint(judge_row["application_seq"]),
		}
		nowUnixTime := time.Now().Unix()
		expert_deadline := application.GetUnixtimeForExpertDeadline()
		if expert_deadline == 0  {  continue  }
		normal_deadline := application.GetUnixtimeForNormalDeadline()
		if normal_deadline == 0  {  continue  }

		application_step := common.ToInt(judge_row["application_step"])
		application_result := common.ToInt(judge_row["application_result"])
		sql := ""
		switch(application_step) {
			case DEF_APP_STEP_PRO  :      //  전문심사 단계
			if nowUnixTime >= expert_deadline {
				if (application_result == DEF_APP_RESULT_JUDGE_DELAY) { // 심사 지연일 경우
					sql = fmt.Sprintf(`SELECT concat(TIMESTAMPDIFF(DAY, FROM_UNIXTIME(%v, '%%Y-%%m-%%d'), NOW()), "일 지연")  AS time_diff`, expert_deadline)
					row := common.DB_fetch_one(sql, nil)
					judge_row["time_diff"] = row["time_diff"]
				}
			}
			case DEF_APP_STEP_NORMAL  :   //  일반심사 단계
			if nowUnixTime >= normal_deadline {
				if (application_result == DEF_APP_RESULT_JUDGE_DELAY) { // 심사 지연일 경우
					sql = fmt.Sprintf(`SELECT concat(TIMESTAMPDIFF(DAY, FROM_UNIXTIME(%v, '%%Y-%%m-%%d'), NOW()), "일 지연")  AS time_diff`, normal_deadline)
					row := common.DB_fetch_one(sql, nil)
					judge_row["time_diff"] = row["time_diff"]
				}
			}
		}
	}

	return judge_rows
}
