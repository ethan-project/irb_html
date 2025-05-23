package ctl

import (
	"ipsap/common"
)

var AllCodes map[string]interface{}

func Init_common_code() {
	AllCodes = make(map[string]interface{})

	sql	 := "SELECT type, id, value, view_order FROM t_code ORDER BY type, id"
	rows := common.DB_fetch_all(sql, nil)
	if nil == rows || 0 == len(rows) {
		return
	}

	oldType := ""
	curMaps := make(map[string]interface{})

	for _, values := range rows {
		code_type := common.ToStr(values["type"])
		code_id := common.ToStr(values["id"])
		code_value := values["value"]
		if code_type != oldType {
			if len(curMaps) > 0 {
				AllCodes[oldType] = curMaps
				curMaps = make(map[string]interface{})
			}
			oldType = code_type
		}
		curMaps[code_id] = code_value
	}
	if len(curMaps) > 0 {
		AllCodes[oldType] = curMaps
	}
}

func GetCommonCodeStr(codetype string, codeid interface{}) string {
	val, exists := AllCodes[codetype]
	if !exists {
		return "Code Error"
	}

	val2, exists2 := (val.(map[string]interface{}))[common.ToStr(codeid)]
	if !exists2 {
		return "Code Error"
	}

	return common.ToStr(val2)
}
