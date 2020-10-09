package main

import (
	"ekgo/api/boot/config"
	"ekgo/api/boot/db"
	"ekgo/api/boot/router"
	"ekgo/api/ek/logger"
	"flag"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
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

	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
	var server = web.NewService(
		web.Registry(etcdReg),
		web.Name("dev"),
		web.Address(":8001"),
		web.Handler(router.Load()),
	)
	server.Init()
	server.Run()
	//运行服务
	/*serve.Default(router.Load(), "95").Run()
	serve.Wait()*/
}
