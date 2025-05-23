package main

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"ipsap/common"
	"ipsap/model"
	"log"
	"os"
	"strings"
	"time"
)

type SmsResult struct {
	Code string `json:"code"`
	Data []struct {
		MessageID string `json:"message_id"`
		Result    string `json:"result"`
		Errorcode string `json:"errorcode"`
		Recvtime  string `json:"recvtime"`
	} `json:"data"`
}

func main() {
	configPath := "/root/go/config/sms.toml"
	argsWithProg := os.Args
	if len(argsWithProg[1:]) != 0 {
		configPath = argsWithProg[1]
	}

	// Config 설정
	if _, err := toml.DecodeFile(configPath, &common.Config); err != nil {
		log.Println(err)
		return
	}
	//  DB 초기화
	common.DB_pool_connect()

	if common.Config.Server.LogPath != "" {
		log.SetOutput(setLog("log_file"))
	}

	log.Println("Send SMS Start!!!!!!!!!!!!")
	go batchGetSmsResult()
	send_code := 0
	sms_result := model.DEF_SMS_RESULT_WAIT
	for {
		sql := `SELECT sms_seq, type, phoneno, msg
              FROM t_sms
             WHERE send_code = 0
               AND result = 0`
		rows := common.DB_fetch_all(sql, nil)
		if len(rows) > 0 {
			for _, row := range rows {
				url := strings.Replace(common.Config.Sms.SendUrl, "{type}", common.ToStr(row["type"]), 1)
				smsSendInfo := getSmsSendInfo(row["sms_seq"], row["phoneno"], row["msg"])
				respData, _ := common.ApiRequest(url, nil, nil, smsSendInfo, "POST")
				if nil != respData {
					respMap := respData.(map[string]interface{})
					send_code = common.ToInt(respMap["code"])
					if common.Api_status_ok != common.ToInt(respMap["code"]) {
						sms_result = model.DEF_SMS_RESULT_FAIL
					}
				} else {
					sms_result = model.DEF_SMS_RESULT_FAIL
				}

				sql = `UPDATE t_sms
                  SET send_code = ?, result = ?, send_dttm = UNIX_TIMESTAMP()
                WHERE sms_seq = ?`
				common.DBconn().Exec(sql, send_code, sms_result, row["sms_seq"])
			}
		}
		time.Sleep(time.Second * 5)
	}
}

func setLog(fileName string) *rotatelogs.RotateLogs {
	basePath := common.Config.Server.LogPath
	_ = os.Mkdir(basePath, os.ModeDir)

	yyyymmdd := "%Y%m%d"

	rl, _ := rotatelogs.New(basePath+fileName+"_"+yyyymmdd+".log",
		rotatelogs.WithLinkName(basePath+fileName),
		rotatelogs.WithMaxAge(time.Duration(common.Config.Server.MaxBackups)*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	return rl
}

func GetSmsResultInfo(sms_seq interface{}) (ret map[string]interface{}) {
	messages := make([]map[string]interface{}, 1)
	message := map[string]interface{}{
		"message_id": sms_seq,
	}
	messages[0] = message
	ret = map[string]interface{}{
		"usercode": common.Config.Sms.Usercode,
		"deptcode": common.Config.Sms.Deptcode,
		"messages": messages,
	}
	return
}

func getSmsSendInfo(sms_seq interface{}, phoneno interface{}, smsText interface{}) (ret map[string]interface{}) {
	messages := make([]map[string]interface{}, 1)
	message := map[string]interface{}{
		"message_id": sms_seq,
		"to":         phoneno,
	}
	messages[0] = message
	ret = map[string]interface{}{
		"usercode": common.Config.Sms.Usercode,
		"deptcode": common.Config.Sms.Deptcode,
		"messages": messages,
		"text":     smsText,
		"from":     common.Config.Sms.From,
	}
	return
}

func batchGetSmsResult() {
	// type이 lms 인 경우 mms 로 요청해야됨!
	for {
		time.Sleep(time.Second * 30)
		sql := `SELECT sms_seq, IF(type = "lms" OR type = "mms", "mms", "sms") type
              FROM t_sms
             WHERE send_dttm > 0
               AND send_code = ?
               AND result = 0`
		rows := common.DB_fetch_all(sql, nil, common.Api_status_ok)
		if len(rows) > 0 {
			for _, row := range rows {
				url := strings.Replace(common.Config.Sms.ResultUrl, "{type}", common.ToStr(row["type"]), 1)
				smsResultInfo := GetSmsResultInfo(row["sms_seq"])
				_, respStrData := common.ApiRequest(url, nil, nil, smsResultInfo, "POST")
				log.Printf("respStrData : %v", respStrData)
				if "" != respStrData {
					smsResult := SmsResult{}
					if err := json.Unmarshal([]byte(respStrData), &smsResult); err != nil {
						log.Println(err)
					}

					if common.Api_status_ok == common.ToUint(smsResult.Code) {
						if "" != smsResult.Data[0].Result && "" != smsResult.Data[0].Errorcode && "" != smsResult.Data[0].Recvtime {
							sql = `UPDATE t_sms
                        SET result = ?, error_code = ?, receive_dttm = ?
                      WHERE sms_seq = ?`
							_, err := common.DBconn().Exec(sql,
								smsResult.Data[0].Result, smsResult.Data[0].Errorcode, smsResult.Data[0].Recvtime,
								row["sms_seq"])
							if nil != err {
								log.Println(err)
							}
						}
					} else {
						log.Printf("smsResult.Code : %v", smsResult.Code)
						updateSmsResultFail(row["sms_seq"])
					}
				} else {
					updateSmsResultFail(row["sms_seq"])
				}
			}
		}
	}
}

func updateSmsResultFail(sms_seq interface{}) {
	sql := `UPDATE t_sms
             SET result = ?
           WHERE sms_seq = ?`
	_, err := common.DBconn().Exec(sql, model.DEF_SMS_RESULT_FAIL, sms_seq)
	if nil != err {
		log.Println(err)
	}
	return
}
