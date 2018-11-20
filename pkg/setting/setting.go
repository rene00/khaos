package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret string
	PrefixUrl string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type         string
	FilePath     string
	Name         string
	TablePrefix  string
	TestUsername string
	TestPassword string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Failed to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo Setting err: %v", err)
	}
}
