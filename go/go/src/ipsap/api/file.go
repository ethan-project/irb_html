package api

import (
	"github.com/gin-gonic/gin"
	"ipsap/common"
	"ipsap/model"
	// "path/filepath"
  // "strings"
	// "log"
  // "net/http"
	// "fmt"
	// "os"
)

// @Tags File
// @Summary 파일 다운로드
// @Description 파일 다운로드
// @Accept  json
// @Produce  json
// @Param filepath_enc path string true "encrypt filepath"
// @Param token query string false "token"
// @Router /file/{filepath_enc} [get]
// @Success 200
func File_Download(c *gin.Context) {
	token := c.Request.URL.Query().Get("token")
	encFilePath	:= c.Param("filepath_enc")
	filePath := common.DecryptToUrl([]byte(common.Config.Program.EncryptionKey), common.ToStr(encFilePath))
	fileName, succ := common.CheckTokenAuthFile(c, token, filePath)
	if !succ {
		return
	}

	common.DownloadFileToS3(c, filePath, fileName)
}

// @Tags File
// @Summary 실험 동물 데이터 엑셀 다운로드
// @Description 실험 동물 데이터 엑셀 다운로드
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param filter.institution_seq query string true "1"
// @Param filter.reg_user_seq query string false "1"
// @Router /file-animal [get]
// @Success 200
func Animal_Data_Download(c *gin.Context) {
	model.AnimalDownload(c, common.ToUint(c.Request.URL.Query().Get("filter.institution_seq")), c.Request.URL.Query().Get("filter.reg_user_seq"))
}
