package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"ipsap/common"
	"ipsap/model"
	"strings"
	"fmt"
	"log"
	// "golang.org/x/crypto/bcrypt"
)

// @Tags Institution
// @Summary 소속기관으로 이동하기
// @Description 소속기관으로 이동하기
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.moveInstitutionModel true "param"
// @Router /move-institution [post]
// @Success 200
func MoveInstitution(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	chkKeys := []interface{}{
		[]string{"user_seq", "tmp_key"},	//	필수 키
		[][]string{},                     //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{}, 											//  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	// Validation : tmp_key 32자 (AES-256)
	if 32 != len(common.ToStr(data["tmp_key"])) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "tmp_key 값을 확인해 주세요")
		return
	}

	user := model.User{}
	user.User_seq = common.ToUint(data["user_seq"])
	if !user.Load(){
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "user_seq 값을 확인해 주세요")
		return
	}

	if tokenMap["email"] != user.Data["email"] {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_id_mismatch)
		return
	}

	if tokenMap["institution_seq"] == user.Data["institution_seq"] {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "현재 로그인된 기관과 이동하려는 기관이 같습니다.")
		return
	}

	moreCondition := fmt.Sprintf(` AND user.user_seq = %d`, user.User_seq)
	succ, ret := user.GetLoginUser(c, common.ToStr(data["tmp_key"]), moreCondition)
	if !succ {
		return
	}

	common.FinishApi(c, common.Api_status_ok, ret)
}

// @Tags Institution
// @Summary 내가 소속된 타 기관 리스트
// @Description 내가 소속된 타 기관 리스트
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /my/other-institution [get]
// @Success 200
func MyOtherInstitutionList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	user := model.User{}
	moreCondition := fmt.Sprintf(` AND user.email = '%v'
																 AND user.user_status = %d
																 AND user.institution_seq <> %v`,
																 tokenMap["email"],
																 model.DEF_USER_STATUS_FINISH,
																 tokenMap["institution_seq"])
	rows := user.GetInstitutionUserList(moreCondition)

	common.FinishApi(c, common.Api_status_ok, rows)
}

// @Tags Institution
// @Summary 내가 소속된 기관 리스트
// @Description 내가 소속된 기관 리스트
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param param body model.myInstitutionModel true "param"
// @Router /my/institution [post]
// @Success 200
func MyInstitutionList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	chkKeys := []interface{}{
		[]string{"tmp_key"},	//	필수 키
		[][]string{},         //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{}, 					//  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	// Validation : tmp_key 32자 (AES-256)
	if 32 != len(common.ToStr(data["tmp_key"])) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "tmp_key 값을 확인해 주세요")
		return
	}

	user := model.User{}
	moreCondition := fmt.Sprintf(` AND user.email = '%v'`, tokenMap["email"])
	succ, ret := user.GetLoginUser(c, common.ToStr(data["tmp_key"]), moreCondition)
	if !succ {
		return
	}

	common.FinishApi(c, common.Api_status_ok, ret)
}

// @Tags Institution
// @Summary 기관 리스트
// @Description 기관 리스트
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Router /institution [get]
// @Success 200
func InstitutionList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	if common.ToUint(tokenMap["user_auth"]) <= model.DEF_USER_AUTH_INSTITUTION {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	ins  := model.Institution{}
	rows := ins.GetInstitutionList(false)
	common.FinishApi(c, common.Api_status_ok, rows)
}

// @Tags Institution
// @Summary 기관 상세 정보
// @Description 기관 상세 정보
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param institution_seq path string true "institution_seq"
// @Router /institution/{institution_seq} [get]
// @Success 200
func InstitutionInfo(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	institution_seq, succ := getInstitutionIdFromPath(c)
	if !succ {
		return
	}

	ins := model.Institution{}
	ins.Institution_seq = common.ToUint(institution_seq)
	row := ins.GetInstitutionList(true)
	common.FinishApi(c, common.Api_status_ok, row)
}

// @Tags Institution
// @Summary 기관에 등록된 회원 리스트
// @Description 기관에 등록된 회원 리스트
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param institution_seq path string true "institution_seq"
// @Param filter.user_type query string false "1"
// @Router /institution/{institution_seq}/user [get]
// @Success 200
func InstitutionUserList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	institution_seq, succ := getInstitutionIdFromPath(c)
	if !succ {
		return
	}

	moreCondition := fmt.Sprintf(" AND user.user_status = %d", model.DEF_USER_STATUS_FINISH)
	userTypeArr := strings.Split(common.ToStr(c.Request.URL.Query().Get("filter.user_type")), ",")
	if len(userTypeArr) > 0 {
		for _, userType := range userTypeArr {
			moreCondition += fmt.Sprintf(` AND user.user_type LIKE '%%%v%%'`, userType)
		}
	}
	moreCondition += fmt.Sprintf(` AND user.institution_seq = %d`, common.ToUint(institution_seq))

	user := model.User{}
	rows := user.GetInstitutionUserList(moreCondition)
	common.FinishApi(c, common.Api_status_ok, rows)
}

// @Tags Institution
// @Summary 기관정보 수정
// @Description 기관정보 수정
// @Accept  mpfd
// @Produce  mpfd
// @Param logo_file formData file false "기관 로고 파일"
// @Security ApiKeyAuth
// @Param institution_seq path string true "institution_seq"
// @Param param formData string true "기관정보 수정(Json String 형태)"
// @Param test body model.institutionPatchModel false "test용 Json Data 실제사용 X"
// @Router /institution/{institution_seq} [patch]
// @Success 200
func InstitutionPatch(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	institution_seq, succ := getInstitutionIdFromPath(c)
	if !succ {
		return
	}

	data, eMsg := common.UnmarshalFormData(c, "param")
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	ins := model.Institution{}
	ins.Institution_seq	= common.ToUint(institution_seq)
	if err := mapstructure.Decode(data, &ins); nil != err {
		log.Println(err)
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

	if !ins.UpdateInstitution(c, tx) {
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

	ins.Load()
	common.FinishApi(c, common.Api_status_ok,
		gin.H{
			"rt": "ok",
			"logo_file_src": ins.Data["logo_file_src"],
		})
}

// @Tags Institution
// @Summary 기관 결제정보 수정 요청
// @Description 기관 결제정보 수정 요청
// @Accept  mpfd
// @Produce  mpfd
// @Param business_file formData file false "사업자 등록증 파일"
// @Security ApiKeyAuth
// @Param institution_seq path string true "institution_seq"
// @Param param formData string true "기관 결제정보 수정(Json String 형태)"
// @Param test body model.institutionPaymentPatchModel false "test용 Json Data 실제사용 X"
// @Router /institution/{institution_seq}/payment [post]
// @Success 200
func InstitutionPaymentPatch(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	institution_seq, succ := getInstitutionIdFromPath(c)
	if !succ {
		return
	}

	data, eMsg := common.UnmarshalFormData(c, "param")
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	ins := model.Institution{}
	ins.Institution_seq	= common.ToUint(institution_seq)
	if err := mapstructure.Decode(data, &ins); nil != err {
		log.Println(err)
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

	if !ins.InsertRequestPaymentChange(c, tx, common.ToUint64(tokenMap["user_seq"])) {
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

// @Tags Institution
// @Summary  기관에서 현재 이용중인 멤버십의 결제 이력
// @Description 기관에서 현재 이용중인 멤버십의 결제 이력
// @Accept  json
// @Produce  json
// @Param institution_seq path string true "institution_seq"
// @Router /institution/{institution_seq}/using-membership/purchased [get]
// @Success 200
func InstitutionUsingMembershipPurchasedList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	institution_seq, succ := getInstitutionIdFromPath(c)
	if !succ { return }

	instt := model.Institution {
		Institution_seq : common.ToUint(institution_seq),
	}

	list := instt.InstitutionPurchasedList()
	result := gin.H{ "value" : list,}
	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Institution
// @Summary 기관 행정간사 인원수
// @Description 기관 행정간사 인원수
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param institution_seq path string true "institution_seq"
// @Router /institution/{institution_seq}/admin-count [get]
// @Success 200
func InstitutionAdminCount(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	institution_seq, succ := getInstitutionIdFromPath(c)
	if !succ {
		return
	}

	ins := model.Institution{}
	ins.Institution_seq = common.ToUint(institution_seq)
	count := ins.GetInstitutionAdminCount()
	result := gin.H{ "value" : count,}
	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Institution
// @Summary IPSAP 기관 탈퇴
// @Description IPSAP 기관 탈퇴
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param institution_seq path string true "institution_seq"
// @Router /institution/{institution_seq} [delete]
// @Success 200
func InstitutionDelete(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	institution_seq, succ := getInstitutionIdFromPath(c)
	if !succ {
		return
	}

	ins := model.Institution{}
	ins.Institution_seq = common.ToUint(institution_seq)
	ins.LoginToken	= tokenMap
	if !ins.Load() {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	status := common.ToUint(ins.Data["service_status"])
	if status != model.DEF_SERVICE_STATUS_BEFORE_USING || status != model.DEF_SERVICE_STATUS_STOPPED {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_membership_cancel_required)
		return
	}

	if !ins.DeleteInstitution() {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}
