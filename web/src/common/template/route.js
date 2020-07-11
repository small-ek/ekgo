import {config} from "../config";

export const routeTemplate = (table_name, humpTable, table_comment) => {
    return `package router

import (
	"` + config.projectName + `/api/app/admin/` + table_name + `"
	"` + config.projectName + `/api/app/middleware"
	"github.com/gin-gonic/gin"
)

//` + table_comment + `
func ` + humpTable + `(Router *gin.RouterGroup) {

	Group := Router.Group("/admin").Use(middleware.AdminAuth()).Use(middleware.CasbinHandler())
	{
		Group.GET("/` + table_name + `", ` + table_name + `.Index)         //分页
		Group.GET("/` + table_name + `/:id", ` + table_name + `.Read)      //查询
		Group.POST("/` + table_name + `", ` + table_name + `.Store)        //添加
		Group.DELETE("/` + table_name + `/:id", ` + table_name + `.Delete) //删除
		Group.PUT("/` + table_name + `/:id", ` + table_name + `.Update)    //修改
	}

}
`
}