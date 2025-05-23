package model

import (
  "ipsap/common"
  "database/sql"
  "strings"
  "log"
//  "fmt"
)

type AppSelect struct {
  Application         *Application
  Datas               map[int]interface{}   //  []map[string]interface{}
}

func (ins *AppSelect)Load(item_name string, tx *sql.Tx) (succ bool) {
  if ins.Application.Application_seq == 0 {
    return
  }

  if ins.Datas != nil {
    return true
  }

  sql := `SELECT item_idx, select_ids
            FROM t_application_select
           WHERE application_seq = ?
             AND item_name = ?`
  filter := func(row map[string]interface{}) {
    row["select_ids"] = strings.Split(common.ToStr(row["select_ids"]), ",")
  }

  funcMakeItemIdxMap := func(item_idx interface{}) {
    _, exists := ins.Datas[common.ToInt(item_idx)]
    if exists {
      return
    }

    ins.Datas[common.ToInt(item_idx)] = map[string]interface{} {}
  }

  ins.Datas = map[int]interface{} {}

  var rows []map[string]interface{}
  if tx != nil {
    rows = common.DB_Tx_fetch_all(tx, sql, filter, ins.Application.Application_seq, item_name)
  } else {
    rows = common.DB_fetch_all(sql, filter, ins.Application.Application_seq, item_name)
  }

  for _, row := range rows {
    item_idx := common.ToInt(row["item_idx"])
    funcMakeItemIdxMap(item_idx);

    data, _ := ins.Datas[item_idx]
    dataMap := data.(map[string]interface{})
    dataMap["select_ids"] = row["select_ids"]

    select_input := AppSelectInput{ Application : ins.Application }
    inputs := select_input.Load(item_name, item_idx)
    if nil != inputs {
      // log.Println("inputs", inputs)
      dataMap["inputs"] = inputs
    }
  }

//  log.Println(ins.Datas)

  return true
}

func (ins *AppSelect)GetJsonData() (ret interface{}) {
  ret = ins.Datas
  return
}


func (ins *AppSelect)UpdateItem(tx *sql.Tx, item_name string, data1 interface{}) (ret bool, err_msg string) {
  //  기존 데이어의 삭제없이 계속 추가되는 아이템 : 하드코딩함!!!!
  append_items := []string{ "expert_review_result", "normal_review_result"}
  del_flag := true
  for _, item := range append_items {
    if item == item_name {
      del_flag = false
      break;
    }
  }

  data := data1.(map[string]interface{})
  if del_flag {
    sql := `DELETE FROM t_application_select
             WHERE application_seq = ?
               AND item_name = ?`
    _, err := tx.Exec(sql, ins.Application.Application_seq, item_name)
    if err != nil {
      log.Println(err)
      return
    }

    sql =  `DELETE FROM t_application_select_input
             WHERE application_seq = ?
               AND item_name = ?`
    _, err = tx.Exec(sql, ins.Application.Application_seq, item_name)
    if err != nil {
      log.Println(err)
      return
    }
  }

  for item_idx, val := range data  {
    select_ids := val.(map[string]interface{})["select_ids"]
    if nil != select_ids {
      ids_arr := make([]string, 0, 0)
      for _, id := range select_ids.([]interface{}) {
        id_str := common.ToStr(id)
        //  wowdolf : id 값 validation 하자.!!!!!
        ids_arr = append(ids_arr, id_str)
      }

      select_string :=  strings.Join(ids_arr, ",")
      sql := `INSERT INTO t_application_select(application_seq, item_name, item_idx, select_ids)
              VALUES(?, ?, ?, ?)
              ON DUPLICATE KEY UPDATE select_ids = ?`
      _, err := tx.Exec(sql, ins.Application.Application_seq, item_name, item_idx, select_string, select_string)
      if err != nil {
        err_msg = "select_ids 값이 잘못 되었습니다."
        log.Println(err)
        return
      }
    }
    inputs := val.(map[string]interface{})["inputs"]
    if nil != inputs {
      inputs_map := inputs.(map[string]interface{})
      for value_id, value_input := range inputs_map {
        value_input_map := value_input.(map[string]interface{})
        for value_no, value_string := range value_input_map {
          sql := `INSERT INTO t_application_select_input(application_seq, item_name, item_idx, id, no, input)
                  VALUES(?, ?, ?, ?, ?, ?)
                  ON DUPLICATE KEY UPDATE input = ?`
          _, err := tx.Exec(sql, ins.Application.Application_seq, item_name, item_idx,
                            value_id, value_no, value_string,
                            value_string)
          if err != nil {
            err_msg = "inputs 값이 잘못 되었습니다."
            log.Println(err)
            return
          }
        }
      }
    }
  }

  ret = true
  return
}
