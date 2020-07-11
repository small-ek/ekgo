package router

import (
	"ekgo/api/app/admin/admin"
	"ekgo/api/app/admin/api"
	"ekgo/api/app/admin/content"
	"ekgo/api/app/admin/menu"
	"ekgo/api/app/admin/role"
	"ekgo/api/app/admin/user"
	"ekgo/api/app/common"
	"ekgo/api/app/middleware"
	"github.com/gin-gonic/gin"
)

//后台管理
func Admin(Router *gin.RouterGroup) {

	Group := Router.Group("/admin").Use(middleware.AdminAuth()).Use(middleware.CasbinHandler())
	{
		//数据库表操作
		Group.GET("/get_all_table", common.GetAllTable) //获取所有表列表
		Group.GET("/get_table/:name", common.GetTable)  //获取单表信息

		//菜单
		Group.POST("/menu", menu.Store)        //创建
		Group.DELETE("/menu/:id", menu.Delete) //删除
		Group.PUT("/menu/:id", menu.Update)    //修改
		Router.GET("/admin/menu", menu.Index)  //分页

		//管理员
		Group.GET("/admin", admin.Index)                    //分页
		Group.POST("/admin", admin.Store)                   //添加管理员
		Group.DELETE("/admin/:id", admin.Delete)            //删除管理员
		Group.PUT("/admin/:id", admin.Update)               //修改管理员
		Group.PUT("/update_password", admin.UpdatePassword) //修改管理员密码
		Group.GET("/check_login", admin.CheckLogin)         //检测登录
		//api
		Group.GET("/api", api.Index)         //分页
		Group.POST("/api", api.Store)        //添加
		Group.DELETE("/api/:id", api.Delete) //删除
		Group.PUT("/api/:id", api.Update)    //修改

		//角色权限
		Group.GET("/role", role.Index)         //分页
		Group.GET("/role/:id", role.Read)      //查询
		Group.POST("/role", role.Store)        //添加
		Group.DELETE("/role/:id", role.Delete) //删除
		Group.PUT("/role/:id", role.Update)    //修改

		//用户
		Group.GET("/user", user.Index)         //分页
		Group.GET("/user/:id", user.Read)      //查询
		Group.POST("/user", user.Store)        //添加
		Group.DELETE("/user/:id", user.Delete) //删除
		Group.PUT("/user/:id", user.Update)    //修改

		Group.GET("/content", content.Index)         //分页
		Group.GET("/content/:id", content.Read)      //查询
		Group.POST("/content", content.Store)        //添加
		Group.DELETE("/content/:id", content.Delete) //删除
		Group.PUT("/content/:id", content.Update)    //修改
	}

}
