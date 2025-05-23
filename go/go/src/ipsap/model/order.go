package model

import (
	"github.com/mitchellh/mapstructure"
	resty "github.com/go-resty/resty/v2"
	"github.com/nleeper/goment"
	"encoding/json"
	"crypto/sha256"
	"encoding/hex"
	"database/sql"
	"ipsap/common"
	"crypto/md5"
	"log"
	"fmt"
)

type Order struct {
	Order_seq					uint		`json:"-"`
	Product_seq				uint		`json:"product_seq"`
	Institution_seq		uint		`json:"institution_seq"`
	User_seq					uint		`json:"user_seq"`
	Order_type				uint		`json:"-"`
	Moid 							string	`json:"-"`
	Tid								string	`json:"-"`
	Pname							string	`json:"-"`
	Amount						string	`json:"-"`
	Pay_method				string	`json:"-"`
	Pay_date					string	`json:"-"`
	The_date					string	`json:"-"`
	Order_status			uint		`json:"-"`
	Order_status_code	string	`json:"-"`
	Mid 							string 	`json:"-"`
	Bid 							string 	`json:"-"`
	MerchantKey 			string 	`json:"-"`
	Sign_data					string	`json:"-"`
	Auth_date					string	`json:"-"`
}

type CancelOrder struct{
	Order Order
	CancelAmt string
	CancelUserSeq uint
	PartialCancelCode string
}

type RequestCancel struct {
	TID string
	MID string
	Moid string
	CancelAmt string
	CancelMsg string
	EdiDate string
	SignData string
	CharSet string
	PartialCancelCode string
}

type RequestApi struct {
	FormData map[string]string
	Url 		 string
}

func (ord *Order) GetAssginDataAndMoid(isNormal bool) {
	md5 := md5.Sum([]byte(common.ToStr(ord.User_seq)))
	ord.Moid = "ipsap" + hex.EncodeToString(md5[:4]) + ord.Pay_date
	signDataStr := ""
	if isNormal{
		signDataStr = ord.Pay_date + ord.Mid + ord.Amount + ord.MerchantKey
	} else {
		// signDataStr = common.Config.Payment.REGULAR_MID + ediDate + moid + amt + bid + common.Config.Payment.REGULAR_MERCHANTKEY
		signDataStr = ord.Mid + ord.Pay_date + ord.Moid + common.ToStr(ord.Amount) + ord.Bid + ord.MerchantKey
	}
	data := sha256.Sum256([]byte(signDataStr))
	var newByte []byte = data[:]
	ord.Sign_data = hex.EncodeToString(newByte)
	return
}

func (ord *Order) InsertOrder(tx *sql.Tx) (succ bool) {
	succ = false
	sql := fmt.Sprintf(`
				  INSERT INTO t_orders (
						moid, tid, institution_seq, product_seq,
						order_type, pname, amount, user_seq,
						pay_method, pay_date
					)
					VALUES
						('%v','%v',%v,%v,
						 '%v','%v','%v',%v,
						 '%v','%v')`,
				 	 ord.Moid, ord.Tid, ord.Institution_seq, ord.Product_seq,
	 				 ord.Order_type, ord.Pname, ord.Amount, ord.User_seq,
	 				 ord.Pay_method, ord.Pay_date)
	var err error
	if nil != tx {
		_, err = tx.Exec(sql)
	} else {
		_, err = common.DBconn().Exec(sql)
	}

  if nil != err {
  	log.Println(err)
  	return
  }

	succ = true
	return
}

func (ord *Order) LoadList() (list interface{}) {
	moreCondition := ""
	if ord.Institution_seq != 0 {
		moreCondition = fmt.Sprintf(` AND ord.institution_seq = %v`, ord.Institution_seq)
	}

	sql := fmt.Sprintf(`
						SELECT ord.order_seq, ord.institution_seq, ord.product_seq,
									 ord.tid, ord.pname, ord.amount, ord.pay_method,
									 ord.order_type, ord.auth_date,
									 instt.institution_code, instt.name_ko as institution_name_ko,
									 user.name AS user_name, user.phoneno,
							 	 	 comm.auth_code, ord.the_date,
									 IF(ord.order_type NOT IN (1, 2),
									 		IF((SELECT COUNT(*)
									 					FROM t_orders ord2
									 				 WHERE ord2.order_type = 4
									 					 AND ord2.moid = ord.moid
									 					 AND ord.auth_date <= ord2.auth_date) <> 0,
									 			CONCAT(ord.moid, "-",(SELECT COUNT(*)
									 															FROM t_orders ord2
									 														 WHERE ord2.order_type = 4
									 													 		 AND ord2.moid = ord.moid
									 														 	 AND ord.auth_date >= ord2.auth_date)),
									 			ord.moid
									 			), ord.moid) AS moid
							FROM t_orders ord, t_institution instt, t_user user, t_resp_pg_pay_api_comm comm
						 WHERE ord.institution_seq = instt.institution_seq
						 	 AND ord.user_seq = user.user_seq
							 AND comm.moid = ord.moid
							 AND comm.tid = ord.tid
						 	 AND ord.order_status = 1
						 	  %v #moreCondition
					ORDER BY auth_date DESC`, moreCondition)
	list = common.DB_fetch_all(sql, nil)
	return
}

func (ord *Order) Load() (info []map[string]interface{}) {
	sql := fmt.Sprintf(`
		SELECT (SELECT plan.desc_text
							FROM t_products prod, t_membership_plan plan
						 WHERE prod.plan_seq = plan.plan_seq
						   AND prod.product_seq = ord.product_seq) AS desc_text,
						(SELECT comm.auth_code
							 FROM t_resp_pg_pay_api_comm comm
							WHERE comm.moid = ord.moid
							  AND comm.tid = ord.tid) AS auth_code,
					  ord.order_seq, ord.institution_seq, ord.product_seq,
					  ord.tid, ord.pname, ord.amount, ord.pay_method,
						ord.order_type, ord.auth_date, ord.pay_date, ord.the_date,
						IF(ord.order_type NOT IN (1, 2),
							 IF((SELECT COUNT(*)
										 FROM t_orders ord2
										WHERE ord2.order_type = 4
											AND ord2.moid = ord.moid
											AND ord.auth_date >= ord2.auth_date) <> 0,
								 CONCAT(ord.moid, "-",(SELECT COUNT(*)
																				 FROM t_orders ord2
																				WHERE ord2.order_type = 4
																					AND ord2.moid = ord.moid
																					AND ord.auth_date >= ord2.auth_date)),
								 ord.moid
								 ), ord.moid) AS moid,
								 ord.moid AS org_moid
			FROM t_orders ord
		 WHERE (SELECT tid FROM t_orders WHERE order_seq = ?) = ord.tid
		   AND ord.order_status = 1
	ORDER BY auth_date ASC`)
	filter := func (row map[string]interface{}) {
		if common.ToInt(row["order_type"]) == DEF_ORDER_TYPE_PARTIAL_CANCEL {
			t, _ := goment.New(common.ToStr(row["auth_date"]), "YYMMDDHHmmss")
			cancelDate := t.Format("YYYYMMDD")
			cancelTime := t.Format("HHmmss")
			sql2 := `SELECT tid
								 FROM t_resp_pg_billing_cancel
								WHERE moid = ?
									AND result_code IN (2001, 2211)
								  AND cancel_date = ?
									AND cancel_time = ?`
			row2 := common.DB_fetch_one(sql2, nil, row["org_moid"], cancelDate, cancelTime)
			if "" != common.ToStr(row2["tid"]) {
				row["tid"] = row2["tid"]
			}
		}
		delete(row, "org_moid")
	}
	info = common.DB_fetch_all(sql, filter, ord.Order_seq)
	return
}

func (ord *Order) InsertCancelOrder() (succ bool) {
	succ = false
	sql := fmt.Sprintf(`
					INSERT INTO t_orders(
						institution_seq, product_seq,
						user_seq, moid, tid, pname,
						amount, order_type, pay_method,
						pay_date, auth_date, the_date,
						order_status, order_status_code)
					VALUES(
						?,?,
						?,?,?,?,
						?,?,?,
						?,?,?,
						?,?)`)
	_, err := common.DBconn().Exec(sql,
																 ord.Institution_seq, ord.Product_seq,
																 ord.User_seq, ord.Moid, ord.Tid, ord.Pname,
																 ord.Amount, ord.Order_type, ord.Pay_method,
																 ord.Pay_date, ord.Auth_date, ord.The_date,
																 ord.Order_status, ord.Order_status_code)
 	if nil != err {
 		log.Println(err)
 		return
 	}

	return true
}

func (cancOrd *CancelOrder) CancelOrder() (succ bool) {
	succ = false
	orderType := 0
	sql := `SELECT institution_seq, product_seq,
								 user_seq, moid, tid, pname, amount,
								 order_type, pay_method, the_date
					  FROM t_orders
					 WHERE order_seq = ?`
	row := common.DB_fetch_one(sql, nil, cancOrd.Order.Order_seq)
	if "0" == cancOrd.PartialCancelCode { // 전체취소
		if cancOrd.CancelAmt != common.ToStr(row["amount"]) {
			return
		}
		orderType = DEF_ORDER_TYPE_ALL_CANCEL
	} else if "1" == cancOrd.PartialCancelCode { // 부분취소
		if common.ToUint(cancOrd.CancelAmt) > common.ToUint(row["amount"]) {
			return
		}
		orderType = DEF_ORDER_TYPE_PARTIAL_CANCEL
	}

	merchantKey := ""
	mid := ""
	paymentOrderType := common.ToUint(row["order_type"])
	if 1 == paymentOrderType { // 일반 결제일때
		merchantKey = common.Config.Payment.GENERAL_MERCHANTKEY
		mid = common.Config.Payment.GENERAL_MID
	} else if 2 == paymentOrderType { // 정기 결제일때
		merchantKey = common.Config.Payment.REGULAR_MERCHANTKEY
		mid = common.Config.Payment.REGULAR_MID
	} else {
		return
	}

	now, _ := goment.New()
	ediDate := now.Format("YYYYMMDDHHmmss")
	signDataStr := mid + cancOrd.CancelAmt + ediDate + merchantKey
	data := sha256.Sum256([]byte(signDataStr))
	var newByte []byte = data[:]
	encSignData := hex.EncodeToString(newByte)

	formData := make(map[string]string)
	formData[DEF_PARAM_TID] = common.ToStr(row["tid"])
	formData[DEF_PARAM_MID] = mid
	formData[DEF_PARAM_Moid] = common.ToStr(row["moid"])
	formData[DEF_PARAM_CancelAmt] = cancOrd.CancelAmt
	formData[DEF_PARAM_CancelMsg] = "IPSAP 결제 취소"
	formData[DEF_PARAM_EdiDate] = ediDate
	formData[DEF_PARAM_SignData] = encSignData
	formData[DEF_PARAM_CharSet] = "utf-8"
	formData[DEF_PARAM_PartialCancelCode] = cancOrd.PartialCancelCode

	paymentCancelReq := PaymentCancelReq{}
	if err := mapstructure.Decode(formData, &paymentCancelReq); err != nil {
		log.Println(err)
		return
	}
	paymentCancelReq.InsertRequest()

	reqApi := RequestApi {
		FormData : formData,
		Url : DEF_URL_CANCEL,
	}

	resp, err := reqApi.Send()
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

	payCancelResp := PaymentCancelResp{}
	if err = mapstructure.Decode(jsonData, &payCancelResp); nil != err {
		log.Println(err)
		return
	}

	payCancelResp.InsertResponse()

	authDate, _ := goment.New(payCancelResp.CancelDate + payCancelResp.CancelTime,"YYYYMMDDHHmmss")

	formData[DEF_PARAM_TID] = common.ToStr(row["tid"])
	formData[DEF_PARAM_MID] = mid
	formData[DEF_PARAM_Moid] = common.ToStr(row["moid"])
	formData[DEF_PARAM_CancelAmt] = cancOrd.CancelAmt
	formData[DEF_PARAM_CancelMsg] = "IPSAP 결제 취소"
	formData[DEF_PARAM_EdiDate] = ediDate
	formData[DEF_PARAM_SignData] = encSignData
	formData[DEF_PARAM_CharSet] = "utf-8"
	formData[DEF_PARAM_PartialCancelCode] = cancOrd.PartialCancelCode

	ord := Order {
		Moid : formData[DEF_PARAM_Moid],
		Tid : formData[DEF_PARAM_TID],
		Amount : formData[DEF_PARAM_CancelAmt],
		Order_type : uint(orderType),
		The_date : common.ToStr(row["the_date"]),
		Pay_method : common.ToStr(row["pay_method"]),
		Pay_date : formData[DEF_PARAM_EdiDate],
		Auth_date : authDate.Format("YYMMDDHHmmss"),
		Order_status_code : payCancelResp.ResultCode,
		User_seq : cancOrd.CancelUserSeq,
		Product_seq : common.ToUint(row["product_seq"]),
		Institution_seq : common.ToUint(row["institution_seq"]),
		Pname : common.ToStr(row["pname"]),
	}

	// 2001	성공코드 취소 성공 , 2211	성공코드 환불 성공
	if payCancelResp.ResultCode == "2001" || payCancelResp.ResultCode == "2211" {
		ord.Order_status = DEF_ORDER_STATUS_COMPLETED
		ord.InsertCancelOrder()
	} else {
		ord.Order_status = DEF_ORDER_STATUS_ERROR
		ord.InsertCancelOrder()
		return
	}

	succ = true
	return
}

func (reqApi *RequestApi) Send() (resp *resty.Response, err error) {
	client := resty.New()
	resp, err = client.R().
	SetHeader("Content-Type", "application/x-www-form-urlencoded").
	SetFormData(reqApi.FormData).
	Post(reqApi.Url)
	return
}
