package model

import (
	"ipsap/common"
	"database/sql"
	"log"
	// "github.com/gin-gonic/gin"
	// "strings"
	// "fmt"
)

type Billing struct {
	GoodsName					string	`form:"GoodsName" binding:"required"`
	Amt								string	`form:"Amt" binding:"required"`
	MID								string	`form:"MID" binding:"required"`
	EdiDate						string	`form:"EdiDate" binding:"required"`
	Moid							string	`form:"Moid" binding:"required"`
	SignData					string	`form:"SignData" binding:"required"`
	CharSet						string	`form:"CharSet" binding:"required"`
	PayMethod					string	`form:"PayMethod" binding:"required"`
	BuyerName					string	`form:"BuyerName" binding:"required"`
	BuyerTel					string	`form:"BuyerTel" binding:"required"`
	BuyerEmail				string	`form:"BuyerEmail" binding:"required"`
	VerifySType				string	`form:"VerifySType" binding:"required"`
	EncGoodsName			string	`form:"EncGoodsName"`
	EncBuyerName			string	`form:"EncBuyerName"`
	NpDirectYn				string	`form:"NpDirectYn" binding:"required"`
	NpDirectLayer			string	`form:"NpDirectLayer" binding:"required"`
	JsVer							string	`form:"JsVer" binding:"required"`
	NpSvcType					string	`form:"NpSvcType" binding:"required"`
	DeployedVer				string	`form:"DeployedVer" binding:"required"`
	DeployedDate			string	`form:"DeployedDate" binding:"required"`
	DeployedFileName	string	`form:"DeployedFileName" binding:"required"`
	AuthResultCode		string	`form:"AuthResultCode" binding:"required"`
	AuthResultMsg			string	`form:"AuthResultMsg" binding:"required"`
	AuthToken					string	`form:"AuthToken" binding:"required"`
	TxTid							string	`form:"TxTid" binding:"required"`
	NextAppURL				string	`form:"NextAppURL" binding:"required"`
	NetCancelURL			string	`form:"NetCancelURL" binding:"required"`
	Signature					string	`form:"Signature" binding:"required"`
	VbankExpDate			string  `form:"VbankExpDate"`
	GoodsCl						string  `form:"GoodsCl"`
	ReqReserved				string  `form:"ReqReserved"`
}

type BillingKey struct {
	Institution_seq uint
	User_seq 				uint
	ResultCode			string	`form:"resultCode"`
	ResultMsg				string	`form:"resultMsg"`
	Billkey					string	`form:"billkey"`
	GoodsName				string	`form:"GoodsName" `
	Mid							string	`form:"MID" `
	Moid						string	`form:"Moid" `
	TID							string	`form:"TID" `
	MallReserved		string	`form:"MallReserved" `
	BuyerEmail			string	`form:"BuyerEmail" `
	BuyerName				string	`form:"BuyerName" `
	BuyerTel				string	`form:"BuyerTel" `
	Amt							string	`form:"Amt" `
	PayMethod				string	`form:"PayMethod" `
	CardNo					string	`form:"CardNo" `
	CardName				string	`form:"cardName" `
	CardCode				string	`form:"cardCode" `
	ExpYY						string	`form:"expYY" `
	ExpMM						string	`form:"expMM" `
	CardCl					string	`form:"CardCl" `
}

func (bill *Billing) InsertBilling(tx *sql.Tx) (succ bool) {
	succ = false
	sql := `INSERT INTO t_req_pg_billing (
						moid, goods_name, amt, mid,
						edi_date, sign_data, buyer_name, buyer_tel,
						buyer_email, pay_method, charset, vbank_exp_date,
						goods_cl, req_reserved
					)
					VALUES
						(?,?,?,?,
						 ?,?,?,?,
						 ?,?,?,?,
						 ?,?)`
	 _, err := tx.Exec(sql,
										 bill.Moid, bill.GoodsName, bill.Amt, bill.MID,
										 bill.EdiDate, bill.SignData, bill.BuyerName, bill.BuyerTel,
										 bill.BuyerEmail, bill.PayMethod, bill.CharSet, bill.VbankExpDate,
										 bill.GoodsCl, bill.ReqReserved)
  if nil != err {
  	log.Println(err)
  	return
  }

	sql	= `INSERT INTO t_resp_pg_billing (
						moid, auth_result_code, auth_result_msg,
						auth_token, pay_method, mid,
						amt, req_reserved, tx_tid,
						next_app_url, net_cancel_url
				 )
				 VALUES
						(?,?,?,
						 ?,?,?,
						 ?,?,?,
						 ?,?)`
	 _, err = tx.Exec(sql,
										bill.Moid, bill.AuthResultCode, bill.AuthResultMsg,
										bill.AuthToken, bill.PayMethod, bill.MID,
										bill.Amt, bill.ReqReserved, bill.TxTid,
										bill.NextAppURL, bill.NetCancelURL)
  if nil != err {
  	log.Println(err)
  	return
  }

  succ = true
	return
}

func (billKey *BillingKey) InsertBillingKey() (succ bool) {
	succ = false
	dbConn	:= common.DBconn()
	tx, err	:= dbConn.Begin()
	if nil != err {
		tx.Rollback()
		return
	}

	sql := `INSERT INTO t_req_pg_bill_key (
						institution_seq, user_seq, mid,
						moid, goods_name, amt, buyer_name,
						buyer_tel, buyer_email, pay_method, mall_reserved
					)
					VALUES
						(?,?,?,
						 ?,?,?,?,
						 ?,?,?,?)`
	 _, err = tx.Exec(sql,
									  billKey.Institution_seq, billKey.User_seq, billKey.Mid,
									  billKey.Moid, billKey.GoodsName, billKey.Amt, billKey.BuyerName,
									  billKey.BuyerTel, billKey.BuyerEmail, billKey.PayMethod, billKey.MallReserved)
	if nil != err {
		log.Println(err)
		return
	}

	sql  = `INSERT INTO t_resp_pg_bill_key (
						institution_seq, user_seq, result_code,
						result_msg, bid, goods_name,
						moid, tid, mall_reserved, buyer_email,
						amt, pay_method, card_no, card_name,
						card_code, exp_yy, exp_mm, card_cl
					)
					VALUES
						(?,?,?,
						 ?,?,?,
						 ?,?,?,?,
						 ?,?,?,?,
						 ?,?,?,?)`
	 _, err = tx.Exec(sql,
										billKey.Institution_seq, billKey.User_seq, billKey.ResultCode,
										billKey.ResultMsg, billKey.Billkey, billKey.GoodsName,
										billKey.Moid, billKey.TID, billKey.MallReserved, billKey.BuyerEmail,
										billKey.Amt, billKey.PayMethod, billKey.CardNo, billKey.CardName,
										billKey.CardCode, billKey.ExpYY, billKey.ExpMM, billKey.CardCl)
	if nil != err {
		log.Println(err)
		return
	}

	err = tx.Commit()
	if nil != err {
		log.Println(err)
		return
	}

	// F100 : 성공 / 그외 실패
	if "F100" != billKey.ResultCode {
		return
	}

	sql = `UPDATE t_institution
	 					SET bid = ?, payment_setting = 2
					WHERE institution_seq = ?`
	_, err = common.DBconn().Exec(sql, billKey.Billkey, billKey.Institution_seq)
  if err != nil {
    log.Println(err)
		return
  }

	succ = true
	return
}
