package ipsap

import (
	"ipsap/api"
	"ipsap/common"
	"ipsap/docs" //docs 파일 못찾는 에러가 날 경우, swag를 통해 doc을 생성한다. (readme.txt 참고)
	"ipsap/model"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "fmt"
)

// 스웨거 설정
// @title IPSAP Service APIs
// @version 1.0
// @BasePath /api/v1.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
func Start() {
	// 로그 출력시, 로그 호출 위치 나오도록 수정
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	port := common.Config.Server.Port
	swag_hostname := common.Config.Server.SwagHostname

	// 스웨거 설정
	docs.SwaggerInfo.Host = swag_hostname + ":" + strconv.Itoa(port)
	//	ctl.Init_common_code()

	//  모니터링 시작
	bgMonitor := model.BackgroundMonitor{}
	bgMonitor.Start()

	realStart(port)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, token, Token, Origin")
		c.Header("Access-Control-Credentials", "true")
		c.Header("Access-Control-Methods", "GET,POST,DELETE,PUT,PATCH")
		c.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setLog(fileName string) *rotatelogs.RotateLogs {
	basePath := common.Config.Server.LogPath
	_ = os.Mkdir(basePath, os.ModeDir)

	yyyymmdd := "%Y%m%d"

	rl, _ := rotatelogs.New(basePath+fileName+"_"+yyyymmdd+".log",
		rotatelogs.WithLinkName(basePath+fileName),
		rotatelogs.WithMaxAge(time.Duration(common.Config.Server.MaxBackups)*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	return rl
}

func setupRouter() *gin.Engine {
	if common.Config.Server.LogPath != "" {
		log.SetOutput(setLog("log_file"))
		gin.DefaultWriter = setLog("log_gin_normal")
		gin.DefaultErrorWriter = setLog("log_gin_error")
	}
	r := gin.Default()
	r.Use(corsMiddleware())
	return r
}

func realStart(port int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	r := setupRouter()

	// 스웨거 설정
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group(common.Config.Server.ApiPath)
	{
		{ // Application
			v1.GET("/application", api.Application_List)                                                   //  신청서 목록 조회
			v1.GET("/application/:application_seq/info", api.Application_Info)                             //  신청서 정보
			v1.GET("/application/:application_seq", api.Application_Get)                                   //  신청서 항목별 내용 조회
			v1.PATCH("/application/:application_seq", api.Application_Patch)                               //  신청서 항목별 내용 임시저장(수정)
			v1.POST("/application/:application_seq", api.Application_Post)                                 //  신청서 제출
			v1.DELETE("/application/:application_seq", api.Application_Delete)                             //  신청서 삭제(임시저장 상태만)
			v1.GET("/application-approved", api.Application_Approved_List)                                 //  승인후 신청서 목록 조회
			v1.GET("/application-possible", api.Application_Possible_List)                                 //  (재승인, 변경 승인) 신청 가능한 리스트
			v1.GET("/application/:application_seq/inspector", api.Application_Inspector_List)              //  승인후 점검위원 리스트
			v1.PATCH("/application/:application_seq/inspector/:user_seq", api.Application_Inspector_Patch) // 승인후 점검위원 지정
			v1.POST("/application/:application_seq/copy", api.Application_Copy)                            // 신청서 복제
			v1.POST("/application/:application_seq/iacuc-copy", api.Application_Copy_Iacuc_For_ibc)        // ibc 에서 iacuc 신청서 복제
			v1.POST("/application/:application_seq/retrial-copy", api.Application_Retrial_Copy)            // 보완후 재심인 경우 copy(우선 iacuc만 적용)
			v1.DELETE("/app/:application_seq", api.App_Delete)                                             //  신청서 삭제(상태 변경)
		}

		{ // Application histroy
			v1.GET("/application/:application_seq/change", api.Application_Change_List) // 신청서 변경 내역 보기
			v1.GET("/application/:application_seq/change/animal", api.Application_Change_Animal_Info)
			v1.GET("/application/:application_seq/change/member", api.Application_Change_Member_Info)
			v1.GET("/application/:application_seq/change/end-date", api.Application_Change_EndDate_Info)
		}

		{
			v1.GET("/file/:filepath_enc", api.File_Download) // 파일 다운로드
			v1.GET("/file-animal", api.Animal_Data_Download) // 동물 파일 다운로드
		}

		{ // RequestService
			v1.GET("/request/service", api.RequestServiceList)                        //  IPSAP 서비스 신청 목록 및 서비스 정보 변경 요청 목록
			v1.POST("/request/service", api.RequestServiceCreate)                     //  IPSAP 서비스 등록
			v1.GET("/request/service/:reqsvc_seq", api.RequestServiceInfo)            //  IPSAP 서비스 등록 정보 및 정보 변경 요청 상세 정보
			v1.PATCH("/request/service/:reqsvc_seq", api.RequestServicePatch)         //  IPSAP 서비스 등록 및 결제 정보 변경 수정
			v1.PATCH("/request/service/:reqsvc_seq/handle", api.RequestServiceHandle) //  IPSAP 서비스 신청 승인, 처리보류, 안내 재발송
		}

		{ // User
			v1.GET("/user/:user_seq", api.UserInfo)                             // User 정보
			v1.PATCH("/user/:user_seq", api.UserPatch)                          // User 수정
			v1.DELETE("/user/:user_seq", api.UserDelete)                        // User 탈퇴
			v1.PATCH("/user/:user_seq/register", api.UserRegister)              // 나의 정보 등록
			v1.PATCH("/user/:user_seq/change-password", api.UserChangePassword) // User비밀번호 변경
			v1.DELETE("/user/:user_seq/institution", api.UserAllDelete)         // User Ipsap 계정 탈퇴
		}

		{ // Admin
			v1.GET("/admin/user", api.AdminUserList)                                     // 사용자 괸리 리스트
			v1.POST("/admin/user", api.AdminUserCreate)                                  // 행정간사가 User 생성
			v1.POST("/admin/user-batch", api.AdminBatchUserCreate)                       // 행정간사가 User 일괄 등록
			v1.PATCH("/admin/user", api.AdminPatchUser)                                  // 행정간사가 User 수정
			v1.PATCH("/admin/user/:user_seq/reset-password", api.AdminResetUserPassword) // 행정간사가 User비밀번호 초기화
			v1.PATCH("/admin/user/:user_seq/withdraw", api.AdminWithdrawUser)            // 행정간사가 회원 강제 탈퇴
			v1.PATCH("/admin/user/:user_seq/resend-msg", api.AdminResendMsg)             // 행정간사가 User 가입 메세지 재전송
			v1.DELETE("/admin/user/:user_seq", api.AdminUserDelete)                      // 행정간사가 등록 대기 상태인 유저 삭제
		}

		{ // Institution
			v1.GET("/institution", api.InstitutionList)                                                                     // 기관 리스트
			v1.GET("/institution/:institution_seq", api.InstitutionInfo)                                                    // 기관 정보
			v1.PATCH("/institution/:institution_seq", api.InstitutionPatch)                                                 // 기관 정보 수정
			v1.DELETE("/institution/:institution_seq", api.InstitutionDelete)                                               // 기관 탈퇴
			v1.POST("/institution/:institution_seq/payment", api.InstitutionPaymentPatch)                                   // 기관 결제정보 변경
			v1.GET("/institution/:institution_seq/user", api.InstitutionUserList)                                           // 기관에 소속된 유저 리스트
			v1.GET("/my/other-institution", api.MyOtherInstitutionList)                                                     // 내가 소속된 타 기관 리스트(현재 로그인된 기관 이외의)
			v1.POST("/my/institution", api.MyInstitutionList)                                                               // 나의 기관 리스트
			v1.POST("/move-institution", api.MoveInstitution)                                                               // 소속기관으로 이동하기
			v1.GET("/institution/:institution_seq/admin-count", api.InstitutionAdminCount)                                  // 기관 행정간사 인원수 가져오기
			v1.GET("/institution/:institution_seq/using-membership/purchased", api.InstitutionUsingMembershipPurchasedList) // 기관에서 사용중인 멤버쉽의 결제 이력
		}

		{ // Board
			v1.GET("/board", api.BoardList)                        // 게시판 리스트
			v1.GET("/institution-board", api.InstitutionBoardList) // 등록 기관 공지사항 리스트
			v1.GET("/board/:board_seq", api.BoardInfo)             // 게시판 상세 정보
			v1.POST("/board", api.BoardCreate)                     // 게시판 등록
			v1.PATCH("/board/:board_seq", api.BoardPatch)          // 게시판 수정
			v1.DELETE("/board/:board_seq", api.BoardDelete)        // 게시판 삭제
		}

		{ // DashBoard
			v1.GET("/dashboard", api.DashBoardList) // DashBoard 리스트
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/login", api.Login)
			auth.POST("/token", api.Token)
			auth.POST("/email/login", api.EmailLogin)                    // 이메일을 통해서 들어 왔을때 로그인
			auth.POST("/email/cancel-withdraw", api.EmailCancelWithDraw) // 탈퇴 철회
			auth.POST("/find-id", api.FindId)                            // User 아이디 찾기
			auth.POST("/find-pwd", api.FindPwd)                          // User 비밀번호 찾기
		}

		platform := v1.Group("/platform")
		{
			platform.GET("/admin-user", api.PlatformAdminUserInfo)
			platform.PATCH("/admin-user", api.PlatformAdminUserPatch)
			platform.GET("/user", api.PlatformUserList)
			platform.PATCH("/user", api.PlatformPatchUser)
			platform.GET("/user/:user_seq", api.PlatformUserInfo)
			platform.POST("/user", api.PlatformUserCreate)
			platform.DELETE("/user/:user_seq", api.PlatformUserDelete)
			platform.PATCH("/user/:user_seq/resend-msg", api.PlatformResendMsg)
			platform.PATCH("/user/:user_seq/reset-password", api.PlatformResetUserPassword)
			platform.POST("/user-batch", api.PlatformBatchUserCreate)
			platform.PATCH("/institution", api.PlatformPatchInstitution)
			platform.PATCH("/user/:user_seq/withdraw", api.PlatformWithdrawUser)
			platform.GET("/other-institution/:user_seq", api.PlatformOtherInstitutionUserList)
			platform.PATCH("/payment/cancel/institution", api.PlatformInstitutionServiceStatusChange)
		}

		{ // membership
			v1.GET("/membership/plan", api.PlanList)                // 요금제 리스트
			v1.GET("/membership/plan/:plan_seq", api.PlanInfo)      // 요금제 정보
			v1.PATCH("/membership/plan/:plan_seq", api.PlanPatch)   // 요금제 수정
			v1.DELETE("/membership/plan/:plan_seq", api.PlanDelete) // 요금제 삭제
			v1.POST("/membership/plan", api.PlanCreate)             // 요금제 생성

			v1.POST("/membership/free", api.MembershipFreeCreate)             //  무료 멤버쉽 지급
			v1.GET("/membership/free", api.MembershipFreeList)                //  무료 멤버쉽 지급 리스트(기관 기준)
			v1.DELETE("/membership/free/:free_seq", api.MembershipFreeDelete) //  무료 멤버쉽 지급 해제

			v1.GET("/membership", api.MembershipInUseAndPaymentInfo)                                       // 이용중인 멤버십 정보 및 결제 정보
			v1.GET("/membership/cancel", api.MembershipCancelInfo)                                         // 이용중인 멤버십 정보 및 결제 정보
			v1.DELETE("/membership", api.MembershipCancel)                                                 // 멤버쉽 해지
			v1.GET("/membership-free/institution", api.InstitutionFreeMembershipPossibleList)              // 무료 멤버십 지급 가능한 리스트(가입비를 지불한 상태여야됨)
			v1.PATCH("/institution/:institution_seq/payment-setting", api.InstitutionPaymentSettingChange) // 이용중인 멤버십 결제 설정(자동, 수동) 변경
			v1.PATCH("/institution/:institution_seq/product/:product_seq", api.InstitutionPlanChange)      // 멤버쉽 변경
		}

		{ // Oreders
			v1.GET("/orders", api.OrderList)                 // 결제 목록
			v1.GET("/orders/:order_seq", api.OrderInfo)      // 결제 정보
			v1.POST("/orders/assign", api.OrderAssgin)       // 상품 결제창 요청
			v1.POST("/orders", api.OrderCreate)              // 상품 결제 생성
			v1.DELETE("/orders/:order_seq", api.OrderCancel) // 결제 취소
		}

		{ // billing
			v1.GET("/billing-key/assign", api.BillingKeyAssgin) // 빌키 생성창 요청
			v1.POST("/billing-key", api.BillingKeyCreate)       // 빌키 생성
		}

		common := v1.Group("/common")
		{
			common.GET("/dup-check/institution-code", api.DupCheckInstitutionCode)
			common.GET("/dup-check/email", api.DupCheckEmail)
		}

	}

	/* 로컬에서 화면 호출시 설정 Start */
	r.Static("/html", "./html/html/")
	r.Static("/assets", "./html/assets/")
	r.Static("/plugins", "./html/plugins/")

	// 🔹 예: index.html을 기본으로 보여주고 싶다면 라우트 추가
	r.GET("/", func(c *gin.Context) {
		c.File("./html/index.html")
	})

	r.GET("/index.js", func(c *gin.Context) {
		c.File("./html/index.js")
	})

	r.GET("/login.js", func(c *gin.Context) {
		c.File("./html/login.js")
	})

	r.GET("/index.html", func(c *gin.Context) {
		c.File("./html/index.html")
	})
	/* 로컬에서 화면 호출시 설정 End */

	port_str := ":" + strconv.Itoa(port)
	/*
		if common.Config.Test.IsTestMode {
			common.Config.S3.BUCKET = common.Config.S3.BUCKET_TEST
			r.Run(port_str)
		} else {
			common.Config.S3.BUCKET = common.Config.S3.BUCKET_REAL
			// r.RunTLS(port_str, common.Config.Server.Cert, common.Config.Server.Key)
		}
	*/
	common.Config.S3.BUCKET = common.Config.S3.BUCKET_TEST
	r.Run(port_str)
}
