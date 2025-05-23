
package model

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
  "database/sql"
  "ipsap/common"
  "fmt"
  "log"
	// "os"
)

type AppFile struct {
  Application	*Application
  Datas				[]map[string]interface{}
}

func (ins *AppFile) Load(item_name string) (succ bool) {
  if ins.Application.Application_seq == 0 {
    return
  }

  if ins.Datas != nil {
    return true
  }

  sql := `SELECT filepath, org_file_name, file_idx
            FROM t_application_file
           WHERE application_seq = ?
             AND item_name = ?
           ORDER BY view_order`
  filter := func(row map[string]interface{}) {
		// row["filepath"] = common.EncryptToUrl([]byte(common.ToStr(ins.Application.LoginToken["tmp_key"])), common.ToStr(row["filepath"]))
		row["filepath"] = common.EncryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(row["filepath"]))
		row["src"] 	= common.MakeDownloadUrl(common.ToStr(row["filepath"]))
	}

  ins.Datas = common.DB_fetch_all(sql, filter, ins.Application.Application_seq, item_name)

  //log.Println(ins.Datas)

  return true
}

func (ins *AppFile) GetJsonData() (ret interface{}) {
  ret = ins.Datas
  return
}

func (ins *AppFile) getAppFileName(item_name string) (int) {
	sql := `SELECT max(file_idx) idx
  					FROM t_application_file
 					 WHERE application_seq = ?
					 	 AND item_name = ?`
	row := common.DB_fetch_one(sql, nil, ins.Application.Application_seq, item_name)
	if nil == row {
		return 1
	} else {
		return common.ToInt(row["idx"]) + 1
	}
}

func (ins *AppFile) getDirPath(subDir string) (string) {
	return fmt.Sprintf("%v/%v/%v/", DIR_APPLICATION, ins.Application.Application_seq, subDir)
}

func (ins *AppFile)UpdateItem(c *gin.Context, tx *sql.Tx, item_name string, data1 interface{}) (ret bool, err_msg string) {
	data := data1.(map[string]interface{})
	// log.Println(data)

	if !ins.uploadFilesAndUpdateItem(c, tx, item_name) {
		err_msg = fmt.Sprintf("[%v] 파일 업로드 실패", item_name);
		return
	}

	if len(data) > 0 {
		// 기존값과 비교해서 DELETE 해준다..
		for encFilePath, useFlag := range data {
			if !useFlag.(bool) {
				// filePath := common.DecryptToUrl([]byte(common.ToStr(ins.Application.LoginToken["tmp_key"])), common.ToStr(encFilePath))
				filePath := common.DecryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(encFilePath))
				ins.deleteAppFile(tx, item_name, filePath)
			}
		}
	}

  ret = true
	return
}

func (ins *AppFile) deleteAppFile(tx *sql.Tx, item_name string, filePath string) {
	sql := `DELETE FROM t_application_file
					 WHERE application_seq = ?
					 	 AND item_name = ?
						 AND filepath = ?`
	_, err := tx.Exec(sql, ins.Application.Application_seq,
                    item_name, filePath)
  if err != nil {
    log.Println(err)
  }


	// err = os.Remove(common.Config.Server.FileUploadPath + filePath)
	err = common.RemoveFileToS3(filePath)
	if nil != err{
		log.Println(err)
	}

	return
}

func (ins *AppFile) uploadFilesAndUpdateItem(c *gin.Context, tx *sql.Tx, param string) (succ bool) {
	newFileName := ins.getAppFileName(param)
	sub_path		:= ins.getDirPath(param)
	// src := common.Config.Server.FileUploadPath + sub_path
	src := sub_path

	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err)
		return
	}

	// if err = os.MkdirAll(src, os.ModePerm); err != nil {
	// 	log.Printf("Mkdir Fail! : %v",err)
	// 	return
	// }

	files := form.File[param]

	// log.Println(param, newFileName, files, len(files))

	if len(files) > 0 {
		for idx, fileHeader := range files {
			file, err := fileHeader.Open()
			if nil != err {
				log.Println(err)
				return
			}
			newFileNameStr := common.ToStr(newFileName + idx)
			extention := filepath.Ext(src + fileHeader.Filename)
			filePath := sub_path + newFileNameStr + extention
			err = common.AddFileToS3(filePath, file, fileHeader)
			if err != nil {
				log.Println(err)
				return
			}
			sql := `INSERT INTO t_application_file(application_seq, item_name, filepath, org_file_name, file_idx, view_order)
							VALUES(?,?,?,?,?,?)`
			_, err2 := tx.Exec(sql, ins.Application.Application_seq, param, filePath, fileHeader.Filename, newFileNameStr, newFileNameStr)
			if nil != err2 {
				log.Println(err2)
				return
			}
		}
	}
	succ = true
	return
}

func (ins *AppFile) UploadFile(c *gin.Context, tx *sql.Tx, item_name string, file_idx int) (succ bool, file_exists bool) {
	newFileName := file_idx
	sub_path		:= ins.getDirPath(item_name)
	// src := common.Config.Server.FileUploadPath + sub_path
	src := sub_path

	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err)
		return
	}

	// if err = os.MkdirAll(src, os.ModePerm); err != nil {
	// 	log.Printf("Mkdir Fail! : %v",err)
	// 	return
	// }

	file_tag_name := item_name + "_" + common.ToStr(file_idx)
	files := form.File[file_tag_name]
	if (len(files) == 0)	{		//	파일 첨부가 없을때 String값이 저장되어 있는지 확인
		data := c.PostForm(file_tag_name)
		if common.ToStr(file_idx) == data {		//	key 값이 일치하면 파일 유지!!!
			file_exists = true
			succ = true
			return
		}
	}

	log.Println(file_tag_name, newFileName, files, len(files))
	if len(files) == 1 {		//	must len(files) = 1
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if nil != err {
				log.Println(err)
				return
			}
			newFileNameStr := common.ToStr(newFileName)
			extention := filepath.Ext(src + fileHeader.Filename)
			filePath := sub_path + newFileNameStr + extention
			err = common.AddFileToS3(filePath, file, fileHeader)
			if err != nil {
				log.Println(err)
				return
			}
			sql := `INSERT INTO t_application_file(application_seq, item_name, filepath, org_file_name, file_idx, view_order)
							VALUES(?,?,?,?,?,?)
							ON DUPLICATE KEY UPDATE org_file_name = ?, file_idx = ?, view_order = ?`
			_, err2 := tx.Exec(sql, ins.Application.Application_seq, item_name, filePath,
												 fileHeader.Filename, newFileNameStr, newFileNameStr,
												 fileHeader.Filename, newFileNameStr, newFileNameStr)
			if nil != err2 {
				log.Println(err2)
				return
			}
			file_exists = true
		}
	}
	succ = true
	return
}
