package common

import (
	"database/sql"
	"runtime"
	"strings"
	"path/filepath"
	"archive/zip"
	"io"
	"os"

	//  "fmt"
	//  "log"
	//  "strconv"
	"encoding/json"
	"github.com/dustin/go-humanize"
	"github.com/nleeper/goment"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"reflect"
)

func FinishApi(c *gin.Context, code int, json ...interface{}) {
	if code == Api_status_not_found {
		c.String(404, "404 page not found")
		c.Abort()
		return
	}
	for _, s := range json {
		c.JSON(code, s)
	}
	/*for _, s := range json {
		vv := s.([]map[string]interface{})
		c.JSON(code, vv[4])

	}*/
	c.AbortWithStatus(code)
}

func FinishApiWithErrCd(c *gin.Context, code int, err_code int, eMsg ...string) {
	for _, em := range eMsg {
		FinishApi(c, code, gin.H{
			"e":  err_code,
			"em": em,
		})
		return
	}

	FinishApi(c, code, gin.H{
		"e":  err_code,
		"em": error_msg[err_code],
	})
}

/*
func end404(c *gin.Context) {
  c.String(404, "404 page not found")
}

func jsonwrite(c *gin.Context, code int, obj interface{}) {
    c.Header("Content-Type", "application/json; charset=utf-8")
    c.JSON(code, obj)
}
*/

func UnmarshalFormData(c *gin.Context, param string) (ret map[string]interface{}, err string) {
	data := c.PostForm(param)
	if "" == data {
		return nil, "데이터가 존재하지 않습니다."
	}

	if err := json.Unmarshal([]byte(data), &ret); err != nil {
		log.Println("Decoding Fail!!!!!!!!!! {" + string(data) + "}")
		return nil, "json 포멧이 아닙니다."
	}
	return
}

func BindCustomBodyToJson(c *gin.Context) (map[string]interface{}, string) {
	//  Json 만 처리한다.
	//  START : body json ---
	data, _ := c.GetRawData()
	var dat map[string]interface{}
	if err := json.Unmarshal(data, &dat); err != nil {
		log.Println("Decoding Fail!!!!!!!!!! {" + string(data) + "}")
		return nil, "json 포멧이 아닙니다."
	}

	return dat, ""
}

func BindCustomBody(c *gin.Context, chkKeys *[]interface{}) (map[string]interface{}, string) {

	if len(*chkKeys) != 3 {
		log.Println("시스템 오류(chk key 값 오류)")
		return nil, "시스템 에러"
	}

	//  Json 만 처리한다.
	//  START : body json ---
	data, _ := c.GetRawData()
	var dat map[string]interface{}
	if err := json.Unmarshal(data, &dat); err != nil {
		log.Println("Decoding Fail!!!!!!!!!! {" + string(data) + "}")
		return nil, "json 포멧이 아닙니다."
	}

	return BindCustomBodyWithDat(chkKeys, dat)
}

func BindCustomBodyWithDat(chkKeys *[]interface{}, dat map[string]interface{}) (map[string]interface{}, string) {
	var newData map[string]interface{}
	newData = make(map[string]interface{})

	// 1. 필수 항목 check
	needObject := reflect.ValueOf((*chkKeys)[0])
	for i := 0; i < needObject.Len(); i++ {
		key := ToStr(needObject.Index(i).Interface())
		if value, ok := dat[key]; ok && ToStr(value) != "" {
			newData[key] = value
		} else {
			log.Println("필수 정보가 (" + key + ") 가 없습니다.")
			return nil, "필수 정보가 (" + key + ") 가 없습니다."
		}
	}

	// 2. 그룹 옵션 항목 check
	optionObjects := reflect.ValueOf((*chkKeys)[1])
	for i := 0; i < optionObjects.Len(); i++ {
		option := optionObjects.Index(i).Interface()
		optionObject := reflect.ValueOf(option)

		bFind := false
		for j := 0; j < optionObject.Len(); j++ {
			key := ToStr(optionObject.Index(j).Interface())
			if value, ok := dat[key]; ok && ToStr(value) != "" {
				newData[key] = value
				bFind = true
			}
		}
		if !bFind {
			log.Println("최소한의 옵션 정보가 한개도 없습니다.")
			return nil, "최소한의 옵션 정보가 한개도 없습니다."
		}
	}

	// 3. 기타 항목 check
	etcObject := reflect.ValueOf((*chkKeys)[2])
	for i := 0; i < etcObject.Len(); i++ {
		key := ToStr(etcObject.Index(i).Interface())
		if value, ok := dat[key]; ok {
			newData[key] = value
		}
	}

	return newData, ""
}

func ConvertArr(interfaceArr []interface{}) (resultArr []string) {
	resultArr = make([]string, len(interfaceArr))
	for i, v := range interfaceArr {
		resultArr[i] = fmt.Sprint(v)
	}
	return
}

func SetDbTimeZone(tx *sql.Tx, userId interface{}) bool {
	_, err := tx.Exec("SET time_zone = (SELECT timezone_gmt_add FROM t_timezone WHERE timezone_seq = (SELECT timezone_seq FROM t_user WHERE user_seq = ?))", userId)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return false
	}
	return true
}

func GetDbCurrentDate(userId interface{}) (currentDate string, succ bool) {

	currentDate = ""
	succ = false

	dbConn := DBconn()
	tx, err := dbConn.Begin()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}

	if !SetDbTimeZone(tx, userId) {
		tx.Rollback()
		return
	}

	sql := "SELECT DATE_FORMAT(NOW(),'%Y%m%d') current_date_str"
	rows := DB_Tx_fetch_one(tx, sql, nil)
	if rows == nil {
		tx.Rollback()
		return
	}

	currentDate = ToStr(rows["current_date_str"])

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}

	succ = true

	return
}

func GetFileSqlVersion() string {

	_, file, _, _ := runtime.Caller(1)

	idx := strings.LastIndex(file, "_")
	runes := []rune(file)
	safeSubstring := string(runes[idx+1 : len(file)])
	safeSubstring = strings.TrimRight(safeSubstring, ".go")
	return safeSubstring
}

// 맵 파싱
func MapParsing(str string, params map[string]interface{}) string {

	for key, val := range params {
		str = strings.ReplaceAll(str, "{{"+key+"}}", ToStr(val))
	}

	return str
}

func Unzip(src string, dest string) ([]string, error) {
    var filenames []string
    r, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer r.Close()

    for _, f := range r.File {
      fpath := filepath.Join(dest, f.Name)

      if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
          return filenames, fmt.Errorf("%s: wrong file path", fpath)
      }

      filenames = append(filenames, fpath)

      if f.FileInfo().IsDir() {
          // Make Folder
          os.MkdirAll(fpath, os.ModePerm)
          continue
      }

      // Make File
      if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
          return filenames, err
      }

      outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
      if err != nil {
          return filenames, err
      }

      rc, err := f.Open()
      if err != nil {
          return filenames, err
      }

      _, err = io.Copy(outFile, rc)

      outFile.Close()
      rc.Close()

      if err != nil {
          return filenames, err
      }
    }

    return filenames, nil
}

func MakeDownloadUrl(path string) (url string) {
	data := Config.Server
	if Config.Test.IsTestMode {
		url = data.Protocol + "://" + "218.232.81.113" + ":" + "17370" + data.ApiPath + "/file/" + path
	} else {
		url = data.Protocol + "://" + data.SwagHostname + ":" + ToStr(data.Port) + data.ApiPath + "/file/" + path
	}
	return
}

func CheckUserTypeAuth(loginUserTypeArr []string, checkUserTypeArr ...int) bool {
	for _, loginUserType := range loginUserTypeArr {
		for _, checkUserType := range checkUserTypeArr {
			if ToInt(loginUserType) == checkUserType {
				return true
			}
		}
	}
	return false
}

func GetDateStr(t *goment.Goment) string{
	g, _ := goment.New(t)
	return g.Format("YYYY년 M월 D일")
}

func GetDateTimeStr(t *goment.Goment) string{
	g, _ := goment.New(t)
	return g.Format("YYYY년 M월 D일 H시 m분 s초")
}

func GetMembershipExpirationDate(t *goment.Goment) string{
	t.Add(1, "years")
	t.EndOf("month")
	return t.Format("YYYY년 M월 D일")
}

func GetStartDateStr(t *goment.Goment ) string{
	t.SetDate(1)
	return GetDateStr(t)
}

func GetCurrency(v int64) string {
	return fmt.Sprintf("%v원", humanize.Comma(v))
}
