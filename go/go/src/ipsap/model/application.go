
package model

import (
	"github.com/nleeper/goment"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
  "ipsap/common"
  "database/sql"
  "strings"
	"time"
	"math"
	"fmt"
  "log"
)

type Application struct {
  LoginToken          map[string]interface{}
  Application_seq     uint
  Items               map[string]interface{}
	Data								map[string]interface{}
	StartIdx						uint
	RowCnt							uint
	SearchWords					string
}

func (ins *Application)Init() (succ bool)  {
	// 21-03-03 donghun : 같은 기관이면 성공
  ins.Items = make(map[string]interface{})
  if ins.Application_seq > 0  {
    sql := `SELECT application_seq, application_result, parent_app_seq, application_type, judge_type
              FROM t_application app
             WHERE application_seq = ?
               AND institution_seq = ?`
    row := common.DB_fetch_one(sql, nil, ins.Application_seq, ins.LoginToken["institution_seq"])
    if nil == row {
      return      //  신청서가 없거나 연구책임자가 다르면 false
    }
		ins.Data = row
  }
  return true
}

func (ins *Application)getAppListQueryAndFilter(moreSelect string, moreJoin string, moreCondition string, appViewType string) (sql string, filter func(map[string]interface{})) {
	//	wowdolf : 로긴 사용자의 권한이 있는 것만 조회 되어야 함.
	sql  = fmt.Sprintf(`
					SELECT app.application_seq, app.parent_app_seq, app.application_no, app.institution_seq,
                 app.judge_type, app.application_type, app.application_step, app.application_result,
                 app.name_ko, app.name_en, app.approved_dttm,
                 app.reg_dttm, app.reg_user_seq, app.submit_dttm,
                 istt.name_ko istt_name_ko, istt.name_en istt_name_en,
                 user.dept user_dept, user.name user_name,
								 user.email user_email, user.phoneno user_phoneno,
								 user.position user_position %v #moreSelct
            FROM t_institution istt,
                 t_application app
                 LEFT OUTER JOIN t_user user ON
                          (app.reg_user_seq = user.user_seq)
							%v #moreJoin
           WHERE app.institution_seq = istt.institution_seq
             AND istt.institution_status = ?
						  %v
					 ORDER BY app.submit_dttm DESC`, moreSelect, moreJoin, moreCondition)
  filter = func(row map[string]interface{}) {
    codeJT := Code {
      Type : "judge_type",
      Id : common.ToUint(row["judge_type"]),
    }
    row["judge_type_str"] = codeJT.GetCodeStrFromTypeAndId()

    codeAT := Code {
      Type : "application_type",
      Id : common.ToUint(row["application_type"]),
    }
    row["application_type_str"] = codeAT.GetCodeStrFromTypeAndId()

    codeAS := Code {
      Type : "application_step",
      Id : common.ToUint(row["application_step"]),
    }
    row["application_step_str"] = codeAS.GetCodeStrFromTypeAndId()

    codeAR := Code {
      Type : "application_result",
      Id : common.ToUint(row["application_result"]),
    }
    row["application_result_str"] = codeAR.GetCodeStrFromTypeAndId()

    dcodeDept := CodeDynamic{
			Institution_seq : INSTITUTION_SHARE_CODE,  //common.ToUint(ins.LoginToken["institution_seq"]),
			DCode_type      : DCODE_TYPE_DEPT,
			Code            : common.ToUint(row["user_dept"]),
		}
		row["user_dept_str"] = dcodeDept.GetValueFromCode()

		dcodePosition := CodeDynamic{
			Institution_seq : INSTITUTION_SHARE_CODE,
			DCode_type      : DCODE_TYPE_POSITION,
			Code            : common.ToUint(row["user_position"]),
		}
		row["user_position_str"] = dcodePosition.GetValueFromCode()

		if nil != row["committee"] {
			committee_judge_type_str := ""
			switch row["committee"] {
				case "committee_ex_member": committee_judge_type_str = "일반심사"
				case "committee_in_member": committee_judge_type_str = "일반심사"
				case "expert_member":  committee_judge_type_str = "전문심사"
			}
			row["committee_judge_type_str"] = committee_judge_type_str
			delete(row, "committee")
		}

		sub_title := ""

		switch common.ToUint(row["judge_type"]) {
			case DEF_APP_JUDGE_TYPE_CODE_IACUC : {
				switch common.ToUint(row["application_type"]) {
					case DEF_APP_TYPE_NEW: sub_title = "동물실험 계획서"
					case DEF_APP_TYPE_CHANGE: sub_title = "동물실험 계획 변경 승인 신청서"
					case DEF_APP_TYPE_RENEW: sub_title = "동물실험 계획 재승인 신청서"
					case DEF_APP_TYPE_BRINGIN: sub_title = "실험동물 반입 보고서"
					case DEF_APP_TYPE_CHECKLIST: sub_title = "동물실험 계획 승인 후 점검표"
					case DEF_APP_TYPE_FINISH: sub_title = "동물실험 종료 보고서"
				}
			}
			case DEF_APP_JUDGE_TYPE_CODE_IBC : {
				switch common.ToUint(row["application_type"]) {
					case DEF_APP_TYPE_NEW: sub_title = "생물안전위원회 심의 신청서"
					case DEF_APP_TYPE_CHANGE: sub_title = "생물안전위원회 변경 심의 신청서"
				}
			}
			case DEF_APP_JUDGE_TYPE_CODE_IRB : {
				switch common.ToUint(row["application_type"]) {
					case DEF_APP_TYPE_NEW: sub_title = "기관생명윤리위원회 심의 신청서"
					case DEF_APP_TYPE_CHANGE: sub_title = "기관생명윤리위원회 변경 심의 신청서"
					case DEF_APP_TYPE_CONTINUE: sub_title = "기관생명윤리위원회 지속심의(중간보고) 신청서"
					case DEF_APP_TYPE_SERIOUS: sub_title = "기관생명윤리위원회 중대한 이상 반응 보고서"
					case DEF_APP_TYPE_VIOLATION: sub_title = "기관생명윤리위원회 연구계획 위반/이탈 보고서"
					case DEF_APP_TYPE_UNEXPECTED: sub_title = "기관생명윤리위원회 예상치 못한 문제발생 보고서"
					case DEF_APP_TYPE_FINISH: sub_title = "기관생명윤리위원회 종료 보고서"
				}
			}

		}

		row["sub_title"] = sub_title
  }
	return;
}

func (ins *Application)getApprovedAppListQueryAndFilter(moreCondition string, appViewType string) (sql string, filter func(map[string]interface{}))  {
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
	sql3 := fmt.Sprintf(`
				SELECT COUNT(app3.application_seq)
					 FROM t_application app3
					WHERE app3.parent_app_seq = app.application_seq
						AND app3.application_type = %d
						AND app3.application_step = %d
						AND app3.application_result IN (%d, %d)`, DEF_APP_TYPE_RENEW, DEF_APP_STEP_FINAL,
								DEF_APP_RESULT_APPROVED, DEF_APP_RESULT_APPROVED_C)

	sql  = fmt.Sprintf(`
					SELECT app.application_seq,		app.name_ko,	app.name_en,
								 IFNULL((%v), etc.contents) AS general_end_date,
								 app.application_step,	app.application_result, app.parent_app_seq, app.check_user_seq,
								 app.approved_dttm as start_time,
								 FROM_UNIXTIME(app.approved_dttm, '%%Y-%%m-%%d') AS approved_dttm,
								 app.institution_seq,	app.judge_type,
								 user.dept user_dept, user.name user_name,
								 user.email user_email, user.phoneno user_phoneno,
								 user.position user_position, app.submit_dttm,
								 (%v) as renew_app_cnt
	          FROM t_application app
						LEFT OUTER JOIN t_application_etc etc ON (app.application_seq = etc.application_seq AND etc.item_name = 'general_end_date')
						LEFT OUTER JOIN t_institution istt ON (app.institution_seq = istt.institution_seq)
					  LEFT OUTER JOIN t_user user ON (app.reg_user_seq = user.user_seq)
	         WHERE 1 = 1
					 	 AND istt.institution_status = %d
						  %v
				ORDER BY app.approved_dttm DESC`, sql2, sql3, DEF_INSTITUTION_STATUS_OK, moreCondition)
  filter = func(row map[string]interface{}) {
		if DEF_APP_VIEW_APPROVED_RESEARCHER_ALL == common.ToInt(appViewType) {
			startUnixT := common.ToInt64(row["start_time"])
			if startUnixT > 0 && nil != row["general_end_date"] {
				row["rquired_renew"] = false
				startT, _ := goment.Unix(startUnixT)
				endDateT, _ := goment.New(row["general_end_date"])
				timeDiffY := math.Abs(cast.ToFloat64(startT.Diff(endDateT, "years")))
				approvedRenewAppCnt := common.ToInt(row["renew_app_cnt"])
				toDay, _ := goment.New()
				if timeDiffY > 0 {
					tempStartT11, _ := goment.Unix(startUnixT)
					tempStartT11.Add(11, "months")
					tempStartT12, _ := goment.Unix(startUnixT)
					tempStartT12.Add(12, "months")
					if toDay.IsAfter(tempStartT11) {
						if approvedRenewAppCnt == 0 {
							if toDay.IsAfter(tempStartT11) && toDay.IsSameOrBefore(tempStartT12) {
								if tempStartT12.Diff(toDay, "days") == 0 {
									row["time_diff"] = "일"
								} else {
									row["time_diff"] = common.ToStr(tempStartT12.Diff(toDay, "days")) + "일 전"
								}
							} else {
								row["time_diff"] = common.ToStr(toDay.Diff(tempStartT12, "days")) + "일 경과"
							}
							row["rquired_renew"] = true
						} else if approvedRenewAppCnt == 1 {
							if timeDiffY == 2 {
								tempStartT23, _ := goment.Unix(startUnixT)
								tempStartT24, _ := goment.Unix(startUnixT)
								tempStartT23.Add(23, "months")
								tempStartT24.Add(24, "months")
								if toDay.IsAfter(tempStartT23) && toDay.IsSameOrBefore(tempStartT24) {
									if tempStartT12.Diff(toDay, "days") == 0 {
										row["time_diff"] = "일"
									} else {
										row["time_diff"] = common.ToStr(tempStartT24.Diff(toDay, "days")) + "일 전"
									}
								} else {
									row["time_diff"] = common.ToStr(toDay.Diff(tempStartT24, "days")) + "일 경과"
								}
								row["rquired_renew"] = true
							}
						}
					}
				}
			}
		}

		codeJT := Code {
			Type : "judge_type",
			Id : common.ToUint(row["judge_type"]),
		}
		row["judge_type_str"] = codeJT.GetCodeStrFromTypeAndId()

		codeAS := Code {
			Type : "application_step",
			Id : common.ToUint(row["application_step"]),
		}
		row["application_step_str"] = codeAS.GetCodeStrFromTypeAndId()

		codeAR := Code {
			Type : "application_result",
			Id : common.ToUint(row["application_result"]),
		}
		row["application_result_str"] = codeAR.GetCodeStrFromTypeAndId()

		dcodeDept := CodeDynamic{
			Institution_seq : INSTITUTION_SHARE_CODE,  //common.ToUint(ins.LoginToken["institution_seq"]),
			DCode_type      : DCODE_TYPE_DEPT,
			Code            : common.ToUint(row["user_dept"]),
		}
		row["user_dept_str"] = dcodeDept.GetValueFromCode()

		dcodePosition := CodeDynamic{
			Institution_seq : INSTITUTION_SHARE_CODE,
			DCode_type      : DCODE_TYPE_POSITION,
			Code            : common.ToUint(row["user_position"]),
		}
		row["user_position_str"] = dcodePosition.GetValueFromCode()

		sql2 := fmt.Sprintf(`
						SELECT app.application_seq, app.parent_app_seq, app.application_no, app.institution_seq,
	                 app.judge_type, app.application_type, app.application_step, app.application_result,
	                 app.name_ko, app.name_en, app.approved_dttm,
	                 app.reg_dttm, app.reg_user_seq, app.submit_dttm,
	                 istt.name_ko istt_name_ko, istt.name_en istt_name_en,
	                 user.dept user_dept, user.name user_name,
									 user.email user_email, user.phoneno user_phoneno,
									 user.position user_position
	            FROM t_institution istt, t_application app
	            LEFT OUTER JOIN t_user user ON (app.reg_user_seq = user.user_seq)
						 WHERE app.institution_seq = istt.institution_seq
						 	 AND app.application_seq = %v
							 AND app.application_result <> %v
							 UNION ALL
					 SELECT app.application_seq, app.parent_app_seq, app.application_no, app.institution_seq,
									app.judge_type, app.application_type, app.application_step, app.application_result,
									app.name_ko, app.name_en, app.approved_dttm,
									app.reg_dttm, app.reg_user_seq, app.submit_dttm,
									istt.name_ko istt_name_ko, istt.name_en istt_name_en,
									user.dept user_dept, user.name user_name,
									user.email user_email, user.phoneno user_phoneno,
									user.position user_position
						 FROM t_institution istt, t_application app
						 LEFT OUTER JOIN t_user user ON (app.reg_user_seq = user.user_seq)
						WHERE app.institution_seq = istt.institution_seq
							AND app.parent_app_seq = %v
							AND app.application_result <> %v
						ORDER BY reg_dttm
						`, row["application_seq"], DEF_APP_RESULT_DELETED, row["application_seq"], DEF_APP_RESULT_DELETED)
		_, filter2 := ins.getAppListQueryAndFilter("", "", "", "")
		row["sub_list"] = common.DB_fetch_all(sql2, filter2)
	}
	return
}

func (ins *Application)getPosibleAppListQueryAndFilter(moreCondition string, appViewType string) (sql string, filter func(map[string]interface{})) {
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
	sql3 := fmt.Sprintf(`
				SELECT COUNT(app3.application_seq)
					 FROM t_application app3
					WHERE app3.parent_app_seq = app.application_seq
						AND app3.application_type = %d
						AND app3.application_step = %d
						AND app3.application_result IN (%d, %d)`, DEF_APP_TYPE_RENEW, DEF_APP_STEP_FINAL,
								DEF_APP_RESULT_APPROVED, DEF_APP_RESULT_APPROVED_C)

	sql  = fmt.Sprintf(`
					SELECT app.application_seq, 	app.parent_app_seq, 		app.application_no,
								 app.institution_seq,		app.judge_type,					app.application_type,
								 app.application_step,	app.application_result,	app.name_ko,
								 app.name_en,					  app.reg_dttm,
								 app.reg_user_seq,			app.submit_dttm,
								 IFNULL((%v), etc.contents) AS general_end_date,
								 app.approved_dttm as start_time,
								 FROM_UNIXTIME(app.approved_dttm, '%%Y-%%m-%%d') AS approved_dttm,
								 (%v) as renew_app_cnt,
                 istt.name_ko istt_name_ko, istt.name_en istt_name_en,
								 user.dept user_dept, user.name user_name,
								 user.email user_email, user.phoneno user_phoneno,
								 user.position user_position
            FROM t_application app
						LEFT OUTER JOIN t_application_etc etc ON (app.application_seq = etc.application_seq AND etc.item_name = 'general_end_date')
						LEFT OUTER JOIN t_institution istt ON (app.institution_seq = istt.institution_seq)
					  LEFT OUTER JOIN t_user user ON (app.reg_user_seq = user.user_seq)
           WHERE 1 = 1
             AND istt.institution_status = ?
						  %v #moreCondition
				ORDER BY app.approved_dttm DESC`, sql2, sql3, moreCondition)
  filter = func(row map[string]interface{}) {
    codeJT := Code {
      Type : "judge_type",
      Id : common.ToUint(row["judge_type"]),
    }
    row["judge_type_str"] = codeJT.GetCodeStrFromTypeAndId()

    codeAT := Code {
      Type : "application_type",
      Id : common.ToUint(row["application_type"]),
    }
    row["application_type_str"] = codeAT.GetCodeStrFromTypeAndId()

    codeAS := Code {
      Type : "application_step",
      Id : common.ToUint(row["application_step"]),
    }
    row["application_step_str"] = codeAS.GetCodeStrFromTypeAndId()

    codeAR := Code {
      Type : "application_result",
      Id : common.ToUint(row["application_result"]),
    }
    row["application_result_str"] = codeAR.GetCodeStrFromTypeAndId()

		dcodeDept := CodeDynamic{
			Institution_seq : INSTITUTION_SHARE_CODE,
			DCode_type      : DCODE_TYPE_DEPT,
			Code            : common.ToUint(row["user_dept"]),
		}
		row["user_dept_str"] = dcodeDept.GetValueFromCode()

		dcodePosition := CodeDynamic{
			Institution_seq : INSTITUTION_SHARE_CODE,
			DCode_type      : DCODE_TYPE_POSITION,
			Code            : common.ToUint(row["user_position"]),
		}
		row["user_position_str"] = dcodePosition.GetValueFromCode()

  }
	return;
}

func (ins *Application)getApprovedAppCountQuery(moreCondition string) (sql string)  {
	sql  = fmt.Sprintf(`
					SELECT COUNT(app.application_seq) AS app_cnt
	          FROM t_application app
						LEFT OUTER JOIN t_application_etc etc ON (app.application_seq = etc.application_seq AND etc.item_name = 'general_end_date')
						LEFT OUTER JOIN t_institution istt ON (app.institution_seq = istt.institution_seq)
					  LEFT OUTER JOIN t_user user ON (app.reg_user_seq = user.user_seq)
	         WHERE 1 = 1
					 	 AND istt.institution_status = %d
						  %v`,DEF_INSTITUTION_STATUS_OK, moreCondition)
	return
}

func (ins *Application)getAppCountQuery(moreJoin string, moreCondition string) (sql string) {
	sql  = fmt.Sprintf(`
					SELECT COUNT(app.application_seq) AS app_cnt
            FROM t_institution istt,
                 t_application app
                 LEFT OUTER JOIN t_user user ON
                          (app.reg_user_seq = user.user_seq)
							%v #moreJoin
           WHERE app.institution_seq = istt.institution_seq
             AND istt.institution_status = ?
						  %v `, moreJoin, moreCondition)
	return;
}

func (ins *Application)Load() (ret interface{}) {
	moreCondition := ` AND app.application_seq = ?`
	if (ins.LoginToken != nil)	{
		moreCondition += fmt.Sprintf(` AND app.institution_seq = %v`,  ins.LoginToken["institution_seq"])
	}

	sql, filter := ins.getAppListQueryAndFilter("", "", moreCondition, "");
  ret = common.DB_fetch_one(sql, filter, DEF_INSTITUTION_STATUS_OK,
																	ins.Application_seq)
  return
}

func (ins *Application) getMoreCondition(appViewType string, approved bool) (moreCondition string) {
	addCondition := ""
	if appViewType ==  common.ToStr(DEF_APP_VIEW_ADMIN_ALL) && approved == false {
		addCondition = fmt.Sprintf(`OR IFNULL((SELECT user.name
																						 FROM t_application_member member, t_user user
																						WHERE member.user_seq = user.user_seq
																							AND member.item_name = 'expert_member'
																							AND app.application_seq = member.application_seq), '') LIKE '%%%v%%'`, ins.SearchWords)
	}

	moreCondition = fmt.Sprintf(` AND(app.application_no LIKE '%%%v%%'
																	OR app.name_ko LIKE '%%%v%%'
																	OR app.name_en LIKE '%%%v%%'
																	OR istt.name_ko LIKE '%%%v%%'
																	OR istt.name_en LIKE '%%%v%%'
																	OR IF(app.approved_dttm <> '', FROM_UNIXTIME(app.approved_dttm, '%%Y-%%m-%%d') LIKE '%%%v%%', false)
																	OR IF(app.reg_dttm <> '', FROM_UNIXTIME(app.reg_dttm, '%%Y-%%m-%%d') LIKE '%%%v%%', false)
																	OR (SELECT value
																				FROM t_code
																			 WHERE type = 'judge_type'
																				 AND id = app.judge_type) LIKE '%%%v%%'
																	OR (SELECT value
																				FROM t_code
																			 WHERE type = 'application_type'
																				 AND id = app.application_type) LIKE '%%%v%%'
																	OR (SELECT value
																				FROM t_code
																			 WHERE type = 'application_step'
																				 AND id = app.application_step) LIKE '%%%v%%'
																	OR (SELECT value
																				FROM t_code
																			 WHERE type = 'application_result'
																				 AND id = app.application_result) LIKE '%%%v%%'
																	OR (SELECT value
																				FROM t_code_dyn
																			 WHERE code = 2
																				 AND institution_seq = istt.institution_seq
																				 AND dcode_type = user.dept) LIKE '%%%v%%'
																	 %v #addCondition
																	)`, ins.SearchWords, ins.SearchWords, ins.SearchWords, ins.SearchWords, ins.SearchWords, ins.SearchWords, ins.SearchWords, ins.SearchWords, ins.SearchWords, ins.SearchWords, ins.SearchWords, ins.SearchWords, addCondition)
	return
}

func (ins *Application) LoadList(moreSelect string, moreJoin string, moreCondition string, appViewType string) (ret interface{})  {
	if (ins.LoginToken != nil)	{
		moreCondition += fmt.Sprintf(` AND app.institution_seq = %v`,  ins.LoginToken["institution_seq"])
	}

	if ins.SearchWords != "" {
		moreCondition += ins.getMoreCondition(appViewType, false)
	}

	sql, filter := ins.getAppListQueryAndFilter(moreSelect, moreJoin, moreCondition, appViewType)
	if ins.RowCnt != 0 {
		sql += fmt.Sprintf(` LIMIT %v, %v`, ins.StartIdx, ins.RowCnt)
	}

	ret = common.DB_fetch_all(sql, filter, DEF_INSTITUTION_STATUS_OK)
  return
}

func (ins *Application) GetAppCnt(moreJoin string, moreCondition string, appViewType string) (ret interface{})  {
	if (ins.LoginToken != nil)	{
		moreCondition += fmt.Sprintf(` AND app.institution_seq = %v`,  ins.LoginToken["institution_seq"])
	}

	if ins.SearchWords != "" {
		moreCondition += ins.getMoreCondition(appViewType, false)
	}

	sql := ins.getAppCountQuery(moreJoin, moreCondition)
	row := common.DB_fetch_one(sql, nil, DEF_INSTITUTION_STATUS_OK)
	ret = row["app_cnt"]
  return
}

func (ins *Application) LoadPossibleList(moreCondition string, appViewType string) (ret interface{})  {
	if (ins.LoginToken != nil) {
		moreCondition += fmt.Sprintf(` AND app.institution_seq = %v`,  ins.LoginToken["institution_seq"])
	}

	sql, filter := ins.getPosibleAppListQueryAndFilter(moreCondition, "");
	rows := common.DB_fetch_all(sql, filter, DEF_INSTITUTION_STATUS_OK)
	if DEF_APP_VIEW_POSSIBLE_IA_RETRY == common.ToUint(appViewType) {
		for i := len(rows) - 1; i >= 0; i-- {
			showFlag := false
			row := rows[i]
			startUnixT := common.ToInt64(row["start_time"])
			if startUnixT > 0 && nil != row["general_end_date"] {
				startT, _ := goment.Unix(startUnixT)
				endDateT, _ := goment.New(row["general_end_date"])
				timeDiffY := math.Abs(cast.ToFloat64(startT.Diff(endDateT, "years")))
				approvedRenewAppCnt := common.ToInt(row["renew_app_cnt"])
				toDay, _ := goment.New()
				if timeDiffY > 0 {
					tempStartT11, _ := goment.Unix(startUnixT)
					tempStartT11.Add(11, "months")
					if toDay.IsAfter(tempStartT11) {
						if approvedRenewAppCnt == 0 {
							showFlag = true
						} else if approvedRenewAppCnt == 1 {
							if timeDiffY == 2 {
								showFlag = true
							}
						}
					}
				}
				if (!showFlag) {
					rows = append(rows[:i], rows[i+1:]...)
				}
			}
		}
	}
	ret = rows
	return
}

func (ins *Application) LoadApprovedList(moreCondition string, appViewType string) (ret interface{})  {
	if (ins.LoginToken != nil)	{
		moreCondition += fmt.Sprintf(` AND app.institution_seq = %v`,  ins.LoginToken["institution_seq"])
	}

	if ins.SearchWords != "" {
		moreCondition += ins.getMoreCondition(appViewType, true)
	}

	sql, filter := ins.getApprovedAppListQueryAndFilter(moreCondition, appViewType);
	if ins.RowCnt != 0 {
		sql += fmt.Sprintf(` LIMIT %v, %v`, ins.StartIdx, ins.RowCnt)
	}

	ret = common.DB_fetch_all(sql, filter)
	return
}

func (ins *Application) GetApprovedAppCnt(moreCondition string) (ret interface{})  {
	if (ins.LoginToken != nil)	{
		moreCondition += fmt.Sprintf(` AND app.institution_seq = %v`,  ins.LoginToken["institution_seq"])
	}

	if ins.SearchWords != "" {
		moreCondition += ins.getMoreCondition("", true)
	}

	sql := ins.getApprovedAppCountQuery(moreCondition)
	row := common.DB_fetch_one(sql, nil)
	ret = row["app_cnt"]
  return
}

func (ins *Application)LoadInspectorList() (ret interface{}) {
	sql := `SELECT IF((SELECT	member.user_seq
											 FROM	t_application_member member
											WHERE	application_seq = ?
												AND item_name IN ('general_director', 'general_expt')
												AND member.user_seq = user.user_seq) = user.user_seq,
											'참여',
											'미참여'
									  ) AS join_str,
								 IF((SELECT	app.check_user_seq
											 FROM	t_application app
											WHERE	app.application_seq = ?
										 ) = user.user_seq,
											true,
											false
									 ) AS selected,
								 user.user_seq,	user.user_type,		user.email,
								 user.name,			user.dept,				user.position,
								 user.phoneno,	user.major_field,	user.edu_course_num
					  FROM t_user user
					 WHERE user.institution_seq = ?
					 	 AND user.user_status = ?`
	filter := GetUserFilter()
	ret = common.DB_fetch_all(sql, filter, ins.Application_seq, ins.Application_seq,
																				 ins.LoginToken["institution_seq"], DEF_USER_STATUS_FINISH)
	return
}

func (ins *Application)LoadItemList(itemList []string, fileOnly bool, judge_type int) (ret map[string]interface{})  {
  ret = make(map[string]interface{})
  for _, item_name := range itemList {
    item_name = strings.TrimSpace(item_name)
    item := Item{}
    if item.GetItem(item_name) {
			if fileOnly &&
				DEF_ITEM_TYPE_ITEM_GROUP != common.ToInt(item.Data["item_type"]) &&
				DEF_ITEM_TYPE_FILE != common.ToInt(item.Data["item_type"])	{
				continue;
			}

			//
      ret[item_name] = ins.getJsonData(item_name, item, fileOnly, judge_type)
    }
  }
  return
}

func (ins *Application)setItem(item_name string, main_select interface{}, item interface{})  {
  if nil == ins.Items {
    //  assert(0)
    return
  }
  item_data := map[string]interface{} {}
  if nil != item {
    item_data["item"] = item
  }
  if nil != main_select {
    item_data["main_select"] = main_select
  }
  ins.Items[item_name] = item_data
}

// item_name의 데이터를 Load 한다.
func (ins *Application) load(item_name string, fileOnly bool, judge_type int) (succ bool)  {
  item := Item{}
  if !item.GetItem(item_name) {
    return
  }

  var main_select interface{}
  if common.ToInt(item.Data["main_select"]) > 0   {   //  Main Select 있음.
    amc := AppMainCheck { Application : ins }
    main_select = amc.Load(item_name)
  }

  switch(common.ToInt(item.Data["item_type"]))  {
  case DEF_ITEM_TYPE_ITEM_GROUP :
    ig := ItemGroup { Item_name : item_name }
    subitems := ig.GetSubItemList()
    sub_item_map := ins.LoadItemList(subitems, fileOnly, judge_type)
    ins.setItem(item_name, main_select, sub_item_map)
    return true
  case DEF_ITEM_TYPE_BASIC :
    app_basic := AppBasic{ Application : ins }
    if app_basic.Load(item_name)  {
      ins.setItem(item_name, main_select, app_basic)
    }
    return true
  case DEF_ITEM_TYPE_SELECT  :
    app_select := AppSelect{ Application : ins }
    if app_select.Load(item_name, nil)  {
      ins.setItem(item_name, main_select, app_select)
    } else {
      ins.setItem(item_name, main_select, nil)
    }
    return true
  case DEF_ITEM_TYPE_KEYWORD  :
    app_keyword := AppKeyword{ Application : ins }
    if app_keyword.Load(item_name)  {
      ins.setItem(item_name, main_select, app_keyword)
    }
    return true
  case DEF_ITEM_TYPE_FILE  :
    app_file := AppFile{ Application : ins }
    if app_file.Load(item_name)  {
      ins.setItem(item_name, main_select, app_file)
    }
    return true
  case DEF_ITEM_TYPE_MEMBER  :
    app_member := AppMember{ Application : ins }
    if app_member.Load(item_name, nil)  {
      ins.setItem(item_name, main_select, app_member)
    }
    return true
  case DEF_ITEM_TYPE_ANIMAL  :
    app_animal := AppAnimal{ Application : ins }
    if app_animal.Load(item_name)  {
      ins.setItem(item_name, main_select, app_animal)
    }
    return true
  case DEF_ITEM_TYPE_ANESTHETIC  :
    app_anesthetic := AppAnesthetic{ Application : ins }
    if app_anesthetic.Load(item_name)  {
      ins.setItem(item_name, main_select, app_anesthetic)
    }
    return true
  case DEF_ITEM_TYPE_STRING  :
    app_etc := AppEtc{ Application : ins }
    if app_etc.Load(item_name)  {
      ins.setItem(item_name, main_select, app_etc)
    }
    return true
	case DEF_ITEM_TYPE_CUSTOM1  :
    app_custom := AppCustom{ Application : ins }
    if app_custom.Load(item_name)  {
      ins.setItem(item_name, main_select, app_custom)
    }
    return true
  default :
  }
  ins.setItem(item_name, main_select, nil)
  return true
}

func (ins *Application)getJsonData(item_name string, item Item, fileOnly bool, judge_type int) (ret map[string]interface{}) {
/*  item := Item{}
  if !item.GetItem(item_name) {
    return
  }*/

  ins.load(item_name, fileOnly, judge_type)   //  데이터가 없어도 다음단계 진행 함. (format은 내려줘야 하기 때문임)

  item_info := item.GetFormatData(judge_type)

  item_data, exists := ins.Items[item_name]
  if exists  {
    out_data := map[string]interface{} {}
    main_select, ms_exists := item_data.(map[string]interface{})["main_select"]
    if ms_exists {  out_data["main_select"] = main_select   }
    data, data_exists := item_data.(map[string]interface{})["item"]
    if data_exists  {
      switch(common.ToInt(item.Data["item_type"]))  {
      case DEF_ITEM_TYPE_ITEM_GROUP  :
        item_info["sub_items"] = data
      case DEF_ITEM_TYPE_BASIC  :
        app_basic := data.(AppBasic)
        out_data["data"] = app_basic.GetJsonData()
      case DEF_ITEM_TYPE_SELECT  :
        app_select := data.(AppSelect)
        out_data["data"] = app_select.GetJsonData()
      case DEF_ITEM_TYPE_KEYWORD  :
        app_keyword := data.(AppKeyword)
        out_data["data"] = app_keyword.GetJsonData()
      case DEF_ITEM_TYPE_FILE  :
        app_file := data.(AppFile)
        out_data["data"] = app_file.GetJsonData()
      case DEF_ITEM_TYPE_MEMBER  :
        app_member := data.(AppMember)
        out_data["data"] = app_member.GetJsonData()
      case DEF_ITEM_TYPE_ANIMAL  :
        app_animal := data.(AppAnimal)
        out_data["data"] = app_animal.GetJsonData()
      case DEF_ITEM_TYPE_ANESTHETIC  :
        app_anesthetic := data.(AppAnesthetic)
        out_data["data"] = app_anesthetic.GetJsonData()
      case DEF_ITEM_TYPE_STRING  :
        app_etc := data.(AppEtc)
        out_data["data"] = app_etc.GetJsonData()
			case DEF_ITEM_TYPE_CUSTOM1  :
        app_custom := data.(AppCustom)
        out_data["data"] = app_custom.GetJsonData()
      default :
//        return    //  nil
      }
    }
    item_info["saved_data"] = out_data
  }

  return item_info
}

func (ins *Application)UpdateItemsFromJson(c *gin.Context,  data map[string]interface{}, submit bool, app_step int, app_result int, submitType uint, reg_user_seq uint) (ret bool, err_item_name string, err_msg string) {
  dbConn	:= common.DBconn()
  tx, err	:= dbConn.Begin()
  if nil != err {
		tx.Rollback()
		return
  }

  defer func() {
    tx.Rollback()
  }()

  //  신청서 번호가 없는 경우 => "application_info"(신청서 정보)가 반듯이 있어야 한다.
  if ins.Application_seq == 0 {
    for item_name, json_data := range data {
      if item_name == "application_info"  {
        app_basic := AppBasic{ Application : ins }
        ret, err_msg =  app_basic.InsertApplicationInfo(tx, json_data)
        if !ret {
          err_item_name = item_name
          return
        }
        log.Println("application_info ok : New Application_seq ->", ins.Application_seq)
        break
      }
  	}
    if !ret {
      err_item_name = "application_info"
      err_msg = "신청서 생성을 위한 기본 정보가 없습니다."
      return
    }
  } else {
    ret = true
  }

  for item_name, json_data := range data {
    if item_name == "application_info"  {
      continue
    }
    ret, err_msg = ins.UpdateItem(c, tx, item_name, json_data)
    if !ret  {
      err_item_name = item_name
      break
    }
	}
  if !ret  {
    return
  }

	log.Println("======= 0")
	//	일반심사 제출과 심사절정 변경인 경우에 => 모든 일반심사 위원이 심사를 완료 했는지 확인한다.
	if (!submit)	{
		log.Println("======= 1")
		if (submitType == DEF_SUBMIT_TYPE_JUDGE_NORMAL || 			//	일반 심사
				submitType == DEF_SUBMIT_TYPE_CHECKING_2)	{					//	심사 설정 변경
			log.Println("======= 2")
			app_select := AppSelect{ Application : ins }
			if app_select.Load("normal_review_result", tx)  {		//	일반심사 결과 로딩
				log.Println("111", app_select.Datas)
				submit = true
				items := []string { "committee_in_member", "committee_ex_member" }
				for _, item_name := range items {
					app_member := AppMember{ Application : ins }
					if !app_member.Load(item_name, tx)	{	//	심사위원 목록 추출
						log.Println("AppMember Load Error!", item_name, ins.Application_seq)
						continue
					}
					for _, user := range app_member.Datas	{
						info := user["info"].(map[string]interface{})
						log.Println(item_name, user["user_seq"], info)

						_, exists := app_select.Datas[common.ToInt(user["user_seq"])]
						if !exists {
							log.Println("NOT OK : ", user["user_seq"])
							submit = false
							break;
						}
					}
					if !submit {
						break;
					}
				}

				if (submit)	{
					app_step		= DEF_APP_STEP_FINAL
					app_result	= DEF_APP_RESULT_DECISION_ING		//	최종심의 단계로 강제 변경
				}
			}
		}
	}

	// 22-02-16 : 유지보수 사항 심사 지연인
	if (submitType == DEF_SUBMIT_TYPE_JUMP_FINAL) {
		app_member := AppMember{ Application : ins }
		if !app_member.DeleteNormalMembers(tx)	{	//	일반심사 미진행자 삭제
			err_msg = "일반심사 미진행자 삭제를 실패 했습니다"
			ret = false
			return
		}
	}

	//  제출 처리
  if submit {
		// 초기 신청서 제출시 제출일 update!
		moreUpdate := ""
		if DEF_APP_STEP_CHECKING == app_step && DEF_APP_RESULT_CHECKING == app_result {
			submit_dttm := ""
			sqlR := fmt.Sprintf(`SELECT contents FROM t_application_etc WHERE application_seq = %d AND item_name = 'reporting_date'`, ins.Application_seq)
			rowR := common.DB_Tx_fetch_one(tx, sqlR, nil)
			if rowR != nil {
				submit_dttm =  fmt.Sprintf("'%v'",common.ToStr(rowR["contents"]))
			}
			moreUpdate = fmt.Sprintf(", submit_dttm = UNIX_TIMESTAMP(%v)", submit_dttm)
		}

		// 최종심의 승인,조건부 승인 일때 승인일, 승인자 update!
		if DEF_SUBMIT_TYPE_JUDGE_FINAL_A == submitType || DEF_SUBMIT_TYPE_JUDGE_FINAL_AC == submitType {
			moreUpdate = fmt.Sprintf(`, approved_dttm = UNIX_TIMESTAMP(),
																	approved_user_seq = %v`,ins.LoginToken["user_seq"])

			application_type := common.ToInt(ins.Data["application_type"])
			// 변경신청서인 경우에, 동물의 종/수량 변경시에는 최종 종/수량상태 저장한다.
			switch application_type {
				case DEF_APP_TYPE_CHANGE : //  변경신청서
					// log.Println("======", ins.Data);
					app_animal := AppAnimal{ Application : ins }
					ret, err_msg = app_animal.UpdateFinalAnimal(tx)
					if !ret  {
						return
					}

					sql1 := `SELECT user_seq,
													(SELECT parent_app_seq FROM t_application WHERE application_seq = ?) as parent_app_seq
										 FROM t_application_member
									  WHERE item_name = 'general_director'
									    AND application_seq = ?`
					row1 := common.DB_Tx_fetch_one(tx, sql1, nil, ins.Application_seq, ins.Application_seq)
					if nil != row1["user_seq"] {
						changeHistory := ChangeHistory{	Application : ins,
																						Item_name : "writer",
																						Org_Item_seq : reg_user_seq,
																						Updated_Item_seq : common.ToUint(row1["user_seq"])}
						ret, err_msg = changeHistory.setHistory(tx)
						if !ret  {
							return
						}

						// 모든 신청서에 대해서 작성자를 update를 해준다.
						sql2 := `UPDATE t_application
						            SET reg_user_seq = ?
										  WHERE parent_app_seq = ?
											   OR application_seq = ?`
						 _, err2 := tx.Exec(sql2, row1["user_seq"], row1["parent_app_seq"], row1["parent_app_seq"])
					   if err2 != nil {
							log.Println(err2)
							err_msg = "신청서 제출을 실패했습니다."
							ret = false
					    return
					  }
					}

				case DEF_APP_TYPE_FINISH :	//  종료 보고서
					sql3 := `UPDATE t_application
					 						 SET application_result = ?
										 WHERE application_seq = ?`
					_, err3 := tx.Exec(sql3, DEF_APP_RESULT_TASK_FINISH, ins.Data["parent_app_seq"])
					if err3 != nil {
						log.Println(err3)
						err_msg = "신청서 제출을 실패했습니다."
						ret = false
						return
					}

				case DEF_APP_TYPE_BRINGIN : //  반입 보고서
					app_animal := AppAnimal{ Application : ins }
					ret, err_msg = app_animal.InsertBringAnimal(tx)
					if !ret  {
						return
					}
			}
		}

		// 최종심의 반려 나 보완후 재심일때
		if DEF_SUBMIT_TYPE_JUDGE_FINAL_REJECT == submitType || DEF_SUBMIT_TYPE_JUDGE_FINAL_REQUIRE_RETRY == submitType {
			moreUpdate = fmt.Sprintf(`, approved_dttm = UNIX_TIMESTAMP(),
																	approved_user_seq = %v`,ins.LoginToken["user_seq"])
		}

	  sql := fmt.Sprintf(`
						UPDATE t_application
	             SET application_step = %v,
							 		 application_result = %v,
									 chg_user_seq = %v
									 %v #moreUpdate
	           WHERE application_seq = %v`, app_step, app_result, ins.LoginToken["user_seq"], moreUpdate,
				 	                    ins.Application_seq)

	  _, err := tx.Exec(sql)
	  if err != nil {
			log.Println(sql)
			log.Println(err)
			err_msg = "신청서 제출을 실패했습니다."
			ret = false
	    return
	  }

		// donghun : 행정 보완 요청일때 자가 점검표를 초기화 해준다.!
		if submitType == DEF_SUBMIT_TYPE_SUPPLEMENT {
			sql2 := `DELETE FROM t_application_etc
							 	WHERE application_seq = ?
							 	 	AND item_name = 'self_inspection'`
			_, err := tx.Exec(sql2, ins.Application_seq)
		  if err != nil {
		    log.Println(err)
				err_msg = "신청서 제출을 실패했습니다."
				ret = false
		    return
		  }
		}
	}

  err = tx.Commit()   //  commit후에 rollback이 다시호출되지만 상관 없음,!!
  if nil != err {
    log.Println(err)
    ret = false
  }

  ret = true
  return
}

func (ins *Application)UpdateItem(c *gin.Context, tx *sql.Tx, item_name string, json interface{}) (ret bool, err_msg string) {
  defer func() {
    if err := recover(); err != nil {
      log.Println(err)
      err_msg = "json 포멧이 잘못되었습니다."
    }
  }()

  item := Item{}
  if !item.GetItem(item_name) {
    err_msg = "item의 정보 로딩을 실패했습니다."
    return
  }

  json_data := json.(map[string]interface{})
  data := json_data["data"]

  if common.ToInt(item.Data["main_select"]) > 0  {   //  Main Select 있음.
    main_select_map := json_data["main_select"]
    if nil == main_select_map {
      err_msg = "main_select 값이 없습니다."
      return
    }
    main_select_map2 := main_select_map.(map[string]interface{})
    amc := AppMainCheck { Application : ins }
    ret, err_msg = amc.UpdateItem(tx, item_name, main_select_map2)
    if !ret { return  }
  }

  switch(common.ToInt(item.Data["item_type"])) {
	  case DEF_ITEM_TYPE_BASIC :
	    app_basic := AppBasic{ Application : ins }
	    ret, err_msg = app_basic.UpdateItem(tx, item_name, data)
	    return
	  case DEF_ITEM_TYPE_SELECT  :
	    app_select := AppSelect{ Application : ins }
	    ret, err_msg = app_select.UpdateItem(tx, item_name, data)
	    return
	  case DEF_ITEM_TYPE_KEYWORD  :
	    app_keyword := AppKeyword{ Application : ins }
	    ret, err_msg = app_keyword.UpdateItem(tx, item_name, data)
	    return
	  case DEF_ITEM_TYPE_MEMBER  :
	    app_member := AppMember{ Application : ins }
	    ret, err_msg = app_member.UpdateItem(tx, item_name, data)
	    return
	  case DEF_ITEM_TYPE_STRING  :
	    app_etc := AppEtc{ Application : ins }
	    ret, err_msg = app_etc.UpdateItem(c, tx, item_name, data)
	    return
		case DEF_ITEM_TYPE_ANIMAL  :
	    app_animal := AppAnimal{ Application : ins }
	    ret, err_msg = app_animal.UpdateItem(tx, item_name, data)
	    return
		case DEF_ITEM_TYPE_ANESTHETIC  :
	    app_anesthetic := AppAnesthetic{ Application : ins }
	    ret, err_msg = app_anesthetic.UpdateItem(tx, item_name, data)
	    return
	  case DEF_ITEM_TYPE_FILE  :
	    app_file := AppFile{ Application : ins }
	    ret, err_msg = app_file.UpdateItem(c, tx, item_name, data)
	    return
  }

  ret = true
  return
}

func (ins *Application)UpdatePrevStep(c *gin.Context) (ret bool) {
	sql := `UPDATE t_application
						 SET application_result = ?, chg_user_seq = ?
					 WHERE application_seq = ?`
	_, err := common.DBconn().Exec(sql, DEF_APP_RESULT_CHECKING, ins.LoginToken["user_seq"], ins.Application_seq)
 	if nil != err {
 		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
 		log.Println(err)
 		return
 	}

	return true
}

func (ins *Application)UpdateInspector(c *gin.Context, check_user_seq uint64) (succ bool){
	succ = false
	sql := `UPDATE t_application
						 SET chg_user_seq = ?,
						 		 check_user_seq = ?
					 WHERE application_seq = ?`
	_, err := common.DBconn().Exec(sql, ins.LoginToken["user_seq"], check_user_seq, ins.Application_seq)
 	if nil != err {
 		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
 		log.Println(err)
 		return
 	}

	return true
}

func (ins *Application)UpdateStepAndResult(app_step int, app_result int) (ret bool)	{
	sql := `UPDATE t_application
	           SET application_step = ?, application_result = ?
	         WHERE application_seq = ?`
	_, err := common.DBconn().Exec(sql, app_step, app_result, ins.Application_seq)
	if err != nil {
		log.Println(err)
		return
	}

	return true;
}

func (ins *Application)GetUnixtimeForExpertDeadline() (usecs int64)	{
	app_etc := AppEtc{ Application : ins }
	if !app_etc.Load("judge_expert_deadline")  {   //  심사종료일
		return;
	}

	value := common.ToStr(app_etc.Data["0"])
	if value == "" {
		return
	}

	addTime := ":00+09:00"

	// donghun : 초가 붙어서 나오는 경우가 있음!
	if len(value) == 19 {
		addTime = "+09:00"
	}

	t1, e := time.Parse(time.RFC3339, (value + addTime))
	if nil != e {
		log.Println(e)
		return;
	}
	usecs = t1.Unix();
	return;
}

func (ins *Application) GetUnixtimeForNormalDeadline() (usecs int64)	{
	// expert_deadline := ins.GetUnixtimeForExpertDeadline();
	// if expert_deadline == 0	{
	// 	return
	// }

	// donghun : 21-04-22 전문 심사 완료후 시간이어야됨!
	app_member := AppMember{ Application : ins }
	if !app_member.Load("expert_member", nil){
		return
	}
	expert_member_seq := app_member.Datas[0]["user_seq"]

	app_etc2 := AppEtc{ Application : ins }
	if !app_etc2.Load("expert_review_dttm")  {   //  전문심사 완료 일시
		return
	}
	expert_review_dttm := common.ToInt64(app_etc2.Data[common.ToStr(expert_member_seq)])

	app_etc := AppEtc{ Application : ins }
	if !app_etc.Load("judge_normal_deadline")  {   //  일반심사 추가 시간
		return
	}

	normal_gap := common.ToStr(app_etc.Data["0"])
	if len(normal_gap) < 0	{
		return
	}

	usecs = expert_review_dttm
	value := common.ToInt64(normal_gap[0:len(normal_gap)-1])
	dayhour := normal_gap[len(normal_gap)-1:]
	if dayhour == "H" || dayhour == "h"	{
		usecs = usecs + (value * 60 * 60)
	} else {
		usecs = usecs + (value * 60 * 60 * 24)
	}

	// log.Println(normal_gap, value, dayhour, usecs)
	return;
}

func (ins *Application) GetTimeFormatForExpertDeadline(format string) (ret string)	{
	deadline := ins.GetUnixtimeForExpertDeadline();
	timeT := time.Unix(deadline, 0)
	loc, _ := time.LoadLocation("Asia/Seoul")
	t := timeT.In(loc)
	ret = t.Format(format);
	return;
}

func (ins *Application) GetTimeFormatForNormalDeadline(format string) (ret string)	{
	deadline := ins.GetUnixtimeForNormalDeadline();
	timeT := time.Unix(deadline, 0)
	loc, _ := time.LoadLocation("Asia/Seoul")
	t := timeT.In(loc)
	ret = t.Format(format);
	return;
}

func (ins *Application)CheckSendMsg(msg_id int) (succ bool) {
  sql := fmt.Sprintf(`
          SELECT application_seq app_seq
            FROM t_application_select
           WHERE application_seq = %v
             AND item_name = 'judge_alarm'
             AND select_ids LIKE '%%%v%%'
             `, ins.Application_seq, msg_id )
  row := common.DB_fetch_one(sql, nil)
  if common.ToUint(row["app_seq"]) == ins.Application_seq {
    return true
  }

  return
}

func (app *Application)GetGeneralDirectorSeq(reg_user_seq uint) (user_seq uint) {
	user_seq = reg_user_seq
 	sql := fmt.Sprintf(`SELECT application_seq, parent_app_seq
											  FROM t_application
											 WHERE application_seq = ?`)
	row := common.DB_fetch_one(sql, nil, app.Application_seq)
	if nil != row {
		app_seq := uint(0)
		if common.ToUint(row["parent_app_seq"]) > 0 {
			app_seq = common.ToUint(row["parent_app_seq"])
		} else {
			app_seq = app.Application_seq
		}

		sql = fmt.Sprintf(`
						SELECT member.user_seq
							FROM t_application app, t_application_member member
						 WHERE app.application_seq = member.application_seq
							 AND app.parent_app_seq = %d
							 AND app.application_type = %d
							 AND app.application_step = %d
							 AND app.application_result IN (%d, %d)
							 AND member.item_name = 'general_director'
					ORDER BY app.approved_dttm DESC
						 LIMIT 1`, app_seq, DEF_APP_TYPE_CHANGE, DEF_APP_STEP_FINAL,
							DEF_APP_RESULT_APPROVED, DEF_APP_RESULT_APPROVED_C)
		row = common.DB_fetch_one(sql, nil)
		if nil != row {
			user_seq = common.ToUint(row["use_seq"])
		}
	}
	return
}
