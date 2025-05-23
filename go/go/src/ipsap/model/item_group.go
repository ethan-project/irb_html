
package model

import (
  "ipsap/common"
)

//  Item is read only!!!
type ItemGroup struct {
  Item_name       string
}

func (ig *ItemGroup)GetSubItemList() (list []string) {
  if ig.Item_name == "" {
    return nil
  }
  sql := `SELECT subitem_name
            FROM t_item_group
           WHERE item_name = ?`

  rows := common.DB_fetch_all(sql, nil, ig.Item_name)
  if nil == rows {
    return nil
  }

  list = make([]string, 0, 0)
  for _, row := range rows {
    list = append(list, common.ToStr(row["subitem_name"]))
  }

  return
}
