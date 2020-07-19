package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type System struct {
	Address           string
	RbacPath          string
	Cors              bool
	ReadTimeout       int64
	WriteTimeout      int64
	ReadHeaderTimeout int64
	IdleTimeout       int64
	MaxHeaderBytes    int
}
type Log struct {
	Close   bool
	Maximum int
	Split   int
}
type Master struct {
	Connection string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
	Logmode    bool
}
type Slave struct {
	Connection string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
	Logmode    bool
}
type Oss struct {
	KeyId     string
	KeySecret string
	Endpoint  string
	Bucket    string
}
type Redis struct {
	Address  string
	Password string
	Db       int
}
type Config struct {
	System System
	Log    Log
	Master Master
	Slave  Slave
	Oss    Oss
	Redis  Redis
}

var Get *Config

func Load(path string) {
	Get = &Config{}
	if _, err := toml.DecodeFile(path, Get); err != nil {
		log.Println("类型不正确或者配置文件路径不正确:" + err.Error())
		os.Exit(1)
	}
}
