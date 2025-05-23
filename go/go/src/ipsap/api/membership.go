package api

import (
	"github.com/mitchellh/mapstructure"
	"github.com/gin-gonic/gin"
	"ipsap/common"
	"ipsap/model"
	"strings"
	"log"
)

// @Tags Membership Plan
// @Summary Membership Plan 등록
// @Description Membership Plan 등록
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.Plan true "plan"
// @Router /membership/plan [post]
// @Success 201
func PlanCreate(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}


	chkKeys := []interface{}{
		[]string{ 	"name", "desc_text",
								"plan_pid", "plan_category",
								"plan_price", "plan_discount_rate", "plan_discount_type",
								"plan_discount_amount", "plan_discounted_amount",
								"membership_pid", "membership_category",
								"membership_price",	"membership_discount_rate",	"membership_discount_type",
								"membership_discount_amount",	 "membership_discounted_amount",
								"usage_limit",
								"plan_available"},	//	필수 키
		[][]string{},                      							//	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},	//  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	plan := model.Plan{}
	if err := mapstructure.Decode(data, &plan); nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	plan.Reg_user_seq = common.ToUint(tokenMap["user_seq"])
	if !plan.InsertPlan() {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Membership Plan
// @Summary Membership Plan 리스트
// @Description Membership Plan  리스트
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Router /membership/plan [get]
// @Success 200
func PlanList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	plan := model.Plan{}
	list := plan.Load()
	result := gin.H{ "value" : list, "@deltaLink" : "{nice pg Url}",}
	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Membership Plan
// @Summary Membership Plan 정보
// @Description Membership Plan 정보
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param plan_seq path uint true "1"
// @Router /membership/plan/{plan_seq} [get]
// @Success 200
func PlanInfo(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	plan_seq, succ := getPlanFromPath(c)
	if !succ {  return  }

	plan := model.Plan{}
	plan.Plan_seq = common.ToUint(plan_seq)
	list := plan.Load()
	result := gin.H{ "value" : list, "@deltaLink" : "{nice pg Url}",}
	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Membership Plan
// @Summary Membership Plan 수정
// @Description Membership Plan 수정
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param plan_seq path uint true "1"
// @Param param body model.Plan true "Membership plan"
// @Router  /membership/plan/{plan_seq} [patch]
// @Success 200
func PlanPatch(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	plan_seq, succ := getPlanFromPath(c)
	if !succ {  return  }

	chkKeys := []interface{}{
		[]string{
			"name", "desc_text", "usage_limit",
			"plan_product_seq",
			"plan_price", "plan_discount_rate", "plan_discount_type",
			"plan_discount_amount", "plan_discounted_amount",
			"membership_product_seq",
			"membership_price",	"membership_discount_rate",	"membership_discount_type",
			"membership_discount_amount",	 "membership_discounted_amount",
			"plan_available" },	//	필수 키
		[][]string{},                      							//	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},	//  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	plan := model.Plan{}
	if err := mapstructure.Decode(data, &plan); nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	plan.Plan_seq = common.ToUint(plan_seq)
	if !plan.UpdatePlan() {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Membership Plan
// @Summary Membership Plan 삭제
// @Description Membership Plan 삭제
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param plan_seq path uint true "1"
// @Router /membership/plan/{plan_seq} [delete]
// @Success 200
func PlanDelete(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	plan_seq, succ := getPlanFromPath(c)
	if !succ {  return  }

	plan := model.Plan{}
	plan.Plan_seq = common.ToUint(plan_seq)
	if !plan.DeletePlan() {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Membership Free
// @Summary 무료 멤버쉽 지급
// @Description 무료 멤버쉽 지급
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param param body model.freeMembershipRegisterModel true "free membership"
// @Router /membership/free [post]
// @Success 201
func MembershipFreeCreate(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	chkKeys := []interface{} {
		[]string{"institution_seqs", "usage_limit", "free_period"},	//	필수 키
		[][]string{},                																//	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{"reason"},																					//	그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	for _, institution_seq := range data["institution_seqs"].([]interface{}) {
		free := model.FreeMembership {
			Institution_seq	: common.ToUint(institution_seq),
			Usage_limit	: common.ToUint(data["usage_limit"]),
			Free_period	: common.ToStr(data["free_period"]),
			Reason	: common.ToStr(data["reason"]),
		}
		if !free.InsertFreeMembership() {
			common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
			return
		}
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Membership Free
// @Summary 무료 멤버쉽 지급 목록
// @Description 무료 멤버쉽 지급 목록
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param filter.start_index query string false	"page start index"
// @Param filter.row_cnt query string false     "row count"
// @Param filter.use_status query string false  "이용여부 (0: 미이용, 1: 이용완료)"
// @Param filter.search_word query string false	"검색어"
// @Router /membership/free [get]
// @Success 200
func MembershipFreeList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	startIdx      := c.Request.URL.Query().Get("filter.start_index")
	rowCnt        := c.Request.URL.Query().Get("filter.row_cnt")
	useStatus     := c.Request.URL.Query().Get("filter.use_status")
	searchWord    := c.Request.URL.Query().Get("filter.search_word")

	free := model.FreeMembership{
		StartIdx : startIdx,
		RowCnt 	 : rowCnt,
		UseStatus : useStatus,
		SearchWord : searchWord,
	}

	ret := free.LoadList()
	common.FinishApi(c, common.Api_status_ok, ret)
}

// @Tags Membership Free
// @Summary 무료 멤버십 지급 회수
// @Description 무료 멤버십 지급 회수
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Router /membership/free/{free_seq} [delete]
// @Success 200
func MembershipFreeDelete(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	free_seq, succ := getMembershipFreeSeqFromPath(c)
	if !succ {  return  }

	free := model.FreeMembership {
		Free_seq : common.ToUint(free_seq),
	}

	if !free.DeleteFreeMembership() {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Membership
// @Summary 멤버쉽 변경
// @Description 멤버쉽 변경
// @Param institution_seq path string true "institution_seq"
// @Param product_seq path string true "product_seq"
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Router /institution/{institution_seq}/product/{product_seq} [patch]
// @Success 200
func InstitutionPlanChange(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	institution_seq, succ1 := getInstitutionIdFromPath(c)
	if !succ1 { return }

	product_seq, succ2 := getProductFromPath(c, "plan")
	if !succ2 { return }

	instt := model.Institution {
		Institution_seq : common.ToUint(institution_seq),
		Product_seq : common.ToUint(product_seq),
	}

  if !instt.PlanChange(common.ToUint(tokenMap["user_seq"])) {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Membership
// @Summary 멤버쉽 결제 방법 변경(수동, 자동)
// @Description 수동 -> 자동, 자동 -> 수동 변경
// @Description bill key를 발급 받은적이 없으면 bill key를 새로 받아야함
// @Param institution_seq path string true "institution_seq"
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Router /institution/{institution_seq}/payment-setting [patch]
// @Success 200
func InstitutionPaymentSettingChange(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	institution_seq, succ := getInstitutionIdFromPath(c)
	if !succ {
		return
	}

	instt := model.Institution {
		Institution_seq : common.ToUint(institution_seq),
	}

  if !instt.PaymentSettingChange() {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Membership
// @Summary 이용중인 멤버십 정보 및 결제 정보
// @Description 이용중인 멤버십 정보 및 결제 정보
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Router /membership [get]
// @Success 200
func MembershipInUseAndPaymentInfo(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	membership := model.Membership {
		Institution_seq : common.ToUint(tokenMap["institution_seq"]),
	}

  info := membership.GetInstitutionMembershipInfo()
	result := gin.H{ "value" : info,}
	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Membership Cancel
// @Summary 멤버십 해지 및 환불 정보 요청
// @Description 멤버십 해지 및 환불(기관 행정간사)
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Router /membership/cancel [get]
// @Success 200
func MembershipCancelInfo(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	mCancel := model.MembershipCancel {
		Institution_seq : common.ToUint(tokenMap["institution_seq"]),
	}

  mCancel.GetMembershipCancelInfo()

	result := gin.H {
		"membership_cancel_amt" : mCancel.MembershipAmt,
		"plan_cancel_amt" : mCancel.PlanAmt,
		"total_amt" : mCancel.PlanAmt + mCancel.MembershipAmt,
	}

	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Membership Cancel
// @Summary 멤버십 해지 및 환불
// @Description 멤버십 해지 및 환불(기관 행정간사)
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Router /membership [delete]
// @Success 200
func MembershipCancel(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	mCancel := model.MembershipCancel {
		Institution_seq : common.ToUint(tokenMap["institution_seq"]),
		CancelUserSeq : common.ToUint(tokenMap["user_seq"]),
	}

	succ := mCancel.MembershipCancel()
	if !succ{
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_payment_fail)
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Membership
// @Summary 무료 멤버십 지급 가능한 리스트
// @Description 무료 멤버십 지급 가능한 리스트(가입비를 지불한 상태여야됨)
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Router /membership-free/institution [get]
// @Success 200
func InstitutionFreeMembershipPossibleList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	instt := model.Institution{}
	list := instt.GetFreeMembershipPossibleList()
	common.FinishApi(c, common.Api_status_ok, list)
}
