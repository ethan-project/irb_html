
package common

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"log"
)

type FileUpload struct {
	Required 				bool
	Param						string
	Src							string
	File_name 			string
	File_extension	string
	New_file_name		string
	Max_size				int64
}

func (fup *FileUpload) UploadFile(c *gin.Context) (succ bool) {
	succ = false;
	if nil == c || "" == fup.Param {
		return
	}

	file, fileHeader, err := c.Request.FormFile(fup.Param)
	if nil == file {
		if true == fup.Required {
			return
		} else {
			fup.New_file_name = ""
			return true
		}
	}

	if nil != err {
		log.Println(err)
		return
	}

	err = AddFileToS3(fup.Src + fup.New_file_name + filepath.Ext(fileHeader.Filename), file, fileHeader)
	if err != nil {
		log.Println(err)
		return
	}

	fup.File_name 			= fileHeader.Filename
	fup.File_extension	= filepath.Ext(fileHeader.Filename)

	return true
}
