package api

import (
	// "github.com/mitchellh/mapstructure"
	"github.com/gin-gonic/gin"
	// "encoding/json"

	"ipsap/common"
	"ipsap/model"
	"log"
	// "fmt"
)

// @Tags Billing
// @Summary 빌링키 창 요청
// @Description 빌링키 창 요청
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Router  /billing-key/assign [get]
// @Success 200
func BillingKeyAssgin(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	// regular-mid = "nictest04m"  # 정기결제용 MID
	// general-mid = "nicepay00m"  # 알반결제용 MID

	common.FinishApi(c, common.Api_status_ok, gin.H{"mid": common.Config.Payment.REGULAR_MID,
		"institution_seq": tokenMap["institution_seq"],
		"user_seq":        tokenMap["user_seq"],
	})
}

// @Tags Billing
// @Summary 빌링키 발급
// @Description 빌링키 발급
// @Accept  json
// @Produce  json
// @Router /billing-key [post]
// @Success 200
func BillingKeyCreate(c *gin.Context) {

	bid := model.BillingKey{}
	if err := c.ShouldBind(&bid); err != nil {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, common.ToStr(err))
		return
	}
	/*
		out1, _ := iconv.ConvertString(string(bid.ResultMsg), "euc-kr", "utf-8")
		out2, _ := iconv.ConvertString(string(bid.BuyerEmail), "euc-kr", "utf-8")
		out3, _ := iconv.ConvertString(string(bid.MallReserved), "euc-kr", "utf-8")
		out4, _ := iconv.ConvertString(string(bid.GoodsName), "euc-kr", "utf-8")
		out5, _ := iconv.ConvertString(string(bid.CardName), "euc-kr", "utf-8")

		data := strings.Split(bid.MallReserved, ",")
		if len(data) != 2 {
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_billing_key_register)
			return
		}

		bid.Institution_seq = common.ToUint(data[0])
		bid.User_seq = common.ToUint(data[1])
		if !bid.InsertBillingKey() {
			log.Println(bid.ResultMsg)
			common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_billing_key_register)
			return
		}

		common.FinishApi(c, common.Api_status_ok, gin.H{"result": bid.ResultMsg})
	*/
}
