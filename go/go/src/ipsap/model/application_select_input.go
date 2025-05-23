
package model

import (
  "ipsap/common"
//  "log"
)

type AppSelectInput struct {
  Application         *Application
  Datas               map[int]interface{}   //  id : values...
}

func (ins *AppSelectInput)Load(item_name string, item_idx int) (ret interface{}) {
  if ins.Application.Application_seq == 0 {
    return
  }
  if ins.Datas != nil && len(ins.Datas) > 0 {
    return ins.Datas
  }

  sql := `SELECT id, no, input
            FROM t_application_select_input
           WHERE application_seq = ?
           AND item_name = ?
           AND item_idx = ?
          ORDER BY id, no`
  rows := common.DB_fetch_all(sql, nil, ins.Application.Application_seq, item_name, item_idx)

  var data map[int]string
  old_id := -1
  for _, row := range rows {
    if nil == ins.Datas {
      ins.Datas = map[int]interface{} {}
    }
    id := common.ToInt(row["id"])
    if old_id != id  {
      if old_id >= 0  {
        ins.Datas[old_id] = data
      }
      old_id = id
      data = map[int]string {}
    }
    data[common.ToInt(row["no"])] = common.ToStr(row["input"])
  }
  if old_id >= 0 {
    ins.Datas[old_id] = data
  }

  if len(rows) == 0 {
    return
  }

  return ins.Datas
}
