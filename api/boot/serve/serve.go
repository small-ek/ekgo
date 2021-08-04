package serve

import (
	"ekgo/boot/db"
	"ekgo/boot/router"
	"flag"
	"github.com/small-ek/antgo/frame/serve"
	"github.com/small-ek/antgo/i18n"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/os/logger"
	"log"
)

func init() {
	//设置打印日志行号和文件名
	log.SetFlags(log.Llongfile | log.LstdFlags)
	flag.Parse()
	//设置配置文件
	config.SetPath(*flag.String("config", "./config/config.toml", "config file"))
	//加载语言包
	i18n.SetPath(*flag.String("lang", "./config/i18n.json", "lang file"))
	//设置日志
	logger.Default(*flag.String("log", "./log/ek.log", "log file")).Register()
	//设置数据库
	db.Register()
	defer db.Close()
}

//Load
func Load() {
	//获取配置地址
	var address = config.Decode().Get("system").Get("address").String()
	//启动服务
	serve.Default(router.Load(), address).Run()
	serve.Wait()
}
