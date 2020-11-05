package router

import (
	"ekgo/api/app/admin/admin"
	"ekgo/api/app/admin/menu"
	"ekgo/api/app/common"
	"ekgo/api/app/middleware"
	"ekgo/api/ek/logger"
	"github.com/gin-gonic/gin"
)

//公共组件
func Common(Router *gin.RouterGroup) {
	Router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<h1 style="text-align: center;margin: 200px auto;color: #343434;text-shadow: -3px -3px #ccc;font-size: 60px;">Hello Ekgo</h1>`)
	})
	Router.GET("/test", func(c *gin.Context) {
		for i:=0;i<10000;i++{
			logger.AsyncInfo("121212",[]string{"test","hello","world"})
		}
		c.String(200,"12121212");
	})
	//登陆不需要权限
	Group := Router.Group("/admin").Use(middleware.AdminAuth())
	{
		Group.GET("/left_menu", menu.LeftMenu) //左侧权限菜单
	}
	Router.POST("/admin/login", admin.Login) //登录接口
	Router.GET("/upload/oss", common.Oss)    //阿里云上传凭证
	Router.GET("/geo/:pid", common.GeoAll)   //中国省市区地理信息
}
