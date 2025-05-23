package common

import (
)

const (
	Api_status_ok                  = 200
	Api_status_created             = 201
	Api_status_accepted            = 202
	Api_status_no_content          = 204

	Api_status_bad_request         = 400
	Api_status_unauthorized        = 401
	Api_status_forbidden           = 403
	Api_status_not_found           = 404
	Api_status_method_not_allowed  = 405
	Api_status_conflict            = 409
	Api_status_too_many_request    = 429
)

const (
	No_error = iota // 0
	Error_invalide_params
	Error_unauthorized
	Error_dup_name
	Error_system_unknown
	Error_upload_fail
	Error_token_mismatch
	Error_id_mismatch
	Error_download_auth
	Error_none_file
	Error_last_admin
	Error_app_step
	Error_info_mismatch
	Error_send_email
	Error_app_seq_mismatch
	Error_not_request_status_finish
	Error_invalid_product
	Error_pay_auth
	Error_pay
	Error_billing_key_register
	Error_payment_fail
	Error_cancel_order_fail
	Error_membership_cancel_required
	Error_cancel_withdraw_over_time
	Error_limit_monthly_utilization

)

var error_msg []string = []string{
	"성공",
	"필수 파라메터가 없습니다.",
	"접근권한이 없습니다.",
	"이미 사용중인 이름입니다.",
	"시스템 에러입니다. 잠시후 다시 시도하세요.",
	"파일 업로드 실패",
	"토큰이 유효하지 않습니다.",
	"아이디가 유효하지 않습니다.",
	"파일 다운로드 권한이 없습니다.",
	"파일이 존재하지 않습니다.",
	"행정간사는 반드시 1명 이상 존재해야합니다.",
	"신청서 상태변경 단계가 맞지 않습니다.",
	"입력한 정보에 맞는 사용자를 찾을 수 없습니다.",
	"이메일 전송 실패입니다.",
	"child_app_seq와 app_seq가 매칭되지 않습니다.",
	"신청서가 처리완료 상태가 아닙니다.",
	"유효하지 않은 결제 상품입니다.",
	"결제 인증에 실패했습니다.",
	"결제를 실패했습니다.",
	"빌링키 등록에 실패했습니다.",
	`결제 시스템의 연결이 원활하지 않습니다.
	 문제가 계속 된다면, 서비스 관리자에 문의해주세요.
	 IPSAP 서비스 관리자 <support@ipsap.co.kr>`,
	"결제 취소를 실패했습니다.",
  "가입된 멥버십이 존재합니다.",
  "탈퇴 철회 할수 있는 기간이 지났습니다. 관리자에게 문의해 주시기 바랍니다.",
  "월 이용 건수를 초과 할수 없습니다.",
  /*
	"패스워드가 일치하지 않습니다.",
	"비밀번호 변경 기간이 지났습니다.",
	"패치 대기 목록이 비어 있습니다.",
	"Patch를 실행 할 수 없습니다. 현재 설치중인 Patch가 있습니다.",
	"Patch를 실행 할 수 없습니다. 현재 스캔이 진행중인 서버가 있습니다.",
	"테스트 서버 등록 오류입니다.",
	"인증오류입니다.",*/
}
