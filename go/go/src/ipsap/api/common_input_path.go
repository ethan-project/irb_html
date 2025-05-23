package api

import (
	"github.com/gin-gonic/gin"
	"ipsap/common"
	"log"
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
/////////////////////////////////////////////////////////////


func getApplicationSeqFromPath(c *gin.Context, institution_seq interface{}) (app_seq uint64,  succ bool, app_step int, app_result int, reg_user_seq uint) {
	app_seq, succ = getPathParamForUint64(c, "application_seq", "Application Seq")
	if !succ {
		return
	}

	if app_seq == uint64(0) {
		succ = true
		return
	}

	sql := `SELECT application_seq, application_step, application_result, reg_user_seq
						FROM t_application
					 WHERE application_seq = ?
					 	 AND institution_seq = ?`
	row := common.DB_fetch_one(sql, nil, app_seq, institution_seq)
	if nil == row {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}

	app_step = common.ToInt(row["application_step"])
	app_result = common.ToInt(row["application_result"])
	reg_user_seq = common.ToUint(row["reg_user_seq"])
	succ = true
	return
}

func getApplicationSeqFromPathForAdmin(c *gin.Context) (app_seq uint64, succ bool, institution_seq int) {
	app_seq, succ = getPathParamForUint64(c, "application_seq", "Application Seq")
	if !succ {
		return
	}

	sql := `SELECT application_seq, institution_seq
						FROM t_application
					 WHERE application_seq = ?`
	row := common.DB_fetch_one(sql, nil, app_seq)
	if nil == row {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}

	institution_seq = common.ToInt(row["institution_seq"])
	succ = true
	return
}

func getUserIdFromPath(c *gin.Context) (user_seq uint64, succ bool) {
	user_seq, succ = getPathParamForUint64(c, "user_seq", "User Seq")
	if !succ {
		return
	}

	sql := `SELECT user_seq FROM t_user WHERE user_seq = ?`
	row := common.DB_fetch_one(sql, nil, user_seq)
	if nil == row {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}

	succ = true
	return
}

func getInstitutionIdFromPath(c *gin.Context) (institution_seq uint64, succ bool) {
	institution_seq, succ = getPathParamForUint64(c, "institution_seq", "Institutio Seq")
	if !succ {
		return
	}

	sql := `SELECT institution_seq FROM t_institution WHERE institution_seq = ?`
	row := common.DB_fetch_one(sql, nil, institution_seq)
	if nil == row {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}

	succ = true
	return
}

func getRequestServiceSeqFromPath(c *gin.Context) (reqsvc_seq uint64, succ bool, request_status uint,  request_type uint, institution_seq uint, user_seq uint) {
	reqsvc_seq, succ = getPathParamForUint64(c, "reqsvc_seq", "Reqsvc Seq")
	if !succ {
		return
	}

	sql := `SELECT reqsvc.reqsvc_seq, reqsvc.request_type,
								 reqsvc.institution_seq, reqsvc.user_seq, reqsvc.request_status
						FROM t_request_service reqsvc, t_institution instt
					 WHERE reqsvc.institution_seq = instt.institution_seq
					 	 AND reqsvc_seq = ?`
	row := common.DB_fetch_one(sql, nil, reqsvc_seq)
	if nil == row {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}
	request_type = common.ToUint(row["request_type"])
	request_status = common.ToUint(row["request_status"])
	institution_seq = common.ToUint(row["institution_seq"])
	user_seq = common.ToUint(row["user_seq"])
	succ = true
	return
}

func getBoardInfoFromPath(c *gin.Context) (board_seq uint64, succ bool, user_seq uint, institution_seq uint) {
	board_seq, succ = getPathParamForUint64(c, "board_seq", "Board Seq")
	if !succ {
		return
	}

	sql := `SELECT board_seq, user_seq, institution_seq
						FROM t_board
					 WHERE board_seq = ?`
	row := common.DB_fetch_one(sql, nil, board_seq)
	if nil == row {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}
	user_seq = common.ToUint(row["user_seq"])
	institution_seq = common.ToUint(row["institution_seq"])
	succ = true
	return
}

func getProductFromPath(c *gin.Context, category string) (product_seq uint64, succ bool) {
	product_seq, succ = getPathParamForUint64(c, "product_seq", "Product Seq")
	if !succ {
		return
	}

	sql := `SELECT product_seq
						FROM t_products
					 WHERE product_seq = ?
					   AND category = ?`
	row := common.DB_fetch_one(sql, nil, product_seq, category)
	if nil == row {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}

	succ = true
	return
}

func getPlanFromPath(c *gin.Context) (plan_seq uint64, succ bool) {
	plan_seq, succ = getPathParamForUint64(c, "plan_seq", "Plan Seq")
	if !succ {
		return
	}

	sql := `SELECT plan_seq
						FROM t_membership_plan
					 WHERE plan_seq = ?`
	row := common.DB_fetch_one(sql, nil, plan_seq)
	if nil == row {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}

	succ = true
	return
}

func getMembershipFreeSeqFromPath(c *gin.Context) (free_seq uint64, succ bool) {
	free_seq, succ = getPathParamForUint64(c, "free_seq", "Membership Free Seq")
	if !succ {
		return
	}

	sql := `SELECT free_seq
						FROM t_institution_free_period
					 WHERE free_seq = ?`
	row := common.DB_fetch_one(sql, nil, free_seq)
	if nil == row {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}

	succ = true
	return
}

func getOrderSeqFromPath(c *gin.Context) (order_seq uint64, succ bool) {
	order_seq, succ = getPathParamForUint64(c, "order_seq", "Order Seq")
	if !succ {
		return
	}

	sql := `SELECT order_seq
						FROM t_orders
					 WHERE order_seq = ?`
	row := common.DB_fetch_one(sql, nil, order_seq)
	if nil == row {
		common.FinishApi(c, common.Api_status_not_found)
		succ = false
		return
	}

	succ = true
	return
}
