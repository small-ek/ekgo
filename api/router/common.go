package router

import (
	"ekgo/app/admin/admins"
	"github.com/gin-gonic/gin"
)

//公共组件
func Common(Router *gin.RouterGroup) {
	Router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<h1 style="text-align: center;margin: 200px auto;color: #343434;text-shadow: -3px -3px #ccc;font-size: 60px;">Hello Ekgo</h1>`)
	})

	//登陆不需要权限
	//Group := Router.Group("/admin").Use(middleware.AdminAuth())
	//{
	//	Group.GET("/left_menu", menu.LeftMenu) //左侧权限菜单
	//}
	Router.POST("/admin/login", admins.Login) //登录接口
	//Router.GET("/upload/oss", common.Oss)    //阿里云上传凭证
	//Router.GET("/geo/:pid", common.GeoAll)   //中国省市区地理信息
}
