package api

import (
	"ipsap/common"
	"ipsap/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	// "strings"
	"fmt"
	"log"
)

// @Tags Auth
// @Summary 로그인
// @Description 로그인
// @Accept  json
// @Produce  json
// @Param param body model.loginModel true "param"
// @Router /auth/login [post]
// @Success 200
func Login(c *gin.Context) {
	chkKeys := []interface{}{
		[]string{"email", "pw", "tmp_key"}, //	필수 키
		[][]string{},                       //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},                         //  그외 읽어야 할 값들.
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

	email := common.ToStr(data["email"])
	if !common.CheckEmail(email) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "email형식이 아닙니다.")
	}

	user := model.User{}
	user.Email = email
	user.Pwd = common.DecryptToStd([]byte(common.ToStr(data["tmp_key"])), common.ToStr(data["pw"]))
	sql := fmt.Sprintf(`
					SELECT pwd, user_auth
						FROM t_user
					 WHERE email = '%v'
					 GROUP BY email`, user.Email)
	row := common.DB_fetch_one(sql, nil)
	log.Println("============sql : ", sql)
	log.Println("============row : ", row)

	if nil == row {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "존재하지 않는 아이디 입니다.")
		return
	}
	/*
		password := "eam1234!"

		// 비밀번호를 bcrypt로 해싱
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		// 저장된 해시된 비밀번호와 사용자가 입력한 비밀번호를 비교
		err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
		if err != nil {
			fmt.Println("비밀번호가 일치하지 않습니다.")
		} else {
			fmt.Println("비밀번호가 일치합니다.")
		}
	*/
	log.Println("222222222222222222222 : ")
	log.Println("PWD 11111 : ", common.ToStr(row["pwd"]))
	log.Println("PWD 22222 : ", user.Pwd)

	err := bcrypt.CompareHashAndPassword([]byte(common.ToStr(row["pwd"])), []byte(user.Pwd))
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "비밀번호가 틀립니다. 다시 입력해 주세요.")
		return
	}

	moreCondition := fmt.Sprintf(` AND user.email = '%v'`, user.Email)
	log.Println("222222222222222222222 : ")
	succ, ret := user.GetLoginUser(c, common.ToStr(data["tmp_key"]), moreCondition)
	if !succ {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	common.FinishApi(c, common.Api_status_ok, ret)

}

// @Tags Auth
// @Summary 이메일을 통해서 들어왔을때 로그인
// @Description 이메일을 통해서 들어왔을때 로그인
// @Accept  json
// @Produce  json
// @Param param body model.EmailloginModel true "param"
// @Router /auth/email/login [post]
// @Success 200
func EmailLogin(c *gin.Context) {
	chkKeys := []interface{}{
		[]string{"email", "pw", "tmp_key", "user_seq"}, //	필수 키
		[][]string{}, //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},   //  그외 읽어야 할 값들.
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

	email := common.ToStr(data["email"])
	if !common.CheckEmail(email) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "email형식이 아닙니다.")
	}

	user := model.User{}
	user.Email = email
	user.User_seq = common.ToUint(data["user_seq"])
	user.Pwd = common.DecryptToStd([]byte(common.ToStr(data["tmp_key"])), common.ToStr(data["pw"]))
	sql := fmt.Sprintf(`
					SELECT pwd
						FROM t_user
					 WHERE email = '%v'
					   AND user_seq = %d`, user.Email, user.User_seq)
	row := common.DB_fetch_one(sql, nil)
	if nil == row {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "존재하지 않는 아이디 입니다.")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(common.ToStr(row["pwd"])), []byte(user.Pwd))
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "비밀번호가 틀립니다. 다시 입력해 주세요.")
		return
	}

	moreCondition := fmt.Sprintf(` AND user.email = '%v'
		                             AND user.user_seq = %d`,
		user.Email,
		user.User_seq)
	succ, ret := user.GetLoginUser(c, common.ToStr(data["tmp_key"]), moreCondition)
	log.Println("33333333333333333 : ")
	if !succ {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	common.FinishApi(c, common.Api_status_ok, ret)

}

// @Tags Auth
// @Summary 탈퇴 철회
// @Description 탈퇴 신청일로 부터 7일이 지나면 탈퇴 철회 불가
// @Accept  json
// @Produce  json
// @Param param body model.loginModel true "param"
// @Router /auth/email/cancel-withdraw [post]
// @Success 200
func EmailCancelWithDraw(c *gin.Context) {
	chkKeys := []interface{}{
		[]string{"email", "pw", "tmp_key"}, //	필수 키
		[][]string{},                       //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},                         //  그외 읽어야 할 값들.
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

	email := common.ToStr(data["email"])
	if !common.CheckEmail(email) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "email형식이 아닙니다.")
	}

	user := model.User{}
	user.Email = email
	user.Pwd = common.DecryptToStd([]byte(common.ToStr(data["tmp_key"])), common.ToStr(data["pw"]))
	sql := fmt.Sprintf(`
					SELECT pwd
						FROM t_user
					 WHERE email = '%v'`, user.Email)
	row := common.DB_fetch_one(sql, nil)
	if nil == row {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_cancel_withdraw_over_time)
		// common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "존재하지 않는 아이디 입니다.")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(common.ToStr(row["pwd"])), []byte(user.Pwd))
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "비밀번호가 틀립니다. 다시 입력해 주세요.")
		return
	}

	if !user.CancelWithdraw() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_system_unknown)
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Auth
// @Summary Id 찾기
// @Description Id 찾기
// @Accept  json
// @Produce  json
// @Param param body model.userIdFindModel true "param"
// @Router /auth/find-id [post]
// @Success 200
func FindId(c *gin.Context) {
	chkKeys := []interface{}{
		[]string{"name", "phoneno"}, //	필수 키
		[][]string{},                //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},                  //	그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	sql := `SELECT email
						FROM t_user
					 WHERE name = ?
					   AND phoneno = ?`
	row := common.DB_fetch_one(sql, nil, data["name"], data["phoneno"])
	if nil == row {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_info_mismatch)
		return
	}

	common.FinishApi(c, common.Api_status_ok, row["email"])

}

// @Tags Auth
// @Summary  Password 찾기
// @Description Password 찾기 새로운 비밀번호 email 전송
// @Accept  json
// @Produce  json
// @Param param body model.userPasswordFindModel true "param"
// @Router /auth/find-pwd [post]
// @Success 200
func FindPwd(c *gin.Context) {
	chkKeys := []interface{}{
		[]string{"email", "name", "phoneno"}, //	필수 키
		[][]string{},                         //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},                           //	그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	sql := `SELECT user_seq, email
						FROM t_user
					 WHERE email = ?
					 	 AND name = ?
					   AND phoneno = ?`
	row := common.DB_fetch_one(sql, nil, data["email"], data["name"], data["phoneno"])
	if nil == row {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_info_mismatch)
		return
	}

	user := model.User{}
	user.User_seq = common.ToUint(row["user_seq"])
	user.Email = common.ToStr(row["email"])

	if !user.SendUserNewPassword(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok"})
}

// @Tags Token
// @Summary Token 발급
// @Description Token 발급
// @Accept  json
// @Produce  json
// @Param param body model.tokenModel true "param"
// @Router /auth/token [post]
// @Success 200
func Token(c *gin.Context) {
	chkKeys := []interface{}{
		[]string{"email", "pw", "institution_name_ko"}, //	필수 키
		[][]string{}, //	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},   //  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	email := common.ToStr(data["email"])
	if !common.CheckEmail(email) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "email형식이 아닙니다.")
	}

	user := model.User{}
	user.Email = email
	user.Pwd = common.ToStr(data["pw"])
	sql := fmt.Sprintf(`
					SELECT pwd, user_auth
						FROM t_user
					 WHERE email = '%v'
					 GROUP BY email`, user.Email)
	row := common.DB_fetch_one(sql, nil)
	if nil == row {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "존재하지 않는 아이디 입니다.")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(common.ToStr(row["pwd"])), []byte(user.Pwd))
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, 0, "비밀번호가 틀립니다. 다시 입력해 주세요.")
		return
	}

	moreCondition := fmt.Sprintf(` AND user.email = '%v'
															   AND ins.name_ko = '%v'`, user.Email, data["institution_name_ko"])

	//log.Println("moreCondition : ", moreCondition)
	log.Println("11111111111111111111 : ")
	succ, ret := user.GetLoginUser(c, common.ToStr(data["tmp_key"]), moreCondition)
	if !succ {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	log.Println("row : ", row)

	result := ""
	for _, row := range ret {
		delete(row, "ins_info")
		delete(row, "tmp_key")
		delete(row, "user_info")
		delete(row, "user_type_all")
		result = common.ToStr(row["token"])

		log.Println("row  token : ", row["token"])
	}

	log.Println("result : ", result)

	common.FinishApi(c, common.Api_status_ok, result)

}
