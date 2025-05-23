package api

import (
	"fmt"
	"ipsap/common"
	"ipsap/model"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// @Tags Admin(행정간사)
// @Summary 사용자 관리 목록
// @Description 사용자 관리 목록
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /admin/user [get]
// @Success 200
func AdminUserList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}
	/*
		if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
			common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
			return
		}
	*/
	moreCondition := fmt.Sprintf(` AND user.institution_seq = %d`, common.ToUint(tokenMap["institution_seq"]))
	user := model.User{}
	rows := user.GetInstitutionUserList(moreCondition)
	common.FinishApi(c, common.Api_status_ok, rows)
}

// @Tags Admin(행정간사)
// @Summary 사용자 관리 User 등록
// @Description 사용자 관리 User 등록
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.userBasicModel true "param"
// @Router /admin/user [post]
// @Success 200
func AdminUserCreate(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	// 행정 간사, 행정 업무 권한이 없는 경우 유저를 등록 할수 없다.!
	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	chkKeys := []interface{}{
		[]string{"email", "phoneno", "user_type"}, //	필수 키
		[][]string{}, //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},   //  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	user := model.User{}
	if err1 := mapstructure.Decode(data, &user); nil != err1 {
		log.Println(err1)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	user.Institution_seq = common.ToUint(tokenMap["institution_seq"])
	if !user.AmdinRegUser(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Admin(행정간사)
// @Summary 사용자 관리 User 일괄 등록
// @Description 사용자 관리 User 일괄 등록
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.userBatchModel true "param"
// @Router /admin/user-batch [post]
// @Success 200
func AdminBatchUserCreate(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	// 행정 간사, 행정 업무 권한이 없는 경우 유저를 등록 할수 없다.!
	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	chkKeys := []interface{}{
		[]string{"userArr"}, //	필수 키
		[][]string{},        //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},          //  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	for _, userMap := range data["userArr"].([]interface{}) {
		user := model.User{}
		if err := mapstructure.Decode(userMap, &user); nil != err {
			log.Println(err)
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
			return
		}
		user.Institution_seq = common.ToUint(tokenMap["institution_seq"])
		if !user.AmdinRegUser(c) {
			return
		}
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Admin(행정간사)
// @Summary 사용자 관리 등록 대기 상태인 유저 가입 메세지 재발송
// @Description 사용자 관리 등록 대기 상태인 유저 가입 메세지 재발송
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /admin/user/{user_seq}/resend-msg [patch]
// @Success 200
func AdminResendMsg(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	if !user.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
		return
	}

	if user.Data["institution_seq"] != tokenMap["institution_seq"] {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "같은 기관 소속이 아닙니다.")
		return
	}

	user.Email = common.ToStr(user.Data["email"])
	if !user.ReSendUserReg(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Admin(행정간사)
// @Summary 등록 대기중인 사용자 삭제
// @Description 등록 대기중인 사용자 삭제
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /admin/user/{user_seq} [delete]
// @Success 200
func AdminUserDelete(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	if !user.DeleteUser(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Admin(행정간사)
// @Summary 사용자 관리 비밀번호 초기화
// @Description 사용자 관리 비밀번호 초기화
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /admin/user/{user_seq}/reset-password [patch]
// @Success 200
func AdminResetUserPassword(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	if !user.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
		return
	}

	if user.Data["institution_seq"] != tokenMap["institution_seq"] {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "같은 기관 소속이 아닙니다.")
		return
	}

	user.Email = common.ToStr(user.Data["email"])
	if !user.SendUserNewPassword(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Admin(행정간사)
// @Summary 행정간사가 사용자 정보 수정
// @Description 행정간사가 사용자 정보 수정
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.adminPatchUserInfoModel true "param"
// @Router /admin/user [patch]
// @Success 200
func AdminPatchUser(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	chkKeys := []interface{}{
		[]string{"user_seq", "user_type", "phoneno", "agree_email", "agree_sms", "agree_pri_open", "name"}, //	필수 키
		[][]string{}, //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{"dept", "position", "major_field", "edu_date", "edu_institution", "edu_course_num"}, //  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	user := model.User{}
	user.Institution_seq = common.ToUint(tokenMap["institution_seq"])
	user.User_status = model.DEF_USER_STATUS_FINISH
	user.Agree_terms_service = 1
	user.Agree_terms_privacy = 1
	if err1 := mapstructure.Decode(data, &user); nil != err1 {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	if !user.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
		return
	}

	if user.Data["institution_seq"] != tokenMap["institution_seq"] {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "같은 기관 소속이 아닙니다.")
		return
	}

	if !user.UpdateUser(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Admin(행정간사)
// @Summary 회원 강제 탈퇴
// @Description 회원 강제 탈퇴(실제 삭제 X 상태만 변경)
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /admin/user/{user_seq}/withdraw [patch]
// @Success 200
func AdminWithdrawUser(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	if !user.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
		return
	}

	if user.Data["institution_seq"] != tokenMap["institution_seq"] {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "같은 기관 소속이 아닙니다.")
		return
	}

	if !user.WithdrawUser(c, model.DEF_USER_STATUS_FORCED_WITHDRAW) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}
