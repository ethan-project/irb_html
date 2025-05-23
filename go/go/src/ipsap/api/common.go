package api

import (
	"github.com/gin-gonic/gin"
	"ipsap/common"
  "ipsap/model"
	// "log"
  // "strings"
  // "fmt"
)

// @Tags Common
// @Summary 기관코드 중복 체크
// @Description 기관코드 중복 체크
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param institution_code query string true "KUN01"
// @Router /common/dup-check/institution-code [get]
// @Success 200
func DupCheckInstitutionCode(c *gin.Context) {
	institution_code := c.Request.URL.Query().Get("institution_code")
	if "" == institution_code {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "institution_code 값이 없습니다.")
		return
	}

	institution	:= model.Institution{}
	institution.Institution_code = institution_code
	if !institution.CheckDuplicateInstitutionCode(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok,
    gin.H{
      "rt": "ok",
    })
}

// @Tags Common
// @Summary Email 중복 체크 (기관내에서만 중복 확인)
// @Description Email 중복 체크 (기관내에서만 중복 확인)
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param email query string true "test@test.com"
// @Param institution_seq query string true "1"
// @Router /common/dup-check/email [get]
// @Success 200
func DupCheckEmail(c *gin.Context) {
	email := c.Request.URL.Query().Get("email")
	if "" == email {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "email 값이 없습니다.")
		return
	}

	user := model.User{}
	user.Email = email
	if !user.CheckDuplicateUserEmail(c) {
		return
	}

	common.FinishApi(c, common.Api_status_ok,
    gin.H{
      "rt": "ok",
    })
}
