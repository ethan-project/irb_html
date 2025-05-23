
package model

//  기타(문자열로 저장이 가능한 정보들) 관리
//  - 자가점검표

import (
//  "github.com/gin-gonic/gin"
  "ipsap/common"
//  "database/sql"
//  "strings"
  "log"
  "fmt"
)

type AppCustom struct {
  Application         *Application
  Datas               []map[string]interface {}
}

func (ins *AppCustom)Load(item_name string) (succ bool) {
  if ins.Application.Application_seq == 0 {
    return
  }

  if ins.Datas != nil {
    return true
  }

//  ins.Datas = []interface{} {}

  switch(item_name) {
  case "child_ca_application" :
    log.Println("====", ins.Application.Data);

    application := Application{
      Application_seq : common.ToUint(ins.Application.Data["application_seq"]),
    }

    moreCondition := fmt.Sprintf(`
      AND app.parent_app_seq = %v
      AND app.application_type = %v
      AND app.application_step = %v
      AND app.application_result in (%v, %v)`, ins.Application.Data["application_seq"],
      DEF_APP_TYPE_CHANGE,   //  변경 신청서
      DEF_APP_STEP_FINAL,    //  최종 심의
      DEF_APP_RESULT_APPROVED, DEF_APP_RESULT_APPROVED_C) //  승인, 조건부 승인
    ins.Datas = application.LoadList("", "", moreCondition, "").([]map[string]interface {})
    log.Println(ins.Datas)
  }
/*
  sql := `SELECT target_item, contents
            FROM t_application_etc
           WHERE application_seq = ?
             AND item_name = ?`
  rows := common.DB_fetch_all(sql, nil, ins.Application.Application_seq, item_name)
  for _, row := range rows {
    ins.Data[common.ToStr(row["target_item"])] = common.ToStr(row["contents"])
  }
*/
//  log.Println(ins.Data)
  return true
}

func (ins *AppCustom)GetJsonData() (ret interface{}) {
  log.Println("=345435=", ins.Datas)
  ret = ins.Datas
  return
}
