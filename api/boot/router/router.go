package router

import (
	"ekgo/app/middleware"
	"ekgo/router"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/request"
)

//Load 初始化路由
func Load() *gin.Engine {
	var app = gin.Default()
	//https配置
	//Router.Use(middleware.LoadTls())
	//请求记录日志
	go func() {
		app.Use(middleware.Logger())
	}()
	//捕获异常
	app.Use(middleware.Recovery())

	var system = config.Decode().Get("system")
	//跨域
	if system.Get("cors").Bool() == true {
		app.Use(request.Cors)
	}
	//设置语言
	app.Use(middleware.Lang())
	//钩子
	//Router.Use(middleware.Hook)
	//swagger swag init
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//静态文件
	app.Static("view", "./resources/view")
	//添加路由组前缀
	Group := app.Group("")
	//注册公共组件路由
	router.Common(Group)
	//注册后台路由
	router.Admin(Group)

	return app
}
