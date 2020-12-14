package router

import (
	"ekgo/app/admin/admins"
	"github.com/gin-gonic/gin"
)

//后台管理
func Admin(Router *gin.RouterGroup) {

	Group := Router.Group("/admin")
	{
		//管理员
		Group.GET("/admins", admins.Index) //分页
		/*Group.POST("/admins", admin.Store)                       //添加管理员
		Group.DELETE("/admins/:id", admin.Delete)                //删除管理员
		Group.PUT("/admins/:id", admin.Update)                   //修改管理员
		Group.PUT("/updateAdminPassword", admin.UpdatePassword) //修改管理员密码
		Group.GET("/checkLogin", admin.CheckLogin)              //检测登录*/
	}

}
