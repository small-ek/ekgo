package main

import (
	"ekgo/api/boot/config"
	"ekgo/api/boot/db"
	"ekgo/api/boot/router"
	"ekgo/api/ek/logger"
	"flag"
	"github.com/micro/go-micro/v2/web"
	"log"
)

func main() {
	//设置打印日志行号和文件名
	log.SetFlags(log.Llongfile | log.LstdFlags)
	flag.Parse()
	//加载配置文件
	config.Load(*flag.String("config", "./config/config.toml", "config file"), &config.Config{})
	//加载请求日志
	logger.Default(*flag.String("log", "./log/ek.log", "log file")).Load()
	//加载主数据库
	db.RegisterMaster()

	var server = web.NewService(
		web.Name("123"),
		web.Address(":10082"),
		web.Handler(router.Load()),
	)
	server.Run()
	//运行服务
	/*serve.Default(router.Load(), "95").Run()
	serve.Wait()*/
}
