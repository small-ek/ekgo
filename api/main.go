package main

import (
	"ekgo/api/boot/config"
	"ekgo/api/boot/db"
	"ekgo/api/boot/logger"
	"ekgo/api/boot/serve"
	"flag"
	"log"
)

// @title 接口文档
// @version 0.0.1
// @description 请求接口
// @securityDefinitions.apikey 权限
// @in header
// @name Authorization
// @BasePath /
func main() {
	//设置打印日志行号和文件名
	log.SetFlags(log.Llongfile | log.LstdFlags)

	//设置配置文件打包路径,不设置默认
	conf := flag.String("config", "./config/config.toml", "config file")

	//设置日志文件打包路径,不设置默认
	log_file := flag.String("log", "./log/", "log file")
	flag.Parse()

	//加载配置文件
	config.Load(*conf)

	//加载请求日志
	logger.Load(*log_file)

	//加载主数据库
	db.RegisterMaster()

	//运行服务,可以添加多个服务
	serve.RunServe()
}
