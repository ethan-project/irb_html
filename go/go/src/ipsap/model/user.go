package model

import (
	"database/sql"
	"fmt"
	"ipsap/common"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/nleeper/goment"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	User_seq            uint   `json:"-"`
	Institution_seq     uint   `json:"-"`
	User_type           string `json:"user_type" example:"1"`
	User_auth           uint   `json:"-"`
	User_status         uint   `json:"-"`
	Email               string `json:"email" example:"test@test.com"`
	Pwd                 string `json:"pwd" example:"tmvmfls!00"`
	Name                string `json:"name" example:"홍길동"`
	Name_en             string `json:"name_en" example:"Hong"`
	Dept                string `json:"dept" example:"소속부서"`
	Position            string `json:"position" example:"직급"`
	Telno               string `json:"telno" example:"021231234"`
	Phoneno             string `json:"phoneno" example:"01012341234"`
	Major_field         string `json:"major_field" example:"수의학"`
	Edu_date            string `json:"edu_date" example:"20210101"`
	Edu_institution     string `json:"edu_institution" example:"건국대학교"`
	Edu_course_num      string `json:"edu_course_num" example:"0"`
	Agree_email         uint   `json:"agree_email" example:"0"`
	Agree_sms           uint   `json:"agree_sms" example:"0"`
	Agree_pri_open      uint   `json:"agree_pri_open" example:"0"`
	Agree_terms_service uint   `json:"-`
	Agree_terms_privacy uint   `json:"-`
	Reg_dttm            uint   `json:"-"`
	Chg_dtt             uint   `json:"-"`

	Data map[string]interface{} `json:"-"`
}

func (user *User) InsertUser(c *gin.Context, tx *sql.Tx) (succ bool) {
	if !user.CheckValidation(c, true, false) {
		return
	}

	sql := `INSERT INTO t_user(institution_seq,		user_type, 						user_auth,
														 user_status,				email,								pwd,
														 name, 							name_en,							dept,
														 position, 					telno,								phoneno,
														 major_field,				edu_date,							edu_institution,
														 edu_course_num,		agree_email,					agree_sms,
														 agree_pri_open,		agree_terms_service,	agree_terms_privacy,
														 reg_dttm)
					VALUES(?,?,?,
								 ?,?,?,
								 ?,?,?,
								 ?,?,?,
								 ?,?,?,
								 ?,?,?,
								 ?,?,?,
								 UNIX_TIMESTAMP())`

	result, err := tx.Exec(sql,
		user.Institution_seq, user.User_type, user.User_auth,
		DEF_USER_STATUS_WAIT, user.Email, user.Pwd,
		user.Name, user.Name_en, user.Dept,
		user.Position, user.Telno, user.Phoneno,
		user.Major_field, user.Edu_date, user.Edu_institution,
		user.Edu_course_num, user.Agree_email, user.Agree_sms,
		user.Agree_pri_open, user.Agree_terms_service, user.Agree_terms_privacy)
	if nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		log.Println(err)
		return
	}

	no, err := result.LastInsertId()
	if nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		log.Println(err)
		return
	}

	user.User_seq = cast.ToUint(no)

	return true
}

func (user *User) CheckValidation(c *gin.Context, isAdmin bool, isPlatFormAdminUpdate bool) (succ bool) {

	if 0 == user.Institution_seq {
		return
	}

	if isAdmin {
		if !isPlatFormAdminUpdate {
			if !user.CheckDuplicateUserEmail(c) {
				return
			}
		}

		if "" == user.Dept {
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "소속(부서)는 필수 입력 사항입니다.")
			return
		}

		if "" == user.Position {
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "소속(직책)는 필수 입력 사항입니다.")
			return
		}

	}

	log.Println("============user.Pwd 11111111111111111111 : ", user.Pwd)
	if "" != user.Pwd {
		user_pwd, err := bcrypt.GenerateFromPassword([]byte(user.Pwd), bcrypt.DefaultCost)
		if nil != err {
			common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
			return
		}

		user.Pwd = cast.ToString(user_pwd)
	}

	codeD := Create_CodeDynamic(INSTITUTION_SHARE_CODE)

	// 소속부서 Dynamic code 값으로 변경한다. (Decode Type 1)
	if "" != user.Dept {
		codeD.DCode_type = DCODE_TYPE_DEPT
		codeD.Value = user.Dept
		user.Dept = common.ToStr(codeD.GetCodeFromValue())
	} else {
		user.Dept = "0"
	}

	// 직급을 Dynamic code 값으로 변경한다. (Decode Type 2)
	if "" != user.Position {
		codeD.DCode_type = DCODE_TYPE_POSITION
		codeD.Value = user.Position
		user.Position = common.ToStr(codeD.GetCodeFromValue())
	} else {
		user.Position = "0"
	}

	// 전공분야 Dynamic code 값으로 변경한다. (Decode Type 4)
	if "" != user.Major_field {
		codeD.DCode_type = DCODE_TYPE_MAJOR
		codeD.Value = user.Major_field
		user.Major_field = common.ToStr(codeD.GetCodeFromValue())
	} else {
		user.Major_field = "0"
	}

	// 교육기관명 Dynamic code 값으로 변경한다. (Decode Type 5)
	if "" != user.Edu_institution {
		codeD.DCode_type = DCODE_TYPE_EDU_ISTT
		codeD.Value = user.Edu_institution
		user.Edu_institution = common.ToStr(codeD.GetCodeFromValue())
	} else {
		user.Edu_institution = "0"
	}

	if !common.CheckBinary(user.Agree_email) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "이메일 수신 동의는 필수 입니다.")
		return
	}

	if !common.CheckBinary(user.Agree_sms) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "SMS 수신 여부 값은 필수 입니다.")
		return
	}

	if !common.CheckBinary(user.Agree_pri_open) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "성명, 연락처등 정보 공개는 필수 입니다.")
		return
	}

	// 휴대폰 번호, 전화번호 체크 필요!!
	return true
}

func (user *User) CheckDuplicateUserEmail(c *gin.Context) (succ bool) {

	if !common.CheckEmail(user.Email) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "Email 형식이 아닙니다.")
		return
	}

	sql := fmt.Sprintf(`SELECT user_seq FROM t_user WHERE email = '%v' AND institution_seq = %d`, user.Email, user.Institution_seq)
	row := common.DB_fetch_one(sql, nil)
	if nil != row {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_dup_name, "이미 사용중인 email 입니다.")
		return
	}
	return true
}

func GetUserFilter() (filter func(map[string]interface{})) {
	filter = func(row map[string]interface{}) {

		var tmpArr []string
		userTypeArr := strings.Split(common.ToStr(row["user_type"]), ",")
		for _, userType := range userTypeArr {
			codeUT := Code{
				Type: "user_type",
				Id:   common.ToUint(userType),
			}
			tmpArr = append(tmpArr, codeUT.GetCodeStrFromTypeAndId())
		}
		row["user_type_str"] = strings.Join(tmpArr, ",")

		codeUA := Code{
			Type: "user_auth",
			Id:   common.ToUint(row["user_auth"]),
		}
		row["user_auth_str"] = codeUA.GetCodeStrFromTypeAndId()

		withdrawTimeDiff := ""
		if common.ToInt(row["withdraw_timediff"]) < 0 && common.ToUint(row["user_status"]) == DEF_USER_STATUS_WITHDRAW {
			withdrawTimeDiff = fmt.Sprintf("(%v일)", common.ToInt(row["withdraw_timediff"]))
		}
		codeST := Code{
			Type: "user_status",
			Id:   common.ToUint(row["user_status"]),
		}
		row["user_status_str"] = codeST.GetCodeStrFromTypeAndId() + withdrawTimeDiff

		dcodeDept := CodeDynamic{
			Institution_seq: INSTITUTION_SHARE_CODE,
			DCode_type:      DCODE_TYPE_DEPT,
			Code:            common.ToUint(row["dept"]),
		}
		row["dept_str"] = dcodeDept.GetValueFromCode()

		dcodePosition := CodeDynamic{
			Institution_seq: INSTITUTION_SHARE_CODE,
			DCode_type:      DCODE_TYPE_POSITION,
			Code:            common.ToUint(row["position"]),
		}
		row["position_str"] = dcodePosition.GetValueFromCode()

		dcodeMajor := CodeDynamic{
			Institution_seq: INSTITUTION_SHARE_CODE,
			DCode_type:      DCODE_TYPE_MAJOR,
			Code:            common.ToUint(row["major_field"]),
		}
		row["major_field_str"] = dcodeMajor.GetValueFromCode()

		dcodeEduIstt := CodeDynamic{
			Institution_seq: INSTITUTION_SHARE_CODE,
			DCode_type:      DCODE_TYPE_EDU_ISTT,
			Code:            common.ToUint(row["edu_institution"]),
		}
		row["edu_institution_str"] = dcodeEduIstt.GetValueFromCode()
	}
	return
}

func getUserQueryAndFilter(moreCondition string) (sql string, filter func(map[string]interface{})) {
	/*
			sql = fmt.Sprintf(`
								SELECT user.user_seq,		user.user_type,				user.user_auth,				user.email,
											 user.pwd,				user.name,						user.name_en,					user.dept,
											 user.position,		user.telno,						user.phoneno,					user.major_field,
											 user.edu_date,		user.edu_institution,	user.edu_course_num,	user.agree_email,
											 user.agree_sms,	user.agree_pri_open,	user.user_status,			ins.institution_code,
											 user.institution_seq, user.withdraw_dttm,
											 IF(user.withdraw_dttm = 0, 99999, TIMESTAMPDIFF(DAY, NOW(), FROM_UNIXTIME(user.withdraw_dttm))) AS withdraw_timediff,
											 ins.name_ko as institution_name_ko,
											 ins.name_en as institution_name_en,
											 ins.membership_payment_date,
											 ins.expiration_date
									FROM t_user user, t_institution ins
								 WHERE user.institution_seq = ins.institution_seq
								 	 	%v	#moreCondition`, moreCondition)
		filter = GetUserFilter()
	*/

	/*
		sql = `SELECT user.user_seq, user.user_type, user.user_auth, user.email, user.pwd,	user.name FROM t_user user`

		filter = func(row map[string]interface{}) {
			row["pwd"] = common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["pwd"]))
		}
	*/

	sql = fmt.Sprintf(`
	SELECT user.user_seq,		user.user_type,				user.user_auth,				user.email,
					user.pwd,				user.name,						user.name_en
		FROM t_user user`)
	filter = GetUserFilter()

	return
}

func (user *User) Load() (succ bool) {
	if user.User_seq == 0 {
		return
	}

	if user.Data != nil {
		return true
	}

	moreCondition := fmt.Sprintf(` AND user_seq = %d`, user.User_seq)
	sql, filter := getUserQueryAndFilter(moreCondition)
	user.Data = common.DB_fetch_one(sql, filter)
	return nil != user.Data
}

func (user *User) GetInstitutionUserList(moreCondition string) (rows []map[string]interface{}) {
	sql, filter := getUserQueryAndFilter(moreCondition)
	return common.DB_fetch_all(sql, filter)
}

func (user *User) GetLoginUser(c *gin.Context, tmp_key string, moreCondition string) (succ bool, ret []map[string]interface{}) {

	sql := fmt.Sprintf(`
						SELECT user.user_seq,					user.user_type,	user.user_auth,
									 user.email,						user.name,			user.name_en,
									 user.dept,							user.position,	user.institution_seq,
									 user.user_status,			user.phoneno,		ins.institution_type,
									 ins.name_ko as institution_name_ko,
									 ins.institution_code,	(SELECT IF(COUNT(app.check_user_seq)> 0,true,false)
																						FROM t_application app
																					 WHERE app.check_user_seq = user.user_seq
																						 AND NOT EXISTS ( SELECT app2.application_seq
																								 								FROM t_application app2
																								 							 WHERE app2.parent_app_seq = app.application_seq
																									 					 		 AND app2.application_type = %v
																									 						 	 AND app2.reg_user_seq = user.user_seq)) as inspector,
									 ins.service_status
						  FROM t_user user, t_institution ins
						 WHERE user.institution_seq = ins.institution_seq
							 AND ins.institution_status = %d
							 AND user.user_status < %d
						    %v #moreCondition`, DEF_APP_TYPE_CHECKLIST, DEF_INSTITUTION_STATUS_OK, DEF_USER_STATUS_WITHDRAW, moreCondition)
	filter := func(row map[string]interface{}) {
		userTypeArr := strings.Split(common.ToStr(row["user_type"]), ",")
		userType := make(map[string]bool)
		for _, data := range userTypeArr {
			userType[common.ToStr(data)] = true
		}

		sql := `SELECT id, value FROM t_code WHERE type = 'user_type'`
		rows := common.DB_fetch_all(sql, nil)
		for _, data2 := range rows {
			_, succ := userType[common.ToStr(data2["id"])]
			if !succ {
				userType[common.ToStr(data2["id"])] = false
			}
		}

		row["user_type_all"] = row["user_type"]
		row["user_type"] = userType

		codeUA := Code{
			Type: "user_auth",
			Id:   common.ToUint(row["user_auth"]),
		}

		row["user_auth_str"] = codeUA.GetCodeStrFromTypeAndId()

		dcodeDept := CodeDynamic{
			Institution_seq: INSTITUTION_SHARE_CODE,
			DCode_type:      DCODE_TYPE_DEPT,
			Code:            common.ToUint(row["dept"]),
		}
		row["dept_str"] = dcodeDept.GetValueFromCode()

		dcodePosition := CodeDynamic{
			Institution_seq: INSTITUTION_SHARE_CODE,
			DCode_type:      DCODE_TYPE_POSITION,
			Code:            common.ToUint(row["position"]),
		}
		row["position_str"] = dcodePosition.GetValueFromCode()

		codeIT := Code{
			Type: "institution_type",
			Id:   common.ToUint(row["institution_type"]),
		}
		row["institution_type_str"] = codeIT.GetCodeStrFromTypeAndId()
	}

	//log.Println("sql : ", sql)

	rows := common.DB_fetch_all(sql, filter)
	if len(rows) > 0 {
		for _, row := range rows {
			data := make(map[string]interface{})
			ins := Institution{}
			ins.Institution_seq = common.ToUint(row["institution_seq"])
			if !ins.Load() {
				return
			}
			uuid, _ := uuid.NewRandom()
			uuidStr := strings.ReplaceAll(uuid.String(), "-", "")
			row["tmp_key"] = uuidStr
			token := common.Make_token(c, row)
			if "" == token {
				common.FinishApi(c, common.Api_status_unauthorized)
				return
			}

			delete(row, "tmp_key")
			delete(row, "institution_type")
			delete(row, "institution_type_str")
			delete(row, "user_auth_str")
			delete(row, "position")
			delete(row, "position_str")
			delete(row, "dept")

			delete(ins.Data, "addr1")
			delete(ins.Data, "addr2")
			delete(ins.Data, "business_file_org_name")
			delete(ins.Data, "business_file_path")
			delete(ins.Data, "business_file_src")
			delete(ins.Data, "business_num")
			delete(ins.Data, "chg_dttm")
			delete(ins.Data, "expiration_date")
			delete(ins.Data, "homepage_url")
			delete(ins.Data, "invoice_flag")
			delete(ins.Data, "invoice_flag_str")
			delete(ins.Data, "logo_file_path")
			delete(ins.Data, "reg_dttm")
			delete(ins.Data, "zipcode")
			delete(ins.Data, "payment_method_str")
			delete(ins.Data, "ia_final_director_str")
			delete(ins.Data, "ia_evaluation_method_str")
			delete(ins.Data, "institution_status_str")
			delete(ins.Data, "institution_type_str")
			delete(ins.Data, "logo_file_org_name")

			data["token"] = token
			data["user_info"] = row
			data["ins_info"] = ins.Data
			data["tmp_key"] = common.EncryptToStd([]byte(tmp_key), uuidStr)

			ret = append(ret, data)
		}
	} else {
		return
	}

	succ = true
	return

}

func (user *User) UpdateUser(c *gin.Context) (succ bool) {
	if !user.CheckValidation(c, false, false) {
		return
	}
	moreUpdate := ""

	if "" != user.User_type {
		moreUpdate += fmt.Sprintf(", user_type = '%v'", user.User_type)
	}

	if "0" != user.Dept {
		moreUpdate += fmt.Sprintf(", dept = '%v'", user.Dept)
	}

	if "0" != user.Position {
		moreUpdate += fmt.Sprintf(", position = '%v'", user.Position)
	}

	if "0" != user.Major_field {
		moreUpdate += fmt.Sprintf(", major_field = '%v'", user.Major_field)
	}

	if "" != user.Name {
		moreUpdate += fmt.Sprintf(", name = '%v'", user.Name)
	}

	sql := fmt.Sprintf(`
					UPDATE t_user
						 SET phoneno = ?,					edu_date = ?,							edu_institution = ?,
						 		 edu_course_num = ?,	agree_email = ?,					agree_sms = ?,
								 agree_pri_open = ?, 	agree_terms_service = ?,	agree_terms_privacy = ?,
								 user_status = ?			%v #moreUpdate
					 WHERE user_seq = ?`, moreUpdate)
	_, err := common.DBconn().Exec(sql,
		user.Phoneno, user.Edu_date, user.Edu_institution,
		user.Edu_course_num, user.Agree_email, user.Agree_sms,
		user.Agree_pri_open, user.Agree_terms_service, user.Agree_terms_privacy,
		user.User_status, user.User_seq)
	if nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		log.Println(err)
		return
	}

	succ = true
	return
}

func (user *User) DeleteUser(c *gin.Context) (succ bool) {

	dbConn := common.DBconn()
	tx, err := dbConn.Begin()
	if nil != err {
		log.Println(err)
		tx.Rollback()
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	sql := fmt.Sprintf(`
	INSERT INTO t_user_rm(user_seq,							institution_seq,	user_type,
												user_auth,						user_status,			email,
												pwd,									name,							name_en,
												dept,									position,					telno,
												phoneno,							major_field,			edu_date,
												edu_institution,			edu_course_num,		agree_email,
												agree_sms,						agree_pri_open,		agree_terms_service,
												agree_terms_privacy,	reg_dttm,					chg_dttm,
												withdraw_dttm,				remove_dttm)
		SELECT *, UNIX_TIMESTAMP()
	  	FROM t_user
	 	 WHERE user_seq = %d`, user.User_seq)
	_, err = tx.Exec(sql)
	if nil != err {
		log.Println(err)
		tx.Rollback()
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	sql = `DELETE FROM t_user WHERE user_seq = ?`
	_, err2 := tx.Exec(sql, user.User_seq)
	if nil != err2 {
		log.Println(err2)
		tx.Rollback()
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	err = tx.Commit()
	if nil != err {
		log.Println(err)
		tx.Rollback()
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	succ = true
	return
}

func (user *User) GetPlatformUserList(moreCondition string) (rows []map[string]interface{}) {
	sql := fmt.Sprintf(`
					SELECT user.user_seq,		GROUP_CONCAT(user.user_type separator '/') AS user_type,
								 user.user_auth,  user.email,
								 user.name,				user.phoneno, user.institution_seq,
								 GROUP_CONCAT(user.user_status separator '/') AS user_status,
								 GROUP_CONCAT((IF(user.withdraw_dttm = 0, 99999, TIMESTAMPDIFF(DAY, NOW(), FROM_UNIXTIME(user.withdraw_dttm)))) separator '/') AS withdraw_timediff,
								 GROUP_CONCAT(ins.name_ko separator ' / ') AS institution_name_ko
						FROM t_user user, t_institution ins
					 WHERE user.institution_seq = ins.institution_seq
							%v	#moreCondition
				GROUP BY user.email`, moreCondition)
	filter := func(row map[string]interface{}) {
		userTypeArr1 := strings.Split(common.ToStr(row["user_type"]), "/")
		var tmpArr2 []string
		for _, userTypeArr2 := range userTypeArr1 {
			var tmpArr []string
			userTypeArr := strings.Split(userTypeArr2, ",")
			for _, userType := range userTypeArr {
				codeUT := Code{
					Type: "user_type",
					Id:   common.ToUint(userType),
				}
				tmpArr = append(tmpArr, codeUT.GetCodeStrFromTypeAndId())
			}
			tmpArr2 = append(tmpArr2, strings.Join(tmpArr, ","))
		}
		row["user_type_str"] = strings.Join(tmpArr2, " / ")

		withdrawTimeDiffArr := strings.Split(common.ToStr(row["withdraw_timediff"]), "/")
		userStatusArr := strings.Split(common.ToStr(row["user_status"]), "/")
		withdrawTimeDiffStr := ""
		var userStatusArr2 []string

		for i, withdrawTimeDiff := range withdrawTimeDiffArr {
			for j, userStatus := range userStatusArr {
				if i == j {
					if common.ToInt(withdrawTimeDiff) < 0 && common.ToUint(userStatus) == DEF_USER_STATUS_WITHDRAW {
						withdrawTimeDiffStr = fmt.Sprintf("(%v일)", common.ToInt(row["withdraw_timediff"]))
					}
					codeST := Code{
						Type: "user_status",
						Id:   common.ToUint(userStatus),
					}
					user_status_str := codeST.GetCodeStrFromTypeAndId()
					user_status_tag := ""
					switch common.ToUint(userStatus) {
					case DEF_USER_STATUS_WAIT, DEF_USER_STATUS_EXISTING_UESR_WAIT:
						user_status_tag = fmt.Sprintf("<span class='%v status'>%v</span>", "unregistered", user_status_str)
					case DEF_USER_STATUS_FINISH:
						user_status_tag = fmt.Sprintf("<span class='%v status'>%v</span>", "registered", user_status_str)
					case DEF_USER_STATUS_WITHDRAW:
						user_status_tag = fmt.Sprintf("<span class='%v status'>%v</span>", "withdrawn", user_status_str+withdrawTimeDiffStr)
					case DEF_USER_STATUS_FORCED_WITHDRAW:
						user_status_tag = fmt.Sprintf("<span class='%v status'>%v</span>", "expelled", user_status_str+withdrawTimeDiffStr)
					case DEF_USER_STATUS_REGIST_FAIL:
						user_status_tag = fmt.Sprintf("<span class='%v status'>%v</span>", "withdrawn", user_status_str)
					}
					userStatusArr2 = append(userStatusArr2, user_status_tag)
					break
				}
			}
		}
		row["user_status_str"] = strings.Join(userStatusArr2, " / ")
	}
	return common.DB_fetch_all(sql, filter)
}

func (user *User) AmdinRegUser(c *gin.Context) (succ bool) {
	dbConn := common.DBconn()
	tx, err := dbConn.Begin()
	if nil != err {
		tx.Rollback()
		return
	}

	defer func() {
		tx.Rollback()
	}()

	exist := false
	sql := `SELECT user_seq
						FROM t_user
				   WHERE email = ?`
	row := common.DB_fetch_one(sql, nil, user.Email)
	if nil != row && nil != row["user_seq"] {
		exist = true
	}

	tmp_pwd := ""
	user_pwd := ""
	if exist {
		sql = fmt.Sprintf(
			`INSERT INTO t_user(institution_seq,	user_type,			user_auth,
													user_status,			email,					pwd,
													name, 						name_en, 				telno,
													phoneno, 					major_field,		edu_date,
													edu_institution,	edu_course_num,	reg_dttm)
			 SELECT %v, '%v', %v,
							%v, '%v', pwd,
							name, name_en, telno,
							phoneno, major_field, edu_date,
							edu_institution, edu_course_num, UNIX_TIMESTAMP()
				 FROM t_user
				WHERE user_seq = %v`,
			user.Institution_seq, user.User_type, DEF_USER_AUTH_NOMARL,
			DEF_USER_STATUS_EXISTING_UESR_WAIT, user.Email,
			row["user_seq"])
	} else {
		succ, user_pwd, tmp_pwd = getUserTmpPassword(c)
		if !succ {
			return
		}
		sql = fmt.Sprintf(`
						INSERT INTO t_user(institution_seq,	user_type,	user_auth,
															 user_status,			email,			pwd,
															 phoneno,					reg_dttm)
						VALUES(%v,'%v',%v,
									 %v,'%v','%v',
									 '%v',UNIX_TIMESTAMP())`,
			user.Institution_seq, user.User_type, DEF_USER_AUTH_NOMARL,
			DEF_USER_STATUS_WAIT, user.Email, user_pwd,
			user.Phoneno)
	}

	result, err := tx.Exec(sql)
	if nil != err {
		log.Println(err)
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == 1062 { //  Duplicate entry
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_dup_name, "이미 사용중인 email 입니다.")
			} else {
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
			}
		}
		return false
	}

	no, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return false
	}

	user.User_seq = common.ToUint(no)
	moreCondition := fmt.Sprintf(` AND user_seq = %d`, user.User_seq)
	sql, filter := getUserQueryAndFilter(moreCondition)
	user.Data = common.DB_Tx_fetch_one(tx, sql, filter)
	if exist {
		// 다른기관에서 새로 등록됨
		user.Data["agree_sms"] = 1
		user.Data["agree_email"] = 1
		succ = user.SendMsg("", DEF_MSG_EXISTING_USER_REGIST)
	} else {
		// 초기 등록시 메일 및 sms 동의하는걸로 처리
		user.Data["agree_sms"] = 1
		user.Data["agree_email"] = 1
		user.Data["name"] = "회원"
		succ = user.SendMsg(tmp_pwd, DEF_MSG_USER_REGIST)
	}

	if !succ {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_send_email)
		return false
	}

	err = tx.Commit() //  commit후에 rollback이 다시호출되지만 상관 없음,!!
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return false
	}

	return true
}

func (user *User) ReSendUserReg(c *gin.Context) (succ bool) {
	user_pwd := ""
	tmp_pwd := ""
	succ, user_pwd, tmp_pwd = getUserTmpPassword(c)
	if !succ {
		return
	}

	sql := `UPDATE t_user
		 				 SET pwd = ?
	 				 WHERE email = ?`
	_, err := common.DBconn().Exec(sql, user_pwd, user.Email)
	user.Data["agree_sms"] = 1
	user.Data["agree_email"] = 1
	user.Data["name"] = "회원"
	if nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		log.Println(err)
		return
	}

	succ = user.SendMsg(tmp_pwd, DEF_MSG_USER_REGIST)
	if !succ {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_send_email)
		return false
	}

	return true
}

func (user *User) SelfRegUser(c *gin.Context) (succ bool) {
	if !user.CheckValidation(c, false, false) {
		return
	}

	sql := `UPDATE t_user
						 SET	user_status = ?,	pwd = ?,			name = ?,
									dept = ?,					position = ?,	phoneno = ?,
									major_field = ?,	edu_date = ?,	edu_institution = ?,
									edu_course_num = ?,	agree_email = ?, agree_sms = ?,
									agree_pri_open = ?,	agree_terms_service = ?, agree_terms_privacy = ?,
									chg_dttm = UNIX_TIMESTAMP()
					  WHERE user_seq = ?`
	_, err := common.DBconn().Exec(sql,
		user.User_status, user.Pwd, user.Name,
		user.Dept, user.Position, user.Phoneno,
		user.Major_field, user.Edu_date, user.Edu_institution,
		user.Edu_course_num, user.Agree_email, user.Agree_sms,
		user.Agree_pri_open, user.Agree_terms_service, user.Agree_terms_privacy,
		user.User_seq)
	if nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		log.Println(err)
		return
	}

	return true
}

func (user *User) WithdrawUser(c *gin.Context, user_status uint) (succ bool) {
	succ = false
	userTypeArr := strings.Split(common.ToStr(user.Data["user_type"]), ",")
	if common.CheckUserTypeAuth(userTypeArr, DEF_USER_TYPE_ADMIN_SECRETARY) {
		sql := fmt.Sprintf(`SELECT COUNT(user_seq) cnt
													FROM t_user
												 WHERE user_type LIKE '%%%v%%'
													 AND institution_seq = %v`, DEF_USER_TYPE_ADMIN_SECRETARY, user.Data["institution_seq"])
		row := common.DB_fetch_one(sql, nil)
		if nil == row {
			common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		} else {
			if common.ToInt(row["cnt"]) == 1 {
				common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_last_admin)
				return
			}
		}
	}

	sql := `UPDATE t_user
						 SET user_status  = ?,
						 		 withdraw_dttm = UNIX_TIMESTAMP()
					 WHERE user_seq = ?`
	_, err := common.DBconn().Exec(sql, user_status, user.User_seq)
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	withdrawUser := User{
		User_seq: user.User_seq,
	}

	withdrawUser.Load()
	unixTime, _ := goment.Unix(common.ToInt64(withdrawUser.Data["withdraw_dttm"]))
	t1, _ := goment.New(unixTime)
	t1.Add(5, "days")
	withdrawUser.Data["retractable_period"] = common.GetDateStr(t1)
	if !user.SendMsg("", DEF_MSG_USER_WITHDRAW) {
		return
	}

	succ = true
	return
}

// 해당 기관의 행정 간사 이름, email, 전화번호, 기관로고 가져오기
func (user *User) GetInstitutionAdminInfo(moreCondition string) (row map[string]interface{}) {
	sql := fmt.Sprintf(`
					SELECT user.email, user.name, user.user_seq,
								 CASE LENGTH(user.telno)
									 WHEN 11 THEN CONCAT(LEFT(user.telno, 3), '-', MID(user.telno, 4, 4), '-', RIGHT(user.telno, 4))
									 WHEN 10 THEN CONCAT(LEFT(user.telno, 3), '-', MID(user.telno, 4, 3), '-', RIGHT(user.telno, 4))
									 WHEN 9  THEN CONCAT(LEFT(user.telno, 2), '-', MID(user.telno, 3, 3), '-', RIGHT(user.telno, 4))
								 END telno,
								 ins.logo_file_path
						FROM t_user user, t_institution ins
					 WHERE user.institution_seq = ins.institution_seq
					 	 AND user.user_status = 2
					   AND user.institution_seq = %v
						 AND (user_type LIKE '%%%v%%')
						 %v #moreCondition`,
		user.Institution_seq,
		DEF_USER_TYPE_ADMIN_SECRETARY,
		moreCondition)
	filter := func(row map[string]interface{}) {
		row["logo_file_src"] = ""
		if "" != row["logo_file_path"] {
			// row["logo_file_path"] = common.EncryptToUrl([]byte(common.ToStr(ins.LoginToken["tmp_key"])), common.ToStr(row["logo_file_path"]))
			encPath := common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["logo_file_path"]))
			row["logo_file_src"] = common.MakeDownloadUrl(encPath)
		}
	}

	return common.DB_fetch_one(sql, filter)
}

func getUserTmpPassword(c *gin.Context) (suuc bool, user_pwd string, tmp_pwd string) {
	// 1. 자릿수, 2. 숫자, 3. 특수문자, 4. 대문자금지, 5 반복문자허용
	tmpPassword, err := password.Generate(10, 2, 0, false, false)
	if err != nil {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	bytePwd, err := bcrypt.GenerateFromPassword([]byte(tmpPassword), bcrypt.DefaultCost)
	if nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	tmp_pwd = tmpPassword
	user_pwd = string(bytePwd)
	suuc = true
	return
}

func (user *User) AdminUpdateUser(c *gin.Context, tx *sql.Tx) (succ bool) {
	succ = false
	if !user.CheckValidation(c, true, true) {
		return
	}

	sql := fmt.Sprintf(`
					UPDATE t_user
						 SET email = ?,			name = ?,		name_en = ?,	dept = ?,
								 position = ?,	telno = ?,	phoneno = ?,	agree_email = ?,
								 agree_sms = ?,	agree_pri_open = ?
					 WHERE user_seq = ?`)
	_, err := tx.Exec(sql,
		user.Email, user.Name, user.Name_en, user.Dept,
		user.Position, user.Telno, user.Phoneno, user.Agree_email,
		user.Agree_sms, user.Agree_pri_open,
		user.User_seq)
	if nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		log.Println(err)
		return
	}

	succ = true
	return
}

func (user *User) SendUserNewPassword(c *gin.Context) (succ bool) {
	if 0 == user.User_seq && "" == user.Email {
		return
	}

	dbConn := common.DBconn()
	tx, err := dbConn.Begin()
	if nil != err {
		tx.Rollback()
		return
	}

	defer func() {
		tx.Rollback()
	}()

	succ, user_pwd, tmp_pwd := getUserTmpPassword(c)
	if !succ {
		return
	}

	sql := `UPDATE t_user
						 SET pwd = ?, chg_dttm = UNIX_TIMESTAMP()
					 WHERE email = ?`
	_, err = tx.Exec(sql, user_pwd, user.Email)
	if nil != err {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		log.Println(err)
		return
	}

	if !user.SendMsg(tmp_pwd, DEF_MSG_USER_PASSWORD_CHANGE) {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_send_email)
		return false
	}

	err = tx.Commit() //  commit후에 rollback이 다시호출되지만 상관 없음,!!
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	return true
}

func (user *User) SendMsg(tmpPassword string, msg_id int) (succ bool) {
	succ = false
	if !user.Load() {
		return
	}

	user.Data["tmp_pwd"] = tmpPassword
	// 21-09-07 donghun 행정간사가 사용자 비밀번호 초기화시 강제로 동의
	if "" != tmpPassword {
		user.Data["agree_sms"] = 1
		user.Data["agree_email"] = 1
	}

	msgMgr := MessageMgr{
		User_info: user.Data,
		Msg_ID:    msg_id,
	}

	succ = msgMgr.SendMessage()
	return
}

func (user *User) GetPlatformAdmin() map[string]interface{} {
	sql := `SELECT user_seq, name, email, telno
					  FROM t_user
					 WHERE user_auth = ?`
	return common.DB_fetch_one(sql, nil, user.User_auth)
}

func (user *User) PlatformPatchAdminUser() (succ bool) {
	succ = false
	moreUpdate := ""
	if "" != user.Phoneno {
		moreUpdate += fmt.Sprintf(`, phoneno = '%v'`, user.Phoneno)
	}

	if "" != user.Email {
		moreUpdate += fmt.Sprintf(`, email = '%v'`, user.Email)
	}

	sql := fmt.Sprintf(`
					UPDATE t_user
						 SET name = '%s'
						  %v #moreUpdate
					 WHERE user_seq = %v`, user.Name, moreUpdate, user.User_seq)
	_, err := common.DBconn().Exec(sql)
	if nil != err {
		return
	}

	return true
}

// 모든 기관에서 탈퇴하기
func (user *User) WithdrawAllInstitutionInUser(c *gin.Context, user_status uint) (succ bool) {
	succ = false
	sql := fmt.Sprintf(`
					SELECT IFNULL(MIN((SELECT COUNT(user_seq)
															 FROM t_user
															WHERE user_type LIKE '%%%v%%'
																AND institution_seq = user.institution_seq)),0) AS cnt
						FROM t_user user
					 WHERE user.email = '%v'
						 AND user.user_type LIKE '%%%v%%'`, DEF_USER_TYPE_ADMIN_SECRETARY, user.Data["email"],
		DEF_USER_TYPE_ADMIN_SECRETARY)
	row := common.DB_fetch_one(sql, nil)
	if common.ToInt(row["cnt"]) == 1 {
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_last_admin)
		return
	}

	sql = `UPDATE t_user
						 SET user_status  = ?,
								 withdraw_dttm = UNIX_TIMESTAMP()
					 WHERE email = ?`
	_, err := common.DBconn().Exec(sql, user_status, user.Data["email"])
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	withdrawUser := User{
		User_seq: user.User_seq,
	}

	withdrawUser.Load()
	unixTime, _ := goment.Unix(common.ToInt64(withdrawUser.Data["withdraw_dttm"]))
	t1, _ := goment.New(unixTime)
	t1.Add(5, "days")
	withdrawUser.Data["retractable_period"] = common.GetDateStr(t1)

	if !withdrawUser.SendMsg("", DEF_MSG_USER_IPSAP_WITHDRAW) {
		return
	}

	go withdrawUser.SendMsgToAdmin()

	succ = true
	return
}

func (withdrawUser *User) SendMsgToAdmin() {
	sql := fmt.Sprintf(`
				 SELECT (SELECT group_concat(user_seq)
									 FROM t_user
									WHERE user.institution_seq = institution_seq
										AND user_type LIKE '%%%v%%'
										AND user_seq <> user.user_seq
									GROUP BY institution_seq) as admin_user_seq,
								user.withdraw_dttm
					 FROM t_user user
					WHERE user.email = '%v'
						AND withdraw_dttm > 0`, DEF_USER_TYPE_ADMIN_SECRETARY, withdrawUser.Data["email"])
	rows := common.DB_fetch_all(sql, nil)
	for _, row := range rows {
		seqArr := strings.Split(common.ToStr(row["admin_user_seq"]), ",")
		for _, user_seq := range seqArr {
			sendAdminUser := User{}
			sendAdminUser.User_seq = common.ToUint(user_seq)
			if !sendAdminUser.Load() {
				continue
			}
			sendAdminUser.Data["target_name"] = withdrawUser.Data["name"]
			sendAdminUser.Data["target_id"] = withdrawUser.Data["email"]
			sendAdminUser.Data["target_type"] = withdrawUser.Data["user_type_str"]
			sendAdminUser.Data["retractable_period"] = withdrawUser.Data["retractable_period"]
			msgMgr := MessageMgr{
				User_info: sendAdminUser.Data,
				Msg_ID:    DEF_MSG_WITHDRAW_NOTICE,
			}
			msgMgr.SendMessage()
		}
	}
}

func GetServiceAdminInfo() (email string, telno string) {
	sql := `SELECT CONCAT(
	                LEFT(user.telno, 2),
	                '-',
	                MID(user.telno, 3, 3),
	                '-',
	                RIGHT(user.telno, 4)
	               ) AS telno,
	               user.email
					  FROM t_user user
					 WHERE user_auth = ?`
	row := common.DB_fetch_one(sql, nil, DEF_USER_AUTH_PLATFORM_SERVICE)
	email = common.ToStr(row["email"])
	telno = common.ToStr(row["telno"])
	return
}

func LoadServiceAdmin() (row map[string]interface{}) {
	sql := `SELECT user.email, user.phoneno,
								 user.agree_sms, user.agree_email,
								 user.user_seq, 0 AS institution_seq
					  FROM t_user user
					 WHERE user_auth = ?`
	row = common.DB_fetch_one(sql, nil, DEF_USER_AUTH_PLATFORM_SERVICE)

	// 21-07-12 : Test code
	if common.Config.Test.IsTestMode {
		row["agree_sms"] = "1"
		row["agree_email"] = "1"
		row["email"] = "alivins@sprintec.co.kr"
	} else {
		row["agree_sms"] = "1"
		row["agree_email"] = "1"
		row["email"] = "jonggil.lee@ipcx.net"
	}

	return
}

func (user *User) CancelWithdraw() (succ bool) {
	succ = false
	sql := `UPDATE t_user
						 SET user_status = ?, withdraw_dttm = 0
					 WHERE email = ?`
	_, err := common.DBconn().Exec(sql, DEF_USER_STATUS_FINISH, user.Email)
	if nil != err {
		log.Println(err)
		return
	}

	succ = true
	return
}
