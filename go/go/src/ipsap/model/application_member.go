
package model

import (
	"github.com/gin-gonic/gin"
  "ipsap/common"
  "database/sql"
  "log"
  "fmt"
)

type AppMember struct {
  Application         *Application
  Datas               []map[string]interface{}   //  []string
}

func (ins *AppMember)Load(item_name string, tx *sql.Tx) (succ bool) {
  if ins.Application.Application_seq == 0 {
    return
  }

  if ins.Datas != nil {
    return true
  }

   app_seq := ins.Application.Application_seq

  // 변경 승인 신청서에서 최종 (승인, 조건부 승인)이 되었는지 확인 하고  실제 내역 반영
  if item_name == "general_director" || item_name == "general_expt" {
    sql2 := fmt.Sprintf(`
            SELECT app.application_seq
              FROM t_application app, t_application_member member
             WHERE app.application_seq = member.application_seq
               AND app.parent_app_seq = %d
               AND app.application_type = %d
               AND app.application_step = %d
               AND app.application_result IN (%d, %d)
               AND member.item_name = '%s'
          ORDER BY app.approved_dttm DESC
             LIMIT 1`, app_seq, DEF_APP_TYPE_CHANGE, DEF_APP_STEP_FINAL,
              DEF_APP_RESULT_APPROVED, DEF_APP_RESULT_APPROVED_C, item_name)
    row := make(map[string]interface{}, 0)
    if tx != nil {
      row = common.DB_Tx_fetch_one(tx, sql2, nil)
    } else {
      row = common.DB_fetch_one(sql2, nil)
    }

    if nil != row["application_seq"] {
      app_seq = common.ToUint(row["application_seq"])
    }
  }

  sql := `SELECT user_seq, animal_mng_flag, exp_year_code, exp_type_code, edu_course, chk_flag
            FROM t_application_member
           WHERE application_seq = ?
             AND item_name = ?`
  filter := func(row map[string]interface{}) {
    switch(item_name) {
      case "expert_member", "committee_in_member", "committee_ex_member" :
        delete(row, "exp_year_code")
        delete(row, "animal_mng_flag")
      default:
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
        row["app_seq"] = app_seq
    }
  }

  if tx != nil {
    ins.Datas = common.DB_Tx_fetch_all(tx, sql, filter, app_seq, item_name)
  } else {
    ins.Datas = common.DB_fetch_all(sql, filter, app_seq, item_name)
  }

  if nil != ins.Datas {
    for _, row := range ins.Datas {
      user := User {  User_seq : common.ToUint(row["user_seq"]) }
      if user.Load()  {
        row["info"] = user.Data
      } else {
        if item_name == "expert_member" || item_name == "committee_in_member" || item_name == "committee_ex_member"{
          row["info"] = gin.H{
            "name" : "탈퇴자",
            "user_seq" : row["user_seq"],
          }
        }
      }
    }

    //  연구 책임자 정보가 없으면 생성해서 내려 줌!
    if (item_name == `general_director` && len(ins.Datas) == 0) {
      data := map[string]interface{} {
        "user_seq" : ins.Application.LoginToken["user_seq"],
        "animal_mng_flag" : 0,
        "exp_year_code" : 0,
      }
      user := User {  User_seq : common.ToUint(data["user_seq"]) }
      if user.Load()  {
        data["info"] = user.Data
      }

      ins.Datas = []map[string]interface{} { data }
    }
  }

  return true
}

// 일반 심사를 완료 하지않은 심사위원 리스트!
func (ins *AppMember) GetToSendNormalMembers() (succ bool){
  succ = false
  if ins.Application.Application_seq == 0 {
    return
  }

  app_seq := ins.Application.Application_seq

  sql := `SELECT user_seq
            FROM t_application_member
           WHERE application_seq = ?
	           AND item_name IN ('committee_in_member', 'committee_ex_member')
	           AND user_seq NOT IN (SELECT target_item
                                    FROM t_application_etc
		                               WHERE application_seq = ?
			                               AND item_name = 'normal_review_dttm')`
  filter := func(row map[string]interface{}) {
    row["app_seq"] = app_seq
  }

  ins.Datas = common.DB_fetch_all(sql, filter, app_seq, app_seq)
  if nil != ins.Datas || 0 > len(ins.Datas) {
    for _, row := range ins.Datas {
      user := User {  User_seq : common.ToUint(row["user_seq"]) }
      if user.Load()  {
        row["info"] = user.Data
      }
    }
  }

  return true
}

func (ins *AppMember) DeleteNormalMembers(tx *sql.Tx) (succ bool) {
  sql := `DELETE FROM t_application_member
           WHERE application_seq = ?
  	         AND item_name IN ('committee_in_member', 'committee_ex_member')
  	         AND user_seq NOT IN (SELECT target_item
                                      FROM t_application_etc
  		                               WHERE application_seq = ?
  			                               AND item_name = 'normal_review_dttm')`
   _, err := tx.Exec(sql, ins.Application.Application_seq, ins.Application.Application_seq)
   if err != nil {
     log.Println(err)
     return false
   }

   return true
}

func (ins *AppMember)GetJsonData() (ret interface{}) {
  ret = ins.Datas
  return
}

func (ins *AppMember)UpdateItem(tx *sql.Tx, item_name string, data1 interface{}) (ret bool, err_msg string) {
  data := data1.([]interface{})

  sql := `UPDATE t_application_member SET chk_flag = chk_flag + 100
          WHERE application_seq = ?
            AND item_name = ?`
  _, err := tx.Exec(sql, ins.Application.Application_seq, item_name)
  if err != nil {
    log.Println(err)
    return
  }

  for _, user_info := range data  {
    user_info_map := user_info.(map[string]interface{})
    var must_has_key []string
    switch(item_name) {
      case "general_expt" : // 21-08-24 donghun 수정 animal_mng_flag 공통이 아님!
        must_has_key = []string { "exp_year_code", "user_seq" }
        if common.ToUint(user_info_map["exp_type_code"]) == 0 {user_info_map["exp_type_code"] = 0}
        if common.ToUint(user_info_map["animal_mng_flag"]) == 0 {user_info_map["animal_mng_flag"] = 0}
      case "general_director" :
        must_has_key = []string { "exp_year_code", "user_seq" }
        user_info_map["animal_mng_flag"] = 0
        user_info_map["exp_type_code"] = 0
      case "expert_member", "committee_in_member", "committee_ex_member" :
        must_has_key = []string { "user_seq" }
        user_info_map["animal_mng_flag"] = 0
        user_info_map["exp_year_code"] = 0
        user_info_map["exp_type_code"] = 0
    }

    if !common.CheckHasMustKey(user_info_map, must_has_key) {
      err_msg = fmt.Sprintf("필수 key가 없습니다.(%v)", must_has_key)
      return
    }

    if ("" == common.ToStr(user_info_map["edu_course"])) {
      user_info_map["edu_course"] = ""
    }

    sql  = `INSERT INTO t_application_member(application_seq, item_name, user_seq, animal_mng_flag, exp_year_code, exp_type_code, edu_course)
            VALUES(?, ?, ?, ?, ?, ?, ?)
            ON DUPLICATE KEY UPDATE animal_mng_flag = ?, exp_year_code = ?, exp_type_code = ?, edu_course = ?, chk_flag = chk_flag - 100`
    _, err := tx.Exec(sql, ins.Application.Application_seq, item_name,
                      user_info_map["user_seq"], user_info_map["animal_mng_flag"], user_info_map["exp_year_code"], user_info_map["exp_type_code"], user_info_map["edu_course"],
                      user_info_map["animal_mng_flag"], user_info_map["exp_year_code"], user_info_map["exp_type_code"], user_info_map["edu_course"])
    if err != nil {
      err_msg = "사용자정보 값이 잘못 되었습니다."
      log.Println(err)
      return
    }
  }

  sql = `DELETE FROM t_application_member
          WHERE application_seq = ?
            AND item_name = ?
            AND chk_flag >= 100`
  _, err = tx.Exec(sql, ins.Application.Application_seq, item_name)
  if err != nil {
    log.Println(err)
    return
  }

  ret = true
  return
}

func (ins *AppMember)SetChkFlag(item_name string, user_seq interface{}, flagId int) (succ bool) {
  flagValue := 1 << flagId;
  sql := `UPDATE t_application_member set chk_flag = chk_flag | ?
          WHERE application_seq = ?
            AND item_name = ?
            AND user_seq = ?`
  _, err := common.DBconn().Exec(sql, flagValue, ins.Application.Application_seq, item_name, user_seq)
  if err != nil {
    log.Println(err)
    return
  }

  return true;
}
