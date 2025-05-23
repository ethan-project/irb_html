
package model

import (
	// "github.com/gin-gonic/gin"
  "ipsap/common"
  "database/sql"
  "strings"
	"fmt"
  "log"
)

func (app *Application)Copy(isRetrial bool) (succ bool, new_app_seq uint) {
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

	data := app.Load().(map[string]interface{})
	if DEF_APP_TYPE_NEW != common.ToUint(data["application_type"]) {
		log.Printf(`application_type : ==> %v`, data["application_type"])
		return
	}

  application_no := "";
  if (isRetrial) {
    sql := `SELECT application_no
              FROM t_application
             WHERE application_seq = ?`
    row := common.DB_Tx_fetch_one(tx, sql, nil, app.Application_seq);
    app_no := common.ToStr(row["application_no"]);
    app_no = app_no[0:len(app_no)-2]
    sql2 := fmt.Sprintf(`SELECT COUNT(application_seq) AS app_cnt
			                     FROM t_application
			                    WHERE application_no LIKE '%%%v%%'
				                    AND application_type = %d`, app_no, DEF_APP_TYPE_NEW)
    row2 := common.DB_Tx_fetch_one(tx, sql2, nil)
    app_cnt := common.ToStr(row2["app_cnt"])
    if (len(app_cnt) == 1) {
      app_cnt = "0" + app_cnt
    }
    application_no = app_no + app_cnt
  } else {
    app_basic := AppBasic{ Application : app }
  	_, application_no = app_basic.GetAppNo(tx, common.ToInt(data["judge_type"]), common.ToInt(data["application_type"]), uint(0))
  }

	sql := `INSERT INTO t_application(
											 application_no,	 	institution_seq, 	judge_type,
											 application_type,	application_step,	application_result,
											 name_ko,						name_en,				 	reg_user_seq,
											 chg_user_seq,		  reg_dttm)
					      VALUES(?, ?, ?,
											 ?, ?, ?,
											 ?, ?, ?,
											 ?, UNIX_TIMESTAMP())`
	result, err := tx.Exec(sql, application_no, data["institution_seq"], data["judge_type"],
															data["application_type"], DEF_APP_STEP_WRITE, DEF_APP_RESULT_TEMP,
															data["name_ko"],	data["name_en"], app.LoginToken["user_seq"],
															app.LoginToken["user_seq"])
	if err != nil {
	  log.Println(err)
	  return
	}

	no, err := result.LastInsertId()
	if err != nil {
	 log.Println(err)
	 return
	}

	new_app_seq = common.ToUint(no)
  if !app.itemCopy(tx, new_app_seq, common.ToUint(data["judge_type"])) {
    return
  }

  err = tx.Commit()
  if nil != err {
    log.Println(err)
    succ = false
  } else {
    succ = true
  }
	return
}

func (app *Application)itemCopy(tx *sql.Tx, new_app_seq uint, judgeType uint) (succ bool) {
	succ = false
  sql := ""
  var err interface{}
  if DEF_APP_JUDGE_TYPE_CODE_IACUC ==  judgeType {
    sql  = `INSERT INTO t_application_anesthetic
            SELECT ?, item_name,	anesthetic_type, anesthetic_type_str, anesthetic_name,
                   injection_mg, injection_route, injection_route_str, injection_time,
                   injection_cnt, view_order
              FROM t_application_anesthetic
             WHERE application_seq = ?
               AND item_name = ?`;
    _, err = tx.Exec(sql, new_app_seq, app.Application_seq, `pain_d_anesthetic`)
    if err != nil {
      log.Printf(`t_application_anesthetic ==> %v`,err)
      return
    }

    sql = `	INSERT INTO t_application_animal
            SELECT ?, item_name, animal_code, animal_code_str,
                   male_cnt, female_cnt, mb_grade, mb_grade_str,
                   breeding_place, breeding_place_str, strain,
                   week_age, age_unit, weight_gram, weight_unit,
                   size, size_unit, supplier_type, supplier_name,
                   lmo_flag, ibc_num, genetic_type, lmo_type, view_order
              FROM t_application_animal
             WHERE application_seq = ?
               AND item_name = ?`;
     _,err = tx.Exec(sql, new_app_seq, app.Application_seq, `animal_type`)
    if err != nil {
      log.Printf(`t_application_animal ==> %v`,err)
      return
    }
  }


	sql = `	INSERT INTO t_application_etc
					SELECT ?, item_name, target_item, contents,
								 UNIX_TIMESTAMP(), ? #reg_user_seq
						FROM t_application_etc
					 WHERE application_seq = ?
						 AND item_name NOT IN ('self_inspection', 'supplement', 'judge_expert_deadline',
																	 'judge_normal_deadline', 'expert_review_opinion',
																	 'expert_review_evaluation', 'expert_review_total_review',
																	 'expert_review_dttm', 'normal_review_dttm')`;
	 _,err = tx.Exec(sql, new_app_seq, app.LoginToken["user_seq"] , app.Application_seq)
	if err != nil {
		log.Printf(`t_application_etc ==> %v`,err)
		return
	}

	sql = `	INSERT INTO t_application_keyword
					SELECT ?, item_name, keyword
						FROM t_application_keyword
					 WHERE application_seq = ?`;
	 _,err = tx.Exec(sql, new_app_seq, app.Application_seq)
	if err != nil {
		log.Printf(`t_application_keyword ==> %v`,err)
		return
	}

	sql = `	INSERT INTO t_application_main_check
					SELECT ?, item_name, item_idx, checked
						FROM t_application_main_check
					 WHERE application_seq = ?`;
	 _,err = tx.Exec(sql, new_app_seq, app.Application_seq)
	if err != nil {
		log.Printf(`t_application_main_check ==> %v`,err)
		return
	}

	// 연구 책임자로만 들어감!
	sql = `	INSERT INTO t_application_member(application_seq, item_name, user_seq,
																					 animal_mng_flag, exp_year_code, exp_type_code, chk_flag)
					VALUES(?,?,?,0,0,0,0)`;
	 _,err = tx.Exec(sql, new_app_seq, `general_director`, app.LoginToken["user_seq"])
	if err != nil {
		log.Printf(`t_application_member ==> %v`,err)
		return
	}

	sql = `	INSERT INTO t_application_select
					SELECT ?, item_name, item_idx, select_ids
						FROM t_application_select
					 WHERE application_seq = ?
					 	 AND item_name NOT IN ('judge_alarm', 'expert_review_result',
							 										 'normal_review_result', 'final_judge_result')`;
	 _,err = tx.Exec(sql, new_app_seq, app.Application_seq)
	if err != nil {
		log.Printf(`t_application_select ==> %v`,err)
		return
	}

	sql = `	INSERT INTO t_application_select_input
					SELECT ?, item_name, item_idx, id, no, input
						FROM t_application_select_input
					 WHERE application_seq = ?
						 AND item_name NOT IN ('judge_alarm', 'expert_review_result',
																	 'normal_review_result', 'final_judge_result')`;
	 _,err = tx.Exec(sql, new_app_seq, app.Application_seq)
	if err != nil {
		log.Printf(`t_application_select ==> %v`,err)
		return
	}

	// 파일이 존재하는 것만 복사!
	sql = `SELECT item_name, filepath, org_file_name,
								file_idx,	 view_order
	         FROM t_application_file
					WHERE application_seq = ?`
	rows := common.DB_Tx_fetch_all(tx, sql, nil, app.Application_seq)
	for _, row := range rows {
		copySrc := common.ToStr(row["filepath"])
		old := fmt.Sprintf(`application/%v`, app.Application_seq)
		new := fmt.Sprintf(`application/%v`, new_app_seq)
		newSrc := strings.Replace(copySrc, old, new, 1)
		copyErr := common.CopyToS3(copySrc, newSrc)
		if nil != copyErr {
			fileCopySql := `INSERT INTO t_application_file(
				   							application_seq, item_name, filepath,
												org_file_name,	 file_idx,	view_order)
											VALUES(?,?,?, ?,?,?)`
			_,copyErr = tx.Exec(fileCopySql,
															new_app_seq, row["item_name"], copySrc,
															row["org_file_name"], row["file_idx"], row["view_order"])
			if copyErr != nil {
				log.Printf(`t_application_file ==> %v`,err)
				return
			}
		}
	}

	return true
}

func (app *Application)IacucCopyForIbc(iacucAppSeq uint) (succ bool) {
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

  if app.Application_seq == 0 {
    if !app.NewIbc(tx) {
      return
    }
  }

  sql := `SELECT name_ko, name_en, application_no
            FROM t_application
           WHERE application_seq = ?`
  row := common.DB_Tx_fetch_one(tx, sql, nil, iacucAppSeq)

  sql = `UPDATE t_application
             SET name_ko = ?, name_en = ?
           WHERE application_seq = ?`
  _, err = tx.Exec(sql, row["name_ko"], row["name_en"], app.Application_seq)
 	if err != nil {
 	  log.Println(err)
 	  return
 	}

  sql = `INSERT INTO t_application_etc(
          application_seq, item_name, target_item,
          contents, reg_dttm, reg_user_seq
         )
         VALUES(?,'ibc_general_animal_iacuc_num','ibc_general_animal_iacuc_num',
                ?,UNIX_TIMESTAMP(),?)
         ON DUPLICATE KEY UPDATE contents = ?`
  _, err = tx.Exec(sql,
                   app.Application_seq, row["application_no"],
                   app.LoginToken["user_seq"], row["application_no"])
  if err != nil {
    log.Println(err)
    return
  }

  sql = `SELECT IF(item_name <> 'general_end_date', concat("ibc_",item_name) , item_name) AS item_name,
                 target_item, contents
            FROM t_application_etc
           WHERE application_seq = ?
             AND item_name IN('general_end_date', 'general_experiment_cnt', 'general_experiment_degree', 'general_fund_org_name')`
  rows1 := common.DB_Tx_fetch_all(tx, sql, nil, iacucAppSeq)
  for _, row1 := range rows1 {
    sql = `INSERT INTO t_application_etc(
          	application_seq, item_name, target_item,
          	contents, reg_dttm, reg_user_seq
           )
           VALUES(?,?,?,
          		    ?,UNIX_TIMESTAMP(),?)
           ON DUPLICATE KEY UPDATE contents = ?`
    _, err = tx.Exec(sql,
                     app.Application_seq, row1["item_name"], row1["target_item"],
                     row1["contents"], app.LoginToken["user_seq"],
                     row1["contents"])
  	if err != nil {
  	  log.Println(err)
  	  return
  	}
  }

  sql = `SELECT concat("ibc_",item_name) AS item_name,
                item_idx, checked
           FROM t_application_main_check
          WHERE application_seq = ?
            AND item_name = 'general_fund_org'`
  row2 := common.DB_Tx_fetch_one(tx, sql, nil, iacucAppSeq)

  sql = `INSERT INTO t_application_main_check(
        	application_seq, item_name,
          item_idx, checked
        )VALUES(?,?,?,?)
        ON DUPLICATE KEY UPDATE item_idx = ?, checked = ?`
  _, err = tx.Exec(sql,
                   app.Application_seq, row2["item_name"],
                   row2["item_idx"], row2["checked"],
                   row2["item_idx"], row2["checked"])
  if err != nil {
    log.Println(err)
    return
  }

  sql = `SELECT concat("ibc_",item_name) AS item_name, item_idx, select_ids
           FROM t_application_select
          WHERE application_seq = ?
            AND item_name = 'general_fund_conflict'`
  row3 := common.DB_Tx_fetch_one(tx, sql, nil, iacucAppSeq)

  sql = `INSERT INTO t_application_select(
  	       application_seq, item_name,
           item_idx, select_ids)
         VALUES(?,?,?,?)
         ON DUPLICATE KEY UPDATE item_idx = ?, select_ids = ?`
  _, err = tx.Exec(sql,
                   app.Application_seq, row3["item_name"],
                   row3["item_idx"], row3["select_ids"],
                   row3["item_idx"], row3["select_ids"])
  if err != nil {
   log.Println(err)
   return
  }

  sql = `SELECT group_concat(DISTINCT animal_code_str) as animal_name
           FROM t_application_animal
          WHERE application_seq = ?
	          AND item_name IN ('animal_type_final', 'animal_type')`
  row4 := common.DB_Tx_fetch_one(tx, sql, nil, iacucAppSeq)
  if "" != common.ToStr(row4["animal_name"]) {
    sql = `INSERT INTO t_application_main_check(
            application_seq, item_name,
            item_idx, checked
          )VALUES(?,'ibc_general_animal_flag',0,0)
          ON DUPLICATE KEY UPDATE item_idx = 0, checked = 0`
    _, err = tx.Exec(sql, app.Application_seq)
    if err != nil {
      log.Println(err)
      return
    }

    sql = `INSERT INTO t_application_etc(
            application_seq, item_name, target_item,
            contents, reg_dttm, reg_user_seq
           )
           VALUES(?,'ibc_general_animal_name','ibc_general_animal_name',
                  ?,UNIX_TIMESTAMP(),?)
           ON DUPLICATE KEY UPDATE contents = ?`
    _, err = tx.Exec(sql,
                     app.Application_seq, row4["animal_name"],
                     app.LoginToken["user_seq"], row4["animal_name"])
    if err != nil {
      log.Println(err)
      return
    }

  }

  err = tx.Commit()
  if nil != err {
    log.Println(err)
    succ = false
  } else {
    succ = true
  }
	return
}

func (app *Application)NewIbc(tx *sql.Tx) (succ bool) {
  succ = false
  app_basic := AppBasic{ Application : app }
	succ2, application_no := app_basic.GetAppNo(tx, DEF_APP_JUDGE_TYPE_CODE_IBC, DEF_APP_TYPE_NEW, uint(0))
  if !succ2 {
    return
  }

	sql := `INSERT INTO t_application(
											 application_no,	 	institution_seq, 	judge_type,
											 application_type,	application_step,	application_result,
											 reg_user_seq,      chg_user_seq,     reg_dttm)
					      VALUES(?, ?, ?,
											 ?, ?, ?,
											 ?, ?, UNIX_TIMESTAMP())`
	result, err := tx.Exec(sql, application_no, app.LoginToken["institution_seq"], DEF_APP_JUDGE_TYPE_CODE_IBC,
															DEF_APP_TYPE_NEW, DEF_APP_STEP_WRITE, DEF_APP_RESULT_TEMP,
															app.LoginToken["user_seq"], app.LoginToken["user_seq"])
	if err != nil {
	  log.Println(err)
	  return
	}

	no, err := result.LastInsertId()
	if err != nil {
	 log.Println(err)
	 return
	}

  app.Application_seq = uint(no)

  succ = true
  return
}
