package sms

import (
  "ipsap/common"
  "ipsap/sms"
  "ipsap/model"
  "fmt"
  "log"
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

func GetSmsResultInfo(sms_seq interface{}) (ret map[string]interface{}) {
  messages := make([]map[string]interface{}, 1)
  message := map[string]interface{}  {
    "message_id"	: sms_seq,
  }
  messages[0] = message
  ret = map[string]interface{} {
    "usercode"  : common.Config.Sms.Usercode,
    "deptcode"  : common.Config.Sms.Deptcode,
    "messages"  : messages,
  }
  return
}

func (msg *MessageMgr)getSmsSendInfo(sms_seq interface{}, smsText string) (ret map[string]interface{}) {
  messages := make([]map[string]interface{}, 1)
  message := map[string]interface{}  {
    "message_id"	: sms_seq,
    "to"					: msg.User_info["phoneno"],
  }
  messages[0] = message
  ret = map[string]interface{} {
    "usercode"  : common.Config.Sms.Usercode,
    "deptcode"  : common.Config.Sms.Deptcode,
    "messages"  : messages,
    "text"      : smsText,
    "from"      : common.Config.Sms.From,
  }
  return
}

func main() {
	configPath := "/root/go/config/sms.toml"
	argsWithProg := os.Args
	if len(argsWithProg[1:]) != 0 {
		configPath = argsWithProg[1]
	}
	// Config 설정
	if _, err := toml.DecodeFile(configPath, &sms.Config); err != nil {
		log.Println(err)
		return
	}
  //  DB 초기화
  common.DB_pool_connect()

  // url := strings.Replace(common.Config.Sms.SendUrl, "{type}", smsType, 1)
  // smsSendInfo := msg.getSmsSendInfo(sms_seq, smsText)
  // respData, _ := common.ApiRequest(url, nil, nil, smsSendInfo, "POST")
  // if nil != respData {
  //   respMap := respData.(map[string]interface{})
  //   send_code = respMap["code"]
  //   if common.Api_status_ok == common.ToInt(respMap["code"]) {
  //     succ = true
  //   } else {
  //     sms_result = model.DEF_SMS_RESULT_FAIL
  //   }
  // } else {
  //   sms_result = model.DEF_SMS_RESULT_FAIL
  // }
  //
  // sql = `UPDATE t_sms
  //           SET send_code = ?, result = ?
  //         WHERE sms_seq = ?`
  // common.DBconn().Exec(sql, send_code, sms_result, sms_seq)

  go batchGetSmsResult()
}

func batchGetSmsResult() {
  // type이 lms 인 경우 mms 로 요청해야됨!
  for {
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
        log.Printf("respStrData : %v",respStrData)
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
              _,err := common.DBconn().Exec(sql,
                       smsResult.Data[0].Result, smsResult.Data[0].Errorcode, smsResult.Data[0].Recvtime,
                       row["sms_seq"])
              if nil != err {
                log.Println(err)
              }
            }
          } else {
            log.Printf("smsResult.Code : %v",smsResult.Code)
            bgm.updateSmsResultFail(row["sms_seq"])
          }
        } else {
          bgm.updateSmsResultFail(row["sms_seq"])
        }
      }
    }

    time.Sleep(time.Second * 30)
  }
}

func updateSmsResultFail(sms_seq interface{}) {
  sql := `UPDATE t_sms
             SET result = ?
           WHERE sms_seq = ?`
  _,err := common.DBconn().Exec(sql, model.DEF_SMS_RESULT_FAIL, sms_seq)
  if nil != err {
    log.Println(err)
  }
  return
}
