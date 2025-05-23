package model

import (
)

// 심사 유형 code
const DEF_APP_JUDGE_TYPE_CODE_IACUC  = 1
const DEF_APP_JUDGE_TYPE_CODE_IBC    = 2
const DEF_APP_JUDGE_TYPE_CODE_IRB    = 3

// APP 번호 심사 유형
const DEF_APP_NO_IACUC  = "IA"
const DEF_APP_NO_IBC    = "IB"
const DEF_APP_NO_IRB    = "IR"

// Directory명
const DIR_APPLICATION                   = "application"
const DIR_GENERAL_REF                   = "general_ref"                   //  1.6 관련자료 첨부
const DIR_ANIMAL_SPECIES_REASON         = "animal_species_reason"         //  3.2 종/계통 선택한 합리적인 이유
const DIR_ANIMAL_CNT_REASON             = "animal_cnt_reason"             //  3.3 동물수에 대한 합리적인 근거
const DIR_ANIMAL_EXP_SUMMARY            = "animal_exp_summary"            //  5.1 동물실험의 개요 및 일정
const DIR_ANIMAL_EXP_SURGICAL_METHOD    = "animal_exp_surgical_method"    //  5.5 처치방법
const DIR_PAIN_RELIEF_PSYCH_M_LICENSE   = "pain_relief_psych_m_license"   //  7.3 사용허가증 첨부
const DIR_PAIN_RELIEF_ANIMAL_M_LICENSE  = "pain_relief_animal_m_license"  //  7.4 처방전 첨부

//  User Type
const DEF_USER_TYPE_ADMIN_SECRETARY = 1 //  행정간사
const DEF_USER_TYPE_CHAIRPERSON     = 2 //  위원장
const DEF_USER_TYPE_RESEARCHER      = 3 //  연구원
const DEF_USER_TYPE_COMMITTEE       = 4 //  심사위원
const DEF_USER_TYPE_ADMIN_WORK      = 5 //  행정업무

//  User Auth
const DEF_USER_AUTH_NOMARL            = 0   //  일반사용자
const DEF_USER_AUTH_INSTITUTION       = 1   //  기관관리자
const DEF_USER_AUTH_PLATFORM          = 9   //  플랫폼관리자
const DEF_USER_AUTH_PLATFORM_SERVICE  = 10  //  플랫폼관리자(서비스관리)
const DEF_USER_AUTH_PLATFORM_SYSTEM   = 11  //  플랫폼관리자(시스템관리)

const INSTITUTION_SHARE_CODE    = 0     //  모든기관 공통 기관코드

//  User status
const DEF_USER_STATUS_WAIT                = 1   //  등록대기
const DEF_USER_STATUS_FINISH              = 2   //  등록완료
const DEF_USER_STATUS_EXISTING_UESR_WAIT  = 3   //  기존등록사용자 다른 기관 새로 등록대기
const DEF_USER_STATUS_WITHDRAW            = 4   //  탈퇴(삭제 X)
const DEF_USER_STATUS_FORCED_WITHDRAW     = 5   //  강제탈퇴(삭제 X)
const DEF_USER_STATUS_REGIST_FAIL         = 6   //  등록실패

//  DCode Type
const DCODE_TYPE_DEPT             = 1     //  부서타입
const DCODE_TYPE_POSITION         = 2     //  직위
//const DCODE_TYPE_ANIMAL           = 3   //  동물종류  (삭제함)
const DCODE_TYPE_MAJOR            = 4     //  전공분야
const DCODE_TYPE_EDU_ISTT         = 5     //  교육기관
//const DCODE_TYPE_MB_GRADE         = 6   //  미생물학적 등급  (삭제함)
//const DCODE_TYPE_BREEDING_PLACE   = 7   //  사육희망장소  (삭제함)
const DCODE_TYPE_ANESTHETIC_TYPE  = 8     //  약물종류

//  ITEM_TYPE                               //  조회     저장
const DEF_ITEM_TYPE_NOTTHING        = 0     //  ok      ok
const DEF_ITEM_TYPE_ITEM_GROUP      = 1     //  ok      ok
const DEF_ITEM_TYPE_BASIC           = 2     //  ok      ok
const DEF_ITEM_TYPE_SELECT          = 3     //  ok      ok
const DEF_ITEM_TYPE_KEYWORD         = 4     //  ok      ok
const DEF_ITEM_TYPE_FILE            = 5     //  ok      ok
const DEF_ITEM_TYPE_MEMBER          = 6     //  ok      ok
const DEF_ITEM_TYPE_ANIMAL          = 7     //  ok      ok
const DEF_ITEM_TYPE_ANESTHETIC      = 8     //  ok      ok
const DEF_ITEM_TYPE_CUSTOM1         = 9
const DEF_ITEM_TYPE_STRING          = 10    //  ok      ok

// Institution Status 기관신청 상태
const DEF_INSTITUTION_STATUS_WRITING  = 1   // 작성중
const DEF_INSTITUTION_STATUS_FINISH   = 2   // 제출완료 (신청 완료)
const DEF_INSTITUTION_STATUS_ING      = 3   // 심사중
const DEF_INSTITUTION_STATUS_REJECT   = 8   // 승인거부
const DEF_INSTITUTION_STATUS_OK       = 9   // 승인완료

//  신청서 종류
const DEF_APP_TYPE_NEW              = 1   //  신규 승인
const DEF_APP_TYPE_CHANGE           = 2   //  변경 신청
const DEF_APP_TYPE_RENEW            = 3   //  재승인
const DEF_APP_TYPE_BRINGIN          = 4   //  반입신청서
const DEF_APP_TYPE_CHECKLIST        = 5   //  승인후 점검표
const DEF_APP_TYPE_FINISH           = 6   //  종료 보고서
const DEF_APP_TYPE_CONTINUE         = 7   //  IRB 지속심의(중간보고)
const DEF_APP_TYPE_SERIOUS          = 8   //  IRB 중대한 이상 반응 보고서
const DEF_APP_TYPE_VIOLATION        = 9   //  IRB 연구계획 위반/이탈 보고
const DEF_APP_TYPE_UNEXPECTED       = 10  //  IRB 예상치 못한 문제발생 보고서

//  신청서 진행 단계
const DEF_APP_STEP_WRITE            = 0   // 신청서 작성중
const DEF_APP_STEP_CHECKING         = 1   // 행점검토
const DEF_APP_STEP_PRO              = 2   // 전문심사
const DEF_APP_STEP_NORMAL           = 3   // 일반심사
const DEF_APP_STEP_FINAL            = 4   // 최종심의
const DEF_APP_STEP_PERFORMANCE      = 5   // 과제수행

//  신청서 진행 상태
const DEF_APP_RESULT_TEMP           = 0   // 임시저장
const DEF_APP_RESULT_SUPPLEMENT     = 1   // 보완중
const DEF_APP_RESULT_CHECKING       = 2   // 검토중
const DEF_APP_RESULT_CHECKING_2     = 3   // 검토중 2 => 행정간사의 경우에 검토중이 2단계로 구성
const DEF_APP_RESULT_JUDGE_ING      = 4   // 심사 중
const DEF_APP_RESULT_JUDGE_ING_2    = 5   // 심사 중 2
const DEF_APP_RESULT_JUDGE_DELAY    = 6   // 심사 지연
//const DEF_APP_RESULT_JUDGE_FINISH   = 7   // 심사 종료    : 삭제됨
const DEF_APP_RESULT_DECISION_ING   = 8   // 심의 중
const DEF_APP_RESULT_REJECT         = 9   // 반려
const DEF_APP_RESULT_REQUIRE_RETRY  = 10  // 보안 후 재심
const DEF_APP_RESULT_EXPER_ING_A    = 11  // 과제 실험 중 (승인)
const DEF_APP_RESULT_EXPER_ING_AC   = 12  // 과제 실험 중 (조건부 승인)
const DEF_APP_RESULT_EXPER_FINISH   = 13  // 실험 종료
const DEF_APP_RESULT_TASK_FINISH    = 14  // 과제 종료 -- > 21-06-04 실험완료 (과제 종료 라는 개념이 사라짐)
const DEF_APP_RESULT_APPROVED       = 15  // 승인
const DEF_APP_RESULT_APPROVED_C     = 16  // 조건부 승인
const DEF_APP_RESULT_DELETED        = 17  // 삭제됨

//  결제 타입
const DEF_REQ_TYPE_SERVICE_REGISTER = 1   //  서비스 등록
const DEF_REQ_TYPE_PAYMENT_CHANGE   = 2   //  정보 변경

//  결제 요청 상태
const DEF_REQ_STATUS_WAIT   = 0   //  미처리
const DEF_REQ_STATUS_ING    = 1   //  처리중
const DEF_REQ_STATUS_HOLD   = 2   //  처리보류
const DEF_REQ_STATUS_FINISH = 9   //  처리완료

// 실험 계획서 권환별 화면
const DEF_APP_VIEW_ADMIN_ALL                = 1   // 행정(간사, 담당자) 전체 계획서 및 보고서
const DEF_APP_VIEW_ADMIN_REVIEW_OFFICE      = 2   // 행정(간사, 담당자) 행정 검토
const DEF_APP_VIEW_ADMIN_REVIEW_CLOSE       = 3   // 행정(간사, 담당자) 심사 종료 설정
const DEF_APP_VIEW_ADMIN_REVIEW_CONFIRM     = 4   // 행정간사 최종 심의
const DEF_APP_VIEW_RESEARCHER_ALL           = 5   // 연구원 실험계획서
const DEF_APP_VIEW_CHAIRMAN_REVIEW_CONFIRM  = 6   // 위원장 심사진행
const DEF_APP_VIEW_COMMITTEE_REVIEW         = 7   // 심사위원 심사진행
const DEF_APP_VIEW_CHAIRPERSON_RECORD       = 8   // 위원장 심사기록
const DEF_APP_VIEW_COMMITTEE_RECORD         = 9   // 심사위원 심사기록
const DEF_APP_VIEW_APPROVED_INSPECT_RECORD  = 10  // 나의 승인후 점검 기록
const DEF_APP_VIEW_MY_INSPECT_CHECK         = 11  // 내가 점검할 실험
const DEF_APP_VIEW_IACUC_FOR_IBC            = 12  // ibc에서 iacuc 데이터 참조 가능한 list
const DEF_APP_VIEW_IBC_FOR_IACUC            = 13  // iacuc에서 ibc 데이터 참조 가능한 list

// 승인후 실험 계획서
const DEF_APP_VIEW_APPROVED_ADMDIN_ALL      = 1   // 행정(간사, 담당자) 승인후 관리 > 전체 실험
const DEF_APP_VIEW_APPROVED_ADMDIN_ING      = 2   // 행정(간사, 담당자) 승인후 관리 > 진행중인 실험
const DEF_APP_VIEW_APPROVED_RESEARCHER_ALL  = 3   // 연구원 승인후 관리 > 실험 수행 및 서류
const DEF_APP_VIEW_APPROVED_CHAIRPERSON_ALL = 4   // 위원장 진행중인 실험 list

// 연구원 > 재승인, 변경승인 가능한 신청서 목록
const DEF_APP_VIEW_POSSIBLE_IA_RETRY     = 1   // IACUC 재승인 신청 가능한 목록
const DEF_APP_VIEW_POSSIBLE_IA_CHANGE    = 2   // IACUC 변경 승인 신청 가능한 목록
const DEF_APP_VIEW_POSSIBLE_IBC_CHANGE   = 3   // IBC  변경승인 신청 가능한 목록
const DEF_APP_VIEW_POSSIBLE_IRB_CHANGE   = 4   // IRB  변경승인 신청 가능한 목록
const DEF_APP_VIEW_POSSIBLE_IRB_RETRY    = 5   // IRB  지속심의 신청 가능한 목록
const DEF_APP_VIEW_POSSIBLE_IA_SUPPLE    = 6   // IACUC 보완후 재심인 목록
const DEF_APP_VIEW_POSSIBLE_IB_SUPPLE    = 7   // IBC 보완후 재심인 목록

// 제출 형식 (없으면 신청서 제출!!)
const DEF_SUBMIT_TYPE_SUPPLEMENT                  = 1   // 보완 요청(행정간사, 행정담당만 가능)
const DEF_SUBMIT_TYPE_CHECKING                    = 2   // 행정 검토 1단계 완료(행정간사, 행정담당만 가능)
const DEF_SUBMIT_TYPE_CHECKING_2                  = 3   // 행정 검토 2단계 완료(행정간사, 행정담당만 가능)
const DEF_SUBMIT_TYPE_JUDGE_PRO                   = 4   // 전문심사 1단계 완료 (심사위원만 가능)
const DEF_SUBMIT_TYPE_JUDGE_PRO_2                 = 5   // 전문심사 2단계 완료 (심사위원만 가능)
const DEF_SUBMIT_TYPE_JUDGE_NORMAL                = 6   // 일반심사 심사(심사위원만 가능)
const DEF_SUBMIT_TYPE_JUDGE_RESUME                = 7   // 일반심사 심사재개(심사위원만 가능) 지연중 -> 심사중
const DEF_SUBMIT_TYPE_JUDGE_FINISH                = 8   // 심사종료설정 완료 (행정간사, 행정담당만 가능)
const DEF_SUBMIT_TYPE_JUDGE_FINAL_A               = 9   // 최종심의 승인(행정간사 또는 위원장 가능)
const DEF_SUBMIT_TYPE_JUDGE_FINAL_AC              = 10  // 최종심의 조건부 승인(행정간사 또는 위원장 가능)
const DEF_SUBMIT_TYPE_JUDGE_FINAL_REJECT          = 11  // 최종심의 반려 (행정간사 또는 위원장 가능)
const DEF_SUBMIT_TYPE_JUDGE_FINAL_REQUIRE_RETRY   = 12  // 최종심의 보완후 재심 (행정간사 또는 위원장 가능)
const DEF_SUBMIT_TYPE_EXPER_FINISH                = 13  // 과제수행 실험 종료(연구원)
const DEF_SUBMIT_TYPE_TASK_FINISH                 = 14  // 과제수행 과제 종료(행정간사 또는 위원장 가능)
const DEF_SUBMIT_TYPE_RETRY_CHECKING              = 15  // 행정 검토 1단계로 돌아가기 (행정간사, 행정담당만 가능)

const DEF_SUBMIT_TYPE_CHILD_SUBMIT                = 20  // 부속 신청서 제출
const DEF_SUBMIT_TYPE_CHECKING_FAST_FINISH        = 21  // 부속 신청서 공통 : 행정검토 신속 완료 (최종 심의로 단계이동)

const DEF_SUBMIT_TYPE_JUMP_FINAL                  = 30  // 일반 심사 지연중 일경우 행정간사 판단하에 최종심의 단계로 이동

//  메세지 종류
const DEF_MSG_INSTITUTION_APPROVED          = 1     //  기관 신청 승인 안내
const DEF_MSG_INSTITUTION_CHANGE_APPROVED   = 2     //  기관 변경신청 승인 안내
const DEF_MSG_USER_REGIST                   = 3     //  계정등록 안내(신규)
const DEF_MSG_EXISTING_USER_REGIST          = 4     //  계정등록 안내(기존유저)
const DEF_MSG_USER_PASSWORD_CHANGE          = 5     //  비밀번호 변경 안내
const DEF_MSG_EXPER_JUDGE_START             = 7     //  전문심사 게시 안내
const DEF_MSG_NORMAL_JUDGE_START            = 8     //  일반심사 게시 안내
const DEF_MSG_BEFORE_24_HOURS               = 9     //  심사종료 1차 안내 (24시간 전), (전문/일반 심사)
const DEF_MSG_BEFORE_4_HOURS                = 10    //  심사종료 2차 안내 (4시간 전), (전문/일반 심사)
const DEF_MSG_JUDGE_DELAYED                 = 11    //  심사 미이행 안내
const DEF_MSG_JUDGE_FINISHED                = 12    //  최종심의 결과 안내
const DEF_MSG_REQUEST_SUPPLEMENT            = 16    //  행정검토 보완 요청(책임연구자에게 전송)
const DEF_MSG_EXPER_JUDGE_START_TO_LEADER   = 17    //  전문심사 게시 안내(책임연구자에게 전송)
const DEF_MSG_NORMAL_JUDGE_START_TO_LEADER  = 18    //  일반심사 게시 안내(책임연구자에게 전송)

// 메세지 탈퇴 관련
const DEF_MSG_USER_WITHDRAW                 = 6     //  계정탈퇴 안내(기관내 탈퇴)
const DEF_MSG_USER_IPSAP_WITHDRAW           = 13    //  계정탈퇴 안내(소속한 모든 기관 탈퇴)
const DEF_MSG_WITHDRAW_NOTICE               = 15    //  IPSAP 서비스 사용자 탈퇴공지 -> 행정간사

const DEF_MSG_IPSAP_WITHDRAW_REQUEST        = 14    // IPSAP 기관 계정 탈퇴 접수 안내 -> 행정간사
const DEF_MSG_IPSAP_WITHDRAW_REQUEST_NOTI   = 19    // IPSAP 기관 계정 탈퇴 신청 알림 -> 관리자
const DEF_MSG_IPSAP_WITHDRAW_FINISHED       = 20    // IPSAP 기관 계정 탈퇴 완료 안내-> 기관 사용자 모두
const DEF_MSG_IPSAP_WITHDRAW_FINISHED_NOTI  = 21    // IPSAP 기관 계정 탈퇴 완료 알림 -> 관리자
const DEF_MSG_USER_WITHDRAW_FINISHED        = 22    // 기관 사용자 탈퇴 완료 안내 -> 신청자 본인


// 재승인 신청 관련
const DEF_MSG_RE_APP_BEFORE                 = 30    // 재승인 신청 1달전
const DEF_MSG_RE_APP_AFTER                  = 31    // 재승인 신청 경과

// 멤버쉽 관련 안내 메일 및 메세지
const DEF_MSG_MEMBERSHIP_JOIN               = 70    //  멤버쉽 가입 안내
const DEF_MSG_MEMBERSHIP_JOIN_NOTI          = 71    //  멤버쉽 가입 알림 -관리자
const DEF_MSG_MEMBERSHIP_CANCEL             = 72    //  멤버쉽 해지 안내
const DEF_MSG_MEMBERSHIP_CANCEL_NOTI        = 73    //  멤버쉽 해지 알림 -관리자
const DEF_MSG_MEMBERSHIP_STOP_EXPECTED      = 74    //  멤버쉽 이용중지 예정 안내
const DEF_MSG_MEMBERSHIP_STOP               = 75    //  멤버쉽 이용중지 알림
const DEF_MSG_MEMBERSHIP_STOP_NOTI          = 76    //  멤버쉽 이용중지 알림 - 관리자
const DEF_MSG_MEMBERSHIP_FREE               = 77    //  멤버쉽 무료이용 안내
const DEF_MSG_MEMBERSHIP_REJOIN             = 78    //  멤버쉽 재가입
const DEF_MSG_MEMBERSHIP_CHANGE             = 80    //  멤버쉽 변경
const DEF_MSG_MEMBERSHIP_EXPIRED            = 81    //  멤버쉽 만료
const DEF_MSG_MEMBERSHIP_PAYMENT_EXPECTED   = 90    //  멤버쉽 결제 예정
const DEF_MSG_MEMBERSHIP_PAYMENT            = 91    //  멤버쉽 결제 완료
const DEF_MSG_MEMBERSHIP_UNPAID             = 92    //  멤버쉽 월 이용료 미납 안내
const DEF_MSG_MEMBERSHIP_UNPAID_NOTI        = 93    //  멤버쉽 월 이용료 미납 알림 - 관리자

// 심사안내 자동 알림설정 종류
const DEF_JUDGE_ALARM_EXPER_JUDGE_START   = 1 //  전문심사 개시 안내
const DEF_JUDGE_ALARM_NORMAL_JUDGE_START  = 2 //  일반심사 개시 안내
const DEF_JUDGE_ALARM_BEFORE_24_HOURS     = 3 //  심사종료 1차 안내 (24시간 전), (전문/일반 심사)
const DEF_JUDGE_ALARM_BEFORE_4_HOURS      = 4 //  심사종료 2차 안내 (4시간 전), (전문/일반 심사)
const DEF_JUDGE_ALARM_JUDGE_DELAYED       = 5 //  심사 미이행 안내

// SMS 결과
const DEF_SMS_RESULT_WAIT  = 0  // 대기
const DEF_SMS_RESULT_SUCC  = 2  // 성공
const DEF_SMS_RESULT_FAIL  = 4  // 실패

// 최종발부승인
const DEF_FINAL_DIRECTOR_CHAIRPERSON  = 1 // 위원장
const DEF_FINAL_DIRECTOR_ALL          = 2 // 위원장, 행정간사 모무 가능

// 신청서 종류
const DEF_REQUEST_TYPE_SERVICE_REGIST = 1 // 서비스 등록
const DEF_REQUEST_TYPE_PAYMENT_CHANGE = 2 // 결제정보 변경

// 신청서 처리 타입
const DEF_REQ_HANDLE_TYPE_APPROVED  = 1  // 승인
const DEF_REQ_HANDLE_TYPE_HOLD      = 2  // 보류
const DEF_REQ_HANDLE_TYPE_RESEND    = 3  // 처리결과 재발송

// 게시판 타입
const DEF_BOARD_TYPE_NOTICE = 1  // 공지사항
const DEF_BOARD_TYPE_FILE   = 2  // 자료실
const DEF_BOARD_TYPE_FAQ    = 3  // FAQ

// DashBoard 타입
const DEF_DASHBOARD_TYPE_ADMIN        = 1  // 행정간사, 행정업무 Dashboard
const DEF_DASHBOARD_TYPE_RESEARCHER   = 2  // 연구원 Dashboard
const DEF_DASHBOARD_TYPE_CHAIRPERSON  = 3  // 위원장 Dashboard

// Order Status
const DEF_ORDER_STATUS_COMPLETED  = 1  // 완료
const DEF_ORDER_STATUS_ERROR      = 2  // 오류

// Order type
const DEF_ORDER_TYPE_NORMAL_PAYMENT   = 1  // 일반결제
const DEF_ORDER_TYPE_REGULAR_PAYMENT  = 2  // 정기결제
const DEF_ORDER_TYPE_ALL_CANCEL       = 3  // 전체취소
const DEF_ORDER_TYPE_PARTIAL_CANCEL   = 4  // 부분취소

// Pay method
const DEF_PAY_METHOD_CARD = "CARD" // 카드
const DEF_PAY_METHOD_BANK = "BANK" // 은행

// Tid 생성시 필요한 데이터
const DEF_TID_CARD = "01" // 지불수단(신용카드)
const DEF_TID_DIVISION = "16" // 매체구분(빌링)

// 기관 서비스 상태
const DEF_SERVICE_STATUS_BEFORE_USING = 0 // 이용전
const DEF_SERVICE_STATUS_IN_USE       = 1 // 이용중
const DEF_SERVICE_STATUS_STOPPED      = 2 // 이용 정지
const DEF_SERVICE_STATUS_WITHDRAWN    = 3 // 서비스 탈퇴

// PG Biiling Const
const DEF_BIILING_CARDINTEREST  = "0" //가맹점 분담 무이자 사용 여부 ( 사용안함 이자 / 1: 사용 무이자)
const DEF_BIILING_CARDQUOTA     = "00" //할부개월00: 일시불 / 02: 2 개월 / 03: 3 개월

// mebership replace name
const DEF_MEMBERSHIP_REFUND           = "${가입비환불금액}"
const DEF_USAGE_REFUND                = "${이용료환불금액}"
const DEF_MEMBERSHIP_JOIN_DATE        = "${멤버십가입일}"
const DEF_MEMBERSHIP_EXPIRATION_DATE  = "${멤버십만료일}"
const DEF_DESIGNATED_PAYMENT_METHOD   = "${지정결제방법}"
const DEF_REFUND_APPLICATION_DATE     = "${환불신청일시}"
const DEF_MONTH_START_DATE            = "${월이용시작일}"
const DEF_MONTH_END_DATE              = "${월이용종료일}"
const DEF_FREE_START_DATE             = "${무료시작일}"
const DEF_FREE_END_DATE               = "${무료만료일}"
const DEF_CANCEL_APPLICATION_DATE     = "${해지신청일}"
const DEF_PAYMENT_DEADLINE            = "${결제마감일}"
const DEF_STOP_DATE                   = "${이용중지일}"
const DEF_PAYMENT_DUE_DATE            = "${결제예정일}"
const DEF_CANCEL_TOTAL_AMOUNT         = "${환불총금액}"
const DEF_MEMBERSHIP_NAME_OLD         = "${기존멤버십명}"
const DEF_MONTHLY_FEE                 = "${월이용료}"
const DEF_MEMBERSHIP_NAME             = "${멤버십명}"
const DEF_PAYMENT_DATE                = "${결제일시}"
const DEF_PAYMENT_NUMBER              = "${결제번호}"
const DEF_PAYMENT_AMOUNT              = "${결제금액}"
const DEF_PAYMENT_METHOD              = "${결제방법}"
const DEF_FREE_DAYS                   = "${무료일수}"
const DEF_CANCEL_DATE                 = "${해지일시}"

// niceApi param
const	DEF_PARAM_TID               = "TID"
const	DEF_PARAM_MID               = "MID"
const	DEF_PARAM_Moid              = "Moid"
const	DEF_PARAM_CancelAmt         = "CancelAmt"
const	DEF_PARAM_CancelMsg         = "CancelMsg"
const	DEF_PARAM_EdiDate           = "EdiDate"
const	DEF_PARAM_SignData          = "SignData"
const	DEF_PARAM_CharSet           = "CharSet"
const	DEF_PARAM_PartialCancelCode = "PartialCancelCode"

// niceApi url
const	DEF_URL_CANCEL  = "https://webapi.nicepay.co.kr/webapi/cancel_process.jsp"
