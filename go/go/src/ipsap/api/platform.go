package api

import (
  "github.com/mitchellh/mapstructure"
  "github.com/gin-gonic/gin"
  "ipsap/common"
  "ipsap/model"
  "fmt"
  "log"
	// "strings"
)

// @Tags Platform
// @Summary 관리자 정보
// @Description 관리자 정보
// @Description 우선 관리자 연락처 관리 정보에서 서비스 관리자로만 호출하면됨!
// @Description filter.user_auth = 9  플랫폼 관리자
// @Description filter.user_auth = 10 서비스 관리자
// @Description filter.user_auth = 11 시스템 관리자
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param filter.user_auth query uint true "1"
// @Router /platform/admin-user [get]
// @Success 200
func PlatformAdminUserInfo(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

  user_auth := common.ToUint(c.Request.URL.Query().Get("filter.user_auth"))
  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM &&  user_auth < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	user := model.User{}
	user.User_auth = user_auth
	row := user.GetPlatformAdmin()
	common.FinishApi(c, common.Api_status_ok, row)
}

// @Tags Platform
// @Summary 등록 회원 관리 리스트
// @Description 등록 회원 관리 리스트
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /platform/user [get]
// @Success 200
func PlatformUserList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	user := model.User{}
  moreCondition := fmt.Sprintf(` AND user.user_auth NOT IN (%d, %d)`, model.DEF_USER_AUTH_PLATFORM, model.DEF_USER_AUTH_PLATFORM)
  list := user.GetPlatformUserList(moreCondition)
	common.FinishApi(c, common.Api_status_ok, list)
}

// @Tags Platform
// @Summary 등록 회원 관리 정보
// @Description 등록 회원 관리 정보
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /platform/user/{user_seq} [get]
// @Success 200
func PlatformUserInfo(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

  user_seq, succ := getUserIdFromPath(c)
  if !succ {
    return
  }

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
  if !user.Load(){
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
    return
  }
  moreCondition := fmt.Sprintf(` AND user.email = '%v'`, user.Data["email"])
  list := user.GetInstitutionUserList(moreCondition)
	common.FinishApi(c, common.Api_status_ok, list)
}

// @Tags Platform
// @Summary 회원 등록
// @Description 회원 등록
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.platformUserRegisterModel true "param"
// @Router /platform/user [post]
// @Success 200
func PlatformUserCreate(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

  chkKeys := []interface{}{
    []string{"institution_seq", "email", "phoneno", "user_type"},	//	필수 키
    [][]string{},                      					//	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
    []string{}, 																//  그외 읽어야 할 값들.
  }

  data, eMsg := common.BindCustomBody(c, &chkKeys)
  if nil == data {
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
    return
  }

  user := model.User{}
  if err := mapstructure.Decode(data, &user); nil != err {
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

  if !user.AmdinRegUser(c) {
    return
  }

  common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Platform
// @Summary 회원 일괄 등록
// @Description 회원 일괄 등록
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.platformUserRegisterBatchModel true "param"
// @Router /platform/user-batch [post]
// @Success 200
func PlatformBatchUserCreate(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	chkKeys := []interface{} {
		[]string{"userArr"},	//	필수 키
		[][]string{},                      					//	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{}, 																//  그외 읽어야 할 값들.
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
		if !user.AmdinRegUser(c) {
			return
		}
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Platform
// @Summary Platform관리자가 사용자 정보 수정
// @Description Platform관리자가 사용자 정보 수정
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.platformPatchUserInfoModel true "param"
// @Router /platform/user [patch]
// @Success 200
func PlatformPatchUser(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	chkKeys := []interface{} {
		[]string{"user_seq", "institution_seq", "user_status", "user_type", "phoneno",
             "agree_email", "agree_sms", "agree_pri_open"},	//	필수 키
		[][]string{},                      							//	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{"position", "dept", "edu_course_num", "edu_date", "edu_institution", "major_field" }, //  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	user := model.User{}
	if err := mapstructure.Decode(data, &user); nil != err {
    log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

  user.Agree_terms_service	= 1
  user.Agree_terms_privacy 	= 1
	if !user.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
		return
	}

	if !user.UpdateUser(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Platform
// @Summary Platform관리자가 기관정보 수정
// @Description Platform관리자가 기관정보 수정
// @Description user_seq institution_seq를 추가로 param을 만들어 줘야함!
// @Description userArr에 행정간사데이터를 담아서 보내준다!
// @Accept  mpfd
// @Produce  mpfd
// @Param logo_file formData file false "기관 로고 파일"
// @Param business_file formData file false "사업자 등록증 파일"
// @Security ApiKeyAuth
// @Param param formData string true "기관정보 수정(Json String 형태)"
// @Router  /platform/institution [patch]
// @Success 200
func PlatformPatchInstitution(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
    return
  }

	data, eMsg := common.UnmarshalFormData(c, "param")
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

  ins 						:= model.Institution{}
  institutionMap 	:= data["institution"].(map[string]interface{})
  if err2 := mapstructure.Decode(institutionMap, &ins); nil != err2 {
    log.Println(err2)
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

  // 기관 계정 탈퇴
  if ins.Service_status == model.DEF_SERVICE_STATUS_WITHDRAWN {
    if !ins.RemoveInstitution() {
      common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_system_unknown)
      return
    }
  } else {
    dbConn	:= common.DBconn()
    tx, err	:= dbConn.Begin()
    if nil != err {
      log.Println(err)
      tx.Rollback()
      common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
      return
    }

    if !ins.PlatformUpdateInstitution(c, tx) {
      tx.Rollback()
      common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_system_unknown)
      return
    }

    for _, userArrMap := range data["userArr"].([]interface{}) {
      user := model.User{}
      if err := mapstructure.Decode(userArrMap, &user); nil != err {
        log.Println(err)
        common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_system_unknown)
        return
      }
      if !user.AdminUpdateUser(c, tx) {
        tx.Rollback()
        common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_system_unknown)
        return
      }
    }

    err = tx.Commit()
    if nil != err {
      log.Println(err)
      common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
      tx.Rollback()
      return
    }
  }

	common.FinishApi(c, common.Api_status_ok,gin.H{"rt": "ok",})
}

// @Tags Platform
// @Summary 관리자 정보 수정
// @Description 관리자 정보 수정
// @Description user_seq 포함해서 데이터 만들어야됨! (연락처, 이메일은 필수 값이 아님!)
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.platformAdminUserPatchModel true "param"
// @Router /platform/admin-user [patch]
// @Success 200
func PlatformAdminUserPatch(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
    return
  }

  chkKeys := []interface{} {
    []string{"user_seq", "name"}, //  필수 키
    [][]string{}, //  옵션 키 (그룹 중 1개이상 반드시 있어야 함)
    []string{"telno", "email"}, //  그외 읽어야 할 값들.
  }

  data, eMsg := common.BindCustomBody(c, &chkKeys)
  if nil == data {
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
    return
  }

  user := model.User{}
  if err := mapstructure.Decode(data, &user); nil != err {
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

  if !user.PlatformPatchAdminUser() {
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

	common.FinishApi(c, common.Api_status_ok,gin.H{"rt": "ok",})
}

// @Tags Platform
// @Summary 등록 대기중인 사용자 삭제
// @Description 등록 대기중인 사용자 삭제
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /platform/user/{user_seq} [delete]
// @Success 200
func PlatformUserDelete(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
    return
  }

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	if !user.DeleteUser(c){
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Platform
// @Summary 등록 대기 상태인 유저 가입 메세지 재발송
// @Description 등록 대기 상태인 유저 가입 메세지 재발송
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /platform/user/{user_seq}/resend-msg [patch]
// @Success 200
func PlatformResendMsg(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
    return
  }

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	if !user.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
		return
	}

	user.Email = common.ToStr(user.Data["email"])
	if !user.ReSendUserReg(c){
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Platform
// @Summary 등록회원 관리 비밀번호 초기화
// @Description 등록회원 관리 비밀번호 초기화
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /platform/user/{user_seq}/reset-password [patch]
// @Success 200
func PlatformResetUserPassword(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
    return
  }

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	if !user.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
		return
	}

	user.Email = common.ToStr(user.Data["email"])
	if !user.SendUserNewPassword(c){
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Platform
// @Summary 회원 강제 탈퇴
// @Description 회원 강제 탈퇴(실제 삭제 X 상태만 변경)
// @Description 전체 기관에서 탈퇴함
// @Description 각 기관의 행정간사에게 메일을 보내야함
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /platform/user/{user_seq}/withdraw [patch]
// @Success 200
func PlatformWithdrawUser(c *gin.Context) {
  tokenMap := common.Check_token(c)
  if nil == tokenMap {
    return
  }

  user_seq, succ := getUserIdFromPath(c)
  if !succ {
    return
  }

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
    return
  }

  user := model.User{}
  user.User_seq = common.ToUint(user_seq)
  if !user.Load() {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
    return
  }

  if !user.WithdrawAllInstitutionInUser(c, model.DEF_USER_STATUS_FORCED_WITHDRAW) {
    return
  }

	common.FinishApi(c, common.Api_status_ok, gin.H{ "rt": "ok",})
}

// @Tags Platform
// @Summary 타 기관 유저 리스트
// @Description 타 기관 유저 리스트
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /platform/other-institution/{user_seq} [get]
// @Success 200
func PlatformOtherInstitutionUserList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}


  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
    return
  }

  user_seq, succ := getUserIdFromPath(c)
  if !succ {
    return
  }

	user := model.User{}
  user.User_seq = common.ToUint(user_seq)
  if !user.Load() {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
    return
  }

	moreCondition := fmt.Sprintf(` AND user.email = '%v'
                                 AND user.user_seq <> %v
																 AND user.institution_seq <> %v`,
																 user.Data["email"],
                                 user.User_seq,
																 user.Data["institution_seq"])
	rows := user.GetInstitutionUserList(moreCondition)

	common.FinishApi(c, common.Api_status_ok, rows)
}

// @Tags Platform
// @Summary 관리자가 결제 취소 후 기관 서비스 이용 상태 변경
// @Description 관리자가 결제 취소 후 기관 서비스 이용 상태 변경
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.InstitutionServiceChangeModel true "param"
// @Router /platform/payment/cancel/institution [patch]
// @Success 200
func PlatformInstitutionServiceStatusChange(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

  if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
    return
  }

  chkKeys := []interface{} {
    []string{"institution_seq", "service_status"}, //  필수 키
    [][]string{}, //  옵션 키 (그룹 중 1개이상 반드시 있어야 함)
    []string{}, //  그외 읽어야 할 값들.
  }

  data, eMsg := common.BindCustomBody(c, &chkKeys)
  if nil == data {
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
    return
  }

  instt := model.Institution {
    Institution_seq : common.ToUint(data["institution_seq"]),
    Service_status : common.ToUint(data["service_status"]),
  }

  if !instt.ServiceStatusChange() {
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

	common.FinishApi(c, common.Api_status_ok,gin.H{"rt": "ok",})
}
