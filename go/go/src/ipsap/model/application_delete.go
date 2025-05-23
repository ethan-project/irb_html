
package model

import (
	// "github.com/gin-gonic/gin"
  "ipsap/common"
  "database/sql"
  // "strings"
	// "fmt"
  "log"
)

func (app *Application)Delete() (succ bool) {
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

	// data := app.Load().(map[string]interface{})
	// switch common.ToUint(data["judge_type"]) {
	// 	case DEF_APP_JUDGE_TYPE_CODE_IACUC:
	// 		if !app.IacucDelete(tx) {
	// 			return
	// 		}
	// 	case DEF_APP_JUDGE_TYPE_CODE_IBC:
  //
	// 	case DEF_APP_JUDGE_TYPE_CODE_IRB:
	// 			// app.IrbDelete(tx, new_app_seq)
	// 	default :
	// 		log.Printf("copy judge_type error --> %v", data["judge_type"])
	// 		return
	// }

  if !app.DeleteApp(tx) {
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

func (app *Application) ChangeToDeleteState() (succ bool) {
  sql := `UPDATE t_application
             SET application_result = ?
           WHERE application_seq = ?`
  _, err := common.DBconn().Exec(sql, DEF_APP_RESULT_DELETED, app.Application_seq)

  if nil != err {
 		log.Println(err)
 		return false
 	}

	succ = true
	return

}

func (app *Application)DeleteApp(tx *sql.Tx) (succ bool) {
	succ = false

  sql := `SELECT filepath FROM t_application_file WHERE application_seq = ?`
  rows := common.DB_Tx_fetch_all(tx, sql, nil, app.Application_seq)
  if len(rows) > 0 {
    for _, row := range rows {
      common.RemoveFileToS3(common.ToStr(row["filepath"]))
    }
  }

  // 210629 donghun : 신청서 삭제 테이블에 insert 후 app no 발급시 같이 count 해준다.
  sql = `INSERT INTO t_application_deleted
         SELECT *, UNIX_TIMESTAMP()
           FROM t_application
          WHERE application_seq = ?`
  _,err := tx.Exec(sql, app.Application_seq)
	if err != nil {
	  log.Printf(`t_application_deleted ==> %v`,err)
	  return
	}

	sql = `DELETE FROM t_application WHERE application_seq = ?`;
	 _,err = tx.Exec(sql, app.Application_seq)
 	if err != nil {
 	  log.Printf(`t_application ==> %v`,err)
 	  return
 	}

  sql = `DELETE FROM t_application_anesthetic WHERE application_seq = ?`;
   _,err = tx.Exec(sql, app.Application_seq)
  if err != nil {
    log.Printf(`t_application_anesthetic ==> %v`,err)
    return
  }

  sql = `DELETE FROM t_application_animal WHERE application_seq = ?`;
	 _,err = tx.Exec(sql, app.Application_seq)
 	if err != nil {
 	  log.Printf(`t_application_animal ==> %v`,err)
 	  return
 	}

  sql = `DELETE FROM t_application_etc WHERE application_seq = ?`;
	 _,err = tx.Exec(sql, app.Application_seq)
 	if err != nil {
 	  log.Printf(`t_application_etc ==> %v`,err)
 	  return
 	}

  sql = `DELETE FROM t_application WHERE application_seq = ?`;
	 _,err = tx.Exec(sql, app.Application_seq)
 	if err != nil {
 	  log.Printf(`t_application_anesthetic ==> %v`,err)
 	  return
 	}

  sql = `DELETE FROM t_application_file WHERE application_seq = ?`;
	 _,err = tx.Exec(sql, app.Application_seq)
 	if err != nil {
 	  log.Printf(`t_application_file ==> %v`,err)
 	  return
 	}

  sql = `DELETE FROM t_application_hist WHERE application_seq = ?`;
   _,err = tx.Exec(sql, app.Application_seq)
  if err != nil {
    log.Printf(`t_application_hist ==> %v`,err)
    return
  }

  sql = `DELETE FROM t_application_keyword WHERE application_seq = ?`;
   _,err = tx.Exec(sql, app.Application_seq)
  if err != nil {
    log.Printf(`t_application_keyword ==> %v`,err)
    return
  }

  sql = `DELETE FROM t_application_main_check WHERE application_seq = ?`;
   _,err = tx.Exec(sql, app.Application_seq)
  if err != nil {
    log.Printf(`t_application_main_check ==> %v`,err)
    return
  }

  sql = `DELETE FROM t_application_member WHERE application_seq = ?`;
   _,err = tx.Exec(sql, app.Application_seq)
  if err != nil {
    log.Printf(`t_application_member ==> %v`,err)
    return
  }

  sql = `DELETE FROM t_application_select WHERE application_seq = ?`;
   _,err = tx.Exec(sql, app.Application_seq)
  if err != nil {
    log.Printf(`t_application_member ==> %v`,err)
    return
  }

  sql = `DELETE FROM t_application_select_input WHERE application_seq = ?`;
   _,err = tx.Exec(sql, app.Application_seq)
  if err != nil {
    log.Printf(`t_application_member ==> %v`,err)
    return
  }

	return true
}
