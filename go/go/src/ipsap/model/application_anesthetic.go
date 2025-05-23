
package model

import (
  "ipsap/common"
  "database/sql"
  "log"
)

type AppAnesthetic struct {
  Application         *Application
  Datas               []map[string]interface{}
}

func (ins *AppAnesthetic)Load(item_name string) (succ bool) {
  if ins.Application.Application_seq == 0 {
    return
  }

  if ins.Datas != nil {
    return true
  }

  search_item_name := item_name
  if search_item_name == "pain_relief_veterinary_mng" {
    search_item_name = "pain_d_anesthetic"    //  wowdolf : 하드 코딩
  }

  sql := `SELECT anesthetic_type, anesthetic_type_str, anesthetic_name,
                 injection_mg,    injection_route,     injection_route_str,
                 injection_time,  injection_cnt
            FROM t_application_anesthetic
           WHERE application_seq = ?
             AND item_name = ?
           ORDER BY view_order`
  filter := func(row map[string]interface{}) {
    // switch (item_name) {
    // case "pain_d_anesthetic"  :
    //   delete(row, "injection_time")
    //   delete(row, "injection_cnt")
    // }

    codeAT := Code {}
    codeAT.Type = "anesthetic_type"
    codeAT.Id = common.ToUint(row[codeAT.Type])
    row[codeAT.Type + "_str"] = codeAT.GetCodeStrFromTypeAndId(row[codeAT.Type + "_str"])

    codeIR := Code {}
    codeIR.Type = "injection_route"
    codeIR.Id = common.ToUint(row[codeIR.Type])
    row[codeIR.Type + "_str"] = codeIR.GetCodeStrFromTypeAndId(row[codeIR.Type + "_str"])
  }
  ins.Datas = common.DB_fetch_all(sql, filter, ins.Application.Application_seq, search_item_name)

  return true
}

func (ins *AppAnesthetic)GetJsonData() (ret interface{}) {
  ret = ins.Datas
  return
}

func (ins *AppAnesthetic)UpdateItem(tx *sql.Tx, item_name string, data1 interface{}) (ret bool, err_msg string) {
  dataArr := data1.([]interface{})
  if item_name == "pain_relief_veterinary_mng" {
    item_name = "pain_d_anesthetic"
  }
	// log.Println(dataArr)
  sql := `DELETE FROM t_application_anesthetic
           WHERE application_seq = ?
             AND item_name = ?`
  _, err := tx.Exec(sql, ins.Application.Application_seq, item_name)
  if err != nil {
    log.Println(err)
    return
  }

  for idx, data2 := range dataArr {
    data := data2.(map[string]interface{})
  	// log.Println(data)
    // 투여횟수
    if nil == data["injection_cnt"] {
      data["injection_cnt"] = ""
    }

    // 투여시점
    if nil == data["injection_time"] {
      data["injection_time"] = ""
    }

    codeAT := Code {}
    codeAT.Type = "anesthetic_type"
    codeAT.Id = common.ToUint(data[codeAT.Type]);
    data[codeAT.Type + "_str"] = codeAT.GetCodeStrFromTypeAndId(data[codeAT.Type + "_str"])

    codeIR := Code {}
    codeIR.Type = "injection_route"
    codeIR.Id = common.ToUint(data[codeIR.Type]);
    data[codeIR.Type + "_str"] = codeIR.GetCodeStrFromTypeAndId(data[codeIR.Type + "_str"])

    sql := `INSERT INTO t_application_anesthetic(	application_seq,  item_name,
                                                  anesthetic_type,  anesthetic_type_str,
                                                  anesthetic_name,  injection_mg,
                                                  injection_route,  injection_route_str,
                                                  injection_time,   injection_cnt,    view_order)
            VALUES(?,?, ?,?, ?,?, ?,?, ?,?,?)
            ON DUPLICATE KEY UPDATE anesthetic_type = ?,  anesthetic_type_str = ?,
                                    anesthetic_name = ?,  injection_mg = ?,
                                    injection_route = ?,  injection_route_str = ?,
                                    injection_time = ?,   injection_cnt = ?,    view_order = ?`
    _, err := tx.Exec(sql,
                      ins.Application.Application_seq,  item_name,
                      data["anesthetic_type"],          data["anesthetic_type_str"],
                      data["anesthetic_name"],          data["injection_mg"],
                      data["injection_route"],          data["injection_route_str"],
                      data["injection_time"],           data["injection_cnt"],        idx,
                      data["anesthetic_type"],          data["anesthetic_type_str"],
                      data["anesthetic_name"],          data["injection_mg"],
                      data["injection_route"],          data["injection_route_str"],
                      data["injection_time"],           data["injection_cnt"],        idx)
    if err != nil {
      log.Println(err)
      err_msg = "data 값이 잘못 되었습니다."
      return
    }
  }

  ret = true
  return
}
