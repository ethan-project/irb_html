
package model

import (
  "ipsap/common"
  "database/sql"
  "log"
  "fmt"
)

type AppMainCheck struct {
  Application         *Application
  Checked             map[int]int
}

func (ins *AppMainCheck)Load(item_name string) (checked interface{}) {
  if ins.Application.Application_seq == 0 {
    return
  }

  ins.Checked = map[int]int {}

  sql := `SELECT item_idx, checked
            FROM t_application_main_check
           WHERE application_seq = ?
             AND item_name = ?
            ORDER BY item_idx`
  rows := common.DB_fetch_all(sql, nil, ins.Application.Application_seq, item_name)
  if nil == rows {
    return
  }

  for _, row := range rows {
    ins.Checked[common.ToInt(row["item_idx"])] = common.ToInt(row["checked"])
  }

  return ins.Checked
}


func (ins *AppMainCheck)UpdateItem(tx *sql.Tx, item_name string, data map[string]interface{}) (ret bool, err_msg string) {
//  log.Println("UpdateItem ", item_name, ": AppMainCheck")

  sql := `DELETE FROM t_application_main_check
           WHERE application_seq = ?
             AND item_name = ?`
  _, err := tx.Exec(sql, ins.Application.Application_seq, item_name)
  if err != nil {
    log.Println(err)
    return
  }

  for key, val := range data  {
    sql := `INSERT INTO t_application_main_check(application_seq, item_name, item_idx, checked)
            VALUES(?, ?, ?, ?)`
    _, err := tx.Exec(sql, ins.Application.Application_seq, item_name, key, val)
    if err != nil {
      err_msg = fmt.Sprintf("main_select의 값이 잘못 되었습니다.")
      log.Println(err)
      return
    }
  }

  ret = true
  return
}
