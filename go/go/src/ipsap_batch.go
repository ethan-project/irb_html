package main

import (
  rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/mitchellh/mapstructure"
	"github.com/nleeper/goment"
	"github.com/mileusna/crontab"
	"github.com/BurntSushi/toml"
	resty "github.com/go-resty/resty/v2"
  "github.com/spf13/cast"
	"encoding/json"
	"ipsap/common"
	"ipsap/model"
  "io/ioutil"
	"math/rand"
	"strings"
  "math"
	"time"
	"fmt"
	"log"
	"os"
)

type reqBillingApprove struct{
	Order model.Order
	BuyerEmail string
	BuyerName string
	BuyerTel string
}

type respBillingApprove struct {
	ResultCode 		string	`json:"ResultCode"`
  ResultMsg  		string	`json:"ResultMsg"`
	Tid						string	`json:"TID"`
	Moid					string	`json:"Moid"`
	Amt						string	`json:"Amt"`
	AuthCode			string	`json:"AuthCode"`
	AuthDate			string	`json:"AuthDate"`
  AcquCardCode	string	`json:"AcquCardCode"`
  AcquCardName	string	`json:"AcquCardName"`
  CardNo				string	`json:"CardNo"`
  CardCode			string	`json:"CardCode"`
  CardName 			string	`json:"CardName"`
  CardQuota			string	`json:"CardQuota"`
  CardCl				string	`json:"CardCl"`
  CardInterest	string	`json:"CardInterest"`
  CcPartCl			string	`json:"CcPartCl"`
}


func main() {
	configPath := "/root/go/config/batch.toml"
	argsWithProg := os.Args
	if len(argsWithProg[1:]) != 0 {
		configPath = argsWithProg[1]
	}

	// Config 설정
	if _, err := toml.DecodeFile(configPath, &common.Config); err != nil {
		log.Println(err)
		return
	}

  common.DB_pool_connect()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
  setupLogOutput()

	if common.Config.Test.IsTestMode {
		common.Config.S3.BUCKET = common.Config.S3.BUCKET_TEST
	} else {
		common.Config.S3.BUCKET = common.Config.S3.BUCKET_REAL
	}


	log.Println("IPSAP Batch Start")

	ctab := crontab.New()
	ctab.MustAddJob("50 7 * * *", startBatch) // 매월 1일 아침 8시 정기결제
	// *     *     *     *     *
	//
	// ^     ^     ^     ^     ^
	// |     |     |     |     |
	// |     |     |     |     +----- day of week (0-6) (Sunday=0)
	// |     |     |     +------- month (1-12)
	// |     |     +--------- day of month (1-31)
	// |     +----------- hour (0-23)
	// +------------- min (0-59)

	for {

	}
}

func setupLogOutput() {
  if common.Config.Server.LogPath != "" {
    log.SetOutput(setLog("log_file"))
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

func sendReAppMsg() {
  // IACUC 승인일을 기준으로 355일 경과시(1년 경과 10일전)에
  // 재승인신청을 하라는 메일/문자 메시지를 연구책임자와 행정간사에게 발송
  // IACUC 승인일을 기준으로 366일 경과시(1년 경과 1일후)에
  // 재승인신청을 하지 않은 경우 재승인신청 독촉 메일/문자 메시지를 연구책임자와 행정간사에게 발송
  sql2 := fmt.Sprintf(`
          SELECT etc2.contents
            FROM t_application app2, t_application_etc etc2
           WHERE app2.application_seq = etc2.application_seq
             AND app2.parent_app_seq = app.application_seq
             AND app2.application_type = %d
             AND app2.application_step = %d
             AND app2.application_result IN (%d, %d)
             AND etc2.item_name = 'general_end_date'
        ORDER BY app2.approved_dttm DESC
           LIMIT 1`, model.DEF_APP_TYPE_CHANGE, model.DEF_APP_STEP_FINAL,
            model.DEF_APP_RESULT_APPROVED, model.DEF_APP_RESULT_APPROVED_C)

  sql3 := fmt.Sprintf(`
				SELECT COUNT(app3.application_seq)
					 FROM t_application app3
					WHERE app3.parent_app_seq = app.application_seq
						AND app3.application_type = %d
						AND app3.application_step = %d
						AND app3.application_result IN (%d, %d)`, model.DEF_APP_TYPE_RENEW, model.DEF_APP_STEP_FINAL,
								model.DEF_APP_RESULT_APPROVED, model.DEF_APP_RESULT_APPROVED_C)

  sql := fmt.Sprintf(`SELECT app.application_seq,
                             app.approved_dttm,
                             IFNULL((%v), etc.contents) AS general_end_date,
                             (%v) as renew_app_cnt
                        FROM t_application app
                        LEFT OUTER JOIN t_application_etc etc ON (app.application_seq = etc.application_seq AND etc.item_name = 'general_end_date')
                       WHERE app.application_type = 1
                         AND app.judge_type = 1
                         AND app.application_step = 5
                         AND app.application_result IN(11, 12)`, sql2, sql3)
  rows := common.DB_fetch_all(sql, nil)
  for _, row := range rows {
    startUnixT := common.ToInt64(row["approved_dttm"])
    if startUnixT > 0 && nil != row["general_end_date"] {
      startT, _ := goment.Unix(startUnixT)
      endDateT, _ := goment.New(row["general_end_date"])
      timeDiffY := math.Abs(cast.ToFloat64(startT.Diff(endDateT, "years")))
      approvedRenewAppCnt := common.ToInt(row["renew_app_cnt"])
      toDay, _ := goment.New()
      if timeDiffY > 0 {
        tempStartT11, _ := goment.Unix(startUnixT)
        tempStartT11.Add(11, "months")
        tempStartT12, _ := goment.Unix(startUnixT)
        tempStartT12.Add(12, "months")
        app_seq := common.ToUint(row["application_seq"])
        // 재승인 신청 기한
        deadline := tempStartT12.Format("YYYY-MM-DD")
        tempStartT12.Add(1, "days")
        if toDay.IsSame(tempStartT11) {
          sendReAppMsg2(app_seq, model.DEF_MSG_RE_APP_BEFORE, deadline)
        }

        if approvedRenewAppCnt == 0 && toDay.IsSame(tempStartT12) {
          sendReAppMsg2(app_seq, model.DEF_MSG_RE_APP_AFTER, deadline)
        }

        if timeDiffY > 2 {
          tempStartT23, _ := goment.Unix(startUnixT)
          tempStartT24, _ := goment.Unix(startUnixT)
          tempStartT23.Add(23, "months")
          tempStartT24.Add(24, "months")
          deadline = tempStartT24.Format("YYYY-MM-DD")

          if toDay.IsSame(tempStartT23) {
            sendReAppMsg2(app_seq, model.DEF_MSG_RE_APP_BEFORE, deadline)
          }

          tempStartT24.Add(1, "days")
          if approvedRenewAppCnt <= 1  && toDay.IsSame(tempStartT24) {
            sendReAppMsg2(app_seq, model.DEF_MSG_RE_APP_AFTER, deadline)
          }
        }
      }
    }
  }
}

func sendReAppMsg2(app_seq uint, msg_id int, deadline string) {
  application := model.Application {
    Application_seq : app_seq,
  }

  data := application.Load().(map[string]interface{})
  user := model.User{
    Institution_seq : common.ToUint(data["institution_seq"]),
  }
  row2 := user.GetInstitutionAdminInfo("")

  // 책임연구원에게 전송
  subject, mailText, smsText := getMsgContent(msg_id, data, row2, common.ToStr(data["user_name"]), deadline)
  sendEmail(common.ToStr(data["user_email"]), subject, mailText)
  sendSms(common.ToUint(data["institution_seq"]), common.ToUint(data["reg_user_seq"]), common.ToStr(data["user_phoneno"]), smsText)

  // 행정간사에게 전송
  subject, mailText, smsText = getMsgContent(msg_id, data, row2, common.ToStr(row2["name"]), deadline)
  sendEmail(common.ToStr(row2["email"]), subject, mailText)
  sendSms(common.ToUint(data["institution_seq"]), common.ToUint(row2["user_seq"]), common.ToStr(row2["telno"]), smsText)
}

func sendEmail(emailAddr string, subject string, mailText string) {
  email := common.Email{}
  email.To = common.ToStr(emailAddr)
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
}

func sendSms(institution_seq uint, user_seq uint, phoneno string, smsText string) {
  sql := `INSERT INTO t_sms(institution_seq, user_seq, type, phoneno, msg, result, reg_dttm)
          VALUES(?,?,?,?,?,?,UNIX_TIMESTAMP())`
  _, err := common.DBconn().Exec(sql, institution_seq,
                                      user_seq,
                                      "lms",
                                      phoneno,
                                      smsText,
                                      model.DEF_SMS_RESULT_WAIT)
  if nil != err {
    log.Println(err)
    return
  }
}

func getMsgContent(msg_id int, data map[string]interface{}, row2 map[string]interface{}, name string, deadline string) (subject string, mailText string, smsText string) {
  msg := model.MessageMgr {
    Msg_ID : msg_id,
  }

  r := strings.NewReplacer("${기관명}", common.ToStr(data["istt_name_ko"]),
                           "${수신자}", name,
                           "${신청서유형}", common.ToStr(data["application_type_str"]),
                           "${과제명}", common.ToStr(data["name_ko"]),
                           "${접수번호}", common.ToStr(data["application_no"]),
                           "${연구책임자}", common.ToStr(data["user_name"]),
                           "${서비스관리자email}", "support@ipsap.co.kr",
                           "${행정간사}", common.ToStr(row2["name"]),
                           "${행정간사전화}", common.ToStr(row2["telno"]),
                           "${행정간사email}", common.ToStr(row2["email"]),
                           "${재승인신청기한}", deadline,
                           "${위원회}", "IACUC",
                           "${메일링크}", "https://www.ipsap.co.kr",
                           "/assets/images/common/IPSAP_logo_top_fitted.png", "cid:IPSAP_logo_top_fitted.png",
                           "/attach/org/sample_org/sample_logo.png", getMailLogoFile(common.ToStr(row2["logo_file_path"])));

  contents, fileName, subject := msg.GetMsgTemplateInfo()
  smsText = r.Replace(contents)
  emailTemplate, err := ioutil.ReadFile(common.Config.Email.TemplatePath + fileName)
  if nil != err {
   log.Println(err)
   return
  }
  subject = r.Replace(subject)
  mailText = r.Replace(string(emailTemplate))

  return
}

// 1. 메일 전송
// 2. 정기 결제
// 3. 기관 상태 업데이트
// 매일 아침 8시경 모든 동작이 된다.
func startBatch() {
  sendReAppMsg()

	// 오늘
	now, _ := goment.New()
	today := now.Format("YYYYMMDD")

	t1, _ := goment.New()
	t1.StartOf("month")
	regularPaymentDate := t1.Format("YYYYMMDD")	// 정기 결제일
	paymentDueDateStr := common.GetDateStr(t1) // 결제 예정일
	t1.Add(1, "days")
	unPaidDate := t1.Format("YYYYMMDD")
	t1.Add(5, "days")
	stopScheduledDate := t1.Format("YYYYMMDD")
	paymentDeadlineStr := common.GetDateStr(t1)

	t1.Add(1, "days")
	stopDate := t1.Format("YYYYMMDD")
	suspensionDateStr := common.GetDateStr(t1) // 이용중지일

	t2, _ := goment.New()
	t2.EndOf("month")
	oneDaysBeforePaymentDate :=  t2.Format("YYYYMMDD")
	t2.Add(-6, "days")
	sevenDaysBeforePaymentDate := t2.Format("YYYYMMDD")
	t2.Add(-1, "days")
	freePeriodApplyDate := t2.Format("YYYYMMDD")

	switch today {
		case regularPaymentDate : // 정기 결제일
			regularPaymentStart()
		case unPaidDate: // 결제일 + 1일
			unPaidSend(paymentDeadlineStr)
		case stopScheduledDate:	// 결제일 + 6일
			stopScheduled(paymentDeadlineStr, suspensionDateStr)
		case stopDate:	// 결제일 + 7일
			stopInstitution(suspensionDateStr)
		case freePeriodApplyDate : // 결제일 -8일(결제 예정 문자 보내기 전)
			freePeriodApply()
		case sevenDaysBeforePaymentDate:	// 결제일 - 7일
			scheduledPayment(paymentDueDateStr)
		case oneDaysBeforePaymentDate:	// 결제일 - 1일
		  scheduledPayment(paymentDueDateStr)
		default :
			return
	}

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

func getInsttitutionPaymentQuery(moreCondition string) (sql string) {
	if "" == moreCondition {
		moreCondition = `AND date_format(now(), '%Y%m%d') > instt.expiration_date #만료일이 지났을때`
	}

	sql = fmt.Sprintf(`
					SELECT instt.institution_seq, plan.name, prod.discounted_amount,
								 (SELECT group_concat(user.user_seq)
										FROM t_user user
									 WHERE user.institution_seq = instt.institution_seq
										 AND user.user_type LIKE '%%%v%%'
										 AND user.user_status = 2) AS user_arr,
								 IF(instt.payment_setting = 1, "신용카드 - 직접", "신용카드 - 자동") as payment_method
						FROM t_institution instt
						LEFT OUTER JOIN t_products prod ON (instt.product_seq = prod.product_seq)
						LEFT OUTER JOIN t_membership_plan plan ON (prod.plan_seq = plan.plan_seq)
					 WHERE 1 = 1
					   %v	#moreCondition
						 AND instt.service_status = %v # 이용중
						 AND instt.membership_fee_status > 0 # 회원가입비 지불
						 AND instt.product_seq > 0 #사용중인 요금제가 있을때`,
						 model.DEF_USER_TYPE_ADMIN_SECRETARY,
 						 moreCondition,
					   model.DEF_SERVICE_STATUS_IN_USE)
  log.Println(sql)
	return
}

func sendMsg(row map[string]interface{}, msgId int, adminMsgId int) {
	membershipInfo := make(map[string]string)
	membershipInfo[model.DEF_MEMBERSHIP_NAME] = common.ToStr(row["name"])
	membershipInfo[model.DEF_MONTHLY_FEE] = common.GetCurrency(common.ToInt64(row["discounted_amount"]))
	membershipInfo[model.DEF_PAYMENT_DEADLINE] = common.ToStr(row[model.DEF_PAYMENT_DEADLINE])
	membershipInfo[model.DEF_STOP_DATE] = common.ToStr(row[model.DEF_STOP_DATE])
	membershipInfo[model.DEF_PAYMENT_DUE_DATE] = common.ToStr(row[model.DEF_PAYMENT_DUE_DATE])
	membershipInfo[model.DEF_PAYMENT_METHOD] = common.ToStr(row["payment_method"])
	userArr := strings.Split(common.ToStr(row["user_arr"]), ",")
	institutionSeq := common.ToUint(row["institution_seq"])
	msg := model.MessageMgr{}
	for _, userSeq := range userArr {
		user := model.User {
			User_seq : common.ToUint(userSeq),
		}
		if !user.Load(){
			continue;
		}
		msg.Msg_ID = msgId
		msg.User_info = user.Data
		msg.Mebership_info = membershipInfo
		msg.Institution_seq = institutionSeq
		msg.SendMessage()
	}

	if adminMsgId > 0 {
		msg.Msg_ID = adminMsgId
		msg.User_info = model.LoadServiceAdmin()
		msg.SendMessage()
	}
	return
}

// 무료 지급일
func freePeriodApply() {
	sql := `SELECT *
  					FROM t_institution_free_period
 					 WHERE deleted_flag = 0
   		 			 AND purchased_date = 0
 				GROUP BY institution_seq
 					HAVING reg_date =  MIN(reg_date)`
	rows := common.DB_fetch_all(sql, nil)
	for _, row := range rows {
	 	freeM := model.FreeMembership {
			Free_seq				: common.ToUint(row["free_seq"]),
			Institution_seq : common.ToUint(row["institution_seq"]),
			Usage_limit 		: common.ToUint(row["usage_limit"]),
			Free_period			: common.ToStr(row["free_period"]),
		}
		freeM.Apply(true)
	}
}

// 결제 예정 안내
func scheduledPayment(paymentDueDateStr string) {
	// 현재월과 만료월이 같을때
	moreCondition := `AND date_format(now(), '%Y%m') = date_format(date(instt.expiration_date), '%Y%m')`
	sql := getInsttitutionPaymentQuery(moreCondition)
	rows := common.DB_fetch_all(sql, nil)
	if len(rows) > 0 {
		for _, row := range rows {
			row[model.DEF_PAYMENT_DUE_DATE] = paymentDueDateStr
			sendMsg(row, model.DEF_MSG_MEMBERSHIP_PAYMENT_EXPECTED, 0)
		}
	}
}

// 이용 중지 알림
func stopInstitution(suspensionDateStr string) {
	sql := getInsttitutionPaymentQuery("")
	rows := common.DB_fetch_all(sql, nil)
	if len(rows) > 0 {
		for _, row := range rows {
			row[model.DEF_STOP_DATE] = suspensionDateStr
			sendMsg(row, model.DEF_MSG_MEMBERSHIP_STOP, model.DEF_MSG_MEMBERSHIP_STOP_NOTI)
			instt := model.Institution{
				Institution_seq : common.ToUint(row["institution_seq"]),
			}
			// 기관 상태 이용중지로 update
			instt.StopInstitution()
		}
	}
}

// 이용 중지 예정 안내
func stopScheduled(paymentDeadlineStr string, suspensionDateStr string) {
	sql := getInsttitutionPaymentQuery("")
	rows := common.DB_fetch_all(sql, nil)
	if len(rows) > 0 {
		for _, row := range rows {
			row[model.DEF_PAYMENT_DEADLINE] = paymentDeadlineStr
			row[model.DEF_STOP_DATE] = suspensionDateStr
			sendMsg(row, model.DEF_MSG_MEMBERSHIP_STOP_EXPECTED, 0)
		}
	}
}

// 미납 안내
func unPaidSend(paymentDeadlineStr string) {
	sql := getInsttitutionPaymentQuery("")
	rows := common.DB_fetch_all(sql, nil)
	if len(rows) > 0 {
		for _, row := range rows {
			row[model.DEF_PAYMENT_DEADLINE] = paymentDeadlineStr
			sendMsg(row, model.DEF_MSG_MEMBERSHIP_UNPAID, model.DEF_MSG_MEMBERSHIP_UNPAID_NOTI)
		}
	}
}

func regularPaymentStart() {
	sql := ` SELECT instt.bid, plan.name AS plan_name, resp_bill.buyer_email,
									prod.discounted_amount, resp_bill.user_seq,
									instt.product_seq, instt.institution_seq
						 FROM t_institution instt, t_products prod, t_membership_plan plan, t_resp_pg_bill_key resp_bill
						WHERE instt.product_seq = prod.product_seq
							AND instt.bid = resp_bill.bid
							AND prod.plan_seq = plan.plan_seq
							AND instt.bid <> ""
							AND instt.membership_fee_status > 0 # 회원가입이 되어 있을때
							AND instt.payment_setting = 2 # 자동
							AND instt.service_status = 1 # 서비스 이용중
							AND plan.deleted_flag	= 0 # 삭제 안됨
							AND plan.plan_available	= 1 # 사용중
							AND date_format(now(), '%%Y%%m%%d') > instt.expiration_date #만료일이 지났을때`
	rows := common.DB_fetch_all(sql, nil)
	if len(rows) > 0 && nil != rows {
		now, _ := goment.New()
		for _, row := range rows {
			ediDate := now.Format("YYYYMMDDHHmmss")
			timeInfo := now.Format("YYMMDDHHmmss")
			rand.Seed(time.Now().UnixNano())
			randomData := common.ToStr(rand.Intn(9999 - 1001) + 1000)
			ord := model.Order {
				Institution_seq : common.ToUint(row["institution_seq"]),
				Product_seq : common.ToUint(row["product_seq"]),
				User_seq : common.ToUint(row["user_seq"]),
				Pay_date : ediDate,
				Pname : common.ToStr(row["plan_name"]),
				Order_type : model.DEF_ORDER_TYPE_REGULAR_PAYMENT,
				Pay_method : model.DEF_PAY_METHOD_CARD,
				Amount : common.ToStr(row["discounted_amount"]),
				Tid : common.Config.Payment.REGULAR_MID + model.DEF_TID_CARD + model.DEF_TID_DIVISION + timeInfo + randomData,
				Bid : common.ToStr(row["bid"]),
				MerchantKey : common.Config.Payment.REGULAR_MERCHANTKEY,
				Mid : common.Config.Payment.REGULAR_MID,
			}

			ord.GetAssginDataAndMoid(false)
			if !ord.InsertOrder(nil){
				log.Printf("정기 결제 Insert Order error")
				return
			}

			reqBill := reqBillingApprove{
				Order : ord,
				BuyerEmail : common.ToStr(row["buyer_email"]),
			}

			succ, authDate := reqBill.requestBilling()
			if ! succ{
				return
			}

			// 기관 상태 업테이트
			instt := model.Institution{
				Institution_seq : common.ToUint(row["institution_seq"]),
			}

			if !instt.UpdateInstitutionRegularPaymentResult() {
				return
			}

			ord.Auth_date = authDate
			go paymentFinishSendMsg(ord)
		}
	}
}

func paymentFinishSendMsg(ord model.Order) {
	user := model.User {
		User_seq : ord.User_seq,
	}

	if !user.Load() {
		return
	}

	membershipInfo := make(map[string]string)
	membershipInfo[model.DEF_MEMBERSHIP_NAME] = ord.Pname
	membershipInfo[model.DEF_MONTHLY_FEE] = common.GetCurrency(common.ToInt64(ord.Amount))
	t1, _ := goment.Unix(common.ToInt64(user.Data["membership_payment_date"]))
	membershipJoinDate, _ := goment.New(t1)
	membershipInfo[model.DEF_MEMBERSHIP_JOIN_DATE] = common.GetDateStr(membershipJoinDate)
	endDate, _ := goment.New(common.ToStr(user.Data["expiration_date"]))
	membershipInfo[model.DEF_MEMBERSHIP_EXPIRATION_DATE] = common.GetMembershipExpirationDate(membershipJoinDate)
	membershipInfo[model.DEF_MONTH_END_DATE] = common.GetDateStr(endDate)
	membershipInfo[model.DEF_MONTH_START_DATE] = common.GetStartDateStr(endDate)
	membershipInfo[model.DEF_PAYMENT_NUMBER] = ord.Moid
	membershipInfo[model.DEF_PAYMENT_METHOD] = "신용카드 - 자동결제"
	membershipInfo[model.DEF_PAYMENT_AMOUNT] = common.GetCurrency(common.ToInt64(ord.Amount))

	paymentDate, _ := goment.New(ord.Auth_date, "YYMMDDHHmmSS")
	membershipInfo[model.DEF_PAYMENT_DATE] = common.GetDateTimeStr(paymentDate)

	msg := model.MessageMgr{
		Msg_ID : model.DEF_MSG_MEMBERSHIP_PAYMENT,
		User_info : user.Data,
		Mebership_info : membershipInfo,
		Institution_seq : ord.Institution_seq,
	}

	msg.SendMessage()
}

// signDataStr := common.Config.Payment.REGULAR_MID + ediDate + moid + amt + bid + common.Config.Payment.REGULAR_MERCHANTKEY
// tid 생성 order_type (2 : 정기 결제)
func (reqBill reqBillingApprove)requestBilling() (succ bool, authDate string) {
	succ = false
	sql := `INSERT INTO t_req_pg_billing_appove(
						bid, mid, tid, edi_date,
						moid, amt, goods_name,
						sign_data, card_interest, card_quota,
						buyer_name, buyer_email, buyer_tel
					)
					VALUES
						(
							?,?,?,?,
							?,?,?,
							?,?,?,
							?,?,?
						)`
	_, err := common.DBconn().Exec(sql,
																reqBill.Order.Bid, reqBill.Order.Mid, reqBill.Order.Tid, reqBill.Order.Pay_date,
																reqBill.Order.Moid, reqBill.Order.Amount, reqBill.Order.Pname,
																reqBill.Order.Sign_data,  model.DEF_BIILING_CARDINTEREST, model.DEF_BIILING_CARDINTEREST,
																reqBill.BuyerEmail, reqBill.BuyerName, reqBill.BuyerTel)
 	if nil != err {
 		log.Println(err)
 		return
 	}

	client := resty.New()
	resp, err := client.R().
	SetHeader("Content-Type", "application/x-www-form-urlencoded").
	SetFormData(map[string]string{
			"BID" : reqBill.Order.Bid,
			"MID": reqBill.Order.Mid,
			"TID" : reqBill.Order.Tid,
			"EdiDate": reqBill.Order.Pay_date,
			"Moid": reqBill.Order.Moid,
			"Amt": reqBill.Order.Amount,
			"GoodsName" : reqBill.Order.Pname,
			"SignData": reqBill.Order.Sign_data,
			"CardInterest" : "0",
			"CardQuota" : "00",
			"CharSet" : "utf-8",
			"BuyerEmail" : reqBill.BuyerEmail,
	}).
	Post("https://webapi.nicepay.co.kr/webapi/billing/billing_approve.jsp")

	var jsonData map[string]interface{}
	if nil == err {
		if err2 := json.Unmarshal([]byte(common.ToStr(resp)), &jsonData); err2 != nil {
			log.Println(err2)
			return
		}
	} else {
		log.Println(err)
		return
	}

	respBilling := respBillingApprove{}
	if err = mapstructure.Decode(jsonData, &respBilling); nil != err {
		log.Println(err)
		return
	}

	if !respBilling.insert() {
		return
	}

	orderStatus := model.DEF_ORDER_STATUS_ERROR
	if respBilling.ResultCode == "3001" {
		orderStatus = model.DEF_ORDER_STATUS_COMPLETED
		succ = true
	}

	now, _ := goment.New()
	theDate := now.Format("YYYYMM")

	sql  = `UPDATE t_orders
						 SET order_status = ?, order_status_code = ?,
								 auth_date = ?, the_date = ?
					 WHERE moid = ? AND tid = ?`
	_, err = common.DBconn().Exec(sql,
									 orderStatus, respBilling.ResultCode,
									 respBilling.AuthDate, theDate,
									 respBilling.Moid, respBilling.Tid)
	if nil != err {
		log.Println(err)
		return
	}

	authDate = respBilling.AuthDate
	return
}

func (respBill respBillingApprove)insert() (succ bool) {
	succ = false
	sql := `INSERT INTO t_resp_pg_billing_appove(
						tid, moid, amt, result_code,
						result_msg, auth_code, auth_date,
						acqu_card_code, acqu_card_name,
						card_no, card_code, card_name,
						card_quota, card_cl, card_interest,
						cc_part_cl
				  )
					VALUES
					(
						?,?,?,?,
						?,?,?,
						?,?,
						?,?,?,
						?,?,?,
						?
					)`
	_, err := common.DBconn().Exec(sql,
																respBill.Tid, respBill.Moid, respBill.Amt, respBill.ResultCode,
																respBill.ResultMsg, respBill.AuthCode, respBill.AuthDate,
																respBill.AcquCardCode, respBill.AcquCardName,
																respBill.CardNo, respBill.CardCode, respBill.CardName,
																respBill.CardQuota, respBill.CardCl, respBill.CardInterest,
																respBill.CcPartCl)
	if nil != err {
 		log.Println(err)
 		return
 	}

	succ = true
	return
}
