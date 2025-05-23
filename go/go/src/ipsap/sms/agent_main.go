package sms

import (
  "ipsap/common"
  "ipsap/model"
  "fmt"
  "log"
  // "io/ioutil"
)

func main() {
	/*
	loc, _ := time.LoadLocation("Asia/Seoul")
// handle err
	time.Local = loc // -> this is setting the global timezone

	//os.Setenv("TZ", "Asia/Seoul")

	{	//	wowdolf time convert test code!!!!
		value := "2021-02-28T01:01"

		t1, e := time.Parse(
	      		"2006-01-02T03:04", //time.RFC3339,  //  20060102150405
	      		value)      //  2013-09-24 오후 5:00:00
	  if nil != e {
	    log.Println(e)
	  }

//		t2 := t1.Unix() - (9 * 60 * 60);

	  log.Println(value, t1, t2)
	}
*/

	log.Println("======= test =====!!!")
	configPath := "/root/go/config/service.toml"
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

	if (false)	{
	  ins := model.CodeDynamic{
	            Institution_seq : 1,
	            DCode_type      : 1,
	            Value           : `토끼`,
	  }
	  log.Println(ins.Value, ins.GetCodeFromValue())

	  ins.Value = `돼지`
	  log.Println(ins.Value, ins.GetCodeFromValue())
		log.Println(ins)
	}

	if (true)	{
		item := model.Item{}
		if item.GetItem("anml_tm_restraint")	{
			log.Println("GetItem success! (anml_tm_restraint)")
		} else {
			log.Println("GetItem Fail!! (anml_tm_restraint)")
		}
	}

	log.Println("configPath : " + configPath)
	ipsap.Start()

}
