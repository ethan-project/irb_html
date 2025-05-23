package common

import (
	"fmt"
	"strconv"
	"strings"
)

func ToValidateUint64(input interface{}) (val uint64, ret bool) {
	val = ToUint64(input)
	valStr := ToStr(val)
	ret = (valStr == ToStr(input))
	return
}

func IsStr(input interface{}) bool {
	if input == nil {
		return false
	}
	switch input.(type) {
	case []interface{}:
		return false
	case map[string]interface{}:
		return false
	default:
		return true
	}
}

func ToStr(input interface{}) string {
	if input == nil {
		return ""
	}
	switch input.(type) {
	case []interface{}:
		return "array"
	case map[string]interface{}:
		return "map"
	case int:
		return strconv.Itoa(input.(int))
	case bool:
		return strconv.FormatBool(input.(bool))
	case float64:
		return strconv.FormatFloat(input.(float64), 'f', -1, 32)
	case int64:
		return strconv.FormatInt(input.(int64), 10)
	case uint64:
		return strconv.FormatUint(input.(uint64), 10)
	default:
		return fmt.Sprintf("%v", input)
	}
}

func ToUint64(input interface{}) uint64 {
	if input == nil {
		return 0
	}

	switch input.(type) {
	case int:
		return uint64(input.(int))
		//  case bool:
		//      return uint64(input.(bool))
	case float64:
		return uint64(input.(float64))
	case int64:
		return uint64(input.(int64))
	case string:
		tmp, _ := strconv.ParseUint(input.(string), 10, 64)
		return tmp
	default:
		return input.(uint64)
	}
}

func ToInt64(input interface{}) int64 {
	if input == nil {
		return 0
	}

	switch input.(type) {
	case int:
		return int64(input.(int))
		//  case bool:
		//      return uint64(input.(bool))
	case float64:
		return int64(input.(float64))
	case int64:
		return int64(input.(int64))
	case string:
		tmp, _ := strconv.ParseInt(input.(string), 10, 64)
		return tmp
	default:
		return input.(int64)
	}
}

func ToInt(input interface{}) int {
	if input == nil {
		return 0
	}

	switch input.(type) {
	case int:
		return int(input.(int))
		//  case bool:
		//      return uint64(input.(bool))
	case float64:
		return int(input.(float64))
	case int64:
		return int(input.(int64))
	case uint64:
		return int(input.(uint64))
	case string:
		tmp, _ := strconv.Atoi(input.(string))
		return tmp
	default:
		return input.(int)
	}
}

func ToUint(input interface{}) uint {
	if input == nil {
		return 0
	}

	switch input.(type) {
	case int:
		return uint(input.(int))
		//  case bool:
		//      return uint64(input.(bool))
	case float64:
		return uint(input.(float64))
	case int64:
		return uint(input.(int64))
	case uint64:
		return uint(input.(uint64))
	case string:
		tmp, _ := strconv.ParseUint(input.(string), 10, 32)
		return uint(tmp)
	default:
		return input.(uint)
	}
}

func Tokens(input string) interface{} {
	parts := strings.Split(input, ",")
	return parts
}

func Contains(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ToConvertStrigBoolToUint8(input interface{}) uint8 {
	if "True" == ToStr(input) {
		return uint8(1)
	} else {
		return uint8(0)
	}
}
