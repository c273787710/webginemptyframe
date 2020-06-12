package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

//#####服务器配置信息
type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

//##### app相关配置
type App struct {
	RuntimeRootPath string
	LoginFailureLock int
	LoginFailureTime int
}
var AppSetting = &App{}

// ##### jwt相关配置
type JWT struct {
	ExpireTime time.Duration
	Issuer string
	Secret string
}
var JWTSetting = &JWT{}

//#####日志系统配置
type Log struct {
	Level string
	FileName string
	MaxAge int
	MaxSize int
	MaxBackups int
}
var LogSetting = &Log{}

//#####数据库相关配置
type Mysql struct {
	Host string
	Port int
	Username string
	Password string
	Databases string
	Prefix string
	Charset string
	MaxOpenConn int
	MaxIdleConn int
}
var MysqlSetting = &Mysql{}

//#####配置文件对象
var cfg *ini.File

func InitConfig(){
	var err error
	cfg,err = ini.Load("config/app.ini")
	if err != nil {
		panic(fmt.Sprintf("Config Error : %s",err))
	}
	getSection("server",ServerSetting)
	getSection("app",AppSetting)
	getSection("log",LogSetting)
	getSection("mysql",MysqlSetting)
	getSection("jwt",JWTSetting)

	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	LogSetting.FileName = AppSetting.RuntimeRootPath + LogSetting.FileName
	JWTSetting.ExpireTime = JWTSetting.ExpireTime * time.Hour
}

func getSection(section string,v interface{}){
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		panic(fmt.Sprintf("Error Conf Get Section : %s",err))
	}
}
