
package model

//  기타(문자열로 저장이 가능한 정보들) 관리
//  - 자가점검표

import (
  "github.com/gin-gonic/gin"
  "ipsap/common"
  "database/sql"
  "strings"
  "log"
  "fmt"
)

type AppEtc struct {
  Application         *Application
  Data                map[string]string
}

func (ins *AppEtc)Load(item_name string) (succ bool) {
  if ins.Application.Application_seq == 0 {
    return
  }

  if ins.Data != nil {
    return true
  }

  ins.Data = map[string]string {}

  // 변경 승인 신청서에서 최종 (승인, 조건부 승인)이 되었는지 확인 하고  실제 내역 반영
  app_seq := ins.Application.Application_seq
  if item_name == "general_end_date" {
    sql2 := fmt.Sprintf(`
            SELECT app.application_seq
              FROM t_application app, t_application_etc etc
             WHERE app.application_seq = etc.application_seq
               AND app.parent_app_seq = %d
               AND app.application_type = %d
               AND app.application_step = %d
               AND app.application_result IN (%d, %d)
               AND etc.item_name = '%s'
          ORDER BY app.approved_dttm DESC
             LIMIT 1`, app_seq, DEF_APP_TYPE_CHANGE, DEF_APP_STEP_FINAL,
              DEF_APP_RESULT_APPROVED, DEF_APP_RESULT_APPROVED_C, item_name)
    row := common.DB_fetch_one(sql2, nil)
    if nil != row["application_seq"] {
      app_seq = common.ToUint(row["application_seq"])
    }
  }

  sql := `SELECT target_item, contents
            FROM t_application_etc
           WHERE application_seq = ?
             AND item_name = ?`
  rows := common.DB_fetch_all(sql, nil, app_seq, item_name)
  for _, row := range rows {
    ins.Data[common.ToStr(row["target_item"])] = common.ToStr(row["contents"])
    if item_name == "general_end_date" {
      ins.Data["cur_app_seq"] = common.ToStr(app_seq)
    }
  }

//  log.Println(ins.Data)
  return true
}

func (ins *AppEtc)GetJsonData() (ret interface{}) {
  ret = ins.Data
  return
}

func (ins *AppEtc)UpdateItem(c *gin.Context, tx *sql.Tx, item_name string, data1 interface{}) (ret bool, err_msg string) {
  append_items := []string{ "expert_review_dttm", "normal_review_dttm", "ibc_review_report"}
  del_flag := true
  for _, item := range append_items {
    if item == item_name {
      del_flag = false
      break;
    }
  }

  data := data1.(map[string]interface{})
  if del_flag {
    sql := `DELETE FROM t_application_etc
             WHERE application_seq = ?
               AND item_name = ?`
    _, err := tx.Exec(sql, ins.Application.Application_seq, item_name)
    if err != nil {
      log.Println(err)
      return
    }
  }

  remain_files := []string {}
  for key, value := range data  {
    sql  := `INSERT INTO t_application_etc(application_seq, item_name, target_item, contents, reg_dttm, reg_user_seq)
            VALUES(?, ?, ?, ?, UNIX_TIMESTAMP(), ?)
            ON DUPLICATE KEY UPDATE contents = ?`
    _, err := tx.Exec(sql, ins.Application.Application_seq, item_name,
                      key, value, ins.Application.LoginToken["user_seq"], value)
    if err != nil {
      err_msg = "data 값이 잘못 되었습니다."
      log.Println(err)
      return
    }
    var file_exists bool

    switch (item_name) {
    case "ca_regular_item" :
      app_file := AppFile{ Application : ins.Application }
	    ret, file_exists = app_file.UploadFile(c, tx, "ca_regular_file", common.ToInt(key))
      if !ret {
        err_msg = fmt.Sprintf("파일 전송을 실패했습니다.(%v)", key);
        log.Println(err_msg)
        return
      }
      if file_exists  { remain_files = append(remain_files, common.ToStr(key))  }
    case "ca_fast_item" :
      app_file := AppFile{ Application : ins.Application }
	    ret, file_exists = app_file.UploadFile(c, tx, "ca_fast_file", common.ToInt(key))
      if !ret {
        err_msg = fmt.Sprintf("파일 전송을 실패했습니다.(%v)", key);
        log.Println(err_msg)
        return
      }
      if file_exists  { remain_files = append(remain_files, common.ToStr(key))  }
    case "ibc_ca_item" :
      app_file := AppFile{ Application : ins.Application }
      ret, file_exists = app_file.UploadFile(c, tx, "ibc_ca_file", common.ToInt(key))
      if !ret {
        err_msg = fmt.Sprintf("파일 전송을 실패했습니다.(%v)", key);
        log.Println(err_msg)
        return
      }
      if file_exists  { remain_files = append(remain_files, common.ToStr(key))  }
    case "irb_ca_item" :
      app_file := AppFile{ Application : ins.Application }
      ret, file_exists = app_file.UploadFile(c, tx, "irb_ca_file", common.ToInt(key))
      if !ret {
        err_msg = fmt.Sprintf("파일 전송을 실패했습니다.(%v)", key);
        log.Println(err_msg)
        return
      }
      if file_exists  { remain_files = append(remain_files, common.ToStr(key))  }
    }
  }

  switch (item_name) {
  case "ca_regular_item", "ca_fast_item", "ibc_ca_item", "irb_ca_item" :
    item_name2 := "ca_regular_file"
    if item_name == "ca_fast_item"  {
      item_name2 = "ca_fast_file"
    } else if item_name == "ibc_ca_item" {
      item_name2 = "ibc_ca_file"
    } else if item_name == "irb_ca_item" {
      item_name2 = "irb_ca_item"
    }

    ret = false
    sql :=  `DELETE FROM t_application_file
             WHERE application_seq = ? AND item_name = ?`
    if len(remain_files) > 0  {
      sql += fmt.Sprintf(` AND file_idx not in (%v)`, strings.Join(remain_files, ","))
    }
    _, err := tx.Exec(sql, ins.Application.Application_seq, item_name2)
    if err != nil {
      log.Println(err)
      return
    }
  }
  ret = true
  return
}
