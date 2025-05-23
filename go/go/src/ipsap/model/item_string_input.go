
package model

import (
  "ipsap/common"
//  "strings"
//  "log"
)

type ItemStringInput struct {
  Datas map[string]interface{}
}

func (ins *ItemStringInput)Load(item_name string) (bool) {
  sql := `SELECT sub_tag, placeholder, max_len, number_only, must
            FROM t_item_string_input
           WHERE item_name = ?`

  rows := common.DB_fetch_all(sql, nil, item_name)
  for _, row := range rows {
    if nil == ins.Datas {
      ins.Datas = map[string]interface{} {}
    }
    sub_tag := common.ToStr(row["sub_tag"])
    delete(row, "sub_tag")
    ins.Datas[sub_tag] = row
  }
  if nil == ins.Datas {
    return false
  }

  return true
}
