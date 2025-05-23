package ctl

import (
	"github.com/gin-gonic/gin"
	"ipsap/common"
	"log"
	"fmt"
)

/////////////////////////////////////////////////////////////
//  Common
func getPathParamForUint64(c *gin.Context, name string, errTag string) (value uint64, succ bool) {
	id_str := c.Param(name)
	value, succ = common.ToValidateUint64(id_str)
	if !succ {
		log.Println(errTag + " : " + id_str)
		common.FinishApi(c, common.Api_status_not_found)
		return
	}
	return
}

// 파라미터 값 가져오기
func GetPathParam(c *gin.Context, name string) (value string, succ bool) {
	value = c.Param(name)
	if "" != value {
		succ = true
		return
	} else {
		log.Println(name + " is null")
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}
}

func GetOsTypeFromPath(c *gin.Context) (osType uint64, succ bool) {
	osType, succ = getPathParamForUint64(c, "os_type", "OS Type")
	if !succ {
		return
	}

	if "Code Error" == GetCommonCodeStr("os_type", osType) {
		succ = false
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params, "os_type이 잘못되었습니다.")
		return
	}

	return
}

func GetPatchIdFromPath(c *gin.Context, chkUseFlag bool) (patch_seq uint64, succ bool) {
	patch_seq, succ = getPathParamForUint64(c, "patch_id", "Patch ID")
	if !succ {
		return
	}

	sql := `SELECT patch_seq FROM t_patch_info WHERE patch_seq = ?`
	if chkUseFlag {
		sql += ` AND (use_flag = 1 OR use_flag = 0)`
	}
	row := common.DB_fetch_one(sql, nil, patch_seq)
	if row == nil {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}

	succ = true
	return
}
