package common

// config > 예하 .toml 파일과 항상 일치해야 한다.

var Config *configInfo

type configInfo struct {
	ServerType string `toml:"server-type"`
	Server     serverInfo
	Database   DatabaseInfo
	Program    programInfo
	Test       testInfo
	Sms    	 	 smsInfo
	Email    	 emailInfo
	S3    		 s3Info
	Payment    paymentInfo
}

type smsInfo struct {
	Usercode	string `toml:"usercode"`
	Deptcode	string `toml:"deptcode"`
	From   		string `toml:"from"`
	SendUrl   string `toml:"send-url"`
	ResultUrl string `toml:"result-url"`
}

type serverInfo struct {
	Port         			int
	SwagHostname 			string `toml:"swag-hostname"`
	LogPath      			string `toml:"log-path"`
	MaxBackups   			int    `toml:"max-backups"`
	WorkingPath  			string `toml:"working-path"`
	FileUploadPath		string `toml:"file-upload-path"`
	Protocol				  string `toml:"protocol"`
	ApiPath				  	string `toml:"api-path"`
	Cert				  		string `toml:"cert"`
	Key				  			string `toml:"key"`
}

type emailInfo struct {
	TemplatePath	string `toml:"template-path"`
	LogoFile			string `toml:"logo-file"`
}

type DatabaseInfo struct {
	DriverName  string `toml:"driver-name"`
	Url         string
	Database    string
	Username    string
	Password    string
	InitVersion string `toml:"init-version"`
}

type programInfo struct {
	InitPassword string `toml:"init-password"`
	EncryptionKey string `toml:"encryption-key"`
}

type s3Info struct {
	BUCKET string
	BUCKET_REAL string `toml:"bucket_real"`
	BUCKET_TEST string `toml:"bucket_test"`
}

type testInfo struct {
	IsTestMode   bool `toml:"is-test-mode"`
	IsQueryDebug bool `toml:"is-query-debug"`
}

type paymentInfo struct {
	GENERAL_MID	string `toml:"general-mid"`
	GENERAL_MERCHANTKEY	string `toml:"general-merchant-key"`
	REGULAR_MID string `toml:"regular-mid"`
	REGULAR_MERCHANTKEY	string `toml:"regular-merchant-key"`
}
