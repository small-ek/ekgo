package main

import (
	"ekgo/api/boot/config"
	"ekgo/api/boot/db"
	"ekgo/api/ek/logger"
	"flag"
	"github.com/micro/go-micro/v2"
	proto "github.com/micro/services/helloworld/proto"
	"log"
	"time"
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

	service := micro.NewService()
	service.Init()

	// create the proto client for helloworld
	client := proto.NewHelloworldService("go.micro.service.helloworld", service.Client())

	// call an endpoint on the service
	rsp, err := client.Call(context.Background(), &proto.Request{
		Name: "John",
	})
	if err != nil {
		log.Println("Error calling helloworld: ", err)
		return
	}

	// print the response
	log.Println("Response: ", rsp.Msg)

	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)
	//运行服务
	/*serve.Default(router.Load(), "95").Run()
	serve.Wait()*/
}
