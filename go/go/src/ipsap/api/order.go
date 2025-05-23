package api

import (
	// "github.com/mitchellh/mapstructure"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"ipsap/common"
	"ipsap/model"
	"log"
	// "fmt"
)

// @Tags Orders
// @Summary 상품 결제창 요청
// @Description 상품 결제창 요청
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param param body model.orderAssginModel true "order assign"
// @Router /orders/assign [post]
// @Success 200
func OrderAssgin(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	chkKeys := []interface{} {
		[]string{"product_seq", "goods_name", "amt", "edi_date"},
		[][]string{},
		[]string{},
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	product := model.Product{}
	product.Product_seq = common.ToUint(data["product_seq"])
	product.Name = common.ToStr(data["goods_name"])
	product.Discounted_amount = common.ToUint(data["amt"])
	if !product.CheckProduct() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_invalid_product)
		return
	}

	order := model.Order{}
	order.Pay_date = common.ToStr(data["edi_date"])
	order.Amount = common.ToStr(data["amt"])
	order.User_seq = common.ToUint(tokenMap["user_seq"])
	order.Mid = common.Config.Payment.GENERAL_MID
	order.MerchantKey = common.Config.Payment.GENERAL_MERCHANTKEY
	order.GetAssginDataAndMoid(true) // 일반 결제

	result := gin.H{ "goods_name": product.Name,
									 "amt": product.Discounted_amount,
									 "mid": common.Config.Payment.GENERAL_MID,
									 "edi_date": order.Pay_date,
									 "moid": order.Moid,
									 "sign_data": order.Sign_data,}

	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Orders
// @Summary 상품 결제
// @Description 상품 결제
// @Accept  json
// @Produce  json
// @Router /orders [post]
// @Success 200
func OrderCreate(c *gin.Context) {
	billing := model.Billing{}
	if err := c.ShouldBind(&billing); err != nil {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, common.ToStr(err))
		return
	}

	order := model.Order {}
	err := json.Unmarshal([]byte(billing.ReqReserved), &order)
	if err != nil {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	dbConn	:= common.DBconn()
	tx, err	:= dbConn.Begin()
	if nil != err {
		tx.Rollback()
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	if !billing.InsertBilling(tx) {
		tx.Rollback()
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	order.Moid = billing.Moid
	order.Tid = billing.TxTid
	order.Pname = billing.GoodsName
	order.Amount = billing.Amt
	order.Pay_method = billing.PayMethod
	order.Pay_date = billing.EdiDate
	order.Order_status_code = billing.AuthResultCode
	order.Order_type = model.DEF_ORDER_TYPE_NORMAL_PAYMENT

  if !order.InsertOrder(tx) {
		tx.Rollback()
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	err = tx.Commit()
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	// 인증 성공일시
	if billing.AuthResultCode == "0000" {
		reqPay := model.ReqPay{
			Institution_seq : order.Institution_seq,
			Product_seq : order.Product_seq,
			User_seq : order.User_seq,
			Tid : billing.TxTid,
			AuthToken : billing.AuthToken,
      MID : billing.MID,
      Amt : billing.Amt,
      CharSet : "utf-8",
			NextAppURL : billing.NextAppURL,
			NetCancelURL :  billing.NetCancelURL,
		}

		resultMsg, succ := reqPay.ApprovePay()
		if !succ {
			common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_pay, resultMsg)
			return
		}
		common.FinishApi(c, common.Api_status_ok, gin.H{ "result": resultMsg,})
	} else {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_pay_auth, billing.AuthResultMsg)
		return
	}
}

// @Tags Orders
// @Summary 결제 이력
// @Description 결제 이력
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param filter.institution_seq query uint false	"기관 seq"
// @Router /orders [get]
// @Success 200
func OrderList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	institution_seq := c.Query("filter.institution_seq")
	ord := model.Order {
		Institution_seq : common.ToUint(institution_seq),
	}

	list := ord.LoadList()
	result := gin.H{ "value" : list,}
	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Orders
// @Summary 결제 상세
// @Description 결제 상세
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param order_seq path uint true "1"
// @Router /orders/{order_seq} [get]
// @Success 200
func OrderInfo(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	order_seq, succ := getOrderSeqFromPath(c)
	if !succ {  return  }

	ord := model.Order{
		Order_seq : common.ToUint(order_seq),
	}

	info := ord.Load()
	common.FinishApi(c, common.Api_status_ok, info)
}

// @Tags Orders
// @Summary 결제 취소 (플랫폼 관리자)
// @Description 결제 취소 (플랫폼 관리자) 기관에 상태 변경 적용 안됨
// @Description partial_cancel_code 0:전체 취소, 1:부분 취소
// @Accept  json
// @Produce  json
// @Param param body model.orderCancelModel true "order cancel"
// @Param order_seq path uint true "1"
// @Router /orders/{order_seq} [delete]
// @Success 200
func OrderCancel(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	chkKeys := []interface{} {
		[]string{"cancel_amt", "partial_cancel_code"},
		[][]string{},
		[]string{},
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	order_seq, succ := getOrderSeqFromPath(c)
	if !succ {  return  }

	cancOrd := model.CancelOrder {
		Order : model.Order{
			Order_seq : common.ToUint(order_seq),
		},
		CancelAmt : common.ToStr(data["cancel_amt"]),
		PartialCancelCode : common.ToStr(data["partial_cancel_code"]),
		CancelUserSeq : common.ToUint(tokenMap["user_seq"]),
	}

	if !cancOrd.CancelOrder() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_cancel_order_fail)
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}
