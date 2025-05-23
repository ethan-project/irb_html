
package model

import (
)

type iacucCopyForIbc struct {
	Application_seq	uint	`json:"application_seq" example:"1"`
}

type requestServiceModel struct {
	user				User
	institution	Institution
}

type institutionPatchModel struct {
	Institution_type     		uint  	`json:"institution_type" example:"1"`
	Institution_code     		string  `json:"institution_code" example:"TET"`
	Name_ko              		string  `json:"name_ko" example:"스프린텍"`
	Name_en              		string  `json:"name_en" example:"sprintec"`
	Zipcode              		string  `json:"zipcode" example:"우편번호"`
	Addr1                		string  `json:"addr1" example:"메인주소"`
	Addr2                		string  `json:"addr2" example:"상세주소"`
	Homepage_url         		string  `json:"homepage_url" example:"http://www.sprintec.co.kr"`
	Judge_type           		string  `json:"judge_type" example:"1,2"`
	Ia_evaluation_method 		uint8   `json:"ia_evaluation_method" example:"2"`
	Ia_base_score        		uint    `json:"ia_base_score" example:"50"`
	Ia_base_item         		uint    `json:"ia_base_item" example:"50"`
	Ia_final_director    		uint8   `json:"ia_final_director" example:"2"`
}

type institutionPaymentPatchModel struct {
	Business_num         		string  `json:"business_num" example:"사업자등록번호:111-231-12"`
	Payment_method       		string  `json:"payment_method" example:"1,2,3"`
	Invoice_flag         		uint8   `json:"invoice_flag" example:"1"`
}

type adminPatchUserInfoModel struct {
	User_seq				uint		`json:"user_seq" example:"1"`
	Name						string	`json:"name" example:"홍길동"`
	User_type				string	`json:"user_type" example:"1"`
	Dept						string	`json:"dept" example:"소속부서"`
	Position				string	`json:"position" example:"직급"`
	Phoneno					string	`json:"phoneno" example:"021231234"`
	Major_field			string	`json:"major_field" example:"수의학"`
	Edu_date				string	`json:"edu_date" example:"20210101"`
	Edu_institution	string	`json:"edu_institution" example:"건국대학교"`
	Edu_course_num	string	`json:"edu_course_num" example:"0"`
	Agree_email			uint		`json:"agree_email" example:"1"`
	Agree_sms				uint		`json:"agree_sms" example:"1"`
	Agree_pri_open	uint		`json:"agree_pri_open" example:"1"`
}

type moveInstitutionModel struct {
	User_seq	uint	`json:"user_seq" example:"1"`
	Tmp_key		string	`json:"tmp_key" example:"12345678901234561234567890123456"`
}

type myInstitutionModel struct {
	Tmp_key		string	`json:"tmp_key" example:"12345678901234561234567890123456"`
}

type approvedCommentModel struct {
	Approved_comment	string  `json:"approved_comment" example:"승인함"`
}

/**********************************************/
// 사용자
/**********************************************/
type userRegisterModel struct {
	Institution_seq	uint    `json:"institution_seq" example:"1"`
	Dept						string	`json:"dept" example:"소속부서"`
	Position				string	`json:"position" example:"직급"`
	Major_field			string	`json:"major_field" example:"수의학"`
	Name						string	`json:"name" example:"홍길동"`
	Pwd							string	`json:"pwd" example:"tmvmfls!00"`
	Phoneno					string	`json:"phoneno" example:"01012341234"`
	Agree_email			uint		`json:"agree_email" example:"1"`
	Agree_sms				uint		`json:"agree_sms" example:"1"`
	Agree_pri_open	uint		`json:"agree_pri_open" example:"1"`
	Edu_date				string	`json:"edu_date" example:"20210101"`
	Edu_institution	string	`json:"edu_institution" example:"건국대학교"`
	Edu_course_num	string	`json:"edu_course_num" example:"0"`
	Tmp_key					string	`json:"tmp_key" example:"12345678901234561234567890123456"`
}

type userIdFindModel struct {
	Name		string	`json:"name" example:"홍길동"`
	Phoneno	string	`json:"phoneno" example:"01012341234"`
}

type userPasswordFindModel struct {
	Email    string `json:"email" example:"test@test.com"`
	Name		string	`json:"name" example:"홍길동"`
	Phoneno	string	`json:"phoneno" example:"01012341234"`
}

type userBasicModel struct {
	User_type	string	`json:"user_type" example:"1"`
	Email			string	`json:"email" example:"test@test.com"`
	Phoneno		string	`json:"phoneno" example:"01012341234"`
}

type userBatchModel struct {
	userArr []userBasicModel `json:"userArr"`
}

type userMyInfoPatchModel struct {
	Dept						string	`json:"dept" example:"소속부서"`
	Position				string	`json:"position" example:"직급"`
	Phoneno					string	`json:"phoneno" example:"021231234"`
	Major_field			string	`json:"major_field" example:"수의학"`
	Edu_date				string	`json:"edu_date" example:"20210101"`
	Edu_institution	string	`json:"edu_institution" example:"건국대학교"`
	Edu_course_num	string	`json:"edu_course_num" example:"0"`
	Agree_email			uint		`json:"agree_email" example:"1"`
	Agree_sms				uint		`json:"agree_sms" example:"1"`
	Agree_pri_open	uint		`json:"agree_pri_open" example:"1"`
}

/**********************************************/
// 플랫폼
/**********************************************/

type platformUserRegisterModel struct {
	Institution_seq uint	  `json:"institution_seq" example:"1"`
	User_type				string	`json:"user_type" example:"1,2"`
	Email						string	`json:"email" example:"test@test.com"`
	Phoneno					string	`json:"phoneno" example:"01012341234"`
}

type platformUserRegisterBatchModel struct {
	userArr []platformUserRegisterModel `json:"userArr"`
}

type platformPatchUserInfoModel struct {
	User_seq				uint		`json:"user_seq"	example:"1"`
	Institution_seq	uint		`json:"institution_seq"	example:"1"`
	User_type				string	`json:"user_type"	example:"1"`
	User_status			uint		`json:"user_status"	example:"1"`
	Dept						string	`json:"dept"	example:"소속부서"`
	Position				string	`json:"position" example:"직급"`
	Phoneno					string	`json:"phoneno" example:"021231234"`
	Major_field			string	`json:"major_field" example:"수의학"`
	Edu_date				string	`json:"edu_date" example:"20210101"`
	Edu_institution	string	`json:"edu_institution" example:"건국대학교"`
	Edu_course_num	string	`json:"edu_course_num" example:"0"`
	Agree_email			uint		`json:"agree_email" example:"1"`
	Agree_sms				uint		`json:"agree_sms" example:"1"`
	Agree_pri_open	uint		`json:"agree_pri_open" example:"1"`
}

type platformAdminUserPatchModel struct {
	User_seq  uint		`json:"user_seq" example:"1"`
	Email    	string	`json:"email" example:"test@test.com"`
	Name			string	`json:"name" example:"홍길동"`
	Telno			string	`json:"telno" example:"0212341234"`
}

type InstitutionServiceChangeModel struct {
	Institution_seq	uint	`json:"institution_seq"	example:"1"`
	Service_status	uint	`json:"service_status" example:"1"`
}

/**********************************************/
// 로그인
/**********************************************/
type loginModel struct {
	email    string `json:"email" example:"test@test.com"`
	pw       string `json:"pw" example:"test"`
	tmp_key	 string `json:"tmp_key" example:"12345678901234561234567890123456"`
}

type EmailloginModel struct {
	email    string `json:"email" example:"test@test.com"`
	pw       string `json:"pw" example:"test"`
	tmp_key	 string `json:"tmp_key" example:"12345678901234561234567890123456"`
	user_seq uint   `json:"user_seq" example:1"`
}

type loginChangePwModel struct {
	old_pw string `json:"old_pw" example:"tester"`
	new_pw string `json:"new_pw" example:"test"`
}

type tokenModel struct {
	email    						 string `json:"email" example:"test@test.com"`
	pw       						 string `json:"pw" example:"test"`
	institution_name_ko  string `json:"institution_name_ko" example:"기관명"`
}

/**********************************************/
// 결제
/**********************************************/
type orderAssginModel struct {
	Product_seq	uint 		`json:"product_seq" example:"1"`
	Goods_name	string 	`json:"goods_name" example:"정기 1개월권"`
	Amt 			 	uint 	 	`json:"amt" example:"1000000"`
	Edi_date	 	string	`json:"edi_date" example:"YYYYMMDDHHMMSS"`
}

type orderCancelModel struct {
	CancelAmt					string	`json:"cancel_amt" example:"1000"`
	PartialCancelCode string	`json:"partial_cancel_code" example:"1"`
}


type freeMembershipRegisterModel struct {
	Institution_seqs	[]uint  `json:"institution_seqs"`
	Usage_limit				uint		`json:"usage_limit" example:"1"`
	Free_period 			string	`json:"free_period" example:"1"`
	Reason 						string	`json:"reason" example:"가입 기념 무료 지급"`
}
