package main

import (
	"ekgo/api/boot/config"
	"ekgo/api/boot/db"
	"ekgo/api/ek/logger"
	"ekgo/api/service"
	"flag"
	transportHttp "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
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
	var users = service.UserService{}
	var endp = service.GenUserEnpoint(users)
	var serviceHttp = transportHttp.NewServer(endp, service.DecodeUser, service.EncodeUser)

	http.ListenAndServe(":8080", serviceHttp)
	//运行服务
	/*serve.Default(router.Load(), "95").Run()
	serve.Wait()*/
}
