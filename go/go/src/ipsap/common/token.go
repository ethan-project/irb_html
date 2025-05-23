package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"

	//	"fmt"
	"io"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

const enc_key_string = "ipsap_key"

func get_key_string(user_seq string) string {
	key_string := user_seq + enc_key_string + "1234567890123456" //  16자리 이상 만들기 위함!
	runes := []rune(key_string)
	safeSubstring := string(runes[0:16])
	return safeSubstring
}

func Make_token(c *gin.Context, userInfo map[string]interface{}) string {
	token := map[string]interface{}{
		"user_seq":        ToStr(userInfo["user_seq"]),
		"institution_seq": ToStr(userInfo["institution_seq"]),
		"user_type":       ToStr(userInfo["user_type"]),
		"user_type_all":   ToStr(userInfo["user_type_all"]),
		"user_auth":       ToInt(userInfo["user_auth"]),
		"email":           ToStr(userInfo["email"]),
		"tmp_key":         ToStr(userInfo["tmp_key"]),
	}

	jsonData, err := json.Marshal(token)
	if nil != err {
		return ""
	}

	key := []byte(get_key_string(c.ClientIP()))
	cryptoText := EncryptToStd(key, string(jsonData))

	log.Println("cryptoText : ", cryptoText)

	return cryptoText
}

func Check_token(c *gin.Context) map[string]interface{} {
	token := c.GetHeader("token")

	//local 테스트트
	if "" == token {
		data := map[string]interface{}{
			"user_seq":        "1",
			"institution_seq": "1",
			"user_type":       "1",
			"user_type_all":   "1,2,3,4",
			"user_auth":       "9",
			"email":           "test@test.com",
			"tmp_key":         Config.Program.EncryptionKey,
		}
		return data
	}

	log.Println("token : ", token)

	if "" == token {
		log.Println("Token Not Exists")
		FinishApiWithErrCd(c, Api_status_bad_request, Error_token_mismatch)
		return nil
	}
	data := make(map[string]interface{})
	key := []byte(get_key_string(c.ClientIP()))
	text := DecryptToStd(key, token)
	byt := []byte(text)

	log.Println("data : ", data)
	log.Println("key : ", key)
	log.Println("text : ", text)
	log.Println("byt : ", byt)

	if err := json.Unmarshal(byt, &data); nil != err {
		log.Println("Decoding Fail!!!!!!!!!!")
		FinishApiWithErrCd(c, Api_status_bad_request, Error_token_mismatch)
		return nil
	}

	sql := "SELECT email FROM t_user WHERE user_seq = ? AND institution_seq = ?"
	userInfo := DB_fetch_one(sql, nil, data["user_seq"], data["institution_seq"])

	log.Println("sql : ", sql)
	log.Println("userInfo : ", userInfo)

	if data["email"] != ToStr(userInfo["email"]) {
		FinishApiWithErrCd(c, Api_status_bad_request, Error_id_mismatch)
		return nil
	}

	return data
}

// /////////////////////////////////////////////////////////
// Encrypt string to base64 crypto using AES
func Encrypt(key []byte, text string, stdFlag bool) string {
	defer func() {
		if err := recover(); nil != err {
			log.Println(err)
		}
	}()

	if 0 == len(text)%aes.BlockSize {
		text = text + " "
	}

	plaintext := []byte(text)

	b, err := aes.NewCipher(key)
	if nil != err {
		log.Println(err)
		return ""
	}

	if mod := len(plaintext) % aes.BlockSize; 0 != mod { // 블록 크기의 배수가 되어야함
		padding := make([]byte, aes.BlockSize-mod) // 블록 크기에서 모자라는 부분을
		plaintext = append(plaintext, padding...)  // 채워줌
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext)) // 초기화 벡터 공간(aes.BlockSize)만큼 더 생성
	iv := ciphertext[:aes.BlockSize]                         // 부분 슬라이스로 초기화 벡터 공간을 가져옴
	if _, err := io.ReadFull(rand.Reader, iv); nil != err {  // 랜덤 값을 초기화 벡터에 넣어줌
		log.Println(err)
		return ""
	}

	mode := cipher.NewCBCEncrypter(b, iv)                   // 암호화 블록과 초기화 벡터를 넣어서 암호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext) // 암호화 블록 모드 인스턴스로

	if stdFlag {
		return base64.StdEncoding.EncodeToString(ciphertext)
	} else {
		return base64.URLEncoding.EncodeToString(ciphertext)
	}
}

func EncryptToStd(key []byte, text string) string {
	return Encrypt(key, text, true)
}

func EncryptToUrl(key []byte, text string) string {
	return Encrypt(key, text, false)
}

// Decrypt from base64 to decrypted string
func Decrypt(key []byte, ciphertext []byte) string {
	defer func() {
		if err := recover(); nil != err {
			log.Println(err)
		}
	}()

	b, err := aes.NewCipher(key)
	if nil != err {
		log.Println(err)
		return ""
	}

	if 0 != len(ciphertext)%aes.BlockSize { // 블록 크기의 배수가 아니면 리턴
		log.Println("암호화된 데이터의 길이는 블록 크기의 배수가 되어야합니다.")
		return ""
	}

	iv := ciphertext[:aes.BlockSize]           // 부분 슬라이스로 초기화 벡터 공간을 가져옴
	ciphertext = ciphertext[aes.BlockSize:]    // 부분 슬라이스로 암호화된 데이터를 가져옴
	plaintext := make([]byte, len(ciphertext)) // 평문 데이터를 저장할 공간 생성
	mode := cipher.NewCBCDecrypter(b, iv)      // 암호화 블록과 초기화 벡터를 넣어서
	// 복호화 블록 모드 인스턴스 생성
	mode.CryptBlocks(plaintext, ciphertext) // 복호화 블록 모드 인스턴스로 복호화

	str1 := strings.TrimSpace(string(bytes.TrimRight(plaintext, "\x00")))
	str := ""
	for i := 0; i < len(plaintext) && i < len(str1); i++ {
		if plaintext[i] < 0x20 {
			break //	Front에서 Encoding된 값은 0x06이 붙어 있음. ???
		}
		if len(str1) == i+1 {
			str += str1[i:]
		} else {
			str += str1[i : i+1]
		}
	}

	return strings.TrimSpace(str)
	// Front에서 16의 배수일 경우, 스페이스를 1개 추가 해서 보내기 때문에 trim을 해준다.
	//	return strings.TrimSpace(string(bytes.TrimRight(plaintext, "\x00")))
}

func DecryptToStd(key []byte, cryptoText string) string {
	ciphertext, _ := base64.StdEncoding.DecodeString(cryptoText)
	return Decrypt(key, ciphertext)
}

func DecryptToUrl(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
	return Decrypt(key, ciphertext)
}

func CheckTokenAuthFile(c *gin.Context, token string, filePath string) (name string, succ bool) {
	checkUserInfo := make(map[string]interface{})
	if "" != token {
		key := []byte(get_key_string(c.ClientIP()))
		text := DecryptToStd(key, token)
		byt := []byte(text)

		if err := json.Unmarshal(byt, &checkUserInfo); nil != err {
			log.Println("Decoding Fail!!!!!!!!!!")
			FinishApiWithErrCd(c, Api_status_bad_request, Error_token_mismatch)
			return
		}
	}

	checkCaseArr := strings.Split(filePath, "/")
	switch checkCaseArr[0] {
	case "application":
		sql := `SELECT app.institution_seq, app_file.org_file_name
							  FROM t_application app, t_application_file app_file
							 WHERE app.application_seq = app_file.application_seq
							 	 AND app_file.filepath = ?
							 	 AND app.application_seq = ?`
		row := DB_fetch_one(sql, nil, filePath, checkCaseArr[1])
		if nil == row {
			FinishApiWithErrCd(c, Api_status_bad_request, Error_none_file)
			return
		} else {
			if "" != token {
				if ToInt(row["institution_seq"]) != ToInt(checkUserInfo["institution_seq"]) {
					FinishApiWithErrCd(c, Api_status_unauthorized, Error_download_auth)
					return
				}
			}
			name = ToStr(row["org_file_name"])
			succ = true
			return
		}
	case "institution":
		sql := `SELECT institution_seq, logo_file_org_name, business_file_org_name
								 FROM t_institution
								WHERE 1 = 1
								 	AND institution_seq = ?`
		row := DB_fetch_one(sql, nil, checkCaseArr[1])
		if "" != token {
			if ToInt(row["institution_seq"]) != ToInt(checkUserInfo["institution_seq"]) {
				FinishApiWithErrCd(c, Api_status_unauthorized, Error_download_auth)
				return
			}
		}

		if "logo" == checkCaseArr[2] {
			name = ToStr(row["logo_file_org_name"])
			succ = true
			return
		} else if "business" == checkCaseArr[2] {
			name = ToStr(row["business_file_org_name"])
			succ = true
			return
		} else {
			return
		}
	case "board":
		sql := `SELECT file_org_name
								FROM t_board
							 WHERE 1 = 1
								 AND file_path = ?
								 AND board_seq = ?`
		row := DB_fetch_one(sql, nil, filePath, checkCaseArr[2])
		if nil == row {
			FinishApiWithErrCd(c, Api_status_bad_request, Error_none_file)
			return
		} else {
			name = ToStr(row["file_org_name"])
			succ = true
			return
		}
	default:
		FinishApiWithErrCd(c, Api_status_not_found, Error_none_file)
		return
	}

	succ = true
	return
}

// func EncryptWithStrKey(key string, text string) string {
// 	newKey := key + "12345678901234567890123456789012"
// 	newKey = newKey[0:32]
// 	return Encrypt([]byte(newKey), text)
// }
//
// func DecryptWithStrKey(key string, text string) string {
// 	newKey := key + "12345678901234567890123456789012"
// 	newKey = newKey[0:32]
// 	return Decrypt([]byte(newKey), text)
// }
