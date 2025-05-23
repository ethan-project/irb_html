package common

import (
	"github.com/mitchellh/mapstructure"
  "gopkg.in/gomail.v2"
  "crypto/tls"
  "log"
)

type Email struct {
  Dial 		*gomail.Dialer
  Msg  		*gomail.Message
	From 		string
	To	 		string
	Subject string
}

type Smtp struct {
  Host string
  Port string
  Id   string
  Pw   string
}

func (email *Email) Connect() (succ bool) {
  smtp := Smtp{}
  sql := `SELECT host, port, id, pw FROM t_smtp`
  row := DB_fetch_one(sql, nil)
  if err := mapstructure.Decode(row, &smtp); nil != err {
    log.Println(err)
    return
  }

  d := gomail.NewDialer(smtp.Host, ToInt(smtp.Port), smtp.Id, smtp.Pw)
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
  s, err := d.Dial()
  if nil != err {
  	log.Println(err)
  	return
  }

  email.Dial = d
  email.Msg  = gomail.NewMessage()
	email.From = smtp.Id
	defer s.Close()
	succ = email.SetHeader()
	return
}

func (email *Email) SetHeader() (succ bool) {
	if "" == email.From || false == CheckEmail(email.To) || "" == email.Subject {	return }
	email.Msg.SetHeaders(map[string][]string{
		"From":    {email.From},
		"To":      {email.To},
		"Subject": {email.Subject},
	})
	return true
}
