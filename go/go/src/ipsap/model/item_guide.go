
package model

import (
  "ipsap/common"
  "strings"
  "fmt"
)

type ItemGuide struct {
}

func (ins *ItemGuide)LoadGuide(itemArr []string) (ret map[string]string) {

  inStr := "";
  for _, item_name := range itemArr {
    item_name = strings.TrimSpace(item_name)
    if inStr != ""  { inStr += ","  }
    inStr += fmt.Sprintf(`"%v"`, item_name)

    item := Item{}
    if item.GetItem(item_name) {
      if DEF_ITEM_TYPE_ITEM_GROUP == common.ToInt(item.Data["item_type"]) {
        ig := ItemGroup { Item_name : item_name }
        subitems := ig.GetSubItemList()
        sub_ret := ins.LoadGuide(subitems);
        if nil == ret  {  ret = map[string]string {}  }

        for key, value := range sub_ret {
          ret[key] = value
        }
      }
    }
  }

  sql := fmt.Sprintf(
          `SELECT item_name, guide
             FROM t_item_guide
            WHERE item_name in (%v)`, inStr);
  rows := common.DB_fetch_all(sql, nil)
  for _, row := range rows {
    if nil == ret  {
        ret = map[string]string {}
    }
    ret[common.ToStr(row["item_name"])] = common.ToStr(row["guide"])
  }

  return
}
