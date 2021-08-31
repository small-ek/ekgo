package db

import (
	"github.com/small-ek/antgo/os/config"
	loggers "github.com/small-ek/antgo/os/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var Master *gorm.DB

var err error

//GetMasterConfig 获取数据库配置
func GetDbConfig() (connection, userName, password, host, port, database, conf string, log bool) {
	var db = config.Decode().Get("database")
	connection = db.Get("connection").String()
	userName = db.Get("username").String()
	password = db.Get("password").String()
	host = db.Get("host").String()
	port = db.Get("port").String()
	database = db.Get("database").String()
	conf = db.Get("config").String()
	log = db.Get("log").Bool()
	return connection, userName, password, host, port, database, conf, log
}

//Mysql 初始化Mysql数据库
func Mysql() *gorm.DB {
	var _, userName, password, host, port, database, conf, log = GetDbConfig()

	var dns = userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?" + conf
	var db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), getConfig(log))
	if err != nil {
		loggers.Error(err)
		os.Exit(0)
	}
	return db
}

//getConfig
func getConfig(log bool) *gorm.Config {
	if log {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
}

//Register 数据库连接注册
func Register() *gorm.DB {
	switch config.Decode().Get("database").Get("connection").String() {
	case "mysql":
		Master = Mysql()
	default:
		Master = Mysql()
	}
	return Master
}

//Close 关闭数据库
func Close() {
	var db, err = Master.DB()
	loggers.Error(err)
	db.Close()
}
