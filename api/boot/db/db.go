package db

import (
	"github.com/small-ek/ginp/os/config"
	"github.com/small-ek/ginp/os/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Master *gorm.DB

var err error

//GetMasterConfig 获取数据库配置
func GetDbConfig() (connection, userName, password, host, port, database, conf string) {
	var db = config.Decode().Get("database")
	connection = db.Get("connection").String()
	userName = db.Get("username").String()
	password = db.Get("password").String()
	host = db.Get("host").String()
	port = db.Get("connection").String()
	database = db.Get("database").String()
	conf = db.Get("config").String()
	return connection, userName, password, host, port, database, conf
}

//Mysql 初始化Mysql数据库
func Mysql() *gorm.DB {
	var _, userName, password, host, port, database, conf = GetDbConfig()

	var dns = userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?" + conf

	var db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
	}
	return db
}

//Register 数据库连接注册
func Register() {
	var connection = config.Decode().Get("database").Get("connection").String()
	switch connection {
	case "mysql":
		Master = Mysql()
	default:
		Master = Mysql()
	}
	defer Close(Master)
}

//Close 关闭数据库
func Close(Master *gorm.DB) {
	var db, err = Master.DB()
	if err != nil {
		logger.Error(err.Error())
	}
	db.Close()
}
