
package model

import (
	"github.com/gin-gonic/gin"
  "ipsap/common"
  "fmt"
	"strings"
//  "log"
)

type Code struct {
  Type            string
  Id              uint
  Value           string
  ViewOrder       int
}

func (cd *Code)GetCodeStrFromTypeAndId(err_str ...interface{}) (string)  {
  sql := `SELECT value
            FROM t_code
           WHERE type = ?
             AND id = ?`
  row := common.DB_fetch_one(sql, nil, cd.Type, cd.Id)
  if nil == row {
		if len(err_str) > 0 {
			return common.ToStr(err_str[0])
		}
    return ""
  }

  cd.Value = common.ToStr(row["value"])
  return cd.Value
}

func (cd *Code)CheckCodeError(c *gin.Context) (succ bool) {
	if "" == cd.GetCodeStrFromTypeAndId() {
		err_cd := fmt.Sprintf("%v 값이 잘못되었습니다.", cd.Type)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, err_cd)
		return
	}
	return true
}


func (cd *Code)GetCodeJsonData(codeNames []string) (ret map[string]interface{}) {
	sum_str := ""
	for _, name := range codeNames {
		if "" != sum_str	{	sum_str += "," }
		sum_str += `'` + strings.TrimSpace(name) + `'`
	}
	sql := fmt.Sprintf(	`SELECT type, id, value
												 FROM t_code
									  	  WHERE type in (%v)
												  AND del_flag = 0
												ORDER BY type, view_order, id`, sum_str)
	rows := common.DB_fetch_all(sql, nil)
	if nil == rows {
		return
	}

	ret = map[string]interface{} {}

	var cur_map map[uint]interface{}
	old_type := ""
	for _, row := range rows {
		type_str := common.ToStr(row["type"])
		id := common.ToUint(row["id"])
		value := common.ToStr(row["value"])
		if old_type != type_str	{
			if "" != old_type	{
				ret[old_type] = cur_map
			}
			cur_map = make(map[uint]interface{})
			old_type = type_str
		}
		cur_map[id] = value
	}
	if "" != old_type	{
		ret[old_type] = cur_map
	}

	return
}
