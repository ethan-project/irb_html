
package model

import (
  "ipsap/common"
  "log"
  "strings"
)

//  Item is read only!!!
type Item struct {
  Item_name       string
  Data            map[string]interface{}
  ItemSelect      *ItemSelect
  ItemStringInput *ItemStringInput
}

func (item *Item)GetItem(item_name string) (bool) {
  if item_name == "" {
    return false
  }

  sql := `SELECT item_type, main_select, multi_select #, reg_dttm
            FROM t_item
           WHERE item_name = ?`

  filter := func(row map[string]interface{}) {
    item_type_code := common.ToUint(row["item_type"])

    code := Code {
      Type : "item_type",
      Id : item_type_code,
    }
    row["item_type_str"] = code.GetCodeStrFromTypeAndId()
  }

  row := common.DB_fetch_one(sql, filter, item_name)
  if nil == row {
    log.Println("Item_name is not defined [", item_name, "]")
    return false
  }

  item.Item_name = item_name
  item.Data = row

  switch(common.ToInt(item.Data["item_type"]))  {
  case DEF_ITEM_TYPE_SELECT   :
    item.ItemSelect = &ItemSelect{}
    item.ItemSelect.Load(item_name)
  case DEF_ITEM_TYPE_BASIC    :
    item.ItemStringInput = &ItemStringInput{}
    item.ItemStringInput.Load(item_name)
  case DEF_ITEM_TYPE_STRING   :
    item.ItemStringInput = &ItemStringInput{}
    item.ItemStringInput.Load(item_name)
  }

  return true
}

func (item *Item)GetFormatData(judge_type int) map[string]interface{}  {
  data := map[string]interface{} {
    "info" : item.Data,
  }

  if nil != item.ItemSelect {
    data["items"] = item.ItemSelect.Datas
  } else if nil != item.ItemStringInput {
    if nil != item.ItemStringInput.Datas {
      data["items"] = item.ItemStringInput.Datas
    }
  }

  code := Code{}
  switch(common.ToInt(item.Data["item_type"]))  {
  case DEF_ITEM_TYPE_ANIMAL  :
    switch(item.Item_name) {
    case "animal_type" :
      codeArr := []string { "animal_code", "mb_grade", "breeding_place",
                            "supplier_type", "lmo_type", "age_unit", "weight_unit", "size_unit"}
      data["codes"] = code.GetCodeJsonData(codeArr)
    }
  case DEF_ITEM_TYPE_MEMBER  :
    codeArr := []string { "exp_year_code", "exp_type_code"}
    data["codes"] = code.GetCodeJsonData(codeArr)
  case DEF_ITEM_TYPE_ANESTHETIC  :
    codeArr := []string { "anesthetic_type", "injection_route"}
    data["codes"] = code.GetCodeJsonData(codeArr)
  case DEF_ITEM_TYPE_STRING  :
    switch(item.Item_name)  {
    case "ibc_risk_bios_infection_chance" :
      codeArr := []string { "exposure_path", "inspection_possibility"}
      data["codes"] = code.GetCodeJsonData(codeArr)
    case "ibc_general_experiment" :
      codeArr := []string { "ibc_experiment_type", "ibc_review_request", "ibc_submitted_doc"}
      data["codes"] = code.GetCodeJsonData(codeArr)
    case "ibc_general_fclty" :
      codeArr := []string { "fclty_type", "fclty_grade", "fclty_lmo"}
      data["codes"] = code.GetCodeJsonData(codeArr)
    case "self_inspection", "expert_review_opinion", "expert_review_evaluation", "approved_inspection" :
      sql := `SELECT contents, link_item_name, is_seq
                FROM t_questionnaire
               WHERE q_item_name = '' AND judge_type = ? AND parent_is_seq = 0
               ORDER BY view_order, is_seq`
      filter := func(row map[string]interface{}) {
        sql := `SELECT is_seq, contents, has_switch, has_check
                  FROM t_questionnaire
                 WHERE judge_type = ? AND parent_is_seq = ? AND q_item_name = ?
                 ORDER BY view_order`
        row["sub_items"] = common.DB_fetch_all(sql, nil, judge_type, common.ToInt(row["is_seq"]), item.Item_name)
      }
      data["codes"] = common.DB_fetch_all(sql, filter, judge_type)
    case "ca_regular_item", "ca_fast_item", "ibc_ca_item", "irb_ca_item"  :
      sql := `SELECT contents, link_item_name, is_seq
                FROM t_questionnaire
               WHERE q_item_name = ?
                 AND judge_type = ?
                 AND parent_is_seq = 0
               ORDER BY view_order, is_seq`
      data["codes"] = common.DB_fetch_all(sql, nil, item.Item_name, judge_type)
    }

  default :
  }

  return data
}

func (item *Item)ValidateItemArrs(itemArr []string) (succ bool, err_item string) {
  for _, item_name := range itemArr {
    item_name = strings.TrimSpace(item_name)
    if "" == item_name  {
      continue;
    }
    sql := `SELECT item_name
              FROM t_item
             WHERE item_name = ?`
    row := common.DB_fetch_one(sql, nil, item_name)
    if nil == row {
      err_item = item_name
      return
    }
  }

  succ = true
  return
}
