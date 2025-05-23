
package model

import (
	"github.com/gin-gonic/gin"
  "ipsap/common"
	"fmt"
	"strings"
	// "log"
)

func (app *Application)ChangeInfo() (result gin.H) {
  sql := `SELECT group_concat(application_seq) as changeAppSeqs,
	 							 group_concat(DISTINCT FROM_UNIXTIME(approved_dttm, '%Y-%m-%d')) as approvedDttms
	          FROM t_application
					 WHERE parent_app_seq = ?
					   AND application_type = ?
						 AND application_step = ?
						 AND application_result IN (?,?)
 						 AND approved_dttm > 0
			  ORDER BY approved_dttm DESC`
	row := common.DB_fetch_one(sql, nil, app.Application_seq, DEF_APP_TYPE_CHANGE, DEF_APP_STEP_FINAL, DEF_APP_RESULT_APPROVED, DEF_APP_RESULT_APPROVED_C)
  if nil != row["changeAppSeqs"] {
    changeAppSeqs := row["changeAppSeqs"]
    sql = fmt.Sprintf(`
      SELECT etc.application_seq, etc.item_name, etc.target_item, etc.contents,
             FROM_UNIXTIME(app.approved_dttm, '%%Y-%%m-%%d') as approved_dttm,
             IFNULL((SELECT filepath FROM t_application_file WHERE application_seq = etc.application_seq AND file_idx  = etc.target_item), '') as filepath,
             IFNULL((SELECT org_file_name FROM t_application_file WHERE application_seq = etc.application_seq AND file_idx  = etc.target_item), '')  as org_file_name,
					 	 (SELECT contents FROM t_questionnaire WHERE is_seq = etc.target_item) as main_title
        FROM t_application_etc etc, t_application app
       WHERE etc.application_seq = app.application_seq
         AND etc.application_seq IN (%v)
         AND etc.item_name IN ('ca_regular_item', 'ca_fast_item', 'ibc_ca_item')
    ORDER BY app.approved_dttm`, changeAppSeqs)
    filter := func(row map[string]interface{}) {
			row["src"] = ""
			if "" != common.ToStr(row["filepath"]) {
				filepath := common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["filepath"]))
				row["src"] = common.MakeDownloadUrl(filepath)
			}
			delete(row, "filepath")

			if common.ToInt(row["target_item"]) == 111 || common.ToInt(row["target_item"]) == 286{
					sql2 := fmt.Sprintf(`
									SELECT user_seq, animal_mng_flag, exp_year_code,
												 IF(item_name = 'general_director', true, false) as general_director
										FROM t_application_member
									 WHERE item_name IN ('general_director', 'general_expt')
										 AND application_seq = ?`)
					rows2 := common.DB_fetch_all(sql2, nil, row["application_seq"])

					for _, row2 := range rows2 {
						user := User {  User_seq : common.ToUint(row2["user_seq"]) }
						if user.Load()  {
							row2["info"] = user.Data
						}
					}

					row["member_after"] = rows2
					tmpArr := strings.Split(common.ToStr(row["contents"]), ",")
					rows3 := common.DB_fetch_all(sql2, nil, strings.Trim(tmpArr[1], "\""))
					for _, row3 := range rows3 {
						user := User {  User_seq : common.ToUint(row3["user_seq"]) }
						if user.Load()  {
							row3["info"] = user.Data
						}
					}

					row["member_org"] = rows3
			}

			if common.ToInt(row["target_item"]) == 112 || common.ToInt(row["target_item"]) == 288{
				sql1 :=  fmt.Sprintf(`
							 SELECT etc.contents
								 FROM t_application_etc etc
								WHERE etc.item_name = 'general_end_date'
									AND etc.application_seq = ?`)
				row1 := common.DB_fetch_one(sql1, nil,  row["application_seq"])
				row["end_date_after"] = row1["contents"]

				tmpArr := strings.Split(common.ToStr(row["contents"]), ",")
				row4 := common.DB_fetch_one(sql1, nil, strings.Trim(tmpArr[1], "\""))
				row["end_date_org"] = row4["contents"]
			}
		}

		change_apps := common.DB_fetch_all(sql, filter)

		result = gin.H {
			"change_apps" : change_apps,
			"approved_dttm" : strings.Split(common.ToStr(row["approvedDttms"]), ","),
		}
	}
	return
}

func (app *Application)Change_Animal_Info() (rows []map[string]interface{}) {
	sql := `SELECT group_concat(application_seq) as changeAppSeqs
						FROM t_application
					 WHERE parent_app_seq = ?
						 AND application_type = ?
						 AND application_step = ?
						 AND application_result IN (?,?)
						 AND approved_dttm > 0`
	row := common.DB_fetch_one(sql, nil, app.Application_seq, DEF_APP_TYPE_CHANGE, DEF_APP_STEP_FINAL, DEF_APP_RESULT_APPROVED, DEF_APP_RESULT_APPROVED_C)
	if nil != row["changeAppSeqs"] {
		changeAppSeqs := row["changeAppSeqs"]
		sql = fmt.Sprintf(`
			SELECT etc.application_seq, etc.item_name, etc.target_item, etc.contents,
						 FROM_UNIXTIME(app.approved_dttm, '%%Y-%%m-%%d') as approved_dttm,
						 IFNULL((SELECT filepath FROM t_application_file WHERE application_seq = etc.application_seq AND file_idx  = etc.target_item), '') as filepath,
						 IFNULL((SELECT org_file_name FROM t_application_file WHERE application_seq = etc.application_seq AND file_idx  = etc.target_item), '')  as org_file_name
				FROM t_application_etc etc, t_application app
			 WHERE etc.application_seq = app.application_seq
				 AND etc.application_seq IN (%v)
				 AND etc.item_name IN ('ca_regular_item')
				 AND etc.target_item = 104
		ORDER BY app.approved_dttm`, changeAppSeqs)
		filter := func(row map[string]interface{}) {
			row["filepath"] = common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["filepath"]))
			row["src"] 	= common.MakeDownloadUrl(common.ToStr(row["filepath"]))
		}
		rows = common.DB_fetch_all(sql, filter)
	}
	return
}

func (app *Application)Change_Member_Info() (rows []map[string]interface{}) {
	sql := `SELECT user_seq, animal_mng_flag, exp_year_code, exp_type_code, IF(item_name = 'general_director', true, false) as general_director, edu_course
						FROM t_application_member
					 WHERE application_seq = ?
						 AND item_name IN ('general_director', 'general_expt')`
	filter := func(row map[string]interface{}) {
		code := Code {
			Type : "exp_year_code",
			Id : common.ToUint(row["exp_year_code"]),
		}

		code2 := Code {
			Type : "exp_type_code",
			Id : common.ToUint(row["exp_type_code"]),
		}
		row["exp_year_code_str"] = code.GetCodeStrFromTypeAndId()
		row["exp_type_code_str"] = code2.GetCodeStrFromTypeAndId()
	}

	rows = common.DB_fetch_all(sql, filter, app.Application_seq)
	if nil != rows {
		for _, row := range rows {
			user := User {  User_seq : common.ToUint(row["user_seq"]) }
			if user.Load()  {
				row["info"] = user.Data
			}
		}
  }
	return
}

func (app *Application)Change_EndDate_Info() (rows []map[string]interface{}) {
	sql := `SELECT target_item, contents
						FROM t_application_etc
					 WHERE application_seq = ?
						 AND item_name = 'general_end_date'`
	rows = common.DB_fetch_all(sql, nil, app.Application_seq)
	for _, row := range rows {
		row[common.ToStr(row["target_item"])] = common.ToStr(row["contents"])
		delete(row, "target_item")
		delete(row, "contents")
	}
	return
}
