package db

import (
	"ekgo/api/boot/config"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Master *gorm.DB

var Slave *gorm.DB

var err error

//获取主库数据配置
func GetMasterConfig() (connection, userName, password, host, port, database string, logmode bool) {
	connection = config.Get.Master.Connection
	userName = config.Get.Master.Username
	password = config.Get.Master.Password
	host = config.Get.Master.Host
	port = config.Get.Master.Port
	database = config.Get.Master.Database
	logmode = config.Get.Master.Logmode
	return connection, userName, password, host, port, database, logmode
}

//主库初始化
func RegisterMaster() *gorm.DB {
	dsn := "user=postgres password=pi3.1415926 dbname=47.103.147.249 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Master, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return Master
}

//获取主库数据配置
func GetSlaveConfig() (connection, userName, password, host, port, database string, logmode bool) {
	connection = config.Get.Slave.Connection
	userName = config.Get.Slave.Username
	password = config.Get.Slave.Password
	host = config.Get.Slave.Host
	port = config.Get.Slave.Port
	database = config.Get.Slave.Database
	logmode = config.Get.Slave.Logmode
	return connection, userName, password, host, port, database, logmode
}