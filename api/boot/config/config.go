package config

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

var Get *ini.File
var CsrfExcept *ini.File

func Load(configFile string) {
	var err error
	Get, err = ini.Load(configFile)
	if err != nil {
		log.Println("无法读取配置文件路径" + err.Error())
		os.Exit(1)
	}
}
