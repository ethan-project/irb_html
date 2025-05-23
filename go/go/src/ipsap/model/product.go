package model

import (
	// "github.com/go-sql-driver/mysql"
	// "github.com/gin-gonic/gin"
	"ipsap/common"
	// "strings"
	"fmt"
	"log"
)

type Product struct {
	Product_seq					uint		`json:"-"`
	Pid 								string	`json:"pid" example:"1"`
	Pcode								string	`json:"pcode" example:"1"`
	Category						string	`json:"category" example:"plans"`
	Name								string	`json:"name" example:"2개월권"`
	Desc_text						string	`json:"desc_text" example:"설명"`
	Usage_period				string	`json:"usage_period" example:"30D"`
	Price								uint		`json:"price" example:"1000000"`
	Fees								uint		`json:"-"`
	Taxes								uint		`json:"-"`
	Discount_rate				float32	`json:"discount_rate" example:"0.1"`
	Discount_amount			uint		`json:"discount_amount" example:"100"`
	Discount_type				string	`json:"discount_type" example:"R"`
	Discounted_amount		uint		`json:"discounted_amount" example:"100000"`
	Rg_expiry_date			uint		`json:"rg_expiry_date" example:"10"`
	Product_available		uint		`json:"product_available" example:"1"`
	Discount_available	uint		`json:"discount_available" example:"0"`
	// RegDate						string	`json:"-"`
	Reg_user_seq				uint		`json:"-"`
	Order_int						uint		`json:"-"`
	Note								string	`json:"-"`
	Deleted_flag				uint		`json:"-"`
}

func (prod *Product) InsertProduct() (succ bool) {
	succ = false
	sql := `INSERT INTO t_products(	pid, pcode, category,
																	name, desc_text, usage_period,
																	price, fees, taxes,
																	discount_rate, discount_amount, discount_type,
																	discounted_amount, rg_expiry_date, product_available,
																	discount_available, note, reg_user_seq,
																	order_int, reg_date)
					VALUES(?,?,?, ?,?,?, ?,?,?,
								 ?,?,?, ?,?,?, ?,?,?,
								 ?,UNIX_TIMESTAMP())`
  _, err := common.DBconn().Exec(sql,
																 prod.Pid, prod.Pcode, prod.Category,
																 prod.Name, prod.Desc_text, prod.Usage_period,
																 prod.Price, prod.Fees, prod.Taxes,
																 prod.Discount_rate, prod.Discount_amount, prod.Discount_type,
																 prod.Discounted_amount, prod.Rg_expiry_date, prod.Product_available,
																 prod.Discount_available, prod.Note, prod.Reg_user_seq,
																 prod.Order_int)
  if nil != err {
    log.Println(err)
    return
  }

  succ = true
	return
}

func (prod *Product) Load() (list interface{}) {
	moreCondition := ""
	if 0 != prod.Product_seq {
		moreCondition = fmt.Sprintf(` AND product_seq = %d`, prod.Product_seq)
	}

	sql := fmt.Sprintf(`
		SELECT product_seq, pid, category,
					 name, desc_text, usage_period, price,
					 fees, taxes, discount_rate, discount_amount,
					 discount_type, discounted_amount, product_available,
					 reg_user_seq, order_int, note,
					 deleted_flag, reg_date
			FROM t_products
		 WHERE deleted_flag = 0
		    %v #moreCondition`, moreCondition)
	// filter := func(row map[string]interface{}) {
	// }
	list = common.DB_fetch_all(sql, nil)
	return
}

func (prod *Product) UpdateProduct() (succ bool)  {
	succ = false
	sql := `UPDATE t_products
						 SET pid= ?, pcode = ?, category = ?,
								 name = ?, desc_text = ?, usage_period = ?,
								 price = ?, fees  = ?, taxes = ?,
								 discount_rate = ?, discount_amount = ?, discount_type = ?,
								 discounted_amount = ?, rg_expiry_date  = ?, product_available = ?,
								 discount_available = ?, note = ?, order_int = ?
					 WHERE product_seq = ?`
	_, err := common.DBconn().Exec(sql,
																 prod.Pid, prod.Pcode, prod.Category,
																 prod.Name, prod.Desc_text, prod.Usage_period,
																 prod.Price, prod.Fees, prod.Taxes,
																 prod.Discount_rate, prod.Discount_amount, prod.Discount_type,
																 prod.Discounted_amount, prod.Rg_expiry_date, prod.Product_available,
																 prod.Discount_available, prod.Note, prod.Order_int,
															 	 prod.Product_seq)
	if nil != err {
		log.Println(err)
		return
	}

	succ = true
	return
}

func (prod *Product) DeleteProduct() (succ bool) {
	succ = false
	sql := `UPDATE t_products
						 SET deleted_flag = 1
					 WHERE product_seq = ?`
	_, err := common.DBconn().Exec(sql, prod.Product_seq)
	if nil != err {
		log.Println(err)
		return
	}

	succ = true
	return
}

func (prod *Product) CheckProduct() (succ bool)  {
	succ = false
	sql := `SELECT product_seq
						FROM t_products
					 WHERE product_seq = ?
						 AND discounted_amount = ?`
	row := common.DB_fetch_one(sql, nil, prod.Product_seq, prod.Discounted_amount)
	if nil == row || nil == row["product_seq"] {
		return
	}

	succ = true
	return
}
