
package model

import (
  "ipsap/common"
  "log"
)

type CodeDynamic struct {
  DCode_seq         uint
  Institution_seq   uint
  DCode_type        uint
  Code              uint
  Value             string
}

func Create_CodeDynamic(institution_seq uint) (*CodeDynamic)  {
  newCodeD := CodeDynamic{
    Institution_seq : institution_seq,
  }
  return &newCodeD
}

func (CodeD *CodeDynamic)GetValueFromCode() (string)  {
  sql := `SELECT dcode_seq, value
            FROM t_code_dyn
           WHERE institution_seq = ?
             AND dcode_type = ?
             AND code = ?`
  row := common.DB_fetch_one(sql, nil, CodeD.Institution_seq, CodeD.DCode_type, CodeD.Code)
  if nil == row {
    return ""
  }

  CodeD.DCode_seq = common.ToUint(row["dcode_seq"])
  CodeD.Value = common.ToStr(row["value"])
  return CodeD.Value
}

func (CodeD *CodeDynamic)GetCodeFromValue() (code uint)  {
  sql := `SELECT dcode_seq, code
            FROM t_code_dyn
           WHERE institution_seq = ?
             AND dcode_type = ?
             AND value = ?`
  row := common.DB_fetch_one(sql, nil, CodeD.Institution_seq, CodeD.DCode_type, CodeD.Value)
  if nil == row {     //  New Insert
    sql = `SELECT IFNULL(max(code), 0) + 1 newcode
             FROM t_code_dyn
            WHERE institution_seq = ?
              AND dcode_type = ?`
    newCode := uint(1)
    row = common.DB_fetch_one(sql, nil, CodeD.Institution_seq, CodeD.DCode_type)
    if nil != row {
      newCode = common.ToUint(row["newcode"])
    }

    sql = `INSERT INTO t_code_dyn(institution_seq, dcode_type, code, value, reg_dttm)
            VALUES(?,?,?,?, UNIX_TIMESTAMP())`
    result, err := common.DBconn().Exec(sql, CodeD.Institution_seq, CodeD.DCode_type, newCode, CodeD.Value)
  	if nil != err {
  		log.Println(err)
  		return 0
  	}

    n, err := result.LastInsertId()
    CodeD.DCode_seq = uint(n)
    CodeD.Code = newCode
  } else {
    CodeD.DCode_seq = common.ToUint(row["dcode_seq"])
    CodeD.Code = common.ToUint(row["code"])
  }

  return CodeD.Code
}



func (CodeD *CodeDynamic)GetCodeJsonData() (ret map[uint]string) {
  sql := `SELECT code, value
            FROM t_code_dyn
           WHERE institution_seq = ?
             AND dcode_type = ?
             AND del_flag = 0
        ORDER BY code`
  rows := common.DB_fetch_all(sql, nil, CodeD.Institution_seq, CodeD.DCode_type)
  if nil == rows {
		return
	}

	ret = map[uint]string {}
	for _, row := range rows {
		code := common.ToUint(row["code"])
		value := common.ToStr(row["value"])
		ret[code] = value
	}

	return
}
