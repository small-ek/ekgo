package router

import (
	"ekgo/api/app/middleware"
	"github.com/small-ek/ginp/os/config"

	"ekgo/api/router"
	"github.com/gin-gonic/gin"
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
	//跨域
	if config.Decode().Get("system").Get("cors").Bool() == true {
		Router.Use(middleware.Cors)
	}
	//钩子
	/*Router.Use(middleware.Hook)*/
	//swagger swag init
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//静态文件
	Router.Static("view", "./resources/view")
	//添加路由组前缀
	Group := Router.Group("")
	//注册公共组件路由
	router.Common(Group)
	//注册后台路由
	//router.Admin(Group)

	return Router
}
