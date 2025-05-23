
package model

import (
  "ipsap/common"
  "database/sql"
  "log"
  "fmt"
)

type AppBasic struct {
  Application         *Application
  Data                map[string]interface{}
}

func (ins *AppBasic)Load(item_name string) (succ bool) {
  if ins.Application.Application_seq == 0 {
    return
  }

  if ins.Data != nil {
    return true
  }

  sql := `SELECT application_no, judge_type, application_type,
                 application_step, application_result, app.name_ko,
	               app.name_en, approved_dttm , IFNULL(user.name,'') as approved_user
            FROM t_application app
            LEFT OUTER JOIN t_user user ON (app.approved_user_seq = user.user_seq)
           WHERE application_seq = ?`
  filter := func(row map[string]interface{}) {
    switch(item_name) {
    case "application_info" :
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

      delete(row, "name_ko")
      delete(row, "name_en")
    case "general_title" :
      delete(row, "judge_type")
      delete(row, "application_no")
      delete(row, "application_type")
      delete(row, "application_result")
      delete(row, "application_step")
    }
  }
  ins.Data = common.DB_fetch_one(sql, filter, ins.Application.Application_seq)
  if nil == ins.Data {
    return
  }

  return true
}

func (ins *AppBasic)GetJsonData() (ret interface{}) {
  ret = ins.Data
  return
}

func (ins *AppBasic)GetAppNo(tx *sql.Tx, judge_type int, application_type int, parent_app_seq uint) (succ bool, app_no string) {
  if parent_app_seq == 0  { //  신규승인
    if DEF_APP_TYPE_NEW != application_type {
      return    //  error
    }
    // donghun : 21-05-14 새로운 년도마다 app_cnt 변경됨!
    // 21-06-29 삭제된거 app_cnt 반영
    sql := `SELECT ins.institution_code,
                  (SELECT LPAD(COUNT(application_seq) + 1 +
                  (SELECT COUNT(application_seq)
                     FROM t_application_deleted
                    WHERE institution_seq = ?
                      AND judge_type = ?
                      AND parent_app_seq = 0
                      AND substring(application_no, INSTR(application_no, "-") + 1, 4) = (SELECT YEAR(CURDATE()))), '4' , '0')
                     FROM t_application
                    WHERE institution_seq = ?
                      AND judge_type = ?
                      AND parent_app_seq = 0
                      AND substring(application_no, INSTR(application_no, "-") + 1, 4) = (SELECT YEAR(CURDATE()))) as app_cnt,
                  (SELECT YEAR(CURDATE())) as cur_year
             FROM t_institution ins
            WHERE ins.institution_seq = ?`
    row := common.DB_Tx_fetch_one(tx, sql, nil,
                                  ins.Application.LoginToken["institution_seq"],
                                  judge_type,
                                  ins.Application.LoginToken["institution_seq"],
                                  judge_type,
                                  ins.Application.LoginToken["institution_seq"])
    if nil == row {
      return
    }
    judge_type_str := ""
    switch judge_type {
    case DEF_APP_JUDGE_TYPE_CODE_IACUC:
      judge_type_str = DEF_APP_NO_IACUC
    case DEF_APP_JUDGE_TYPE_CODE_IBC:
      judge_type_str = DEF_APP_NO_IBC
    case DEF_APP_JUDGE_TYPE_CODE_IRB:
      judge_type_str = DEF_APP_NO_IRB
    }
    app_no = fmt.Sprintf("%v-%v-%v%v-00", row["institution_code"], row["cur_year"], judge_type_str, row["app_cnt"])
  } else {
    sql := `SELECT application_no, judge_type
              FROM t_application
             WHERE application_seq = ?
               AND application_type = ?`
    row := common.DB_Tx_fetch_one(tx, sql, nil, parent_app_seq, DEF_APP_TYPE_NEW);
    if row == nil {
      return
    }
    if common.ToInt(row["judge_type"]) != judge_type  {
      return    //  심의유형 다름 : Error
    }
    app_no = common.ToStr(row["application_no"]);
    app_no = app_no[0:len(app_no)-3]

    sql = `SELECT count(*) + 1 as cnt
             FROM t_application
            WHERE parent_app_seq = ?
              AND judge_type = ?
              AND application_type = ?`
    row = common.DB_Tx_fetch_one(tx, sql, nil, parent_app_seq, judge_type, application_type);
    if row == nil {
      return
    }
    serial_no := common.ToStr(row["cnt"])

    switch application_type {
    case DEF_APP_TYPE_CHANGE     : //  변경신청서
      app_no = app_no + "-M" + serial_no
    case DEF_APP_TYPE_RENEW, DEF_APP_TYPE_CONTINUE : //  재승인 or IRB 지속심의(중간보고)
      app_no = app_no + "-R" + serial_no
    case DEF_APP_TYPE_BRINGIN    : //  반입신청서
      app_no = app_no + "-P" + serial_no
    case DEF_APP_TYPE_CHECKLIST  : //  승인후 점검표
      app_no = app_no + "-E" + serial_no
    case DEF_APP_TYPE_SERIOUS    : //  IRB 중대한 이상 반응 보고서
      app_no = app_no + "-S" + serial_no
    case DEF_APP_TYPE_VIOLATION  : //  IRB 연구계획 위반/이탈 보고
      app_no = app_no + "-V" + serial_no
    case DEF_APP_TYPE_UNEXPECTED : //  IRB 예상치 못한 문제발생 보고서
      app_no = app_no + "-U" + serial_no
    case DEF_APP_TYPE_FINISH     : //  종료 보고서
      app_no = app_no + "-FF"
    default :
      return      //  Error
    }
  }

  succ = true
  return
}

func (ins *AppBasic)InsertApplicationInfo(tx *sql.Tx, json interface{}) (ret bool, err_msg string) {
  defer func() {
    if err := recover(); err != nil {
      log.Println(err)
      err_msg = "json 포멧이 잘못되었습니다."
    }
  }()

  data := json.(map[string]interface{})["data"].(map[string]interface{})
  log.Println("data :", data)

  must_has_key := []string { "application_type", "judge_type" }

  if !common.CheckHasMustKey(data, must_has_key)  {
    err_msg = fmt.Sprintf("필수 key가 없습니다.(%v)", must_has_key)
    return
  }
  parent_app_seq := 0;
  name_ko := ``;
  name_en := ``;

  switch(common.ToInt(data["application_type"])) {
  case DEF_APP_TYPE_CHANGE, DEF_APP_TYPE_RENEW, DEF_APP_TYPE_BRINGIN, DEF_APP_TYPE_CHECKLIST, DEF_APP_TYPE_FINISH,
       DEF_APP_TYPE_CONTINUE, DEF_APP_TYPE_SERIOUS, DEF_APP_TYPE_VIOLATION, DEF_APP_TYPE_UNEXPECTED :
    more_chk_key := []string { "parent_app_seq" }
    if !common.CheckHasMustKey(data, more_chk_key)  {
      err_msg = fmt.Sprintf("필수 key가 없습니다.(%v)", more_chk_key)
      return
    }
    parent_app_seq = common.ToInt(data["parent_app_seq"])

    p_app := Application {  Application_seq : uint(parent_app_seq) }
    row := p_app.Load().(map[string]interface{})
    if (row == nil) {
      err_msg = fmt.Sprintf("신규 신청서 정보가 잘못되었습니다.(%v)", parent_app_seq)
      return
    }

    //  신청서 : parent_app_seq가 신규승인 신청서 이고, 과제 진행중이어야만 한다.
    if (common.ToInt(row["application_type"]) != DEF_APP_TYPE_NEW ||
        !(common.ToInt(row["application_result"]) == DEF_APP_RESULT_EXPER_ING_A ||
          common.ToInt(row["application_result"]) == DEF_APP_RESULT_EXPER_ING_AC))  {
      err_msg = fmt.Sprintf("신규 신청서의 상태가 진행중이 아닙니다.(%v)", parent_app_seq)
      return
    }

    name_ko = common.ToStr(row["name_ko"])
    name_en = common.ToStr(row["name_en"])
  }

  succ, application_no := ins.GetAppNo(tx,  common.ToInt(data["judge_type"]),
                                            common.ToInt(data["application_type"]),
                                            uint(parent_app_seq))
  if !succ {
    return
  }
  sql := `INSERT INTO t_application(application_no, parent_app_seq, institution_seq, judge_type, application_type,
                                    name_ko, name_en, reg_dttm, reg_user_seq, chg_user_seq)
          VALUES(?, ?, ?, ?, ?,
                 ?, ?, UNIX_TIMESTAMP(), ?, ?)`
  result, err := tx.Exec(sql, application_no, parent_app_seq,
                         ins.Application.LoginToken["institution_seq"],
                         data["judge_type"], data["application_type"],
                         name_ko, name_en,
                         ins.Application.LoginToken["user_seq"],
                         ins.Application.LoginToken["user_seq"])
  if err != nil {
    log.Println(err)
    return
  }

  no, err := result.LastInsertId()
  if err != nil {
   log.Println(err)
   return
  }

  ins.Application.Application_seq = common.ToUint(no)

  ret = true
  return
}

func (ins *AppBasic)UpdateItem(tx *sql.Tx, item_name string, data1 interface{}) (ret bool, err_msg string) {
  data := data1.(map[string]interface{})

  switch(item_name) {
  case "general_title"  :
    must_has_key := []string { "name_en", "name_ko" }
    if !common.CheckHasMustKey(data, must_has_key)  {
      err_msg = fmt.Sprintf("필수 key가 없습니다.(%v)", must_has_key)
      return
    }
    sql := `UPDATE t_application
            SET name_ko = ?, name_en = ?
            WHERE application_seq = ?`
    _, err := tx.Exec(sql, data["name_ko"], data["name_en"], ins.Application.Application_seq)
    if err != nil {
      log.Println(err)
      return
    }
    ret = true
    return
  }

  err_msg = "정의되지 않은 item_name입니다."
  return
}
