package router

import (
	"ekgo/app/middleware"
	"ekgo/router"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/ginp/os/config"
	"github.com/small-ek/ginp/request"
)

//初始化路由
func Load() *gin.Engine {
	var Router = gin.Default()
	//https配置
	//Router.Use(middleware.LoadTls())
	//请求记录日志
	go func() {
		Router.Use(middleware.Logger())
	}()
	var system = config.Decode().Get("system")
	//跨域
	if system.Get("cors").Bool() == true {
		Router.Use(request.Cors)
	}
	//设置语言
	Router.Use(middleware.Lang())
	//钩子
	//Router.Use(middleware.Hook)
	//swagger swag init
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//静态文件
	Router.Static("view", "./resources/view")
	//添加路由组前缀
	Group := Router.Group("")
	//注册公共组件路由
	router.Common(Group)
	//注册后台路由
	router.Admin(Group)

	return Router
}
