package main

import (
	"ipsap"
	"ipsap/common"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	//"ipsap/model"
)

func main() {
	//configPath := "/root/go/config/service.toml"
	configPath := "C:\\ipsap\\go\\go\\config\\win_service.toml"
	argsWithProg := os.Args
	if len(argsWithProg[1:]) != 0 {
		configPath = argsWithProg[1]
	}

	// Config 설정
	if _, err := toml.DecodeFile(configPath, &common.Config); err != nil {
		log.Println(err)
		return
	}

	//  DB 초기화
	common.DB_pool_connect()
	ipsap.Start()
}
