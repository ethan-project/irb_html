
package model

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"database/sql"
  "ipsap/common"
	// "strings"
	"log"
	"fmt"
)

type Board struct {
  Board_seq				uint   `json:"-"`
	Institution_seq uint   `json:"institution_seq" example:"1"`
	User_seq				uint   `json:"-"`
	Board_no 			  uint   `json:"-"`
	Board_type 			uint   `json:"board_type" example:"1"`
  View_order      uint   `json:"view_order" example:"1"`
	Title 			    string `json:"title" example:"제목"`
	Text 			      string `json:"text" example:"내용"`
  File_path 			string `json:"-"`
  File_org_name   string `json:"-"`
  File_idx 			  string `json:"-"`
	Chg_user_seq		uint   `json:"-"`
	File_delete			bool   `json:"file_delete" example:"true"`
	Data 						map[string]interface{}  `json:"-"`
}

func (board *Board)getBoardQueryAndFilter(moreCondition string, infoFlag bool) (sql string, filter func(map[string]interface{}))  {
  orderBy := ""
  if DEF_BOARD_TYPE_FAQ == board.Board_type {
  	orderBy = `board.view_order ASC`
  } else {
		orderBy = `board.board_no DESC`
	}
  sql = fmt.Sprintf(`
					SELECT board.board_seq,					board.institution_seq,	board.user_seq,
								 user.name AS user_name,	board.title,						board.view_cnt,
								 board.file_path,         board.file_org_name,    board.text,
                 board.board_type,				board.board_no,         board.view_order,
                 FROM_UNIXTIME(board.reg_dttm, '%%Y-%%m-%%d') AS reg_dttm
						FROM t_board board, t_user user
					 WHERE board.user_seq = user.user_seq
							%v #moreCondtion
        ORDER BY %v #orderBy`, moreCondition, orderBy)
	filter = func(row map[string]interface{}) {
		if infoFlag {
			sql2 := fmt.Sprintf(`
							 SELECT board_seq, board.board_no, board.title,
											user.name AS user_name,	board.view_cnt,
											FROM_UNIXTIME(board.reg_dttm, '%%Y-%%m-%%d') AS reg_dttm
								 FROM t_board board, t_user user
								WHERE board.user_seq = user.user_seq
									AND board.board_no < ?
									AND board.board_type =?
									AND board.institution_seq = ?
						 ORDER BY board.board_no DESC LIMIT 1`)
		  row["prev_info"] = common.DB_fetch_one(sql2, nil, row["board_no"], row["board_type"], row["institution_seq"])

			sql2 = fmt.Sprintf(`
							SELECT board_seq, board.board_no, board.title,
										 user.name AS user_name,	board.view_cnt,
										 FROM_UNIXTIME(board.reg_dttm, '%%Y-%%m-%%d') AS reg_dttm
							  FROM t_board board, t_user user
							 WHERE board.user_seq = user.user_seq
								 AND board_no > ?
								 AND board_type =?
								 AND board.institution_seq = ?
					  ORDER BY board.board_no LIMIT 1`)
		  row["next_info"] = common.DB_fetch_one(sql2, nil, row["board_no"], row["board_type"], row["institution_seq"])
		}
		if "" != common.ToStr(row["file_org_name"]) {
			row["file_path"]	= common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["file_path"]))
			row["file_src"]	= common.MakeDownloadUrl(common.ToStr(row["file_path"]))
		}
	}

  return
}

func (board *Board)LoadList(moreCondition string) (rows [] map[string]interface{}) {
	sql, filter := board.getBoardQueryAndFilter(moreCondition, false)
	rows = common.DB_fetch_all(sql, filter)
	return
}

func (board *Board)Load(countFlag bool) (succ bool) {
  if board.Board_seq == 0 {
    return
  }

  moreCondition := fmt.Sprintf(` AND board.board_seq = %d`, board.Board_seq)
	sql, filter := board.getBoardQueryAndFilter(moreCondition, true)
  board.Data = common.DB_fetch_one(sql, filter)

  if countFlag && nil != board.Data {
    sql = `UPDATE t_board SET view_cnt = view_cnt + 1 WHERE board_seq = ?`
    _, err := common.DBconn().Exec(sql, board.Board_seq)
    if err != nil {
      log.Println(err)
    }
  }

  return nil != board.Data
}

func (board *Board)UpdateBoard(c *gin.Context) (succ bool) {
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

	sql := `SELECT file_idx + 1 as file_idx
						FROM t_board
					 WHERE board_seq = ?`
	row := common.DB_Tx_fetch_one(tx, sql, nil, board.Board_seq)
	board.File_idx = common.ToStr(row["file_idx"])

  sql  = `UPDATE t_board
						 SET title = ?,				 	text = ?,	view_order = ?,
						 		 chg_user_seq = ?,	chg_dttm = UNIX_TIMESTAMP()
					 WHERE board_seq = ?
				 `
	_, err = tx.Exec(sql, board.Title,				board.Text,		board.View_order,
                        board.Chg_user_seq,	board.Board_seq)
	if err != nil {
		log.Println(err)
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

  if !board.BoardFileUpload(c, tx){
    return
  }

  err = tx.Commit()
  if nil != err {
    log.Println(err)
    common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
    return
  }

	return true
}

func (board *Board)InsertBoard(c *gin.Context) (succ bool) {
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

  board.getBoardNo(tx)

  sql := `INSERT INTO t_board(institution_seq, user_seq,       board_no,
                              board_type,      title,          text,
	                            view_order,      reg_dttm)
					VALUES(?, ?, ?,
								 ?, ?, ?,
								 ?, UNIX_TIMESTAMP())
				 `
	result, err := tx.Exec(sql, board.Institution_seq, board.User_seq,       board.Board_no,
                         board.Board_type,      board.Title,          board.Text,
                         board.View_order)
	if err != nil {
		log.Println(err)
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

  no, err := result.LastInsertId()
  if err != nil {
    log.Println(err)
    common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

  board.Board_seq = cast.ToUint(no)

  if !board.BoardFileUpload(c, tx){
    return
  }

  err = tx.Commit()
  if nil != err {
    log.Println(err)
    common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
    return
  }

	return true
}

func (board *Board)getBoardNo(tx *sql.Tx) (succ bool) {
  succ = false
  sql := `SELECT IFNULL(max(board_no), 0) + 1 AS board_no
            FROM t_board
           WHERE institution_seq = ?
             AND board_type = ?`
  row := common.DB_Tx_fetch_one(tx, sql, nil, board.Institution_seq, board.Board_type)
  if nil == row {
    return
  }

  board.Board_no = common.ToUint(row["board_no"])
  return true
}

func (board *Board)GetDirPath(subDir string) (string) {
	return fmt.Sprintf("%v/%v/%v/", "board", board.Institution_seq, subDir)
}

func (board *Board)BoardFileUpload(c *gin.Context, tx *sql.Tx) (succ bool) {
	fup := common.FileUpload{}
	fup.Required = false
	fup.Param	= "board_file"
	fup.New_file_name	= board.File_idx
	sub_path := board.GetDirPath(common.ToStr(board.Board_seq))
	// fup.Src = common.Config.Server.FileUploadPath + sub_path
	fup.Src = sub_path

	if !fup.UploadFile(c) {
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_upload_fail)
		return
	}

	if "" == fup.New_file_name {
		if !board.File_delete {
			return true
		}
		board.File_idx = common.ToStr(common.ToUint(board.File_idx) - 1)
	}

	file_path := sub_path + fup.New_file_name + fup.File_extension

	sql := `UPDATE t_board
						 SET file_org_name = ?,
						 		 file_path = ?,
								 file_idx = ?
					 WHERE board_seq = ?`
  _, err := tx.Exec(sql, fup.File_name, file_path, board.File_idx, board.Board_seq)
	if err != nil {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_bad_request, common.Error_invalide_params)
    return
  }

	board.File_path 		 = file_path
	board.File_org_name  = fup.File_name
	return true
}

func (board *Board)DeleteBoard(c *gin.Context) (succ bool) {
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

	board.Load(false)

	sql := `DELETE FROM t_board WHERE board_seq = ?`
	_, err = tx.Exec(sql, board.Board_seq)
	if nil != err {
		log.Println(err)
		common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
		return
	}

	sql = `SELECT ROW_NUMBER() OVER (ORDER BY reg_dttm) AS row_num, board_seq
					 FROM t_board
				  WHERE institution_seq = ?
					  AND board_type = ?`
	rows := common.DB_Tx_fetch_all(tx, sql, nil, board.Data["institution_seq"], board.Data["board_type"])
	if len(rows) > 0 {
		for _, row := range rows {
			sql = `UPDATE t_board
						    SET board_no = ?
						  WHERE board_seq = ?`
			_, err = tx.Exec(sql, row["row_num"], row["board_seq"])
			if nil != err {
				log.Println(err)
				common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
				return
			}
		}
	}

	err = tx.Commit()
  if nil != err {
    log.Println(err)
    common.FinishApiWithErrCd(c, common.Api_status_forbidden, common.Error_system_unknown)
    return
  }

	return true
}
