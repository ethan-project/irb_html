
package model

import (
	"github.com/gin-gonic/gin"
	"database/sql"
  "ipsap/common"
	"strings"
	"log"
	"fmt"
)

type RequestService struct {
  Reqsvc_seq          uint
  Institution_seq     uint
  User_seq            uint
	Request_type				uint
	Request_status			uint
  Approved_comment    string
	Handle_type					uint
	LoginToken					map[string]interface{}
	Data								map[string]interface{}
}

func (reqSvc *RequestService)LoadList() (rows [] map[string]interface{}) {
	sql := `SELECT reqsvc.reqsvc_seq, reqsvc.institution_seq, reqsvc.user_seq,
								 reqsvc.request_type, reqsvc.request_status, reqsvc.reg_dttm,
								 instt.institution_code, instt.name_ko,
								 user.name as user_name, user.phoneno
						FROM t_request_service reqsvc
						LEFT OUTER JOIN t_institution instt ON (reqsvc.institution_seq = instt.institution_seq)
						LEFT OUTER JOIN t_user user ON(reqsvc.user_seq = user.user_seq)
					 WHERE 1 = 1`
  filter := func(row map[string]interface{}) {
		codeRT := Code {
			Type : "request_type",
			Id : common.ToUint(row["request_type"]),
		}
		row["request_type_str"] = codeRT.GetCodeStrFromTypeAndId()

		codeRS := Code {
			Type : "request_status",
			Id : common.ToUint(row["request_status"]),
		}
		row["request_status_str"] = codeRS.GetCodeStrFromTypeAndId()
	}

	rows = common.DB_fetch_all(sql, filter)
	return
}

func (reqSvc *RequestService)InsertRequestService (c *gin.Context, tx *sql.Tx) (succ bool) {
	if 0 == reqSvc.Institution_seq {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	if 0 == reqSvc.User_seq {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	sql := `INSERT INTO t_request_service(institution_seq,	user_seq, request_type, request_status, reg_dttm)
				  VALUES(?,?,?,?,UNIX_TIMESTAMP())`
	_, err := tx.Exec(sql, reqSvc.Institution_seq, reqSvc.User_seq, DEF_REQ_TYPE_SERVICE_REGISTER, DEF_REQ_STATUS_WAIT)
	if err != nil {
    log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
    return
  }

  return true
}

func (reqSvc *RequestService)Load() (succ bool) {
	moreSelect := ""
	switch reqSvc.Request_type {
		case DEF_REQUEST_TYPE_SERVICE_REGIST:
			moreSelect = `, instt.business_num,	instt.business_file_path,
			instt.business_file_org_name,	instt.payment_method,instt.invoice_flag`
		case DEF_REQUEST_TYPE_PAYMENT_CHANGE:
			moreSelect = ` , reqsvc.business_num,	reqsvc.business_file_path,
			reqsvc.business_file_org_name,	reqsvc.payment_method, reqsvc.invoice_flag`
		default :
		return false
	}

	sql := fmt.Sprintf(`
					SELECT instt.institution_seq,			instt.institution_type,				instt.institution_status,
								 instt.institution_code,		instt.name_ko,								instt.name_en,
								 instt.zipcode,							instt.addr1,									instt.addr2,
								 instt.homepage_url,				instt.logo_file_org_name,			instt.logo_file_path,
								 instt.judge_type,					instt.ia_evaluation_method,		instt.ia_base_score,
								 instt.ia_base_item,				instt.ia_final_director,			user.name as user_name,
								 user.name_en user_name_en, user.email,										user.dept,
								 user.position,							user.telno,										user.phoneno,
								 user.agree_email,					user.agree_sms,								user.agree_sms,
								 reqsvc.reqsvc_seq,					reqsvc.request_type,					reqsvc.request_status,
								 reqsvc.approved_comment,		reqsvc.approved_dttm					%v #moreSelct
						FROM t_request_service reqsvc, t_institution instt, t_user user
					 WHERE reqsvc.institution_seq = instt.institution_seq
						 AND reqsvc.user_seq = user.user_seq
						 AND reqsvc.reqsvc_seq = %d`, moreSelect, reqSvc.Reqsvc_seq)
 	filter :=  func(row map[string]interface{}) {
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

		if "" != row["logo_file_path"] && nil != reqSvc.LoginToken["tmp_key"] {
			// row["logo_file_path"] = common.EncryptToUrl([]byte(common.ToStr(reqSvc.LoginToken["tmp_key"])), common.ToStr(row["logo_file_path"]))
			row["logo_file_path"] = common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["logo_file_path"]))
			row["logo_file_src"] = common.MakeDownloadUrl(common.ToStr(row["logo_file_path"]))
		}

		if "" != row["business_file_path"]  && nil != reqSvc.LoginToken["tmp_key"] {
			// row["business_file_path"]	= common.EncryptToUrl([]byte(common.ToStr(reqSvc.LoginToken["tmp_key"])), common.ToStr(row["business_file_path"]))
			row["business_file_path"]	= common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["business_file_path"]))
			row["business_file_src"]	= common.MakeDownloadUrl(common.ToStr(row["business_file_path"]))
		}

		dcodeDept := CodeDynamic{
			Institution_seq : INSTITUTION_SHARE_CODE,
			DCode_type      : DCODE_TYPE_DEPT,
			Code            : common.ToUint(row["dept"]),
		}
		row["dept_str"] = dcodeDept.GetValueFromCode()

		dcodePosition := CodeDynamic{
			Institution_seq : INSTITUTION_SHARE_CODE,
			DCode_type      : DCODE_TYPE_POSITION,
			Code            : common.ToUint(row["position"]),
		}
		row["position_str"] = dcodePosition.GetValueFromCode()

		codeRT := Code {
		 Type : "request_type",
		 Id : common.ToUint(row["request_type"]),
		}
		row["request_type_str"] = codeRT.GetCodeStrFromTypeAndId()

		codeRS := Code {
		 Type : "request_status",
		 Id : common.ToUint(row["request_status"]),
		}
		row["request_status_str"] = codeRS.GetCodeStrFromTypeAndId()
	}
	reqSvc.Data = common.DB_fetch_one(sql, filter)
	if nil == reqSvc.Data {
		return false
	}
	succ = true
	return
}

func (reqSvc *RequestService)getUpdateRequestServiceQuery() (sql string){
	sql = fmt.Sprintf(`UPDATE t_request_service
												SET request_status = %d,		approved_comment = '%v',
														approved_user_seq = %v,	approved_dttm = UNIX_TIMESTAMP()
											WHERE reqsvc_seq = %d`,
											reqSvc.Request_status, reqSvc.Approved_comment,
											reqSvc.LoginToken["user_seq"],	reqSvc.Reqsvc_seq)
	return
}

func (reqSvc *RequestService)UpdateRequestService(c *gin.Context) (succ bool)  {
	succ = false
	dbConn	:= common.DBconn()
	tx, err	:= dbConn.Begin()
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	defer func() {
		tx.Rollback()
	}()

	switch reqSvc.Handle_type {
		case DEF_REQ_HANDLE_TYPE_APPROVED:	// 승인
			switch reqSvc.Request_type {
				case DEF_REQUEST_TYPE_SERVICE_REGIST:
						// 기관, 유저, 신청서 등록상태 update -> 이메일전송
					if !reqSvc.ApproveRequsetService(tx) {
						log.Println("신청서 기관, 유저 승인 실패!!!!!!")
						common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_invalide_params)
						return
					}
				case DEF_REQUEST_TYPE_PAYMENT_CHANGE:
						// 기관 결제정보변경 및 신청서 등록상태 update  -> 이메일전송
					if !reqSvc.ApproveRequsetPaymentChange(tx) {
						log.Println("기관 결제정보변경 승인 실패!!!!!!")
						common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_invalide_params)
						return
					}
				default :
				common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_invalide_params)
				return
			}
			reqSvc.Request_status = DEF_REQ_STATUS_FINISH
			sql := reqSvc.getUpdateRequestServiceQuery()
			_, err := tx.Exec(sql)
			if err != nil {
				log.Println(err)
				common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
				return
			}
			if !reqSvc.SendEmail(){
				common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_send_email)
				return
			}
		case DEF_REQ_HANDLE_TYPE_HOLD:	//처리보류
			reqSvc.Request_status = DEF_REQ_STATUS_HOLD
			sql := reqSvc.getUpdateRequestServiceQuery()
			_, err := tx.Exec(sql)
			if err != nil {
		    log.Println(err)
				common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		    return
		  }
		case DEF_REQ_HANDLE_TYPE_RESEND:	// 안내 재발송
			if reqSvc.Request_status != DEF_REQ_STATUS_FINISH { // 처리완료 상태가 아닐때!
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_not_request_status_finish)
				return
			}
			if !reqSvc.SendEmail() {
				common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_send_email)
				return
			}
			succ = true
			return
		default :
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
			return
	}

	err = tx.Commit()
	if nil != err {
		log.Println(err)
		return
	}

	succ = true
	return
}

func (reqSvc *RequestService)ApproveRequsetService(tx *sql.Tx) (succ bool) {
	succ = false
	if 0 == reqSvc.Institution_seq || 0 == reqSvc.User_seq {
		return
	}

	sql := `UPDATE t_institution
						 SET institution_status = ?
					 WHERE institution_seq = ?`
	_, err := tx.Exec(sql, DEF_INSTITUTION_STATUS_OK, reqSvc.Institution_seq)
	if err != nil {
		log.Println(err)
		return
	}

	sql = `UPDATE t_user
						SET user_status = ?
				  WHERE user_seq = ?`
	_, err = tx.Exec(sql, DEF_USER_STATUS_FINISH, reqSvc.User_seq)
	if err != nil {
		log.Println(err)
		return
	}

	succ = true
	return
}

func (reqSvc *RequestService)ApproveRequsetPaymentChange(tx *sql.Tx) (succ bool) {
	succ = false
	if 0 == reqSvc.Reqsvc_seq {
		return
	}

	sql := `UPDATE t_institution instt, t_request_service reqsvc
						 SET instt.business_num = reqsvc.business_num,
						 		 instt.business_file_path = reqsvc.business_file_path,
						 		 instt.business_file_org_name = reqsvc.business_file_org_name,
						 		 instt.payment_method = reqsvc.payment_method,
						 		 instt.invoice_flag = reqsvc.invoice_flag
					 WHERE instt.institution_seq = reqsvc.institution_seq
					 	 AND reqsvc.reqsvc_seq = ?`
	_, err := tx.Exec(sql, reqSvc.Reqsvc_seq)
	if err != nil {
		log.Println(err)
		return
	}

	succ = true
	return
}

func (reqSvc *RequestService)SendEmail() (succ bool) {
	succ = false
	user := User{}
	user.User_seq = reqSvc.User_seq
	if !user.Load(){
		return
	}

	msgMgr := MessageMgr	{
	 	User_info : user.Data,
	}

	switch reqSvc.Request_type {
		case DEF_REQUEST_TYPE_SERVICE_REGIST:
			msgMgr.Msg_ID = DEF_MSG_INSTITUTION_APPROVED
		case DEF_REQUEST_TYPE_PAYMENT_CHANGE:
			msgMgr.Msg_ID = DEF_MSG_INSTITUTION_CHANGE_APPROVED
		default :
			return
	}

	if !msgMgr.SendMessage() {
		return
	}

	succ = true
	return
}
