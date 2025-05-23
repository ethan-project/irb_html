package main

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	resty "github.com/go-resty/resty/v2"
  "ipsap/common"
	"os/exec"
	"time"
	"log"
	"os"
)

func setLog(fileName string) *rotatelogs.RotateLogs {
	basePath := "/root/api/status_log/"
	_ = os.Mkdir(basePath, os.ModeDir)

	yyyymmdd := "%Y%m%d"

	rl, _ := rotatelogs.New(basePath+fileName+"_"+yyyymmdd+".log",
		rotatelogs.WithLinkName(basePath+fileName),
		rotatelogs.WithMaxAge(time.Duration(365)*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	return rl
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
  log.SetOutput(setLog("log_file"))
	log.Println("START")
	for {
		client := resty.New()
		resp, err := client.R().
						SetQueryParams(map[string]string{
								"filter.board_type": "1",
								"filter.institution_seq": "0",
						}).
						SetHeader("Accept", "application/json").
						Get("https://www.ipsap.co.kr:7375/api/v1.0/board")
		if err != nil || common.ToStr(resp) == "" {
			log.Println(resp)
			log.Println(err)
			cmd := exec.Command("systemctl", "restart", "ipsap")
			if err := cmd.Run(); err != nil {
				log.Println(err)
			}
		}

		time.Sleep(time.Second * 60)   //  1분단위로 작업
	}
}
