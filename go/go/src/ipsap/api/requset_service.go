package api

import (
	"github.com/mitchellh/mapstructure"
	"github.com/gin-gonic/gin"
	"ipsap/common"
	"ipsap/model"
	// "strings"
	"log"
)

// @Tags RequsetService
// @Summary 신청서 관리 목록
// @Description 신청서 관리 목록
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /request/service [get]
// @Success 200
func RequestServiceList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	if common.ToUint(tokenMap["user_auth"]) <= model.DEF_USER_AUTH_INSTITUTION {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	reqSvc := model.RequestService{}
	rows := reqSvc.LoadList()

	common.FinishApi(c, common.Api_status_ok, rows)
}

// @Tags RequsetService
// @Summary 온라인 심의 시스템 서비스 신청서 등록
// @Description 온라인 심의 시스템 서비스 신청서 등록
// @Accept  mpfd
// @Produce  mpfd
// @Param logo_file formData file false "기관 로고 파일"
// @Param business_file formData file false "사업자 등록증 파일"
// @Param request_service formData string true "서비스 신청(Json String 형태)"
// @Param test body model.requestServiceModel false "test용 Json Data 실제사용 X"
// @Router /request/service [post]
// @Success 200
func RequestServiceCreate(c *gin.Context) {

	data, eMsg := common.UnmarshalFormData(c, "request_service")
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	user 						:= model.User{}
	institution 		:= model.Institution{}
	userMap 				:= data["user"].(map[string]interface{})
	institutionMap 	:= data["institution"].(map[string]interface{})

	if err1 := mapstructure.Decode(userMap, &user); nil != err1 {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	if err2 := mapstructure.Decode(institutionMap, &institution); nil != err2 {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_system_unknown)
		return
	}

	dbConn	:= common.DBconn()
	tx, err	:= dbConn.Begin()
	if nil != err {
		log.Println(err)
		tx.Rollback()
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	if !institution.InsertInstitution(c, tx) {
		tx.Rollback()
		return
	}

	user.Institution_seq			= institution.Institution_seq
	user.Agree_terms_service	= 1 // 약관동의
	user.Agree_terms_privacy 	= 1 // 개인정보처리방침 동의
	user.User_status 					= model.DEF_USER_STATUS_WAIT // 유저 등록대기
	user.User_type						= common.ToStr(model.DEF_USER_TYPE_ADMIN_SECRETARY)
	user.User_auth						= model.DEF_USER_AUTH_INSTITUTION
	if !user.InsertUser(c, tx) {
		tx.Rollback()
		return
	}

	reqSvc := model.RequestService{}
	reqSvc.Institution_seq = institution.Institution_seq
	reqSvc.User_seq				 = user.User_seq
	if !reqSvc.InsertRequestService(c, tx) {
		tx.Rollback()
		return
	}

	institution.Logo_file_idx			=	"1"
	institution.Business_file_idx	=	"1"
	if !institution.FileUpload(c, tx) {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		tx.Rollback()
		return
	}

	common.FinishApi(c, common.Api_status_ok,
    gin.H{
      "rt": "ok",
    })
}

// @Tags RequsetService
// @Summary 신청서 상세 정보
// @Description 신청서 상세 정보
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param reqsvc_seq path uint true "1"
// @Router /request/service/{reqsvc_seq} [get]
// @Success 200
func RequestServiceInfo(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	reqsvc_seq, succ, _, request_type, institution_seq, _ := getRequestServiceSeqFromPath(c)
	if !succ {  return  }

	if common.ToUint(tokenMap["user_auth"]) <= model.DEF_USER_AUTH_INSTITUTION {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	reqSvc := model.RequestService{}
	reqSvc.LoginToken = tokenMap
	reqSvc.Reqsvc_seq = common.ToUint(reqsvc_seq)
	reqSvc.Request_type = request_type
	reqSvc.Institution_seq = institution_seq
	if !reqSvc.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	common.FinishApi(c, common.Api_status_ok, reqSvc.Data)
}

// @Tags RequsetService
// @Summary 신청서 수정
// @Description 신청서 수정(플랫폼 관리자만 수정 가능하다)
// @Description request_type이 1 인경우 (서비스 등록) 결제정보 내용 뺴고 수정 가능하다
// @Description request_type이 2 인경우 (결제정보 변경) 결제정보 만 수정 가능하다
// @Accept  mpfd
// @Produce  mpfd
// @Security ApiKeyAuth
// @Param logo_file formData file false "기관 로고 파일"
// @Param business_file formData file false "사업자 등록증 파일"
// @Param reqsvc_seq path uint true "1"
// @Param param formData string true "기관정보 수정(Json String 형태)"
// @Param test1 body model.requestServiceModel false "test용 Json Data 실제사용 X 서비스 등록 내용 수정시 Sample"
// @Param test2 body model.institutionPaymentPatchModel false "test용 Json Data 실제사용 X 결제 내용 변경시 Sample"
// @Router /request/service/{reqsvc_seq} [patch]
// @Success 200
func RequestServicePatch(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	reqsvc_seq, succ, _, request_type, institution_seq, user_seq := getRequestServiceSeqFromPath(c)
	if !succ {  return  }

	if common.ToUint(tokenMap["user_auth"]) <= model.DEF_USER_AUTH_INSTITUTION {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	data, eMsg := common.UnmarshalFormData(c, "param")
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	dbConn	:= common.DBconn()
	tx, err	:= dbConn.Begin()
	if nil != err {
		log.Println(err)
		tx.Rollback()
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	switch request_type {
		case model.DEF_REQUEST_TYPE_SERVICE_REGIST:
			user 						:= model.User{}
			ins 						:= model.Institution{}
			userMap 				:= data["user"].(map[string]interface{})
			institutionMap 	:= data["institution"].(map[string]interface{})
			if err1 := mapstructure.Decode(userMap, &user); nil != err1 {
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
				return
			}
			if err2 := mapstructure.Decode(institutionMap, &ins); nil != err2 {
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
				return
			}

			ins.Institution_seq	= institution_seq
			if !ins.UpdateInstitution(c, tx) {
				tx.Rollback()
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
				return
			}

			user.User_seq = user_seq
			user.Institution_seq = institution_seq
			if !user.AdminUpdateUser(c, tx) {
				tx.Rollback()
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
				return
			}
		case model.DEF_REQUEST_TYPE_PAYMENT_CHANGE:
			ins := model.Institution{}
			if err := mapstructure.Decode(data, &ins); nil != err {
				log.Println(err)
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
				return
			}
			ins.Institution_seq	= institution_seq
			if !ins.UpdateRequestPaymentChange(c, tx, reqsvc_seq) {
				tx.Rollback()
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
				return
			}
		default :
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
			return
	}

	err = tx.Commit()
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		tx.Rollback()
		return
	}
	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok",})
}

// @Tags RequsetService
// @Summary 신청서 처리
// @Description 신청서 승인, 보류, 안내 재발송
// @Description 승인 handle_type 1
// @Description 보류 handle_type 2
// @Description 재발송 handle_type 3
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param filter.handle_type query string true "1"
// @Param reqsvc_seq path uint true "1"
// @Param param body model.approvedCommentModel false "승인 및 보류시 comment가 있을때"
// @Router /request/service/{reqsvc_seq}/handle [patch]
// @Success 200
func RequestServiceHandle(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	reqsvc_seq, succ, request_status, request_type, institution_seq, user_seq := getRequestServiceSeqFromPath(c)
	if !succ {  return  }

	if common.ToUint(tokenMap["user_auth"]) <= model.DEF_USER_AUTH_INSTITUTION {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	handle_type := common.ToUint(c.Request.URL.Query().Get("filter.handle_type"))

	chkKeys := []interface{}{
		[]string{},										//	필수 키
		[][]string{},                 //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{"approved_comment"},	//	그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	reqSvc := model.RequestService{}
	reqSvc.LoginToken = tokenMap
	reqSvc.Reqsvc_seq = common.ToUint(reqsvc_seq)
	reqSvc.Request_status = request_status
	reqSvc.Request_type = request_type
	reqSvc.Handle_type 	= handle_type
	reqSvc.Institution_seq = institution_seq
	reqSvc.User_seq = user_seq
	if nil != data["approved_comment"] {
		reqSvc.Approved_comment = common.ToStr(data["approved_comment"])
	}

	if !reqSvc.UpdateRequestService(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok",})
}
