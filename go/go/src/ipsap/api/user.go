package api

import (
	"ipsap/common"
	"ipsap/model"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"

	// "strings"
	// "fmt"
	"log"
)

// @Tags User
// @Summary User 정보
// @Description User 정보
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /user/{user_seq} [get]
// @Success 200
func UserInfo(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	if !user.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	common.FinishApi(c, common.Api_status_ok, user.Data)
}

// @Tags User
// @Summary 나의 정보 수정
// @Description 나의 정보 수정
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Param param body model.userMyInfoPatchModel true "param"
// @Router /user/{user_seq} [patch]
// @Success 200
func UserPatch(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

	chkKeys := []interface{}{
		[]string{"phoneno", "agree_email", "agree_sms", "agree_pri_open"}, //	필수 키
		[][]string{}, //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{"dept", "position", "edu_institution", "edu_course_num", "major_field", "edu_date"}, //  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	user.Institution_seq = common.ToUint(tokenMap["institution_seq"])
	user.Agree_terms_service = 1                    // 약관동의
	user.Agree_terms_privacy = 1                    // 개인정보처리방침 동의
	user.User_status = model.DEF_USER_STATUS_FINISH // 유저 등록대기

	if err1 := mapstructure.Decode(data, &user); nil != err1 {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	if !user.UpdateUser(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags User
// @Summary User 탈퇴 (현재 로그인한 기관내에서만 탈퇴!)
// @Description User 탈퇴(현재 로그인한 기관내에서만 탈퇴!)
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /user/{user_seq} [delete]
// @Success 200
func UserDelete(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
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

	if !user.WithdrawUser(c, model.DEF_USER_STATUS_WITHDRAW) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags User
// @Summary User 탈퇴 (소속된 기관 전체에서 탈퇴!)
// @Description User 탈퇴(소속된 기관 전체에서 탈퇴!)
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Router /user/{user_seq}/institution [delete]
// @Success 200
func UserAllDelete(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
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

	if !user.WithdrawAllInstitutionInUser(c, model.DEF_USER_STATUS_WITHDRAW) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})

}

// @Tags User
// @Summary User 비밀번호변경
// @Description User 비밀번호변경
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param user_seq path string true "user_seq"
// @Param param body model.loginChangePwModel true "param"
// @Router /user/{user_seq}/change-password [patch]
// @Success 200
func UserChangePassword(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

	chkKeys := []interface{}{
		[]string{"new_pw", "old_pw"}, //  필수 키
		[][]string{},                 //  옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},                   //  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	if !user.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "아이디가 존재하지 않습니다.")
	}

	newPw := common.DecryptToStd([]byte(common.ToStr(tokenMap["tmp_key"])), common.ToStr(data["new_pw"]))
	oldPw := common.DecryptToStd([]byte(common.ToStr(tokenMap["tmp_key"])), common.ToStr(data["old_pw"]))
	err := bcrypt.CompareHashAndPassword([]byte(common.ToStr(user.Data["pwd"])), []byte(oldPw))
	if err != nil {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_invalide_params, "현재 비밀번호가 틀렸습니다.")
		return
	}

	new_pw, err := bcrypt.GenerateFromPassword([]byte(newPw), bcrypt.DefaultCost)
	if nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	sql := ` UPDATE t_user SET pwd = ?, chg_dttm = UNIX_TIMESTAMP() WHERE email = ?`
	_, err2 := common.DBconn().Exec(sql, string(new_pw), user.Data["email"])
	if nil != err2 {
		log.Println(err2)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags User
// @Summary  나의 정보 등록
// @Description 나의 정보 등록(User가 직접입력)
// @Accept  json
// @Produce  json
// @Param user_seq path string true "user_seq"
// @Param param body model.userRegisterModel true "param"
// @Router /user/{user_seq}/register [patch]
// @Success 200
func UserRegister(c *gin.Context) {
	user_seq, succ := getUserIdFromPath(c)
	if !succ {
		return
	}

	chkKeys := []interface{}{
		[]string{"name", "pwd", "phoneno", "agree_email", "agree_email",
			"agree_sms", "agree_pri_open", "institution_seq", "tmp_key"}, //  필수 키
		[][]string{}, //  옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{"dept", "position", "major_field", "edu_date", "edu_institution", "edu_course_num"}, //  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	sql := `SELECT user_status
	 					FROM t_user
					 WHERE user_seq = ?
					   AND institution_seq = ?`
	row := common.DB_fetch_one(sql, nil, user_seq, data["institution_seq"])
	if nil == row {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "변경할 유저가 없습니다.")
		return
	} else {
		if common.ToInt(row["user_status"]) != model.DEF_USER_STATUS_WAIT {
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "등록대기 상태가 아닙니다.")
			return
		}
	}

	user := model.User{}
	user.User_seq = common.ToUint(user_seq)
	user.User_status = model.DEF_USER_STATUS_FINISH
	user.Agree_terms_service = 1
	user.Agree_terms_privacy = 1
	if err1 := mapstructure.Decode(data, &user); nil != err1 {
		log.Println(err1)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	user.Pwd = common.DecryptToStd([]byte(common.ToStr(data["tmp_key"])), common.ToStr(data["pwd"]))
	if "" == user.Pwd {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "비밀번호가 올바르지 않습니다.")
		return
	}

	if !user.SelfRegUser(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})

}
