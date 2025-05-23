package common

import (
	"github.com/spf13/cast"
	"net/url"
	"regexp"
	"net"
)

// ipv4 형식체크
func CheckIPV4(ip interface{}) bool {
	if net.ParseIP(ip.(string)) == nil {
		return false
	} else {
		return true
	}
//	re := regexp.MustCompile(`/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`)
//	return re.MatchString(ToStr(ip))
}

// url 체크
func CheckUrl(page_url interface{})  bool {
	_, err := url.ParseRequestURI(ToStr(page_url))
	if nil != err {
		return false
	}
	return true
}

 // 영문만 입력 체크
 func CheckEn(en interface{}) bool {
 	 re := regexp.MustCompile("[a-zA-Z]")
	 return re.MatchString(ToStr(en))
 }

 // 한글만 입력 체크
 func CheckKr(kr interface{}) bool {
	 re := regexp.MustCompile("[가-힣]")
	 return re.MatchString(ToStr(kr))
 }

 // 숫자만 입력 체크
 func CheckNum(num interface{}) bool {
	 re := regexp.MustCompile("/^[0-9]+$/")
	 return re.MatchString(ToStr(num))
 }

// 이메일 형식 체크
func CheckEmail(email interface{}) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(ToStr(email))
}

// 날짜 형식 체크 yyyy-mm-dd
func CheckDate(date interface{}) bool {
	re := regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)
	return re.MatchString(ToStr(date))
}

// 8자리 날자 포멧인지 확인
func CheckYYYYMMDD(input interface{}) bool {
	_, succ := ToValidateUint64(input)
	if !succ {
		return false
	}

	valStr := ToStr(input)

	if len(valStr) != 8 {
		return false
	}

	return true
}

// 0, 1 만 가능
func CheckBinary (binary interface{}) bool {
	result, err := cast.ToUint8E(binary)
	if nil != err {
		return false
	}

	if 0 == result || 1 == result {
		return true
	} else {
		return false
	}
}


func CheckHasMustKey(data map[string]interface{}, must_keys []string) (ret bool)	{
	for _, key := range must_keys	{
		_, ok := data[key]
		if !ok	{
			return false
		}
	}
	return true
}
