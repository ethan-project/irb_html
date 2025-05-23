
package model

import (
  "ipsap/common"
  "database/sql"
  "strings"
  "log"
)

type AppKeyword struct {
  Application         *Application
  Datas               []interface{}   //  []string
}

func (ins *AppKeyword)Load(item_name string) (succ bool) {
  if ins.Application.Application_seq == 0 {
    return
  }

  if ins.Datas != nil {
    return true
  }

  ins.Datas = make([]interface{}, 0, 0)
  sql := `SELECT keyword
            FROM t_application_keyword
           WHERE application_seq = ?
             AND item_name = ?`
  rows := common.DB_fetch_all(sql, nil, ins.Application.Application_seq, item_name)
  if nil != rows {
    for _, row := range rows {
      ins.Datas = append(ins.Datas, common.ToStr(row["keyword"]))
    }
  }

  // log.Println(ins.Datas)

  return true
}

func (ins *AppKeyword)GetJsonData() (ret interface{}) {
  ret = ins.Datas
  return
}


func (ins *AppKeyword)UpdateItem(tx *sql.Tx, item_name string, data1 interface{}) (ret bool, err_msg string) {
  data := data1.([]interface{})

  sql := `DELETE FROM t_application_keyword
           WHERE application_seq = ?
             AND item_name = ?`
  _, err := tx.Exec(sql, ins.Application.Application_seq, item_name)
  if err != nil {
    log.Println(err)
    return
  }

  for _, keyword := range data  {
    keyword := strings.TrimSpace(common.ToStr(keyword))
    sql  = `INSERT INTO t_application_keyword(application_seq, item_name, keyword)
            VALUES(?, ?, ?)`
    _, err := tx.Exec(sql, ins.Application.Application_seq, item_name, keyword)
    if err != nil {
      err_msg = "select_ids 값이 잘못 되었습니다."
      log.Println(err)
      return
    }
  }

  ret = true
  return
}
