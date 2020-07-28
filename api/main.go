package main

import (
	"ekgo/api/boot/config"
	"ekgo/api/boot/db"
	"ekgo/api/ek/logger"

	"ekgo/api/boot/router"
	"ekgo/api/ek/frame/serve"
	"flag"
	"log"
	"net/http"
	"time"
)

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

	//运行服务,也可以添加多个服务
	var service = serve.Option{Server: &http.Server{
		Addr:              config.Get.System.Address,
		ReadTimeout:       time.Duration(config.Get.System.ReadTimeout) * time.Second,       //设置秒的读超时
		WriteTimeout:      time.Duration(config.Get.System.WriteTimeout) * time.Second,      //设置秒的写超时
		ReadHeaderTimeout: time.Duration(config.Get.System.ReadHeaderTimeout) * time.Second, //读取头超时
		IdleTimeout:       time.Duration(config.Get.System.IdleTimeout) * time.Second,       //空闲超时
		MaxHeaderBytes:    config.Get.System.MaxHeaderBytes,
		Handler:           router.Load(),
	}}
	service.Run()
	service.Wait()
}
