package model

import (
  "ipsap/common"
  "io/ioutil"
  "strings"
	"net/url"
//  "database/sql"
  "fmt"
  "log"
)

func (msg *MessageMgr)SendMessage() (succ bool) {
  smsSucc   := false
  emailSucc := false

  // donghun : 2021-08-18 멤버십 가입 알림 메세지는 전송 안한다!
  if msg.Msg_ID == DEF_MSG_MEMBERSHIP_JOIN {
    return true
  }

  // 서비스 이용 승인 안내시 sms를 보내지 않는다!
  if (common.ToInt(msg.User_info["agree_sms"]) > 0 && msg.Msg_ID != DEF_MSG_INSTITUTION_APPROVED) {
    smsSucc = msg.sendSms()
    if !smsSucc{
      log.Println("sms 전송대기 실패 !!")
    }
  }

  if (common.ToInt(msg.User_info["agree_email"]) > 0) {
    emailSucc = msg.sendEmail()
    if !emailSucc{
      log.Println("email 전송 실패 !!")
    }
  }

  if smsSucc || emailSucc {
    succ = true
    return
  }

  return
}

func (msg *MessageMgr)sendSms() (succ bool) {
  smsText := ""
  if msg.Msg_ID < 70 {
    _, _, smsText = msg.getMsgContent(true)
  } else {
    _, _, smsText = msg.getMembershipMsgContent(true)
  }

  smsType := "sms"
  // 80 byte 이상일시 lms로 전송!
  if len(smsText) > 80 {
    smsType = "lms"
  }

  sql := `INSERT INTO t_sms(institution_seq, user_seq, type, phoneno, msg, result, reg_dttm)
          VALUES(?,?,?,?,?,?,UNIX_TIMESTAMP())`
  _, err := common.DBconn().Exec(sql, msg.User_info["institution_seq"], msg.User_info["user_seq"], smsType,
                                      msg.User_info["phoneno"], smsText, DEF_SMS_RESULT_WAIT)
  if nil != err {
    log.Println(err)
    return
  }

  succ = true
  return
}

func (msg *MessageMgr)sendEmail() (succ bool) {
  subject := ""
  mailText := ""
  if msg.Msg_ID < 70 {
    subject, mailText, _ = msg.getMsgContent(false)
  } else {
    subject, mailText, _ = msg.getMembershipMsgContent(false)
  }

  email := common.Email{}
  email.To = common.ToStr(msg.User_info["email"])
  email.Subject = subject
  if !email.Connect() {
    log.Println("SMTP 계정 연결 실패!!")
    return
  }

  email.Msg.Embed(common.Config.Email.LogoFile)
  email.Msg.SetBody("text/html", mailText)
  err := email.Dial.DialAndSend(email.Msg)
  if nil != err {
    log.Printf("메일전송 실패 !!!!! : %v",err)
    return
  }

  succ = true
  return
}

func (msg *MessageMgr) getMsgContent(isSms bool) (subject string, mailText string, smsText string) {
  application := Application {
    Application_seq : msg.Application_seq,
  }

  data := application.Load().(map[string]interface{})
  sql := `SELECT IFNULL(item_select.value, "") as final_judge_result
            FROM t_application_select app_select, t_item_select item_select
           WHERE item_select.item_name = app_select.item_name
             AND app_select.item_name = 'final_judge_result'
             AND application_seq = ?
             AND app_select.select_ids = item_select.id`
  row := common.DB_fetch_one(sql, nil, msg.Application_seq)

  admin_email, admin_telno := GetServiceAdminInfo()
  remainingTime := ""
  deadline      := ""
  round         := ""
  target_name := common.ToStr(msg.User_info["name"])
  target_id :=  common.ToStr(msg.User_info["email"])
  target_type := common.ToStr(msg.User_info["user_type_str"])
  status := ""
  user := User{}
  user.Institution_seq = common.ToUint(msg.User_info["institution_seq"])
  moreCondition := ""
  if common.ToUint(msg.User_info["admin_user_seq"]) > 0 {
    moreCondition = fmt.Sprintf(` AND user.user_seq = %d`, common.ToUint(msg.User_info["admin_user_seq"]))
  }

  row2 := user.GetInstitutionAdminInfo(moreCondition)
  reqUrl := ""
  // withdarwCancelUrl := ""
  if common.Config.Test.IsTestMode {
    reqUrl = "http://218.232.81.113:16999"
  } else {
    reqUrl = "https://ipsap.co.kr"
  }

  params := url.Values{}
  params.Set("user", common.ToStr(msg.User_info["user_seq"]))
  params.Set("instt", common.ToStr(msg.User_info["institution_seq"]))
  switch msg.Msg_ID {
    case DEF_MSG_USER_WITHDRAW : //  계정탈퇴 안내(기관내 탈퇴)
      reqUrl += "/cancel_withdrawal.html?" + params.Encode()
    case DEF_MSG_USER_IPSAP_WITHDRAW : //  계정탈퇴 안내(소속한 모든 기관 탈퇴)
      reqUrl += "/cancel_withdrawal.html?" + params.Encode()
    case DEF_MSG_EXPER_JUDGE_START :
      deadline = application.GetTimeFormatForExpertDeadline("2006년01월02일 15시04분")
      status = "시작"
      params.Set("page" ,"pro")
      params.Set("app" , common.ToStr(msg.Application_seq))
      reqUrl += "/login3.html?" + params.Encode()
    case DEF_MSG_EXPER_JUDGE_START_TO_LEADER :
      deadline = application.GetTimeFormatForExpertDeadline("2006년01월02일 15시04분")
      status = "시작"
      params.Set("page" ,"app")
      params.Set("app" , common.ToStr(msg.Application_seq))
      reqUrl += "/login3.html?" + params.Encode()
    case DEF_MSG_NORMAL_JUDGE_START:
      deadline = application.GetTimeFormatForNormalDeadline("2006년01월02일 15시04분")
      status = "시작"
      params.Set("page" ,"normal")
      params.Set("app" , common.ToStr(msg.Application_seq))
      reqUrl += "/login3.html?" + params.Encode()
    case DEF_MSG_NORMAL_JUDGE_START_TO_LEADER:
      deadline = application.GetTimeFormatForNormalDeadline("2006년01월02일 15시04분")
      status = "시작"
      params.Set("page" ,"normal")
      params.Set("app" , common.ToStr(msg.Application_seq))
      reqUrl += "/login3.html?" + params.Encode()
    case DEF_MSG_BEFORE_24_HOURS:
      round = "1차"
      remainingTime = "24시간"
      status = "종료"
      params.Set("app" , common.ToStr(msg.Application_seq))
      if common.ToUint(data["application_step"]) == DEF_APP_STEP_PRO {
        deadline = application.GetTimeFormatForExpertDeadline("2006년01월02일 15시04분")
        params.Set("page" ,"pro")
      } else if common.ToUint(data["application_step"]) == DEF_APP_STEP_NORMAL {
        deadline = application.GetTimeFormatForNormalDeadline("2006년01월02일 15시04분")
        params.Set("page" ,"normal")
      }
      reqUrl += "/login3.html?" + params.Encode()
    case DEF_MSG_BEFORE_4_HOURS:
      round = "2차"
      remainingTime = "4시간"
      status = "종료"
      params.Set("app" , common.ToStr(msg.Application_seq))
      if common.ToUint(data["application_step"]) == DEF_APP_STEP_PRO {
        deadline = application.GetTimeFormatForExpertDeadline("2006년01월02일 15시04분")
        params.Set("page" ,"pro")
      } else if common.ToUint(data["application_step"]) == DEF_APP_STEP_NORMAL {
        deadline = application.GetTimeFormatForNormalDeadline("2006년01월02일 15시04분")
        params.Set("page" ,"normal")
      }
      reqUrl += "/login3.html?" + params.Encode()
    case DEF_MSG_JUDGE_DELAYED:
      status = "미이행"
      params.Set("app" , common.ToStr(msg.Application_seq))
      if common.ToUint(data["application_step"]) == DEF_APP_STEP_PRO {
        deadline = application.GetTimeFormatForExpertDeadline("2006년01월02일 15시04분")
        params.Set("page" ,"pro")
      } else if common.ToUint(data["application_step"]) == DEF_APP_STEP_NORMAL {
        deadline = application.GetTimeFormatForNormalDeadline("2006년01월02일 15시04분")
        params.Set("page" ,"normal")
      }
      reqUrl += "/login3.html?" + params.Encode()
    case DEF_MSG_REQUEST_SUPPLEMENT: // 신청서 행정보완요청!
      data["application_step_str"] = "보완요청"
      status = "접수"
      params.Set("page" ,"app")
      params.Set("app" , common.ToStr(msg.Application_seq))
      reqUrl += "/login3.html?" + params.Encode()
    case DEF_MSG_WITHDRAW_NOTICE: // 탈퇴 공지
      target_name = common.ToStr(msg.User_info["target_name"])
      target_id   = common.ToStr(msg.User_info["target_id"])
      target_type = common.ToStr(msg.User_info["target_type"])
    case DEF_MSG_IPSAP_WITHDRAW_REQUEST_NOTI : // IPSAP 기관 계정 탈퇴 신청 알림 -> 관리자
      target_id   = common.ToStr(msg.User_info["target_id"])
    case DEF_MSG_IPSAP_WITHDRAW_FINISHED_NOTI : // IPSAP 기관 계정 탈퇴 완료 알림 -> 관리자
      target_id   = common.ToStr(row2["email"])
  }

  r := strings.NewReplacer("${기관명}", common.ToStr(msg.User_info["institution_name_ko"]),
                           "${임시비밀번호}", common.ToStr(msg.User_info["tmp_pwd"]),
                           "${수신자}", common.ToStr(msg.User_info["name"]),
                           "${대상자}", target_name,
                           "${아이디}", target_id,
                           "${권한}", target_type,
                           "${위원회}", common.ToStr(data["judge_type_str"]),
                           "${신청서유형}", common.ToStr(data["application_type_str"]),
                           "${과제명}", common.ToStr(data["name_ko"]),
                           "${접수번호}", common.ToStr(data["application_no"]),
                           "${연구책임자}", common.ToStr(data["user_name"]),
                           "${진행단계}", common.ToStr(data["application_step_str"]),
                           "${진행상태}", status,
                           "${최종결과}", common.ToStr(row["final_judge_result"]),
                           "${서비스관리자전화}", admin_telno,
                           "${서비스관리자email}", admin_email,
                           "${행정간사}", common.ToStr(row2["name"]),
                           "${행정간사전화}", common.ToStr(row2["telno"]),
                           "${행정간사email}", common.ToStr(row2["email"]),
                           "${심사기한}", deadline,
                           "${차수}", round,
                           "${심사잔여시간}", remainingTime,
                           "${철회가능기간}", common.ToStr(msg.User_info["retractable_period"]),
                           "${탈퇴일}", common.ToStr(msg.User_info["remove_dttm"]),
                           "${탈퇴처리일}", common.ToStr(msg.User_info["remove_dttm"]),
                           "${탈퇴신청접수일}", common.ToStr(msg.User_info["submit_withdraw_dttm"]),
                           "${메일링크}", reqUrl,
                           "/assets/images/common/IPSAP_logo_top_fitted.png", "cid:IPSAP_logo_top_fitted.png",
                           "/attach/org/sample_org/sample_logo.png", getMailLogoFile(common.ToStr(row2["logo_file_path"])));

  contents, fileName, subject := msg.GetMsgTemplateInfo()
  if isSms {
    smsText = r.Replace(contents)
  } else {
    emailTemplate, err := ioutil.ReadFile(common.Config.Email.TemplatePath + fileName)
    if nil != err {
     log.Println(err)
     return
    }
    subject = r.Replace(subject)
    mailText = r.Replace(string(emailTemplate))
  }

  return
}

func (msg *MessageMgr)getMembershipMsgContent(isSms bool) (subject string, mailText string, smsText string) {
  // 멤버십 만료	멤버십 플랜 만료월 21일 ~ 만료월 말일
  admin_email, admin_telno := GetServiceAdminInfo()
  instt := Institution{
    Institution_seq : msg.Institution_seq,
  }
  instt.GetInstitutionPaymentMsgData()

  r := strings.NewReplacer("${기관명}", common.ToStr(instt.Data["name_ko"]),
                           "${수신자}", common.ToStr(msg.User_info["name"]),
                           DEF_MEMBERSHIP_REFUND, msg.Mebership_info[DEF_MEMBERSHIP_REFUND],
                           DEF_USAGE_REFUND, msg.Mebership_info[DEF_USAGE_REFUND],
                           DEF_MEMBERSHIP_JOIN_DATE, msg.Mebership_info[DEF_MEMBERSHIP_JOIN_DATE],
                           DEF_MEMBERSHIP_EXPIRATION_DATE, msg.Mebership_info[DEF_MEMBERSHIP_EXPIRATION_DATE],
                           DEF_DESIGNATED_PAYMENT_METHOD, msg.Mebership_info[DEF_DESIGNATED_PAYMENT_METHOD],
                           DEF_REFUND_APPLICATION_DATE, msg.Mebership_info[DEF_REFUND_APPLICATION_DATE],
                           DEF_MONTH_START_DATE, msg.Mebership_info[DEF_MONTH_START_DATE],
                           DEF_MONTH_END_DATE, msg.Mebership_info[DEF_MONTH_END_DATE],
                           DEF_FREE_START_DATE, msg.Mebership_info[DEF_FREE_START_DATE],
                           DEF_FREE_END_DATE, msg.Mebership_info[DEF_FREE_END_DATE],
                           DEF_CANCEL_APPLICATION_DATE, msg.Mebership_info[DEF_CANCEL_APPLICATION_DATE],
                           DEF_PAYMENT_DEADLINE, msg.Mebership_info[DEF_PAYMENT_DEADLINE],
                           DEF_STOP_DATE, msg.Mebership_info[DEF_STOP_DATE],
                           DEF_PAYMENT_DUE_DATE, msg.Mebership_info[DEF_PAYMENT_DUE_DATE],
                           DEF_MONTHLY_FEE, msg.Mebership_info[DEF_MONTHLY_FEE],
                           DEF_MEMBERSHIP_NAME, msg.Mebership_info[DEF_MEMBERSHIP_NAME],
                           DEF_PAYMENT_DATE, msg.Mebership_info[DEF_PAYMENT_DATE],
                           DEF_PAYMENT_NUMBER, msg.Mebership_info[DEF_PAYMENT_NUMBER],
                           DEF_PAYMENT_AMOUNT, msg.Mebership_info[DEF_PAYMENT_AMOUNT],
                           DEF_PAYMENT_METHOD, msg.Mebership_info[DEF_PAYMENT_METHOD],
                           DEF_FREE_DAYS, msg.Mebership_info[DEF_FREE_DAYS],
                           DEF_CANCEL_DATE, msg.Mebership_info[DEF_CANCEL_DATE],
                           DEF_CANCEL_TOTAL_AMOUNT, msg.Mebership_info[DEF_CANCEL_TOTAL_AMOUNT],
                           DEF_MEMBERSHIP_NAME_OLD, msg.Mebership_info[DEF_MEMBERSHIP_NAME_OLD],
                           "${서비스관리자전화}", admin_telno,
                           "${서비스관리자email}", admin_email,
                           "/assets/images/common/IPSAP_logo_top_fitted.png", "cid:IPSAP_logo_top_fitted.png",
                           "/attach/org/sample_org/sample_logo.png", getMailLogoFile(common.ToStr(instt.Data["logo_file_path"])));

  contents, fileName, subject := msg.GetMsgTemplateInfo()
  if isSms {
    smsText = r.Replace(contents)
  } else {
    emailTemplate, err := ioutil.ReadFile(common.Config.Email.TemplatePath + fileName)
    if nil != err {
     log.Println(err)
     return
    }
    subject = r.Replace(subject)
    mailText = r.Replace(string(emailTemplate))
  }
  return
}

func (msg *MessageMgr)GetMsgTemplateInfo() (contents string, fileName string, subject string) {
  sql := `SELECT temp.subject, temp.file_name, temp.contents
           FROM t_msg_template temp
          WHERE temp.msg_id = ?`
  row := common.DB_fetch_one(sql, nil, msg.Msg_ID)
  contents  = common.ToStr(row["contents"])
  fileName  = common.ToStr(row["file_name"])
  subject   = common.ToStr(row["subject"])
  return
}

func getMailLogoFile(logo_file_path string) (url string) {
  if "" == logo_file_path {
    url = "cid:IPSAP_logo_top_fitted.png"
    return
  }

  if common.Config.Test.IsTestMode {
    url = common.GetPresignUrl(logo_file_path)
  } else {
    url = common.MakeDownloadUrl(common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), logo_file_path))
  }
  return
}
