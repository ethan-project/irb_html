
package model

import (
  "ipsap/common"
  "strings"
//  "log"
)

type ItemSelect struct {
  Datas  []map[string]interface{}
}

func arrayData(data interface{}, default_value interface{}, max_cnt int) (ret []interface{}) {
  isBool := false
  isNumber := false
  switch default_value.(type) {
	case int:
    isNumber = true
	case bool:
    isBool = true
  }

  ret = make([]interface{}, 0, 0)
  if common.ToStr(data) != ""  {
    str_arr := strings.Split(common.ToStr(data), "|")
    for i := 0 ; i < len(str_arr) ; i++ {
      if isBool {
        ret = append(ret, (str_arr[i] != "0"))
      } else if isNumber {
        ret = append(ret, common.ToInt(str_arr[i]))
      } else {
        ret = append(ret, str_arr[i])
      }
    }

    if len(ret) > 0 {
      default_value = ret[len(ret)-1]
    }
  }
  for i := len(ret) ; i < max_cnt ; i++ {
    ret = append(ret, default_value)
  }
  return
}

func (ins *ItemSelect)Load(item_name string) (bool) {
  sql := `SELECT id, value,
                 w_class, input_cnt, iw_class, input_prefix, input_suffix, input_placeholder, input_max_len,
                 number_only, input_only, must  #, view_order, reg_dttm, reg_user_seq
            FROM t_item_select
           WHERE item_name = ?
          ORDER BY view_order`

  filter := func(row map[string]interface{}) {
    input_cnt := common.ToInt(row["input_cnt"])
    if input_cnt > 0  {
      input_data := map[string]interface{} {
        "cnt" : row["input_cnt"],
        "w_class" : row["w_class"],
        "max_len" : row["input_max_len"],
      }
      input_data["iw_class"] = arrayData(row["iw_class"], "", input_cnt)
      input_data["prefix"] = arrayData(row["input_prefix"], "", input_cnt)
      input_data["suffix"] = arrayData(row["input_suffix"], "", input_cnt)
      input_data["placeholder"] = arrayData(row["input_placeholder"], "", input_cnt)
      input_data["input_only"] = arrayData(row["input_only"], false, input_cnt)
      input_data["number_only"] = arrayData(row["number_only"], false, input_cnt)
      input_data["max_len"] = arrayData(row["input_max_len"], 0, input_cnt)
      input_data["must"] = arrayData(row["must"], 0, input_cnt)
      row["input"] = input_data
    }

    delete(row, "input_cnt")
    delete(row, "w_class")
    delete(row, "iw_class")
    delete(row, "input_max_len")
    delete(row, "input_only")
    delete(row, "input_prefix")
    delete(row, "input_suffix")
    delete(row, "input_placeholder")
    delete(row, "number_only")
    delete(row, "must")
  }

  ins.Datas = common.DB_fetch_all(sql, filter, item_name)
  if nil == ins.Datas {
    return false
  }

  return true
}
