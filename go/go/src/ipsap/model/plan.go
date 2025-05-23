package model

import (
	// "github.com/go-sql-driver/mysql"
	// "github.com/gin-gonic/gin"
	"ipsap/common"
	// "strings"
	"fmt"
	"log"
)

type Plan struct {
	Plan_seq											uint		`json:"-"`
	// Planid 												string	`json:"planid" example:"1"`
	Plan_product_seq							uint		`json:"plan_product_seq" example:"1"`
	Membership_product_seq				uint		`json:"membership_product_seq" example:"1"`
	Plan_category									string	`json:"plan_category" example:"plan"`
	Plan_pid 											string	`json:"plan_pid" example:"1"`
	Membership_category						string	`json:"membership_category" example:"membership"`
	Membership_pid 								string	`json:"membership_pid" example:"1"`
	Name													string	`json:"name" example:"2개월권"`
	Desc_text											string	`json:"desc_text" example:"설명"`
	Usage_limit										uint		`json:"usage_limit" example:"30"`
	Plan_price										uint		`json:"plan_price" example:"1000000"`
	Plan_discount_rate						float32	`json:"plan_discount_rate" example:"0.1"`
	Plan_discount_amount					uint		`json:"plan_discount_amount" example:"100"`
	Plan_discount_type						string	`json:"plan_discount_type" example:"R"`
	Plan_discounted_amount				uint		`json:"plan_discounted_amount" example:"100000"`
	Membership_price							uint		`json:"membership_price" example:"1000000"`
	Membership_discount_rate			float32	`json:"membership_discount_rate" example:"0.1"`
	Membership_discount_amount		uint		`json:"membership_discount_amount" example:"100"`
	Membership_discount_type			string	`json:"membership_discount_type" example:"R"`
	Membership_discounted_amount	uint		`json:"membership_discounted_amount" example:"100000"`
	Plan_available								uint		`json:"plan_available" example:"1"`
	Reg_user_seq									uint		`json:"-"`
	Order_int											uint		`json:"-"`
	Note													string	`json:"-"`
	Deleted_flag									uint		`json:"-"`
}

func (plan *Plan) InsertPlan() (succ bool) {
	succ = false
	dbConn	:= common.DBconn()
	tx, err	:= dbConn.Begin()
	if nil != err {
		tx.Rollback()
		return
	}

	defer func() {
		tx.Rollback()
	}()

	sql := `INSERT INTO t_membership_plan(name, desc_text, usage_limit, reg_user_seq,
																				plan_available, reg_date)
					VALUES(?,?,?,?,
								 ?,UNIX_TIMESTAMP())`
  result, err := tx.Exec(sql,
												 plan.Name, plan.Desc_text, plan.Usage_limit, plan.Reg_user_seq,
												 plan.Plan_available)
	if err != nil {
	 log.Println(err)
	 return
	}

	no, err2 := result.LastInsertId()
	if err2 != nil {
	 log.Println(err2)
	 return
	}

	plan.Plan_seq = common.ToUint(no)

	sql = `INSERT INTO t_products( plan_seq, pid, category,
																 discount_rate, discount_amount, discount_type,
																 discounted_amount, price)
					VALUES(?,?,?,
								 ?,?,?,
								 ?,?)`
  _, err = tx.Exec(sql,
									 plan.Plan_seq, plan.Plan_pid, plan.Plan_category,
									 plan.Plan_discount_rate, plan.Plan_discount_amount, plan.Plan_discount_type,
									 plan.Plan_discounted_amount, plan.Plan_price)
	if err != nil {
		log.Println(err)
		return
	}

	sql = `INSERT INTO t_products( plan_seq, pid, category,
																 discount_rate, discount_amount, discount_type,
																 discounted_amount, price)
						VALUES(?,?,?,
									 ?,?,?,
									 ?,?)`
	_, err = tx.Exec(sql,
									 plan.Plan_seq, plan.Membership_pid, plan.Membership_category,
									 plan.Membership_discount_rate, plan.Membership_discount_amount, plan.Membership_discount_type,
									 plan.Membership_discounted_amount, plan.Membership_price)
	if err != nil {
		log.Println(err)
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

func (plan *Plan) Load() (list interface{}) {
	moreCondition := ""
	if 0 != plan.Plan_seq {
		moreCondition = fmt.Sprintf(` AND plan.plan_seq = %d`, plan.Plan_seq)
	}

	sql := fmt.Sprintf(`
		SELECT *
			FROM t_membership_plan plan
		 WHERE 1 = 1
		   AND plan.deleted_flag = 0
		    %v #moreCondition`, moreCondition)
	filter := func(row map[string]interface{}) {
		sql2 := `SELECT *
							 FROM t_products prod
							WHERE prod.plan_seq = ?`
		row["products"] = common.DB_fetch_all(sql2, nil, row["plan_seq"])
	}
	list = common.DB_fetch_all(sql, filter)
	return
}

func (plan *Plan) UpdatePlan() (succ bool)  {
	succ = false
	dbConn	:= common.DBconn()
	tx, err	:= dbConn.Begin()
	if nil != err {
		tx.Rollback()
		return
	}

	defer func() {
		tx.Rollback()
	}()

	sql := `UPDATE t_membership_plan
						 SET name = ?, desc_text = ?, usage_limit = ?, plan_available = ?
					 WHERE plan_seq = ?`
	_, err = tx.Exec(sql,
									 plan.Name, plan.Desc_text, plan.Usage_limit, plan.Plan_available,
									 plan.Plan_seq)
	if err != nil {
		log.Println(err)
		return
	}

	sql = `UPDATE t_products
					 SET discount_rate = ?, discount_amount = ?, discount_type = ?,
							 discounted_amount = ?, price = ?
				 WHERE product_seq = ?`
	_, err = tx.Exec(sql,
									 plan.Plan_discount_rate, plan.Plan_discount_amount, plan.Plan_discount_type,
									 plan.Plan_discounted_amount, plan.Plan_price,
								 	 plan.Plan_product_seq)
	if err != nil {
		log.Println(err)
		return
	}

	sql = `UPDATE t_products
					  SET discount_rate = ?, discount_amount = ?, discount_type = ?,
							  discounted_amount = ?, price = ?
				  WHERE product_seq = ?`
	_, err = tx.Exec(sql,
									 plan.Membership_discount_rate, plan.Membership_discount_amount, plan.Membership_discount_type,
									 plan.Membership_discounted_amount, plan.Membership_price,
									 plan.Membership_product_seq)
	if err != nil {
		log.Println(err)
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

func (plan *Plan) DeletePlan() (succ bool) {
	succ = false
	sql := `UPDATE t_membership_plan
						 SET deleted_flag = 1
					 WHERE plan_seq = ?`
	_, err := common.DBconn().Exec(sql, plan.Plan_seq)
	if nil != err {
		log.Println(err)
		return
	}

	succ = true
	return
}
