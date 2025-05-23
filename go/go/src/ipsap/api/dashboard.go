package api

import (
	"github.com/gin-gonic/gin"
	"ipsap/common"
	"ipsap/model"
	"strings"
)

// @Tags DashBoard
// @Summary DashBoard List
// @Description DashBoard List
// @Description dash_view_type 1 = 행정간사	dashboard(기관 전체)
// @Description dash_view_type 2 = 연구원	 dashboard(자신이 참여한 실험만)
// @Description dash_view_type 3 = 위원장	 dashboard(기관 전체)
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param filter.dashboard_type query string false "1"
// @Router /dashboard [get]
// @Success 200
func DashBoardList(c *gin.Context) {
	tokenMap := common.Check_token(c)
	if nil == tokenMap {
		return
	}

	userTypeArr := strings.Split(common.ToStr(tokenMap["user_type_all"]), ",")
	if 0 == len(userTypeArr) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	dashboardType := common.ToStr(c.Request.URL.Query().Get("filter.dashboard_type"))
	if "" == dashboardType {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
		return
	}

	switch common.ToUint(dashboardType) {
		case model.DEF_DASHBOARD_TYPE_ADMIN:
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_ADMIN_SECRETARY, model.DEF_USER_TYPE_ADMIN_WORK) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
		case model.DEF_DASHBOARD_TYPE_RESEARCHER:
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_RESEARCHER) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
		case model.DEF_DASHBOARD_TYPE_CHAIRPERSON:
			if !common.CheckUserTypeAuth(userTypeArr, model.DEF_USER_TYPE_CHAIRPERSON) {
				common.FinishApiWithErrCd(c, common.Api_status_unauthorized, common.Error_unauthorized)
				return
			}
		default :
			common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
			return
	}

	app := model.Application{LoginToken :tokenMap,}
	dashBoard := model.DashBoard{App : app, DashBoard_type :common.ToUint(dashboardType),}
	result := dashBoard.GetDashBoardContent()
	common.FinishApi(c, common.Api_status_ok, result)
}
