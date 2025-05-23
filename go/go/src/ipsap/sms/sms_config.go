package sms

var Config *configInfo

type configInfo struct {
	ServerType string `toml:"server-type"`
	Sms    	 	 smsInfo
	Database   DatabaseInfo
}

type smsInfo struct {
	Usercode	string `toml:"usercode"`
	Deptcode	string `toml:"deptcode"`
	From   		string `toml:"from"`
	SendUrl   string `toml:"send-url"`
	ResultUrl string `toml:"result-url"`
}

type DatabaseInfo struct {
	DriverName  string `toml:"driver-name"`
	Url         string
	Database    string
	Username    string
	Password    string
	InitVersion string `toml:"init-version"`
}
