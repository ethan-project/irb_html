package model

import (
	// "github.com/gin-gonic/gin"
  "ipsap/common"
  "database/sql"
	"fmt"
	// "strings"
	"log"
)
//
type ChangeHistory struct {
	Application			*Application
	Item_name 	 		 string
	Org_Item_seq		 uint
	Updated_Item_seq uint
	Ref_seq 		 		 uint
}

func (change *ChangeHistory) setHistory(tx *sql.Tx) (ret bool, err_msg string) {
	change.getLastRefSeq(tx)
	if 0 == change.Ref_seq {
		ret = change.insertHistroy(tx, "origin")
		if (!ret) {
			err_msg = fmt.Sprintf("item_name[%s] origin histroy insert Fail", change.Item_name)
			return
		}
	}

	ret = change.insertHistroy(tx, "updated")
	if (!ret) {
		err_msg = fmt.Sprintf("item_name[%s] updated histroy insert Fail", change.Item_name)
		return
	}

	return
}

func (change *ChangeHistory) getLastRefSeq(tx *sql.Tx) {
	sql := `SELECT change_seq
						FROM t_change_mgr_histories
					 WHERE application_seq = ?
					   AND item_name = ?
						 AND chg_code = 'updated'
					 ORDER BY date DESC
					 LIMIT 1`
	row := common.DB_Tx_fetch_one(tx, sql, nil, change.Application.Application_seq, change.Item_name)
	if nil != row["chagne_seq"] {
		change.Ref_seq = common.ToUint(row["chagne_seq"])
	}
	return
}

func (change *ChangeHistory) insertHistroy(tx *sql.Tx, chg_code string) (ret bool){
	itme_seq := uint(0)
	ret = false
	if "origin" == chg_code {
		itme_seq = change.Org_Item_seq
	} else {
		itme_seq = change.Updated_Item_seq
	}

	sql:= `INSERT INTO t_change_mgr_histories(
										 application_seq, chg_code, item_name,
										 item_seq, ref_seq)
			 	 VALUES(?, ?, ?,
					 			?, ?)
			  `
	result, err := tx.Exec(sql,
												 change.Application.Application_seq, chg_code, change.Item_name,
												 itme_seq, change.Ref_seq)
	if nil != err {
		log.Println(err)
		return
	}

	no, err := result.LastInsertId()
	if err != nil {
	 log.Println(err)
	 return
	}

	if change.Ref_seq == 0 && "origin" == chg_code {
		change.Ref_seq = common.ToUint(no)
	}

	ret = true
	return
}
