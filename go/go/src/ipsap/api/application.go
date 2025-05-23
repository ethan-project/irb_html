package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"ipsap/common"
	"ipsap/model"
  "strings"
	"math"
  "fmt"
	"log"
)

// @Tags Application
// @Summary Application 신청서 목록 조회
// @Description Application 신청서 목록 조회/ 실험 계획서 권환별 화면
// @Description app_view_type(없으면 전체 리스트 출력)
// @Description 1 행정(간사, 담당자) 전체 계획서 및 보고서 (권한 체크함 !)
// @Description 2 행정(간사, 담당자) 행정 검토
// @Description 3 행정(간사, 담당자) 심사 종료 설정
// @Description 4 행정간사 최종 심의
// @Description 5 연구원 실험계획서
// @Description 6 위원장 심사진행
// @Description 7 심사위원 심사진행
// @Description 8 위원장 심사기록
// @Description 9 심사위원 심사기록
// @Description 10 나의 승인후 점검 기록
// @Description 11 내가 점걸할 실험
// @Description 12 IBC에서 IACUC 참조 가능한 리스트
// @Description 13 IACUC에서 IBC 참조 가능한 리스트
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param filter.app_view_type query string false "1"
// @Param filter.start_index query string false   "page start index"
// @Param filter.row_cnt query string false       "row count"
// @Param search_words  query string false        "Search Words"
// @Router /application [get]
// @Success 200
func Application_List(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	appViewType	:= common.ToStr(c.Request.URL.Query().Get("filter.app_view_type"))
	startIdx 		:= c.Request.URL.Query().Get("filter.start_index")
	rowCnt   		:= c.Request.URL.Query().Get("filter.row_cnt")
	searchWords := c.Request.URL.Query().Get("search_words")

	if "" == startIdx {
	 startIdx = "0"
	} else if _, ret := common.ToValidateUint64(startIdx); false == ret {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "filter.start_index값이 잘못되었습니다.")
		return
	}

	if "" == rowCnt {
		rowCnt = "10"
	} else if _, ret := common.ToValidateUint64(rowCnt); false == ret {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "filter.page_count값이 잘못되었습니다.")
		return
	}

	if appViewType == "" {
		if common.ToUint(tokenMap["user_auth"]) >= model.DEF_USER_AUTH_PLATFORM {
			app := model.Application {
				StartIdx : common.ToUint(startIdx),
				RowCnt : common.ToUint(rowCnt),
				SearchWords : searchWords,
			}
			list2 := app.LoadList("","","AND app.application_step > 0","")
			pageInfo := gin.H {
				"startIdx": app.StartIdx,
				"totalCnt": app.GetAppCnt("", "AND app.application_step > 0",appViewType),
			}

			common.FinishApi(c, common.Api_status_ok,
				gin.H{
					"rt": "ok",
					"list" : list2,
					"pageInfo" : pageInfo,
				})

			return
		}
	}

	application := model.Application{
		LoginToken : tokenMap,
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	moreSelct := ""
	moreJoin := ""
	moreCondition := ""
	succ := true
	if "" != appViewType {
		succ, moreSelct, moreJoin, moreCondition = checkAppViewAuthAndGetCondition(c, userTypeArr, common.ToUint(appViewType), common.ToUint(tokenMap["user_seq"]))
	 	if !succ {
			return
	  }
	}

	application.StartIdx = common.ToUint(startIdx)
	application.RowCnt = common.ToUint(rowCnt)
	application.SearchWords = searchWords
	list := application.LoadList(moreSelct, moreJoin, moreCondition, appViewType)

	pageInfo := gin.H {
		"startIdx": application.StartIdx,
		"totalCnt": application.GetAppCnt(moreJoin, moreCondition, appViewType),
	}

	common.FinishApi(c, common.Api_status_ok,
    gin.H{
      "rt": "ok",
      "list" : list,
			"pageInfo" : pageInfo,
    })
}

// @Tags Application
// @Summary Application 신청서 정보
// @Description Application 신청서 정보
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Router /application/{application_seq}/info [get]
// @Success 200
func Application_Info(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq, succ, _, _, _ := getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	if !succ {  return  }

	application := model.Application{
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	info := application.Load()
	common.FinishApi(c, common.Api_status_ok, info)
}

// @Tags Application
// @Summary Application 승인후 신청서 목록 조회
// @Description Application 승인후 신청서 목록 조회
// @Description app_view_type(없으면 전체 리스트 출력)
// @Description 1 행정(간사, 담당자) 승인후 관리 > 전체 실험
// @Description 2 행정(간사, 담당자) 승인후 관리 > 진행중인 실험
// @Description 3 연구원 승인 후 관리 > 실험 수행 및 서류
// @Description 4 위원장 진행중인 실험
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param filter.app_view_type query string false "1"
// @Param filter.start_index query string false   "page start index"
// @Param filter.row_cnt query string false       "row count"
// @Param search_words  query string false        "Search Words"
// @Router /application-approved [get]
// @Success 200
func Application_Approved_List(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	appViewType := common.ToStr(c.Request.URL.Query().Get("filter.app_view_type"))
	startIdx := c.Request.URL.Query().Get("filter.start_index")
	rowCnt   := c.Request.URL.Query().Get("filter.row_cnt")
	searchWords := c.Request.URL.Query().Get("search_words")

	if "" == startIdx {
	 startIdx = "0"
	} else if _, ret := common.ToValidateUint64(startIdx); false == ret {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "filter.start_index값이 잘못되었습니다.")
		return
	}

	if "" == rowCnt {
		rowCnt = "10"
	} else if _, ret := common.ToValidateUint64(rowCnt); false == ret {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "filter.page_count값이 잘못되었습니다.")
		return
	}

	moreCondition := ""
	succ := true
	if "" != appViewType {
		succ, moreCondition = checkApprovedAppViewAuthAndGetCondition(c, userTypeArr, common.ToUint(appViewType), common.ToUint(tokenMap["user_seq"]))
	 	if !succ {
			return
	  }
	}

	application := model.Application{
		LoginToken : tokenMap,
		StartIdx : common.ToUint(startIdx),
		RowCnt : common.ToUint(rowCnt),
		SearchWords : searchWords,
	}

	list := application.LoadApprovedList(moreCondition, appViewType)

	pageInfo := gin.H {
		"startIdx": application.StartIdx,
		"totalCnt": application.GetApprovedAppCnt(moreCondition),
	}

	common.FinishApi(c, common.Api_status_ok,
    gin.H{
      "rt": "ok",
      "list" : list,
			"pageInfo" : pageInfo,
    })
}

// @Tags Application
// @Summary Application (재승인, 변경 승인) 가능한 신청서 목록 조회
// @Description Application (재승인, 변경 승인) 가능한 신청서 목록 조회
// @Description app_view_type(없으면 전체 리스트 출력)
// @Description 1 연구원 승인 신청서 관리  > IACUC 재승인 가능한 신청서 리스트
// @Description 2 연구원 승인 신청서 관리  > IACUC 변경 승인 가능한 신청서 리스트
// @Description 3 연구원 승인 신청서 관리  > IBC 변경 심의 가능한 신청서 리스트
// @Description 4 연구원 승인 신청서 관리  > IRB 변경 심의 가능한 신청서 리스트
// @Description 5 연구원 승인 신청서 관리  > IRB 지속심의(중간보고) 가능한 신청서 리스트
// @Description 6 연구원 승인 신청서 관리  > IACUC 보완후 재심 리스트
// @Description 7 연구원 승인 신청서 관리  > IBC 보완후 재심 리스트
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param filter.app_view_type query string false "1"
// @Router /application-possible [get]
// @Success 200
func Application_Possible_List(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	appViewType := common.ToStr(c.Request.URL.Query().Get("filter.app_view_type"))
	moreCondition := ""
	succ := true
	if "" != appViewType {
		succ, moreCondition = checkPossibleAppViewAuthAndGetCondition(c, userTypeArr, common.ToUint(appViewType), common.ToUint(tokenMap["user_seq"]))
		if !succ {
			return
		}
	}

	application := model.Application{
		LoginToken : tokenMap,
	}

	list := application.LoadPossibleList(moreCondition, appViewType)
	common.FinishApi(c, common.Api_status_ok,
		gin.H{
			"rt": "ok",
			"list" : list,
		})
}

// @Tags Application
// @Summary Application 신청서 페이지 정보 조회
// @Description Application 신청서 페이지 정보 조회
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Param filter.query_items query string true   "item_name,item_name"
// @Param filter.guide_items query string false		"item_name,item_name"
// @Param filter.parent_items query string false  "item_name,item_name"
// @Param filter.child_app_seq query string false   "2"
// @Param filter.child_items query string false   "item_name,item_name"
// @Router /application/{application_seq} [get]
// @Success 200
func Application_Get(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq := uint64(0);
	succ := false;
	institution_seq := 0;
	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		app_seq, succ, _, _, _ = getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	  if !succ {  return  }
	} else {
		app_seq, succ, institution_seq = getApplicationSeqFromPathForAdmin(c)
		if !succ {  return  }
		tokenMap["institution_seq"] = institution_seq
	}

  items := c.Request.URL.Query().Get("filter.query_items")
  itemArr := strings.Split(common.ToStr(items), ",")

  tmpItem := model.Item{}
  succ, err_item := tmpItem.ValidateItemArrs(itemArr)
  if !succ {
    msg := fmt.Sprintf("query_items의 [%v]는 없는 item_name입니다.", err_item)
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, msg)
    return
  }

	items_guide := c.Request.URL.Query().Get("filter.guide_items")
	itemGuideArr := strings.Split(common.ToStr(items_guide), ",")
  succ, err_item = tmpItem.ValidateItemArrs(itemGuideArr)
  if !succ {
    msg := fmt.Sprintf("guide_items의 [%v]는 없는 item_name입니다.", err_item)
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, msg)
    return
  }

	items_parent := c.Request.URL.Query().Get("filter.parent_items")
	itemParentArr := strings.Split(common.ToStr(items_parent), ",")
	succ, err_item = tmpItem.ValidateItemArrs(itemParentArr)
	if !succ {
		msg := fmt.Sprintf("parent_items의 [%v]는 없는 item_name입니다.", err_item)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, msg)
		return
	}

	items_child := c.Request.URL.Query().Get("filter.child_items")
	itemChildArr := strings.Split(common.ToStr(items_child), ",")
	succ, err_item = tmpItem.ValidateItemArrs(itemChildArr)
	if !succ {
		msg := fmt.Sprintf("child_items의 [%v]는 없는 item_name입니다.", err_item)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, msg)
		return
	}

	application := model.Application{
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	if !application.Init() {
    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
    return
	}

	judge_type := common.ToInt(application.Data["judge_type"])
	if 0 == judge_type {
		judge_type = 1;
	}

	data := application.LoadItemList(itemArr, false, judge_type)

	if len(itemParentArr) > 0 && nil != application.Data &&
	 	 common.ToUint(application.Data["parent_app_seq"]) > 0		{
		 appParent := model.Application{
	 		LoginToken : tokenMap,
	 		Application_seq : common.ToUint(application.Data["parent_app_seq"]),
	 	}

		if !appParent.Init() {
			common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
			return
		}
		data2 := appParent.LoadItemList(itemParentArr, false, judge_type)
		data["parent_item"] = data2;
	}

	items_app_seq := common.ToUint(c.Request.URL.Query().Get("filter.child_app_seq"))
	if len(itemChildArr) > 0 && items_app_seq > 0 {
		app_Child := model.Application{
			LoginToken : tokenMap,
			Application_seq : items_app_seq,
		}

		if !app_Child.Init() {
			common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
			return
		}

		if app_seq != common.ToUint64(app_Child.Data["parent_app_seq"])	{
			common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_app_seq_mismatch)
			return;
		}

		data2 := app_Child.LoadItemList(itemChildArr, false, judge_type)
		data["child_item"] = data2;
	}

	retJson :=  gin.H{
								"rt": "ok",
								"data" : data,
							}

	guideIns := model.ItemGuide {}
	guide_map := guideIns.LoadGuide(itemGuideArr);
	if nil != guide_map	{
		retJson["guide"] = guide_map
	}

	common.FinishApi(c, common.Api_status_ok, retJson)
}

// @Tags Application
// @Summary Application 신청서 페이지 정보 수정 및 임시저장
// @Description Application 신청서 페이지 정보 수정 및 임시저장
// @Accept  mpfd
// @Produce  mpfd
// @Security ApiKeyAuth
// @Param general_ref formData file false "1.6 관련자료첨부"
// @Param animal_species_reason formData file false "3.2 종/계통 선택한 합리적인 이유"
// @Param animal_cnt_reason formData file false "3.3 동물수에 대한 합리적인 근거"
// @Param animal_exp_summary formData file false "5.1 동물실험의 개요 및 일정"
// @Param animal_exp_surgical_method formData file false "5.5 처치방법"
// @Param pain_relief_psych_m_license formData file false "7.3 사용허가증 첨부"
// @Param pain_relief_animal_m_license formData file false "7.4 처방전 첨부"
// @Param application_seq path uint true "1"
// @Param param formData string true "json format"
// @Router /application/{application_seq} [patch]
// @Success 200
func Application_Patch(c *gin.Context) {
	application_Patch_or_Post(c, false)
}

// @Tags Application
// @Summary Application 신청서 제출
// @Description Application 신청서 제출
// @Description filter.submit_type(없으면 신청서 제출)
// @Description 1 보완 요청(행정간사, 행정담당만 가능)
// @Description 2 행정 검토 1단계 완료(행정간사, 행정담당만 가능)
// @Description 3 행정 검토 2단계 완료(행정간사, 행정담당만 가능)
// @Description 4 전문심사 1단계 완료 (심사위원만 가능)
// @Description 5 전문심사 2단계 완료 (심사위원만 가능)
// @Description 6 일반심사 (심사위원만 가능)
// @Description 7 심사종료설정 완료 (행정간사,담당만 가능)
// @Description 8 최종심의 승인(행정간사 또는 위원장 가능)
// @Description 9 최종심의 조건부 승인(행정간사 또는 위원장 가능)
// @Description 10 최종심의 반려 (행정간사 또는 위원장 가능)
// @Description 11 최종심의 보완후 재심 (행정간사 또는 위원장 가능)
// @Description 12 과제수행 실험 종료(연구원)
// @Description 13 과제수행 과제 종료(행정간사 또는 위원장 가능)
// @Description 신청서 제출 -> 행정검토 -> 전문심사 -> 일반심사 -> 심사종료설정 -> 최종심의 -> 과제수행
// @Description 이전상태가 안 맞으면 에러
// @Accept  mpfd
// @Produce  mpfd
// @Security ApiKeyAuth
// @Param general_ref formData file false "1.6 관련자료첨부"
// @Param animal_species_reason formData file false "3.2 종/계통 선택한 합리적인 이유"
// @Param animal_cnt_reason formData file false "3.3 동물수에 대한 합리적인 근거"
// @Param animal_exp_summary formData file false "5.1 동물실험의 개요 및 일정"
// @Param animal_exp_surgical_method formData file false "5.5 처치방법"
// @Param pain_relief_psych_m_license formData file false "7.3 사용허가증 첨부"
// @Param pain_relief_animal_m_license formData file false "7.4 처방전 첨부"
// @Param filter.submit_type query string false "1"
// @Param filter.iacuc_seq query string false "1"
// @Param application_seq path uint true "1"
// @Param param formData string true "json format"
// @Router /application/{application_seq} [post]
// @Success 200
func Application_Post(c *gin.Context) {
	application_Patch_or_Post(c, true)
}

func application_Patch_or_Post(c *gin.Context, submit bool) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	app_seq, succ, now_app_step, _, reg_user_seq := getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	if !succ {  return  }

	application := model.Application{
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	if !application.Init() {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	app_step		:= model.DEF_APP_STEP_WRITE
	app_result	:= model.DEF_APP_RESULT_TEMP
	succ				= true
	if submit {
		app_step		= model.DEF_APP_STEP_CHECKING
		app_result	= model.DEF_APP_RESULT_CHECKING
	}

	submitType := common.ToStr(c.Request.URL.Query().Get("filter.submit_type"))
	if "" != submitType {
		succ, app_step, app_result = checkSubmitTypeAuthAndGetAppStepAndResult(c, userTypeArr, common.ToUint(submitType), application)
		if !succ {
			return
		}
	} else {
		if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_RESEARCHER) {
			common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
			return
		}
	}

	if submit {
		if (common.ToUint(submitType) != model.DEF_SUBMIT_TYPE_CHECKING_FAST_FINISH)	{
			//	행정검토 신속 단계는 확인을 안한다.
			if math.Abs(cast.ToFloat64(app_step - now_app_step)) > 1 {
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_app_step)
				return
			}
		}

		if app_step == model.DEF_APP_STEP_CHECKING && app_result == model.DEF_APP_RESULT_CHECKING && common.ToUint(submitType) != model.DEF_SUBMIT_TYPE_CHILD_SUBMIT{
			// 건수 counting: 승인신청서 작성 및 자가점검을 완료하여 신청서를 행정간사에 제출완료한 경우 1건으로 계산
			instt := model.Institution{
				Institution_seq : common.ToUint(tokenMap["institution_seq"]),
			}
			if !instt.CheckPossibleAppSubmit() {
				common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_limit_monthly_utilization)
				return
			}
		}

	}

	data, eMsg := common.UnmarshalFormData(c, "param")
	if data == nil {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	itemArr := make([]string, 0, 0)
	for item_name, _ := range data {
		itemArr = append(itemArr, item_name)
	}

	tmpItem := model.Item{}
	succ, err_item := tmpItem.ValidateItemArrs(itemArr)
	if !succ {
		msg := fmt.Sprintf("[%v]는 없는 item_name입니다.(1)", err_item)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, msg)
		return
	}

	succ, err_item, err_msg := application.UpdateItemsFromJson(c, data, submit, app_step, app_result, common.ToUint(submitType), reg_user_seq)
	if !succ {
		msg := fmt.Sprintf("[%v] %v", err_item, err_msg)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, msg)
		return
	}

	//	성공
	if (application.Application_seq > 0 && submit)	{
		msgMgr := model.MessageMgr	{
			Application_seq : application.Application_seq,
		}
		user := model.User{}
		// 책임 연구자에게 전송
		user.User_seq = reg_user_seq
		// user.User_seq = application.GetGeneralDirectorSeq(reg_user_seq)
		user.Load()
		switch(common.ToUint(submitType))	{
			case model.DEF_SUBMIT_TYPE_SUPPLEMENT : // 행정검토에서 보완 요청시 작성자에게 메세지 전송!
				msgMgr.Msg_ID = model.DEF_MSG_REQUEST_SUPPLEMENT
				msgMgr.User_info = user.Data
				go msgMgr.SendMessage();
			case model.DEF_SUBMIT_TYPE_CHECKING_2 :
				msgMgr.Msg_ID = model.DEF_MSG_EXPER_JUDGE_START	// 전문심사 게시
				go msgMgr.SendMsgExpert();
				// 전문심사 개시 안내(책임연구자에게 전송)
				msgMgr2 := model.MessageMgr {
					Application_seq : application.Application_seq,
					Msg_ID : model.DEF_MSG_EXPER_JUDGE_START_TO_LEADER,
					User_info : user.Data,
				}
				go msgMgr2.SendMessage();
			case model.DEF_SUBMIT_TYPE_JUDGE_PRO_2 : // 전문심사 2단계 종료! 일반심사위원 들에게 메세지 전송
				msgMgr.Msg_ID = model.DEF_MSG_NORMAL_JUDGE_START
				go msgMgr.SendMsgNormal();
				// 일반심사 개시 안내(책임연구자에게 전송)
				msgMgr2 := model.MessageMgr {
					Application_seq : application.Application_seq,
					Msg_ID : model.DEF_MSG_NORMAL_JUDGE_START_TO_LEADER,
					User_info : user.Data,
				}
				go msgMgr2.SendMessage();
			case model.DEF_SUBMIT_TYPE_JUDGE_FINAL_A,
			 		 model.DEF_SUBMIT_TYPE_JUDGE_FINAL_AC,
					 model.DEF_SUBMIT_TYPE_JUDGE_FINAL_REJECT,
					 model.DEF_SUBMIT_TYPE_JUDGE_FINAL_REQUIRE_RETRY : // 최종심의 일때
				msgMgr.Msg_ID = model.DEF_MSG_JUDGE_FINISHED
				msgMgr.User_info = user.Data
				go msgMgr.SendMessage();

				user2 := model.User{}
				moreCondition := fmt.Sprintf(`AND user.user_seq != %v
																			AND ins.institution_seq = %v
																			AND user.user_type LIKE '%%%v%%'`,
																			user.Data["user_seq"],
																			user.Data["institution_seq"],
																			model.DEF_USER_TYPE_CHAIRPERSON);
				rows := user2.GetInstitutionUserList(moreCondition);
				for _, row := range rows {
					msgMgr2 := model.MessageMgr {
						Application_seq : application.Application_seq,
						Msg_ID : model.DEF_MSG_JUDGE_FINISHED,
						User_info : row,
					}
					go msgMgr2.SendMessage();
				}

		}
	}

	retJson := gin.H{	"rt": "ok"}

	if (uint(app_seq) == uint(0))	{
		//	신규 생성 : 생성된 기본 정보 내려 줌.
		log.Println("new app_seq : ", application.Application_seq)
		retJson["new_info"] = application.Load()
	}
	retJson["files"] = application.LoadItemList(itemArr, true, common.ToInt(application.Data["judge_type"]));

	common.FinishApi(c, common.Api_status_ok, retJson)
}

func checkSubmitTypeAuthAndGetAppStepAndResult(c *gin.Context, userTypeArr []string, submitType uint, application model.Application) (succ bool, app_step int, app_result int) {
	switch submitType {
		case model.DEF_SUBMIT_TYPE_SUPPLEMENT:	// 보완 요청(행정간사, 행정담당만 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_WRITE
			app_result	= model.DEF_APP_RESULT_SUPPLEMENT
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_CHECKING:	// 행정 검토 1단계 완료(행정간사, 행정담당만 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_CHECKING
			app_result	= model.DEF_APP_RESULT_CHECKING_2
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_CHECKING_2:	// 행정 검토 2단계 완료(행정간사, 행정담당만 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_PRO
			app_result	= model.DEF_APP_RESULT_JUDGE_ING
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_JUDGE_PRO:	// 전문심사 1단계 완료 (심사위원만 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_COMMITTEE) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_PRO
			app_result	= model.DEF_APP_RESULT_JUDGE_ING_2
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_JUDGE_PRO_2:	 // 전문심사 2단계 완료 (심사위원 또는 행정간사, 행정담당 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_COMMITTEE, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_NORMAL
			app_result	= model.DEF_APP_RESULT_JUDGE_ING
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_JUDGE_NORMAL:	// 일반심사 심사 완료(심사위원만 가능) 모든사람이 제출 했을때만 완료 상태가 된다
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_COMMITTEE, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_FINAL
			app_result	= model.DEF_APP_RESULT_DECISION_ING
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_JUDGE_RESUME:	// 일반심사 심사재개 (심사위원만 가능) 모든사람이 제출 했을때만 완료 상태가 된다
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_COMMITTEE) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_NORMAL
			app_result	= model.DEF_APP_RESULT_JUDGE_ING
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_JUDGE_FINISH:	// 심사종료설정 완료 (행정 간사,업무 만 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_FINAL
			app_result	= model.DEF_APP_RESULT_DECISION_ING
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_JUDGE_FINAL_A:	// 최종심의 승인(행정간사 또는 위원장 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_CHAIRPERSON) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_PERFORMANCE
			app_result	= model.DEF_APP_RESULT_EXPER_ING_A
			//	신규 승인이 아니면 : 최종심의의 승인 단계로 설정
			switch(common.ToInt(application.Data["application_type"]))	{
			case model.DEF_APP_TYPE_CHANGE, model.DEF_APP_TYPE_RENEW, model.DEF_APP_TYPE_BRINGIN, model.DEF_APP_TYPE_CHECKLIST, model.DEF_APP_TYPE_FINISH :
					app_step		= model.DEF_APP_STEP_FINAL
					app_result	= model.DEF_APP_RESULT_APPROVED
					break;
			}
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_JUDGE_FINAL_AC:	// 최종심의 조건부 승인(행정간사 또는 위원장 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_CHAIRPERSON) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_PERFORMANCE
			app_result	= model.DEF_APP_RESULT_EXPER_ING_AC
			//	신규 승인이 아니면 : 최종심의의 조건부 승인 단계로 설정
			switch(common.ToInt(application.Data["application_type"]))	{
				case model.DEF_APP_TYPE_CHANGE, model.DEF_APP_TYPE_RENEW, model.DEF_APP_TYPE_BRINGIN, model.DEF_APP_TYPE_CHECKLIST, model.DEF_APP_TYPE_FINISH :
					app_step		= model.DEF_APP_STEP_FINAL
					app_result	= model.DEF_APP_RESULT_APPROVED_C
					break;
			}
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_JUDGE_FINAL_REJECT:	// 최종심의 반려 (행정간사 또는 위원장 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_CHAIRPERSON) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_FINAL
			app_result	= model.DEF_APP_RESULT_REJECT
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_JUDGE_FINAL_REQUIRE_RETRY:	// 최종심의 보완후 재심 (행정간사 또는 위원장 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_CHAIRPERSON) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_FINAL
			app_result	= model.DEF_APP_RESULT_REQUIRE_RETRY
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_EXPER_FINISH:	// 과제수행 실험 종료(연구원)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_RESEARCHER) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_PERFORMANCE
			app_result	= model.DEF_APP_RESULT_EXPER_FINISH
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_TASK_FINISH:	 // 과제수행 과제 종료(행정간사 또는 위원장 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK, model.DEF_USER_TYPE_CHAIRPERSON) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_PERFORMANCE
			app_result	= model.DEF_APP_RESULT_TASK_FINISH
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_RETRY_CHECKING:	//	행정 검토 1단계로 돌아가기 (행정간사, 행정담당만 가능)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_CHAIRPERSON) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_CHECKING
			app_result	= model.DEF_APP_RESULT_CHECKING
			succ = true
			return

		// 부속 신청서 공통
		case model.DEF_SUBMIT_TYPE_CHECKING_FAST_FINISH:	// 부속 신청서들 : 행정검토 신속 완료 (최종 심의로 단계이동)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_FINAL
			app_result	= model.DEF_APP_RESULT_DECISION_ING
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_CHILD_SUBMIT:	//	부속 신청서 제출 (변경신청서 등)
			// donghun : 21-04-23 승인후 점검 신청서는 모든인원 이 할수 있다..
			app_step		= model.DEF_APP_STEP_CHECKING
			app_result	= model.DEF_APP_RESULT_CHECKING
			succ = true
			return
		case model.DEF_SUBMIT_TYPE_JUMP_FINAL:	//	최종심의 단계로 이동(일반심사 지연 일때)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			app_step		= model.DEF_APP_STEP_FINAL
			app_result	= model.DEF_APP_RESULT_DECISION_ING
			succ = true
			return
		default :
			common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_invalide_params, "submit_type이 존재하지 않습니다.")
			return
	}
	return
}

func checkAppViewAuthAndGetCondition(c *gin.Context, userTypeArr []string, appViewType uint, user_seq uint) (succ bool, moreSelect string, moreJoin string, moreCondition string) {
	switch appViewType {
		case model.DEF_APP_VIEW_ADMIN_ALL:	// 행정(간사, 담당자) 전체 계획서 및 보고서
		 	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
		    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
		 	}
			// 21-10-07 수정 요청 사항 : 전문 위원을 리스트에 표시
			moreSelect = fmt.Sprintf(` , IFNULL((SELECT user.name
																						 FROM t_application_member member, t_user user
																						WHERE member.user_seq = user.user_seq
																							AND member.item_name = 'expert_member'
																							AND app.application_seq = member.application_seq), '-') as expert_member`)

			// 전체목록 (임시저장, 보완중 제외!)
			moreCondition = fmt.Sprintf(` AND app.application_result NOT IN(%d, %d)`,
																		model.DEF_APP_RESULT_TEMP, model.DEF_APP_RESULT_DELETED)
			succ = true
			return
		case model.DEF_APP_VIEW_ADMIN_REVIEW_OFFICE:	// 행정(간사, 담당자) 행정 검토
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			moreCondition = fmt.Sprintf(` AND app.application_step = %d
																		AND app.application_result IN(%d, %d)`,
																		model.DEF_APP_STEP_CHECKING, model.DEF_APP_RESULT_CHECKING, model.DEF_APP_RESULT_CHECKING_2)
			succ = true
			return
		case model.DEF_APP_VIEW_ADMIN_REVIEW_CLOSE:	// 행정(간사, 담당자) 심사 종료 설정
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			moreCondition = fmt.Sprintf(` AND ((app.application_step = %d AND app.application_result IN(%d, %d, %d)) or
																				 (app.application_step = %d AND app.application_result IN(%d, %d, %d)) )`,
																	  model.DEF_APP_STEP_PRO,
																			model.DEF_APP_RESULT_JUDGE_ING,
																			model.DEF_APP_RESULT_JUDGE_ING_2,
																			model.DEF_APP_RESULT_JUDGE_DELAY,
																		model.DEF_APP_STEP_NORMAL,
																			model.DEF_APP_RESULT_JUDGE_ING,
																			model.DEF_APP_RESULT_JUDGE_ING_2,
																			model.DEF_APP_RESULT_JUDGE_DELAY)
			succ = true
			return
		case model.DEF_APP_VIEW_ADMIN_REVIEW_CONFIRM:	// 행정간사 최종 심의
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			// moreCondition = fmt.Sprintf(` AND app.application_step =  %d
			// 															AND app.application_result = %d
			// 													    AND NOT app.judge_type = IF(istt.ia_final_director = %d, %d, 0)`,
			// 															model.DEF_APP_STEP_FINAL, model.DEF_APP_RESULT_DECISION_ING,
			// 															model.DEF_FINAL_DIRECTOR_CHAIRPERSON, model.DEF_APP_JUDGE_TYPE_CODE_IACUC)
			moreCondition = fmt.Sprintf(` AND app.application_step =  %d
																		AND app.application_result = %d
																		AND istt.ia_final_director = %d`,
																		model.DEF_APP_STEP_FINAL, model.DEF_APP_RESULT_DECISION_ING,
																		model.DEF_FINAL_DIRECTOR_ALL)
			succ = true
			return
		case model.DEF_APP_VIEW_RESEARCHER_ALL:	// 연구원 실험계획서
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_RESEARCHER) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			moreCondition = fmt.Sprintf(` AND app.reg_user_seq = %d
																		AND app.application_result <> %d`, user_seq, model.DEF_APP_RESULT_DELETED)
			succ = true
			return
		case model.DEF_APP_VIEW_CHAIRMAN_REVIEW_CONFIRM:	// 위원장 심사진행(최종심의 심의중)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_CHAIRPERSON) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			moreCondition = fmt.Sprintf(` AND app.application_step =  %d
																		AND app.application_result = %d`,
																		model.DEF_APP_STEP_FINAL, model.DEF_APP_RESULT_DECISION_ING)
			succ = true
			return
		case model.DEF_APP_VIEW_COMMITTEE_REVIEW:	// 심사위원 심사진행 (자기 자신것만 봐야됨!!)
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_COMMITTEE) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
			moreJoin = ` LEFT OUTER JOIN t_application_member member ON (app.application_seq = member.application_seq)`
			// 21-03-03 donghun : 우선 test를 위해서 expert_review_dttm, normal_review_dttm 만 조건 추가!
			//	wowdolf :
			//	normal_review_dttm 와 normal_review_result 값이 없는 경우만 노출되어야 함.!
			moreCondition = fmt.Sprintf(` AND IF(app.application_step = %d,  member.item_name IN ('expert_member'), member.item_name IN ( 'committee_ex_member', 'committee_in_member'))
																		AND app.application_step IN (%d, %d)
																		AND app.application_result IN (%d, %d, %d)
																		AND member.user_seq = %v
																		AND IF(app.application_step = %d,
																						IF(
																							IFNULL(
																								(
																									SELECT
																										application_seq
																									FROM
																										t_application_etc
																									WHERE
																										item_name = "expert_review_dttm"
																										AND application_seq = app.application_seq
																								),
																								0
																							) > 0,
																							1,
																							0
																						),
																						IF(
																							IFNULL(
																								(
																									SELECT
																										application_seq
																									FROM
																										t_application_etc
																									WHERE
																										item_name = "normal_review_dttm"
																										AND target_item = %v
																										AND application_seq = app.application_seq
																								),
																								0
																							) > 0,
																							1,
																							0
																						)
																					) = 0 `,
																		model.DEF_APP_STEP_PRO,
																		model.DEF_APP_STEP_PRO, model.DEF_APP_STEP_NORMAL,
																		model.DEF_APP_RESULT_JUDGE_ING, model.DEF_APP_RESULT_JUDGE_ING_2, model.DEF_APP_RESULT_JUDGE_DELAY,
																		user_seq, model.DEF_APP_STEP_PRO, user_seq)
			succ = true
			return
		case model.DEF_APP_VIEW_CHAIRPERSON_RECORD:	// 위원장 심사기록(최종 승인 기록)
				if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_CHAIRPERSON) {
					common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
					return
				}
				moreCondition = fmt.Sprintf(` AND app.approved_user_seq = %d
																			AND app.application_result <> %d`, user_seq, model.DEF_APP_RESULT_DELETED)
				succ = true
				return
		case model.DEF_APP_VIEW_COMMITTEE_RECORD:	// 심사위원 심사기록
				if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_COMMITTEE) {
					common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
					return
				}
				moreSelect = fmt.Sprintf(` , member.item_name as committee,
																		 IFNULL((SELECT etc.contents
																							 FROM t_application_etc etc
																							WHERE app.application_seq = etc.application_seq
																								AND etc.target_item = %d
																								AND etc.item_name = 'expert_review_dttm'), 0) as expert_review_dttm,
																		 IFNULL((SELECT etc.contents
																						 	 FROM t_application_etc etc
																						  WHERE app.application_seq = etc.application_seq
																								AND etc.target_item = %d
																							  AND etc.item_name = 'normal_review_dttm'), 0) as normal_review_dttm`, user_seq, user_seq)
				moreJoin = ` LEFT OUTER JOIN t_application_member member ON (app.application_seq = member.application_seq)`
				moreCondition = fmt.Sprintf(` AND member.user_seq = %d
																			AND member.item_name IN ('committee_ex_member','committee_in_member','expert_member')
																			AND (
																					IF(
																						IFNULL(
																							(
																								SELECT
																									application_seq
																								FROM
																									t_application_etc
																								WHERE
																									(
																										IF (
																											member.item_name = 'expert_member',
																											"expert_review_dttm", "normal_review_dttm"
																										)
																									) = item_name
																									AND target_item = member.user_seq
																									AND application_seq = app.application_seq
																							),
																							0
																						) > 0,
																						1,
																						0
																					)
																				) = 1
																			AND app.application_result <> %d
																				`,user_seq, model.DEF_APP_RESULT_DELETED)
				succ = true
				return
		case model.DEF_APP_VIEW_APPROVED_INSPECT_RECORD:	// 승인후 점검 기록
				// 따로 권한 체크 안함!
				moreCondition = fmt.Sprintf(` AND app.application_type = %d
																			AND app.reg_user_seq = %d
																			AND app.application_result <> %d
																		`, model.DEF_APP_TYPE_CHECKLIST, user_seq, model.DEF_APP_RESULT_DELETED)
				succ = true
				return
		case model.DEF_APP_VIEW_MY_INSPECT_CHECK:	// 내가 점검할 실험
				// 따로 권한 체크 안함!
				moreSelect = " , user2.name AS check_user_name, app.check_user_seq"
				moreJoin = ` LEFT OUTER JOIN t_user user2 ON (app.check_user_seq = user2.user_seq)`
				moreCondition = fmt.Sprintf(` AND app.check_user_seq = %d
																			AND NOT EXISTS (SELECT app2.application_seq
																 												FROM t_application app2
				 																							 WHERE app2.parent_app_seq = app.application_seq
				 																						 		 AND app2.application_type = %d)
																		`, user_seq, model.DEF_APP_TYPE_CHECKLIST)
				succ = true
				return
		case model.DEF_APP_VIEW_IACUC_FOR_IBC: // ibc 에서 iacuc 참조 가능한 리스트
			moreCondition = fmt.Sprintf(` AND app.judge_type = %d
																		AND app.application_type = %d
																		AND app.application_step =  %d
																		AND app.application_result IN (%d, %d, %d, %d) `,
																		model.DEF_APP_JUDGE_TYPE_CODE_IACUC,
																		model.DEF_APP_TYPE_NEW,
																		model.DEF_APP_STEP_PERFORMANCE,
																		model.DEF_APP_RESULT_EXPER_ING_A, model.DEF_APP_RESULT_EXPER_ING_AC,
																		model.DEF_APP_RESULT_EXPER_FINISH, model.DEF_APP_RESULT_TASK_FINISH)
			succ = true
			return
		case model.DEF_APP_VIEW_IBC_FOR_IACUC: // iacuc 에서 ibc 참조 가능한 리스트
			moreCondition = fmt.Sprintf(` AND app.judge_type = %d
																		AND app.application_type = %d
																		AND app.application_step =  %d
																		AND app.application_result IN (%d, %d, %d, %d) `,
																		model.DEF_APP_JUDGE_TYPE_CODE_IBC,
																		model.DEF_APP_TYPE_NEW,
																		model.DEF_APP_STEP_PERFORMANCE,
																		model.DEF_APP_RESULT_EXPER_ING_A, model.DEF_APP_RESULT_EXPER_ING_AC,
																		model.DEF_APP_RESULT_EXPER_FINISH, model.DEF_APP_RESULT_TASK_FINISH)
			succ = true
			return
		default :
			common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_invalide_params, "app_view_type이 존재하지 않습니다.")
			return
	}
}

func checkApprovedAppViewAuthAndGetCondition(c *gin.Context, userTypeArr []string, appViewType uint, user_seq uint) (succ bool, moreCondition string) {
	switch appViewType {
			case model.DEF_APP_VIEW_APPROVED_ADMDIN_ALL:	// 승인후 관리 > 전체 실험
				if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
			    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
					return
			 	}

				moreCondition = fmt.Sprintf(` AND app.application_step =  %d
																			AND app.application_result IN(%d, %d, %d, %d)`,
																			model.DEF_APP_STEP_PERFORMANCE,
																			model.DEF_APP_RESULT_EXPER_ING_A, model.DEF_APP_RESULT_EXPER_ING_AC,
																			model.DEF_APP_RESULT_EXPER_FINISH, model.DEF_APP_RESULT_TASK_FINISH)
				succ = true
				return
			case model.DEF_APP_VIEW_APPROVED_ADMDIN_ING:	// 승인후 관리 > 진행중인 실험
				if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
			    common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
					return
			 	}

				moreCondition = fmt.Sprintf(` AND app.application_step = %d
																			AND app.application_result IN(%d, %d, %d)`,
																			model.DEF_APP_STEP_PERFORMANCE,
																			model.DEF_APP_RESULT_EXPER_ING_A, model.DEF_APP_RESULT_EXPER_ING_AC,
																			model.DEF_APP_RESULT_EXPER_FINISH)
				succ = true
				return
			case model.DEF_APP_VIEW_APPROVED_RESEARCHER_ALL:	// 승인후 관리 > 실험 수행 및 서류
				// donghun : 21-05-17 general_director, general_expt 참여한 유저들 전부 보이게 처리 해야됨
				if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_RESEARCHER) {
					common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
					return
				}

				moreCondition = fmt.Sprintf(` AND app.application_step =  %d
																			AND app.application_result IN(%d, %d, %d, %d)
																			AND IFNULL((SELECT DISTINCT member.user_seq
																										FROM t_application_member member
																								   WHERE member.application_seq = app.application_seq
																									   AND member.item_name IN ('general_director', 'general_expt')
																										 AND member.user_seq = %d), app.reg_user_seq) = %d`,

																			model.DEF_APP_STEP_PERFORMANCE,
																			model.DEF_APP_RESULT_EXPER_ING_A, model.DEF_APP_RESULT_EXPER_ING_AC,
																			model.DEF_APP_RESULT_EXPER_FINISH, model.DEF_APP_RESULT_TASK_FINISH,
																			user_seq, user_seq)
				succ = true
				return
			case model.DEF_APP_VIEW_APPROVED_CHAIRPERSON_ALL:	// 진행중인 실험
				if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_CHAIRPERSON) {
					common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
					return
				}

				moreCondition = fmt.Sprintf(` AND app.application_step =  %d
																			AND app.application_result IN(%d, %d, %d)`,
																			model.DEF_APP_STEP_PERFORMANCE,
																			model.DEF_APP_RESULT_EXPER_ING_A, model.DEF_APP_RESULT_EXPER_ING_AC,
																			model.DEF_APP_RESULT_EXPER_FINISH)
				succ = true
				return
		default :
			common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_invalide_params, "app_view_type이 존재하지 않습니다.")
			return
	}
}

func checkPossibleAppViewAuthAndGetCondition(c *gin.Context, userTypeArr []string, appViewType uint, user_seq uint) (succ bool, moreCondition string) {
	// 연구원 권한이 있는 사람만 가능함
	if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_RESEARCHER) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	switch appViewType {
			case model.DEF_APP_VIEW_POSSIBLE_IA_RETRY:	// IACUC 재승인 가능한 신청서 리스트
				moreCondition = fmt.Sprintf(` AND app.judge_type = %d
																			AND app.parent_app_seq = 0
																			AND app.application_step = %d
																			AND app.application_result IN (%d, %d)
																			AND app.reg_user_seq = %d`,
																			model.DEF_APP_JUDGE_TYPE_CODE_IACUC,
																			model.DEF_APP_STEP_PERFORMANCE,
																			model.DEF_APP_RESULT_EXPER_ING_A,
																			model.DEF_APP_RESULT_EXPER_ING_AC,
																		  user_seq)
				succ = true
				return
			case model.DEF_APP_VIEW_POSSIBLE_IA_CHANGE:	// IACUC 변경 승인 가능한 신청서 리스트
				moreCondition = fmt.Sprintf(` AND app.judge_type = %d
																			AND app.parent_app_seq = 0
																			AND app.application_step = %d
																			AND app.application_result IN (%d, %d)
																			AND app.reg_user_seq = %d`,
																		  model.DEF_APP_JUDGE_TYPE_CODE_IACUC,
																			model.DEF_APP_STEP_PERFORMANCE,
																			model.DEF_APP_RESULT_EXPER_ING_A,
																			model.DEF_APP_RESULT_EXPER_ING_AC,
																			user_seq)
				succ = true
				return
			case model.DEF_APP_VIEW_POSSIBLE_IBC_CHANGE:	// IBC 변경 승인 가능한 신청서 리스트
				moreCondition = fmt.Sprintf(` AND app.judge_type = %d
																			AND app.parent_app_seq = 0
																			AND app.application_step = %d
																			AND app.application_result IN (%d, %d)
																			AND app.reg_user_seq = %d`,
																			model.DEF_APP_JUDGE_TYPE_CODE_IBC,
																			model.DEF_APP_STEP_PERFORMANCE,
																			model.DEF_APP_RESULT_EXPER_ING_A,
																			model.DEF_APP_RESULT_EXPER_ING_AC,
																			user_seq)
				succ = true
				return
			case model.DEF_APP_VIEW_POSSIBLE_IRB_CHANGE:	// IRB 변경 승인 가능한 신청서 리스트
				moreCondition = fmt.Sprintf(` AND app.judge_type = %d
																			AND app.parent_app_seq = 0
																			AND app.application_step = %d
																			AND app.application_result IN (%d, %d)
																			AND app.reg_user_seq = %d`,
																			model.DEF_APP_JUDGE_TYPE_CODE_IRB,
																			model.DEF_APP_STEP_PERFORMANCE,
																			model.DEF_APP_RESULT_EXPER_ING_A,
																			model.DEF_APP_RESULT_EXPER_ING_AC,
																			user_seq)
				succ = true
				return
			case model.DEF_APP_VIEW_POSSIBLE_IRB_RETRY:	// IRB 변경 승인 가능한 신청서 리스트
				moreCondition = fmt.Sprintf(` AND app.judge_type = %d
																			AND app.parent_app_seq = 0
																			AND app.application_step = %d
																			AND app.application_result IN (%d, %d)
																			AND app.reg_user_seq = %d`,
																			model.DEF_APP_JUDGE_TYPE_CODE_IRB,
																			model.DEF_APP_STEP_PERFORMANCE,
																			model.DEF_APP_RESULT_EXPER_ING_A,
																			model.DEF_APP_RESULT_EXPER_ING_AC,
																			user_seq)
				succ = true
				return
			case model.DEF_APP_VIEW_POSSIBLE_IA_SUPPLE:	// IACUC 보완후 재심인 목록
				moreCondition = fmt.Sprintf(` AND app.judge_type = %d
																			AND app.parent_app_seq = 0
																			AND app.application_step = %d
																			AND app.application_result = %d
																			AND app.reg_user_seq = %d`,
																			model.DEF_APP_JUDGE_TYPE_CODE_IACUC,
																			model.DEF_APP_STEP_FINAL,
																			model.DEF_APP_RESULT_REQUIRE_RETRY,
																			user_seq)
				succ = true
				return
			case model.DEF_APP_VIEW_POSSIBLE_IB_SUPPLE:	// IBC 보완후 재심인 목록
				moreCondition = fmt.Sprintf(` AND app.judge_type = %d
																			AND app.parent_app_seq = 0
																			AND app.application_step = %d
																			AND app.application_result = %d
																			AND app.reg_user_seq = %d`,
																			model.DEF_APP_JUDGE_TYPE_CODE_IBC,
																			model.DEF_APP_STEP_FINAL,
																			model.DEF_APP_RESULT_REQUIRE_RETRY,
																			user_seq)
				succ = true
				return
		default :
			common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_invalide_params, "app_view_type이 존재하지 않습니다.")
			return
	}
}

// @Tags Application
// @Summary Application 승인후 점검위원 리스트
// @Description Application 승인후 점검위원 리스트
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Router /application/{application_seq}/inspector [get]
// @Success 200
func Application_Inspector_List(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq, succ, _, _, _ := getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	if !succ {  return  }

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	rows := application.LoadInspectorList()

	common.FinishApi(c, common.Api_status_ok, rows)

}

// @Tags Application
// @Summary Application 승인후 점검위원 지정
// @Description Application 승인후 점검위원 지정
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Param user_seq path uint true "1"
// @Router /application/{application_seq}/inspector/{user_seq} [patch]
// @Success 200
func Application_Inspector_Patch(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq, succ, now_app_step, now_app_result, _ := getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	if !succ {  return  }
	if model.DEF_APP_STEP_PERFORMANCE != now_app_step {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "과제수행 단계가 아닙니다.")
		return
	}

	if model.DEF_APP_RESULT_EXPER_ING_A != now_app_result && model.DEF_APP_RESULT_EXPER_ING_AC != now_app_result  {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "실험중 상태가 아닙니다.")
		return
	}

	check_user_seq, succ2 := getUserIdFromPath(c)
	if !succ2{  return  }

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	if !application.UpdateInspector(c, check_user_seq) {
		return
	}

	common.FinishApi(c, common.Api_status_ok,
		gin.H{
			"rt": "ok",
		})

}

// @Tags Application
// @Summary Application 복제
// @Description Application 복제
// @Description 임시저장된 상태로 복제를 해준다!
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Router /application/{application_seq}/copy [post]
// @Success 200
func Application_Copy(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq, succ, _, now_app_result, _ := getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	if !succ {  return  }

	if model.DEF_APP_RESULT_TEMP == now_app_result {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "임시 저장은 복제 할수 없습니다.")
		return
	}

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	succ, new_app_seq := application.Copy(false)
	if !succ{
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	application.Application_seq = new_app_seq
	list := application.Load()
	common.FinishApi(c, common.Api_status_ok,
		gin.H{
			"rt": "ok",
			"list": list,
		})
}

// @Tags Application
// @Summary Application 보완후 재심 신청
// @Description Application 보완후 재심 신청서 선택시 복제
// @Description 새로운 app_no 가 발급됨
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Router /application/{application_seq}/retrial-copy [post]
// @Success 200
func Application_Retrial_Copy(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq, succ, now_app_step, now_app_result, _ := getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	if !succ {  return  }

	if model.DEF_APP_RESULT_REQUIRE_RETRY != now_app_result {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "해당 신청서는 보완후 재심 상태가 아닙니다.")
		return
	}

	if model.DEF_APP_STEP_FINAL != now_app_step {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "해당 신청서는 최종 심의 상태가 아닙니다.")
		return
	}

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	succ, new_app_seq := application.Copy(true)
	if !succ{
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	application.Application_seq = new_app_seq
	list := application.Load()
	common.FinishApi(c, common.Api_status_ok,
		gin.H{
			"rt": "ok",
			"list": list,
		})
}

// @Tags Application
// @Summary Application 삭제
// @Description Application 삭제
// @Description 임시저장된 상태만 삭제 가능함!
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Router /application/{application_seq} [delete]
// @Success 200
func Application_Delete(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq, succ, _, now_app_result, _ := getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	if !succ {  return  }

	if model.DEF_APP_RESULT_TEMP != now_app_result {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "임시 저장 상태만 삭제할수 있습니다.")
		return
	}

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	succ = application.Delete()
	if !succ {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "신규 신청서 삭제 실패")
		return
	}

	common.FinishApi(c, common.Api_status_ok,
		gin.H{
			"rt": "ok",
		})
}

// @Tags Application
// @Summary Application 삭제 (상태값만 변경)
// @Description Application 삭제 (상태값만 변경)
// @Description 22-03월 유지보수 사항
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Router /app/{application_seq} [delete]
// @Success 200
func App_Delete(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq, succ, _, _, _ := getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	if !succ {  return  }

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	succ = application.ChangeToDeleteState()
	if !succ {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "신청서 삭제 실패")
		return
	}

	common.FinishApi(c, common.Api_status_ok, gin.H{"rt": "ok",})
}

// @Tags Application Change
// @Summary Application 변경 내역
// @Description Application 변경 내역
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Router /application/{application_seq}/change [get]
// @Success 200
func Application_Change_List(c *gin.Context)  {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq := uint64(0);
	succ := false;
	institution_seq := 0;
	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		app_seq, succ, _, _, _ = getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	  if !succ {  return  }
	} else {
		app_seq, succ, institution_seq = getApplicationSeqFromPathForAdmin(c)
		if !succ {  return  }
		tokenMap["institution_seq"] = institution_seq
	}

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	result := application.ChangeInfo()
	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Application Change
// @Summary Application 변경 내역 동물
// @Description Application 변경 내역 동물
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Router /application/{application_seq}/change/animal [get]
// @Success 200
func Application_Change_Animal_Info(c *gin.Context)  {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq := uint64(0);
	succ := false;
	institution_seq := 0;
	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		app_seq, succ, _, _, _ = getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	  if !succ {  return  }
	} else {
		app_seq, succ, institution_seq = getApplicationSeqFromPathForAdmin(c)
		if !succ {  return  }
		tokenMap["institution_seq"] = institution_seq
	}

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	result := application.Change_Animal_Info()
	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Application Change
// @Summary Application 변경 승인 신청 내역
// @Description Application 연구 책임자 또는 실험 수행자 변경 내역
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Router /application/{application_seq}/change/member [get]
// @Success 200
func Application_Change_Member_Info(c *gin.Context)  {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq := uint64(0);
	succ := false;
	institution_seq := 0;
	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		app_seq, succ, _, _, _ = getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	  if !succ {  return  }
	} else {
		app_seq, succ, institution_seq = getApplicationSeqFromPathForAdmin(c)
		if !succ {  return  }
		tokenMap["institution_seq"] = institution_seq
	}

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	result := application.Change_Member_Info()
	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Application Change
// @Summary Application 변경 승인 신청 내역
// @Description Application 연구 종료일
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Router /application/{application_seq}/change/end-date [get]
// @Success 200
func Application_Change_EndDate_Info(c *gin.Context)  {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq := uint64(0);
	succ := false;
	institution_seq := 0;
	if common.ToUint(tokenMap["user_auth"]) < model.DEF_USER_AUTH_PLATFORM {
		app_seq, succ, _, _, _ = getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	  if !succ {  return  }
	} else {
		app_seq, succ, institution_seq = getApplicationSeqFromPathForAdmin(c)
		if !succ {  return  }
		tokenMap["institution_seq"] = institution_seq
	}

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	result := application.Change_EndDate_Info()
	common.FinishApi(c, common.Api_status_ok, result)
}

// @Tags Application
// @Summary Application	IBC -> IACUC신청서에서 선택하기
// @Description Application	IACUC신청서에서 선택하기
// @Description body에 복사할 iacuc app seq를 담아서 보낸다.
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param application_seq path uint true "1"
// @Param param body model.iacucCopyForIbc true "param"
// @Router /application/{application_seq}/iacuc-copy` [post]
// @Success 200
func Application_Copy_Iacuc_For_ibc(c *gin.Context)  {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	app_seq, succ, _, _, _ := getApplicationSeqFromPath(c, tokenMap["institution_seq"])
	if !succ {  return  }

	chkKeys := []interface{}{
		[]string{"application_seq"},	//	필수 키
		[][]string{},	//	옵션 키 (그룹 중 1개이상 반드시 있어야 함)
		[]string{},	//  그외 읽어야 할 값들.
	}

	data, eMsg := common.BindCustomBody(c, &chkKeys)
	if nil == data {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, eMsg)
		return
	}

	iacucAppSeq := common.ToUint(data["application_seq"])
	iacucApp := model.Application {
		LoginToken : tokenMap,
		Application_seq : iacucAppSeq,
	}

	iacucApp.Data = iacucApp.Load().(map[string]interface{})

	if common.ToUint(iacucApp.Data["judge_type"]) != model.DEF_APP_JUDGE_TYPE_CODE_IACUC || common.ToUint(iacucApp.Data["application_step"]) != model.DEF_APP_STEP_PERFORMANCE {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	application := model.Application {
		LoginToken : tokenMap,
		Application_seq : uint(app_seq),
	}

	if !application.IacucCopyForIbc(iacucAppSeq) {
		common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
		return
	}

	retJson := gin.H{	"rt": "ok"}
	retJson["new_info"] = application.Load()

	common.FinishApi(c, common.Api_status_ok, retJson)
}
