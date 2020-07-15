package db

import (
	"ekgo/api/boot/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
	"time"
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
	connection, userName, password, host, port, database, logmode := GetMasterConfig()

	switch connection {
	case "mysql":
		Master, err = gorm.Open(connection, ""+userName+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local&timeout=10s")
	case "postgres":
		Master, err = gorm.Open(connection, "host="+host+" port="+port+" user="+userName+" dbname="+database+" sslmode=disable"+" password="+password)
	}
	Master.LogMode(logmode)
	// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	Master.SingularTable(true)
	//连接池闲置连接数
	Master.DB().SetMaxIdleConns(10)
	//最大打开连接数0不限制
	Master.DB().SetMaxOpenConns(10)
	//超时
	Master.DB().SetConnMaxLifetime(time.Second * 5)
	if err != nil {
		log.Println("数据库启动异常,请检查连接状态:" + err.Error())
		os.Exit(1)
	}

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

//从库初始化
func RegisterSlave() *gorm.DB {
	connection, userName, password, host, port, database, logmode := GetSlaveConfig()

	switch connection {
	case "mysql":
		Slave, err = gorm.Open(connection, ""+userName+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local&timeout=10s")
	case "postgres":
		Slave, err = gorm.Open(connection, "host="+host+" port="+port+" user="+userName+" dbname="+database+" sslmode=disable"+" password="+password)
	}
	Slave.LogMode(logmode)
	// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	Slave.SingularTable(true)
	//连接池闲置连接数
	Slave.DB().SetMaxIdleConns(20)
	//最大打开连接数0不限制
	Slave.DB().SetMaxOpenConns(100)

	if err != nil {
		log.Println("数据库启动异常,请检查连接状态:" + err.Error())
		os.Exit(1)
	}

	return Slave
}
