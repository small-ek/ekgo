package main

import (
	"ekgo/api/boot/config"
	"ekgo/api/boot/db"
	"ekgo/api/boot/router"
	"ekgo/api/ek/logger"
	"flag"
	"github.com/small-ek/ginp/frame/serve"
	"log"
	"net/http"
	"time"
)

func main() {
	//设置打印日志行号和文件名
	log.SetFlags(log.Llongfile | log.LstdFlags)

	flag.Parse()

	//加载配置文件
	config.Load(*flag.String("config", "./config/config.toml", "config file"), &config.Config{})
	//加载请求日志
	logger.Load(*flag.String("log", "./log/", "log file"), "debug")

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
