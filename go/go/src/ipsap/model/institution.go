package model

import (
	"github.com/nleeper/goment"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"database/sql"
  "ipsap/common"
	"strings"
  "fmt"
	// "os"
  "log"
)

type Institution struct {
	Institution_seq      		uint		`json:"-"`
	Product_seq   					uint  	`json:"-"`
	Institution_type     		uint  	`json:"institution_type" example:"1"`
	Institution_status   		uint  	`json:"-"`
	Institution_code     		string  `json:"institution_code" example:"TET"`
	Name_ko              		string  `json:"name_ko" example:"스프린텍"`
	Name_en              		string  `json:"name_en" example:"sprintec"`
	Zipcode              		string  `json:"zipcode" example:"우편번호"`
	Addr1                		string  `json:"addr1" example:"메인주소"`
	Addr2                		string  `json:"addr2" example:"상세주소"`
	Homepage_url         		string  `json:"homepage_url" example:"http://www.sprintec.co.kr"`
	Logo_file_path			 		string	`json:"-"`
	Logo_file_idx						string	`json:"-"`
	Logo_file_org_name	 		string	`json:"-"`
	Judge_type           		string  `json:"judge_type" example:"1,2"`
	Ia_evaluation_method 		uint8   `json:"ia_evaluation_method" example:"2"`
	Ia_base_score        		uint    `json:"ia_base_score" example:"50"`
	Ia_base_item         		uint    `json:"ia_base_item" example:"50"`
	Ia_final_director    		uint8   `json:"ia_final_director" example:"2"`
	Business_num         		string  `json:"business_num" example:"tester"`
	Business_file_path			string	`json:"-"`
	Business_file_idx				string	`json:"-"`
	Business_file_org_name	string	`json:"-"`
	Payment_method       		string  `json:"payment_method" example:"1,2,3"`
	Invoice_flag         		uint8   `json:"invoice_flag" example:"1"`
	Service_status					uint 		`json:"service_status"`
  LoginToken							map[string]interface{}	`json:"-"`
	Data										map[string]interface{}	`json:"-"`
}

func (ins *Institution) InsertInstitution(c *gin.Context, tx *sql.Tx) (succ bool) {
	if !ins.CheckValidation(c) {
		return
	}

	sql := `INSERT INTO t_institution(institution_type,	institution_status,		institution_code,
																		name_ko, 					name_en,							zipcode,
																		addr1, 						addr2,								homepage_url,
																		judge_type,				ia_evaluation_method,	ia_base_score,
																		ia_base_item,			ia_final_director,		business_num,
																		payment_method,		invoice_flag,
																		reg_dttm,					chg_dttm)
					VALUES(?, ?, ?,
								 ?, ?, ?,
								 ?, ?, ?,
								 ?, ?, ?,
								 ?, ?, ?,
								 ?, ?,
								 UNIX_TIMESTAMP(), UNIX_TIMESTAMP()
							  )
				 `
	result, err := tx.Exec(sql,
												 ins.Institution_type,	ins.Institution_status,		ins.Institution_code,
												 ins.Name_ko,					 	ins.Name_en,							ins.Zipcode,
												 ins.Addr1,							ins.Addr2,						 		ins.Homepage_url,
												 ins.Judge_type,				ins.Ia_evaluation_method,	ins.Ia_base_score,
												 ins.Ia_base_item,			ins.Ia_final_director,		ins.Business_num,
												 ins.Payment_method,		ins.Invoice_flag)
	if err != nil {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

 	no, err := result.LastInsertId()
 	if err != nil {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
   	return
 	}

	ins.Institution_seq = cast.ToUint(no)

	return true
}

func (ins *Institution) FileUpload(c *gin.Context, tx *sql.Tx) (succ bool) {
	if 0 == ins.Institution_seq {
		return
	}

	if !ins.LogoFileUpload(c, tx) {
		return
	}

	if !ins.BusinessFileUpload(c, tx, true) {
		return
	}

	return true
}

func (ins *Institution) GetDirPath(subDir string) (string) {
	return fmt.Sprintf("%v/%v/%v/", "institution", ins.Institution_seq, subDir)
}

func (ins *Institution) LogoFileUpload(c *gin.Context, tx *sql.Tx) (succ bool) {
	fup := common.FileUpload{}
	fup.Required = false
	fup.Param	= "logo_file"
	fup.New_file_name	= ins.Logo_file_idx
	sub_path := ins.GetDirPath("logo")
	// fup.Src = common.Config.Server.FileUploadPath + sub_path
	fup.Src = sub_path

	if !fup.UploadFile(c) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_upload_fail, "기관로고 업로드 실패")
		return
	}

	if "" == fup.New_file_name {
		return true
	}

	logo_file_path := sub_path + fup.New_file_name + 	fup.File_extension
	sql := `UPDATE t_institution
						 SET logo_file_org_name = ?,
						 		 logo_file_path = ?,
								 logo_file_idx = ?
					 WHERE institution_seq = ?`
  _, err := tx.Exec(sql, fup.File_name, logo_file_path, ins.Logo_file_idx, ins.Institution_seq)
	if err != nil {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

	ins.Logo_file_path 			= logo_file_path
	ins.Logo_file_org_name  = fup.File_name
	return true
}

func (ins *Institution) BusinessFileUpload(c *gin.Context, tx *sql.Tx, isNew bool) (succ bool) {
	fup := common.FileUpload{}
	fup.Required = false
	fup.Param	= "business_file"
	fup.New_file_name = ins.Business_file_idx
	sub_path := ins.GetDirPath("business")
	// fup.Src = common.Config.Server.FileUploadPath + sub_path
	fup.Src = sub_path

	if !fup.UploadFile(c) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_upload_fail, "사업자 등록증 업로드 실패")
		return
	}

	if "" == fup.New_file_name {
		return true
	}

	business_file_path := sub_path + fup.New_file_name + 	fup.File_extension
	if isNew {
		sql := `UPDATE t_institution
							 SET business_file_org_name = ?,
							 		 business_file_path = ?,
							 		 business_file_idx = ?
						 WHERE institution_seq = ?`
	  _, err := tx.Exec(sql, fup.File_name, business_file_path, ins.Business_file_idx, ins.Institution_seq)
		if err != nil {
			log.Println(err)
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
			return
		}
	} else {
		sql := `UPDATE t_institution
							 SET business_file_idx = ?
						 WHERE institution_seq = ?`
		_, err := tx.Exec(sql, ins.Business_file_idx, ins.Institution_seq)
		if err != nil {
			log.Println(err)
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
			return
		}
	}

	ins.Business_file_path 			= business_file_path
	ins.Business_file_org_name	= fup.File_name

	return true
}

func (ins *Institution) CheckValidation(c *gin.Context) (succ bool) {
	cd := Code {}
	cd.Type = "institution_type"
	cd.Id 	=	ins.Institution_type
  if !cd.CheckCodeError(c) {
    return
  }

	if !ins.CheckDuplicateInstitutionCode(c){
		return
	}

	if !common.CheckKr(ins.Name_ko) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "기관명(국문)값이 잘못 되었습니다.")
		return
	}

	if !common.CheckEn(ins.Name_en) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "기관명(영문)값이 잘못 되었습니다.")
		return
	}

	judgeTypeArr := strings.Split(ins.Judge_type, ",")
	for _, judge_type := range judgeTypeArr {
		cd.Type = "judge_type"
		cd.Id 	=	cast.ToUint(judge_type)
		if !cd.CheckCodeError(c) {
	    return
	  }
	}

	cd.Type = "evaluation_method"
	cd.Id 	=	cast.ToUint(ins.Ia_evaluation_method)
	if !cd.CheckCodeError(c) {
		return
	}

	cd.Type = "final_director"
	cd.Id 	=	cast.ToUint(ins.Ia_final_director)
	if !cd.CheckCodeError(c) {
		return
	}

	PaymentMethodArr := strings.Split(ins.Payment_method, ",")
	for _, payment_method := range PaymentMethodArr {
		cd.Type =	"payment_method"
		cd.Id 	=	cast.ToUint(payment_method)
		if !cd.CheckCodeError(c) {
			return
		}
	}

	if "" == ins.Zipcode {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "우편 번호가 입력되지 않았습니다.")
		return
	}

	if "" == ins.Addr1 {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "주소가 입력되지 않았습니다.")
		return
	}

	if "" == ins.Addr2 {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "주소가 입력되지 않았습니다.")
		return
	}

	cd.Type = "invoice_flag"
	cd.Id 	=	cast.ToUint(ins.Invoice_flag)
	if !cd.CheckCodeError(c) {
		return
	}

	ins.Institution_status = DEF_INSTITUTION_STATUS_FINISH
  return true
}

func (ins *Institution) CheckDuplicateInstitutionCode(c *gin.Context) (succ bool) {
	if "" == ins.Institution_code {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "기관 코드는 필수 입력 사항입니다.")
		return
	}
	// 정규식 체크 필요!! 알파벳 대문자 + 숫자 조합으로 3~5자리로 지정 가능합니다.
	sql := `SELECT institution_code FROM t_institution WHERE institution_code = ?`
	row := common.DB_fetch_one(sql, nil, ins.Institution_code)
	if nil != row {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_dup_name, "중복된 기관코드 입니다.")
		return
	}
	return true
}

func (ins *Institution) GetInstitutionQueryAndFilter2(moreCondition string, infoFlag bool)(sql string, filter func(map[string]interface{})) {
	sql = fmt.Sprintf(`
						SELECT ins.institution_seq,			ins.institution_type,					ins.institution_status,
									 ins.institution_code,  	ins.name_ko, 									ins.name_en,
									 ins.zipcode,							ins.addr1,										ins.addr2,
									 ins.judge_type,					ins.ia_evaluation_method,			ins.ia_base_score,
									 ins.homepage_url,				ins.logo_file_org_name,				ins.logo_file_path,
									 ins.ia_base_item,				ins.ia_final_director,				ins.business_num,
									 ins.business_file_path,	ins.business_file_org_name,		ins.payment_method,
									 ins.invoice_flag,				ins.reg_dttm,									ins.chg_dttm,
									 IF(COUNT(user.user_seq) <= 1 , user.name, CONCAT(MAX(user.name)," 외 ", COUNT(user.user_seq) - 1, "명")) as user_name,
									 user.telno, ins.payment_setting, ins.service_status, ins.expiration_date,
									 IFNULL((SELECT MAX(auth_date)
									 					 FROM t_orders ord
									 					WHERE ord.institution_seq = ins.institution_seq
															AND ord.order_status = 1
															AND ord.order_type IN (1, 2)), '-') AS last_payment_date,
									 membership_payment_date
					    FROM t_institution ins, t_user user
					   WHERE ins.institution_seq = user.institution_seq
							 AND user.user_type LIKE '%%%v%%'
							  %v	#moreCondition
					GROUP BY ins.institution_seq`, DEF_USER_TYPE_ADMIN_SECRETARY, moreCondition)

	filter =  func(row map[string]interface{}) {
		if infoFlag {
			sql2 := fmt.Sprintf(`
				SELECT user.name user_name,	user.name_en user_name_en,	user.telno,
							 user.email,					user.phoneno,								user.dept,
							 user.position,			  user.agree_sms,							user.agree_email,
							 user.agree_pri_open, user.user_seq,							user.institution_seq
				  FROM t_user user
				 WHERE user.user_type LIKE '%%%v%%'
					 AND user.institution_seq = %v`, DEF_USER_TYPE_ADMIN_SECRETARY, row["institution_seq"])
			filter2 :=  func(row2 map[string]interface{}) {
				dcodeDept := CodeDynamic{
					Institution_seq : INSTITUTION_SHARE_CODE,
					DCode_type      : DCODE_TYPE_DEPT,
					Code            : common.ToUint(row2["dept"]),
				}
				row2["dept_str"] = dcodeDept.GetValueFromCode()

				dcodePosition := CodeDynamic{
					Institution_seq : INSTITUTION_SHARE_CODE,
					DCode_type      : DCODE_TYPE_POSITION,
					Code            : common.ToUint(row2["position"]),
				}
				row2["position_str"] = dcodePosition.GetValueFromCode()
			}
			row["user_list"] = common.DB_fetch_all(sql2, filter2)
		}

		codeSS := Code {
			Type : "service_status",
			Id : common.ToUint(row["service_status"]),
		}
		row["service_status_str"] = codeSS.GetCodeStrFromTypeAndId()

		codeIT := Code {
      Type : "institution_type",
      Id : common.ToUint(row["institution_type"]),
    }
    row["institution_type_str"] = codeIT.GetCodeStrFromTypeAndId()

		codeIS := Code {
      Type : "institution_status",
      Id : common.ToUint(row["institution_status"]),
    }
    row["institution_status_str"] = codeIS.GetCodeStrFromTypeAndId()

		codeES := Code {
      Type : "evaluation_method",
      Id : common.ToUint(row["ia_evaluation_method"]),
    }
    row["ia_evaluation_method_str"] = codeES.GetCodeStrFromTypeAndId()

		codeFD := Code {
			Type : "final_director",
			Id : common.ToUint(row["ia_final_director"]),
		}
		row["ia_final_director_str"] = codeFD.GetCodeStrFromTypeAndId()

		codeIF := Code {
			Type : "invoice_flag",
			Id : common.ToUint(row["invoice_flag"]),
		}
		row["invoice_flag_str"] = codeIF.GetCodeStrFromTypeAndId()

		var tmpArr []string
		judgeTypeArr := strings.Split(common.ToStr(row["judge_type"]), ",")
		for _, judeType := range judgeTypeArr {
			codeJT := Code {
				Type : "judge_type",
				Id : common.ToUint(judeType),
			}
			tmpArr = append(tmpArr, codeJT.GetCodeStrFromTypeAndId())
		}
		row["judge_type_str"] = strings.Join(tmpArr, ",")
		tmpArr = nil

		paymentMethodeArr := strings.Split(common.ToStr(row["payment_method"]), ",")
		for _, payment_method := range paymentMethodeArr {
			codePM := Code {
				Type : "payment_method",
				Id : common.ToUint(payment_method),
			}
			tmpArr = append(tmpArr, codePM.GetCodeStrFromTypeAndId())
		}
		row["payment_method_str"] = strings.Join(tmpArr, ",")

		row["logo_file_src"] = ""
		row["business_file_src"] = ""

		if "" != row["logo_file_path"] {
			// row["logo_file_path"] = common.EncryptToUrl([]byte(common.ToStr(ins.LoginToken["tmp_key"])), common.ToStr(row["logo_file_path"]))
			row["logo_file_path"] = common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["logo_file_path"]))
			row["logo_file_src"] = common.MakeDownloadUrl(common.ToStr(row["logo_file_path"]))
		}

		if "" != row["business_file_path"] {
			// row["business_file_path"]	= common.EncryptToUrl([]byte(common.ToStr(ins.LoginToken["tmp_key"])), common.ToStr(row["business_file_path"]))
			row["business_file_path"]	= common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["business_file_path"]))
			row["business_file_src"]	= common.MakeDownloadUrl(common.ToStr(row["business_file_path"]))
		}
	}
	return
}

func (ins *Institution) GetInstitutionQueryAndFilter(moreCondition string)(sql string, filter func(map[string]interface{})) {
	sql = fmt.Sprintf(`
						SELECT ins.institution_seq,			ins.institution_type,					ins.institution_status,
									 ins.institution_code,  	ins.name_ko, 									ins.name_en,
									 ins.zipcode,							ins.addr1,										ins.addr2,
									 ins.judge_type,					ins.ia_evaluation_method,			ins.ia_base_score,
									 ins.homepage_url,				ins.logo_file_org_name,				ins.logo_file_path,
									 ins.ia_base_item,				ins.ia_final_director,				ins.business_num,
									 ins.business_file_path,	ins.business_file_org_name,		ins.payment_method,
									 ins.invoice_flag,				ins.expiration_date,					ins.reg_dttm,
									 ins.chg_dttm,						ins.service_status
					    FROM t_institution ins
					   WHERE 1 = 1
					 	    %v	#moreCondition`, moreCondition)

	filter =  func(row map[string]interface{}) {
		codeIT := Code {
      Type : "institution_type",
      Id : common.ToUint(row["institution_type"]),
    }
    row["institution_type_str"] = codeIT.GetCodeStrFromTypeAndId()

		codeIS := Code {
      Type : "institution_status",
      Id : common.ToUint(row["institution_status"]),
    }
    row["institution_status_str"] = codeIS.GetCodeStrFromTypeAndId()

		codeES := Code {
      Type : "evaluation_method",
      Id : common.ToUint(row["ia_evaluation_method"]),
    }
    row["ia_evaluation_method_str"] = codeES.GetCodeStrFromTypeAndId()

		codeFD := Code {
			Type : "final_director",
			Id : common.ToUint(row["ia_final_director"]),
		}
		row["ia_final_director_str"] = codeFD.GetCodeStrFromTypeAndId()

		codeIF := Code {
			Type : "invoice_flag",
			Id : common.ToUint(row["invoice_flag"]),
		}
		row["invoice_flag_str"] = codeIF.GetCodeStrFromTypeAndId()

		var tmpArr []string
		judgeTypeArr := strings.Split(common.ToStr(row["judge_type"]), ",")
		for _, judeType := range judgeTypeArr {
			codeJT := Code {
				Type : "judge_type",
				Id : common.ToUint(judeType),
			}
			tmpArr = append(tmpArr, codeJT.GetCodeStrFromTypeAndId())
		}
		row["judge_type_str"] = strings.Join(tmpArr, ",")
		tmpArr = nil

		paymentMethodeArr := strings.Split(common.ToStr(row["payment_method"]), ",")
		for _, payment_method := range paymentMethodeArr {
			codePM := Code {
				Type : "payment_method",
				Id : common.ToUint(payment_method),
			}
			tmpArr = append(tmpArr, codePM.GetCodeStrFromTypeAndId())
		}
		row["payment_method_str"] = strings.Join(tmpArr, ",")

		row["logo_file_src"] = ""
		row["business_file_src"] = ""

		if "" != row["logo_file_path"] {
			// row["logo_file_path"] = common.EncryptToUrl([]byte(common.ToStr(ins.LoginToken["tmp_key"])), common.ToStr(row["logo_file_path"]))
			row["logo_file_path"] = common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["logo_file_path"]))
			row["logo_file_src"] = common.MakeDownloadUrl(common.ToStr(row["logo_file_path"]))
		}

		if "" != row["business_file_path"] {
			// row["business_file_path"]	= common.EncryptToUrl([]byte(common.ToStr(ins.LoginToken["tmp_key"])), common.ToStr(row["business_file_path"]))
			row["business_file_path"]	= common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["business_file_path"]))
			row["business_file_src"]	= common.MakeDownloadUrl(common.ToStr(row["logo_file_path"]))
		}
	}
	return
}

func (ins *Institution) Load() (succ bool) {
  if 0 == ins.Institution_seq {
    return
  }

  if nil != ins.Data {
    return true
  }

	moreCondition := fmt.Sprintf(` AND ins.institution_seq = %d`, ins.Institution_seq)
	sql, filter := ins.GetInstitutionQueryAndFilter(moreCondition)

  ins.Data = common.DB_fetch_one(sql, filter)

  return nil != ins.Data
}

func (ins *Institution) GetInstitutionPaymentMsgData () {
	sql := `SELECT ins.name_ko, ins.logo_file_path
						FROM t_institution ins
					 WHERE institution_seq = ?`
	 ins.Data = common.DB_fetch_one(sql, nil, ins.Institution_seq)
	 return
}

func (ins *Institution) GetInstitutionList(infoFlag bool) (ret interface{}) {
	moreCondition := ""
	if ins.Institution_seq != 0 {
		moreCondition = fmt.Sprintf( ` AND ins.institution_seq = %d`, ins.Institution_seq)
	}
	sql, filter := ins.GetInstitutionQueryAndFilter2(moreCondition, infoFlag)
	if infoFlag {
		ret = common.DB_fetch_one(sql, filter)
	} else {
		ret = common.DB_fetch_all(sql, filter)
	}
	return
}

// 기관정보중에 결제정보 제외하고 수정!
func (ins *Institution) UpdateInstitution(c *gin.Context, tx *sql.Tx) (succ bool) {
	sql := `SELECT logo_file_idx + 1 as logo_file_idx  FROM t_institution WHERE institution_seq = ?`
	row := common.DB_Tx_fetch_one(tx, sql, nil, ins.Institution_seq)
	ins.Logo_file_idx = common.ToStr(row["logo_file_idx"])
	if !ins.LogoFileUpload(c, tx) {
		return
	}

	sql	= `UPDATE t_institution
						SET institution_code = ?,	name_ko = ?,			name_en = ?,
								zipcode = ?,					addr1 = ?,				addr2 = ?,
								homepage_url = ?,			judge_type = ?,		ia_evaluation_method = ?,
								ia_base_score = ?,		ia_base_item = ?,	ia_final_director = ?,
								institution_type = ?, chg_dttm = UNIX_TIMESTAMP()
					WHERE institution_seq = ?`
  _, err := tx.Exec(sql,
										ins.Institution_code,	ins.Name_ko,			ins.Name_en,
										ins.Zipcode,					ins.Addr1,				ins.Addr2,
										ins.Homepage_url,			ins.Judge_type,		ins.Ia_evaluation_method,
										ins.Ia_base_score,		ins.Ia_base_item,	ins.Ia_final_director,
										ins.Institution_type, ins.Institution_seq)
	if err != nil {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

	return true
}

// 기관 모든정보 수정가능
func (ins *Institution) PlatformUpdateInstitution(c *gin.Context, tx *sql.Tx) (succ bool) {
	sql := `SELECT logo_file_idx + 1 as logo_file_idx, business_file_idx + 1 as business_file_idx FROM t_institution WHERE institution_seq = ?`
	row := common.DB_Tx_fetch_one(tx, sql, nil, ins.Institution_seq)
	ins.Logo_file_idx = common.ToStr(row["logo_file_idx"])
	if !ins.LogoFileUpload(c, tx) {
		return
	}

	ins.Business_file_idx = common.ToStr(row["business_file_idx"])
	if !ins.BusinessFileUpload(c, tx, false) {
		return
	}

	sql	= `UPDATE t_institution
						SET institution_code = ?,	name_ko = ?,			name_en = ?,
								zipcode = ?,					addr1 = ?,				addr2 = ?,
								homepage_url = ?,			judge_type = ?,		ia_evaluation_method = ?,
								ia_base_score = ?,		ia_base_item = ?,	ia_final_director = ?,
								service_status = ?,		chg_dttm = UNIX_TIMESTAMP()
					WHERE institution_seq = ?`
  _, err := tx.Exec(sql,
										ins.Institution_code,	ins.Name_ko,			ins.Name_en,
										ins.Zipcode,					ins.Addr1,				ins.Addr2,
										ins.Homepage_url,			ins.Judge_type,		ins.Ia_evaluation_method,
										ins.Ia_base_score,		ins.Ia_base_item,	ins.Ia_final_director,
										ins.Service_status,		ins.Institution_seq)
	if err != nil {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

	return true
}

func (ins *Institution) InsertRequestPaymentChange(c *gin.Context, tx *sql.Tx, user_seq uint64) (succ bool) {
	sql := `SELECT business_file_idx + 1 as business_file_idx,
								 business_file_path, business_file_org_name
						FROM t_institution
					 WHERE institution_seq = ?`
	row := common.DB_fetch_one(sql, nil, ins.Institution_seq)
	ins.Business_file_idx = common.ToStr(row["business_file_idx"])
	if !ins.BusinessFileUpload(c, tx, false) {
		return
	}

	if "" == ins.Business_file_path || "" == ins.Business_file_org_name {
		ins.Business_file_path = common.ToStr(row["business_file_path"])
		ins.Business_file_org_name = common.ToStr(row["business_file_org_name"])
	}

	sql = `INSERT INTO t_request_service(institution_seq,					user_seq,				request_type,
																			 request_status,					business_num,		business_file_path,
																			 business_file_org_name,	payment_method,	invoice_flag,
																		 	 reg_dttm)
					VALUES(?, ?, ?,
								 ?, ?, ?,
								 ?,	?, ?,
							 	 UNIX_TIMESTAMP())
				`
	 _, err := tx.Exec(sql,
										 ins.Institution_seq,					user_seq,						DEF_REQ_TYPE_PAYMENT_CHANGE,
										 DEF_REQ_STATUS_WAIT,					ins.Business_num,		ins.Business_file_path,
										 ins.Business_file_org_name,	ins.Payment_method,	ins.Invoice_flag)
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

	return true
}

func (ins *Institution) UpdateRequestPaymentChange(c *gin.Context, tx *sql.Tx, reqsvc_seq uint64) (succ bool) {
	succ = false
	if 0 == ins.Institution_seq {
		return
	}

	sql := `SELECT business_file_idx + 1 as business_file_idx,
								 business_file_path, business_file_org_name
						FROM t_institution
					 WHERE institution_seq = ?`
	row := common.DB_fetch_one(sql, nil, ins.Institution_seq)
	ins.Business_file_idx = common.ToStr(row["business_file_idx"])
	if !ins.BusinessFileUpload(c, tx, false) {
		return
	}

	if "" == ins.Business_file_path || "" == ins.Business_file_org_name {
		ins.Business_file_path = common.ToStr(row["business_file_path"])
		ins.Business_file_org_name = common.ToStr(row["business_file_org_name"])
	}

	sql = `UPDATE t_request_service
				 		SET business_num = ?,		business_file_path = ?, business_file_org_name = ?,
				 		 		payment_method = ?,	invoice_flag = ?
					WHERE reqsvc_seq = ?`
	 _, err := tx.Exec(sql,
										 ins.Business_num,		ins.Business_file_path,	ins.Business_file_org_name,
										 ins.Payment_method,	ins.Invoice_flag,
										 reqsvc_seq)
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

	return true
}

func (ins *Institution) GetInstitutionApplicationCnt() (appCnt uint){
	// 당월의 플랜별 이용건수를 초과할수 없음 (IACUC, IRB, IBC 합산)
	// 건수 counting: 승인신청서 작성 및 자가점검을 완료하여 신청서를
	// 행정간사에 제출완료한 경우 1건으로 기산
	sql := `SELECT COUNT(application_seq) as app_cnt
						FROM t_application
					 WHERE institution_seq = ?
						 AND application_type = ?
						 AND application_step >= ?
					   AND reg_dttm > UNIX_TIMESTAMP(LAST_DAY(NOW() - interval 1 month))
					 	 AND reg_dttm <= UNIX_TIMESTAMP(LAST_DAY(NOW()))`
	row := common.DB_fetch_one(sql, nil, ins.Institution_seq, DEF_APP_TYPE_NEW, DEF_APP_STEP_CHECKING)
	if nil != row {
		appCnt = common.ToUint(row["app_cnt"])
	}
	return
}

func (ins *Institution) PaymentSettingChange() (succ bool) {
	succ = false
	sql := `SELECT bid, payment_setting
						FROM t_institution
					 WHERE institution_seq = ?`
	row := common.DB_fetch_one(sql, nil, ins.Institution_seq)

	paymentSetting := common.ToUint(row["payment_setting"])
	chgPaymentSetting := 2
	switch paymentSetting {
		case 1:	// 수동
			if common.ToStr(row["payment_setting"]) == "" {
				return
			}
		case 2: // 자동
			chgPaymentSetting = 1;
		default :
		 return
	}

	sql = `UPDATE t_institution
						SET payment_setting = ?
					WHERE institution_seq = ?`
 	_, err := common.DBconn().Exec(sql, chgPaymentSetting, ins.Institution_seq)
	if nil != err {
		return
	}

	succ = true
	return
}

func (ins *Institution) PlanChange(chg_user_seq uint) (succ bool) {
	succ = false

	sql := `UPDATE t_institution
						 SET product_seq = ?,
						 		 chg_dttm = UNIX_TIMESTAMP()
					 WHERE institution_seq = ?`
	_, err := common.DBconn().Exec(sql, ins.Product_seq, ins.Institution_seq)
	if nil != err {
		return
	}

	succ = true
	return
}

func (ins *Institution) CheckPossibleAppSubmit() (succ bool) {
	succ = false
	if ins.Institution_seq > 0 {
		sql := `SELECT IFNULL(IF(instt.usage_limit > 0, instt.usage_limit,
												 		(SELECT usage_limit
															 FROM t_products prod, t_membership_plan plan
															WHERE prod.plan_seq = plan.plan_seq
															  AND prod.product_seq = instt.product_seq)),0) as usage_limit
							FROM t_institution instt
						 WHERE instt.institution_seq = ?`
		row := common.DB_fetch_one(sql, nil, ins.Institution_seq)
		curAppCnt := ins.GetInstitutionApplicationCnt()
		limitAppCnt := common.ToUint(row["usage_limit"])
		if limitAppCnt > curAppCnt {
			succ = true
		}
	}
	return
}

func (ins *Institution) MambershipCancel() (succ bool) {
	succ = false
	sql := `UPDATE t_institution
						 SET expiration_date = "",
						 		 product_seq = 0,
						 		 payment_setting = 1,
								 service_status = ?,
								 bid = "",
								 membership_fee_status = 0,
								 stop_dttm = UNIX_TIMESTAMP()
					 WHERE institution_seq = ?`
	_, err := common.DBconn().Exec(sql, DEF_SERVICE_STATUS_STOPPED, ins.Institution_seq)
	if nil != err {
		return
	}

	succ = true
	return
}

func (ins *Institution) InstitutionPurchasedList() (list interface{}) {
	sql := `SELECT ord.order_seq, ord.tid, ord.pname,
								 ord.amount, ord.auth_date, ord.the_date,
								 ord.pay_method, ord.order_type, (SELECT category FROM t_products WHERE product_seq = ord.product_seq) as category,
								 IF(ord.order_type NOT IN (1, 2),
										IF((SELECT COUNT(*)
													FROM t_orders ord2
												 WHERE ord2.order_type = 4
													 AND ord2.moid = ord.moid
													 AND ord.auth_date <= ord2.auth_date) <> 0,
											CONCAT(ord.moid, "-",(SELECT COUNT(*)
																							FROM t_orders ord2
																						 WHERE ord2.order_type = 4
																							 AND ord2.moid = ord.moid
																							 AND ord.auth_date >= ord2.auth_date)),
											ord.moid
											), ord.moid) AS moid
  					FROM t_orders ord, t_institution instt, t_products prod
					 WHERE ord.institution_seq = instt.institution_seq
					   AND prod.product_seq = instt.product_seq
					   AND instt.institution_seq = ?
					   AND ord.order_status = 1
						 AND ord.product_seq IN((SELECT	prod2.product_seq
						 													 FROM t_products prod2
						 												 	WHERE prod2.plan_seq = (SELECT plan_seq
																												 				FROM t_products
																												 			 WHERE product_seq = instt.product_seq)))`
	list = common.DB_fetch_all(sql, nil, ins.Institution_seq)
	return
}

func (ins *Institution) GetInstitutionAdminCount() (count uint) {
	sql := fmt.Sprintf(`
					SELECT count(user_seq) as cnt
						FROM t_user user
					 WHERE user.institution_seq = %v
						 AND (user_type LIKE '%%%v%%')`,
						 ins.Institution_seq,
						 DEF_USER_TYPE_ADMIN_SECRETARY)
	row := common.DB_fetch_one(sql, nil)
	count = common.ToUint(row["cnt"])
	return
}

func (ins *Institution) UpdateInstitutionRegularPaymentResult() (succ bool) {
	succ = false
	sql := `UPDATE t_institution
		  			 SET expiration_date = date_format(LAST_DAY(NOW()), '%Y%m%d'), #이번달 말일까지 만료일 업데이트
						 		 free_start_date = "",
								 free_end_date = "",
 								 usage_limit = 0
					 WHERE institution_seq = ?`
	_, err := common.DBconn().Exec(sql, ins.Institution_seq)
	if nil != err {
		log.Println(err)
		return
	}

	succ = true
	return
}

func (ins *Institution) StopInstitution() (succ bool) {
	succ = false
	sql := `UPDATE t_institution
						 SET stop_dttm = UNIX_TIMESTAMP(),
						 		 service_status = ?
					 WHERE institution_seq = ?`
	_, err := common.DBconn().Exec(sql, DEF_SERVICE_STATUS_STOPPED, ins.Institution_seq)
	if nil != err {
		log.Println(err)
		return
	}

	succ = true
	return
}

func (ins *Institution) GetFreeMembershipPossibleList() (list interface{}){
	moreCondition := `AND ins.membership_fee_status > 0`
	sql, filter := ins.GetInstitutionQueryAndFilter2(moreCondition, false)
  list = common.DB_fetch_all(sql, filter)
  return
}

func (ins *Institution) ServiceStatusChange() (succ bool) {
	succ = false
	sql := `UPDATE t_institution
					   SET service_status = ?
					 WHERE institution_seq = ?`
  _, err := common.DBconn().Exec(sql, ins.Service_status, ins.Institution_seq)
	if nil != err {
		log.Println(err)
		return
	}

	succ = true
	return
}

func (ins *Institution) DeleteInstitution() (succ bool) {
	succ = false
	sql := `UPDATE t_institution
						 SET submit_withdraw_dttm = UNIX_TIMESTAMP() # 탈퇴 신청일
					 WHERE institution_seq = ?`
	_, err := common.DBconn().Exec(sql, ins.Institution_seq)
	if nil != err {
		log.Println(err)
		return
	}

	go ins.InstitutionWithdrawExpectedSendMsg()

	succ = true
	return
}

func (ins *Institution) RemoveInstitution() (succ bool) {
	succ = false

	// 메일을 먼저 보낸뒤 삭제 함
	ins.RemoveInstitutionSendMsg()

	sql := `INSERT t_institution_rm
					SELECT *, UNIX_TIMESTAMP()
					 	FROM t_institution
					 WHERE institution_seq = ?`
	_, err := common.DBconn().Exec(sql, ins.Institution_seq)
	if err != nil {
		log.Println(err)
    return
  }

	sql = `INSERT t_user_rm
				 SELECT *, UNIX_TIMESTAMP()
				 	 FROM t_user
				  WHERE institution_seq = ?`
	_, err = common.DBconn().Exec(sql, ins.Institution_seq)
	if err != nil {
		log.Println(err)
    return
  }

	sql = `DELETE FROM t_institution WHERE institution_seq = ?`
	_, err = common.DBconn().Exec(sql, ins.Institution_seq)
	if err != nil {
		log.Println(err)
		return
	}

	sql = `DELETE FROM t_user WHERE institution_seq = ?`
	_, err = common.DBconn().Exec(sql, ins.Institution_seq)
	if err != nil {
		log.Println(err)
		return
	}

	succ = true
	return
}

func (ins *Institution) RemoveInstitutionSendMsg() {
	sql := `SELECT user_seq
						FROM t_user
					 WHERE institution_seq = ?`
  rows := common.DB_fetch_all(sql, nil, ins.Institution_seq)
	now, _ := goment.New()
	removeDate := common.GetDateStr(now)

	for _, row := range rows {
		user := User{
			User_seq : common.ToUint(row["user_seq"]),
		}

		if !user.Load(){
			continue
		}

		user.Data["remove_dttm"] = removeDate
		msgMgr := MessageMgr {
			User_info : user.Data,
			Msg_ID : DEF_MSG_IPSAP_WITHDRAW_FINISHED,
		}
		msgMgr.SendMessage()
	}

	sql = `SELECT submit_withdraw_dtt
					 FROM t_institution
					WHERE institution_seq = ?`
  row := common.DB_fetch_one(sql, nil, ins.Institution_seq)
	unixTime, _ := goment.Unix(common.ToInt64(row["submit_withdraw_dttm"]))
	t, _ := goment.New(unixTime)
	submitWithdrawDttm := common.GetDateStr(t)

	adminData := LoadServiceAdmin()
	adminData["institution_seq"] = ins.Institution_seq
	adminData["submit_withdraw_dttm"] = submitWithdrawDttm
	adminData["remove_dttm"] = removeDate


	msgMgr2 := MessageMgr {
		User_info : adminData,
		Msg_ID : DEF_MSG_IPSAP_WITHDRAW_FINISHED_NOTI,
	}

	msgMgr2.SendMessage()

}

func (ins *Institution) InstitutionWithdrawExpectedSendMsg() {
	sql := fmt.Sprintf(`
				 SELECT group_concat(user.user_seq) user_arr, instt.submit_withdraw_dttm
					 FROM t_user user, t_institution instt
					WHERE user.institution_seq = %v
						AND user.user_type LIKE '%%%v%%'`,
						 ins.Institution_seq, DEF_USER_TYPE_ADMIN_SECRETARY)
	row := common.DB_fetch_one(sql, nil)
	unixTime, _ := goment.Unix(common.ToInt64(row["submit_withdraw_dttm"]))
	t, _ := goment.New(unixTime)
	submitWithdrawDttm := common.GetDateStr(t)

	userSeqArr := strings.Split(common.ToStr(row["user_arr"]), ",")
	for _, userSeq := range userSeqArr {
		user := User {
			User_seq : common.ToUint(userSeq),
		}

		if !user.Load(){
			continue;
		}

		user.Data["submit_withdraw_dttm"] = submitWithdrawDttm
		msgMgr := MessageMgr {
			User_info : user.Data,
			Msg_ID : DEF_MSG_IPSAP_WITHDRAW_REQUEST,
		}

		msgMgr.SendMessage()
	}

	adminData := LoadServiceAdmin()
	adminData["institution_seq"] = ins.Institution_seq
	adminData["target_id"] = ins.LoginToken["email"]
	adminData["admin_user_seq"] = ins.LoginToken["user_seq"]
	adminData["submit_withdraw_dttm"] = submitWithdrawDttm

	msgMgr2 := MessageMgr {
		User_info : adminData,
		Msg_ID : DEF_MSG_IPSAP_WITHDRAW_REQUEST_NOTI,
	}

	msgMgr2.SendMessage()

	return
}
