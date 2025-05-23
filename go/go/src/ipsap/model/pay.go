package model

import (
	resty "github.com/go-resty/resty/v2"
	"github.com/nleeper/goment"
	"crypto/sha256"
	"encoding/json"
	"ipsap/common"
	"encoding/hex"
	"database/sql"
	"math/rand"
	"log"
	"fmt"
	"time"
)

type ReqPay struct {
	Tid							string
	Institution_seq uint
	Product_seq			uint
	User_seq				uint
	AuthToken				string
	MID							string
	Amt							string
	EdiDate					string
	SignData				string
	CharSet					string
	NextAppURL			string
	NetCancelURL		string
}

func (reqPay *ReqPay) ApprovePay() (resultMsg string, succ bool) {
	succ = false
	orderStatus := DEF_ORDER_STATUS_ERROR
	resultMsg = "시스템 오류"
	payMethod := ""
	resultCode := uint(0)

	dbConn	:= common.DBconn()
	tx, err	:= dbConn.Begin()
	if nil != err {
		tx.Rollback()
		return
	}

	defer func() {
		tx.Rollback()
	}()

	t, _ := goment.New()
	reqPay.EdiDate = t.Format("YYYYMMDDHHmmss")
	signDataStr := reqPay.AuthToken + reqPay.MID + reqPay.Amt + reqPay.EdiDate + common.Config.Payment.GENERAL_MERCHANTKEY
	data := sha256.Sum256([]byte(signDataStr))
	var newByte []byte = data[:]
  reqPay.SignData = hex.EncodeToString(newByte)

	sql := `INSERT INTO t_req_pg_pay_api (
						tid, auth_token, mid,
						amt, edi_date, sign_data,
						char_set
					)
					VALUES
						(?,?,?,
						 ?,?,?,
						 ?)`
	 _, err = tx.Exec(sql,
										reqPay.Tid, reqPay.AuthToken, reqPay.MID,
										reqPay.Amt, reqPay.EdiDate, reqPay.SignData,
										reqPay.CharSet)
  if nil != err {
  	log.Println(err)
  	return
  }

	client := resty.New()
	resp, err := client.R().
	SetHeader("Content-Type", "application/x-www-form-urlencoded").
	SetFormData(map[string]string{
			"TID": reqPay.Tid,
			"AuthToken": reqPay.AuthToken,
			"MID": reqPay.MID,
			"Amt": reqPay.Amt,
			"EdiDate": reqPay.EdiDate,
			"SignData": reqPay.SignData,
			"CharSet" : reqPay.CharSet,
	}).
	Post(reqPay.NextAppURL)

	var jsonData map[string]interface{}
	if nil == err {
		if err2 := json.Unmarshal([]byte(common.ToStr(resp)), &jsonData); err2 != nil {
			log.Println(err2)
			return
		} else {
			// log.Println(jsonData)
			resultMsg = common.ToStr(jsonData["ResultMsg"])
			payMethod = common.ToStr(jsonData["PayMethod"])
			resultCode = common.ToUint(jsonData["ResultCode"])

			if 3001 == resultCode || 4000 == resultCode {
				orderStatus = DEF_ORDER_STATUS_COMPLETED
			}

			sql := `SELECT instt.membership_fee_status, prod.category,
										 date_format(NOW(), '%Y%m') as this_month,
									   date_format(LAST_DAY(IF(instt.expiration_date <> "", DATE(instt.expiration_date), NOW())) + interval 1 month, '%Y%m') as next_month,
										 date_format(LAST_DAY(NOW()), '%Y%m%d') as this_month_last_date,
									   date_format(LAST_DAY(IF(instt.expiration_date <> "", DATE(instt.expiration_date), NOW())) + interval 1 month, '%Y%m%d') as next_month_last_date,
										 (SELECT product_seq
											  FROM t_products
											 WHERE (SELECT plan_seq FROM t_products WHERE product_seq = ?) = plan_seq
											   AND category = 'plan') as plan_product_seq
								FROM t_orders ord
							  LEFT OUTER JOIN t_institution instt ON (instt.institution_seq = ord.institution_seq)
								LEFT OUTER JOIN t_products prod ON (prod.product_seq = ord.product_seq)
							 WHERE moid = ?
							 	 AND tid = ?
								 AND ord.order_type IN (1, 2)`
			row := common.DB_Tx_fetch_one(tx, sql, nil, reqPay.Product_seq, jsonData["Moid"], jsonData["TID"])
			moreUpdate := ""
			expiration_date := common.ToStr(row["this_month_last_date"])
			the_date 	:= common.ToStr(row["this_month"])
			jsonData["Category"] = row["category"]

			if common.ToStr(row["category"]) == "membership" {
				moreUpdate = fmt.Sprintf(`membership_fee_status = 1,
																	membership_payment_date = UNIX_TIMESTAMP(),
																	product_seq = %v,`, row["plan_product_seq"])
			} else {
				if common.ToStr(row["expiration_date"]) != expiration_date {
					the_date = common.ToStr(row["next_month"])
					expiration_date = common.ToStr(row["next_month_last_date"])
				}
			}

			sql = fmt.Sprintf(`
						 UPDATE t_institution
								SET %v #moreUpdate
										expiration_date = ?,
										service_status = ?,
										usage_limit = 0,
										free_start_date = '',
										free_end_date = ''
							WHERE institution_seq = ?`, moreUpdate)
			_, err := tx.Exec(sql, expiration_date, DEF_SERVICE_STATUS_IN_USE, reqPay.Institution_seq)
			if nil != err {
				log.Println(err)
				return
			}

			cardName := common.ToStr(jsonData["CardName"])
			if "" != cardName {
				payMethod = "CARD"
			} else {
				payMethod = "BANK"
			}

			sql  = `UPDATE t_orders
								 SET order_status = ?, order_status_code = ?,
								 		 auth_date = ?, the_date = ?, pay_method = ?
							 WHERE moid = ? AND tid = ?`
			_, err = tx.Exec(sql,
											 orderStatus, resultCode,
											 jsonData["AuthDate"], the_date, payMethod,
											 jsonData["Moid"], jsonData["TID"])
		  if nil != err {
		  	log.Println(err)
		  	return
		  }

			if !insertRespPayApi(tx, jsonData) {
				return
			}
		}
	} else {
		resultMsg = "결제 요청 실패"
		return
	}

	err = tx.Commit()
	if nil != err {
		log.Println(err)
		return
	}

	go reqPay.SendMsgMembership(jsonData)

	if orderStatus == DEF_ORDER_STATUS_COMPLETED {
		succ = true
	}

	return
}

func (reqPay *ReqPay)SendMsgMembership(data map[string]interface{}) {
	user := User {
		User_seq : reqPay.User_seq,
	}

	if !user.Load() {
		return
	}

	membershipInfo := make(map[string]string)
	membershipInfo[DEF_MEMBERSHIP_NAME] = common.ToStr(data["GoodsName"])
	t1, _ := goment.Unix(common.ToInt64(user.Data["membership_payment_date"]))
	membershipJoinDate, _ := goment.New(t1)
	membershipInfo[DEF_MEMBERSHIP_JOIN_DATE] = common.GetDateStr(membershipJoinDate)
	endDate, _ := goment.New(common.ToStr(user.Data["expiration_date"]))
	d1 := membershipJoinDate.Get("dates")
	d2 := endDate.Get("dates")
	freeDays := d2 - d1 + 1
	membershipInfo[DEF_MEMBERSHIP_EXPIRATION_DATE] = common.GetMembershipExpirationDate(membershipJoinDate)
	paymentDate, _ := goment.New(common.ToStr(data["AuthDate"]), "YYMMDDHHmmSS")
	membershipInfo[DEF_PAYMENT_DATE] = common.GetDateTimeStr(paymentDate)
	membershipInfo[DEF_PAYMENT_NUMBER] = common.ToStr(data["Moid"])
	membershipInfo[DEF_PAYMENT_AMOUNT] = common.GetCurrency(common.ToInt64(reqPay.Amt))
	membershipInfo[DEF_FREE_END_DATE] = common.GetDateStr(endDate)

	// 멤버십 회원 가입일 경우
	if common.ToStr(data["Category"]) ==  "membership" {
		msg := MessageMgr{
			Msg_ID : DEF_MSG_MEMBERSHIP_JOIN,
			User_info : user.Data,
			Mebership_info : membershipInfo,
			Institution_seq : reqPay.Institution_seq,
		}
		msg.SendMessage()

		msg.Msg_ID = DEF_MSG_MEMBERSHIP_JOIN_NOTI
		msg.User_info = LoadServiceAdmin()
		membershipInfo[DEF_FREE_DAYS] = common.ToStr(freeDays)
		msg.SendMessage()
	} else {
		membershipInfo[DEF_MONTHLY_FEE] = common.GetCurrency(common.ToInt64(reqPay.Amt))
		membershipInfo[DEF_MONTH_END_DATE] = common.GetDateStr(endDate)
		membershipInfo[DEF_MONTH_START_DATE] = common.GetStartDateStr(endDate)
		membershipInfo[DEF_PAYMENT_METHOD] = "신용카드 - 직접결제)"
		msg := MessageMgr{
			Msg_ID : DEF_MSG_MEMBERSHIP_PAYMENT,
			User_info : user.Data,
			Mebership_info : membershipInfo,
			Institution_seq : reqPay.Institution_seq,
		}
		msg.SendMessage()
	}

	return
}

func insertRespPayApi(tx *sql.Tx, data map[string]interface{}) (succ bool) {
	succ = false
	sql := `INSERT INTO t_resp_pg_pay_api_comm(
						tid, result_code, result_msg,
						amt, mid, moid, buyer_email,
						buyer_tel, buyer_name, goods_name,
						auth_code, auth_date, pay_method
					)
					VALUES
					(
						?,?,?,
						?,?,?,?,
						?,?,?,
						?,?,?
					)`
	_, err := tx.Exec(sql,
										data["TID"], data["ResultCode"], data["ResultMsg"],
										data["Amt"], data["MID"], data["Moid"], data["BuyerEmail"],
										data["BuyerTel"], data["BuyerName"], data["GoodsName"],
										data["AuthCode"], data["AuthDate"], data["PayMethod"])
  if nil != err {
  	log.Println(err)
  	return
  }

	cardName := common.ToStr(data["CardName"])
	if "" != cardName {
		sql = fmt.Sprintf(`
					 INSERT INTO t_resp_pg_pay_api_card (
						tid, card_code, card_name, card_no,
						card_quota, card_interest, acqu_card_code,
						acqu_card_name, card_cl, cc_part_cl,
						clickpay_cl, point_app_amt
					 )
					 VALUES
					 (
						'%v','%v','%v','%v',
						'%v','%v','%v',
						'%v','%v','%v',
						'%v','%v'
					 )`,
					 data["TID"], data["CardCode"], data["CardName"], data["CardNo"],
					 data["CardQuota"], data["CardInterest"], data["AcquCardCode"],
					 data["AcquCardName"], data["CardCl"], data["CcPartCl"],
					 data["ClickpayCl"], data["PointAppAmt"])
	} else {
		sql = fmt.Sprintf(`
					 INSERT INTO t_resp_pg_pay_api_bank (
						tid, bank_code, bank_name,
						rcpt_type, rcpt_tid, rcpt_auth_code
					 )
					 VALUES
					 (
						'%v','%v','%v',
						'%v','%v','%v'
					 )`,
					 data["TID"], data["BankCode"], data["BankName"],
					 data["RcptType"], data["RcptTID"], data["RcptAuthCode"])
	}

	_, err = tx.Exec(sql)
	if nil != err {
		log.Println(err)
		return
	}

	succ = true
	return
}

// 모든 결제건은 규칙에 따라 유니크한 TID를 생성합니다.
// 따라서, 빌키 승인 요청 시 가맹점에서 생성하는 TID도 다른 모든 TID에 대하여 유니크해야 하며,
// 동일한 TID로 결제 요청 시 실패로 처리됩니다.
func GetTid() (tid string) {
	t := time.Now()
	timeInfo := t.Format("20060102150405") // 시간정보 yyMMddHHmmss, 12byte
	rand.Seed(time.Now().UnixNano())
	randomData := common.ToStr(rand.Intn(9999 - 1001) + 1000)

	tid = common.Config.Payment.REGULAR_MID + DEF_TID_CARD + DEF_TID_DIVISION + timeInfo + randomData
	return
}
