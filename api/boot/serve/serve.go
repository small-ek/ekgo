package serve

import (
	"ekgo/boot/db"
	"ekgo/boot/router"
	"flag"
	"github.com/small-ek/ginp/frame/serve"
	"github.com/small-ek/ginp/os/config"
	"github.com/small-ek/ginp/os/logger"
	"log"
)

func Load() {
	//设置打印日志行号和文件名
	log.SetFlags(log.Llongfile | log.LstdFlags)
	flag.Parse()
	//设置配置文件
	config.SetPath(*flag.String("config", "./config/config.toml", "config file"))
	//设置日志
	logger.Default(*flag.String("log", "./log/ek.log", "log file")).Register()
	//设置数据库
	db.Register()
	//获取配置地址
	var address = config.Decode().Get("system").Get("address").String()
	//启动服务
	serve.Default(router.Load(), address).Run()
	serve.Wait()
}
