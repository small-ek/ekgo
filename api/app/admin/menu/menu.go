package menu

import (
	"ekgo/api/app/middleware"
	"ekgo/api/app/model"
	"ekgo/api/app/service"
	"ekgo/api/boot/db"
	"github.com/gin-gonic/gin"
)

//接口服务
var Interface service.MenuInterface

//查询菜单
func Index(this *gin.Context) {
	Interface = &service.Menu{Db: db.Master}
	this.SecureJSON(200, Interface.Index())
}

//创建
func Store(this *gin.Context) {
	var data = model.Menu{}
	this.ShouldBind(&data)
	Interface = &service.Menu{Model: data, Db: db.Master}
	this.SecureJSON(200, Interface.Store())
}

//修改
func Update(this *gin.Context) {
	var data = model.Menu{}
	this.ShouldBind(&data)
	Interface = &service.Menu{Model: data, Db: db.Master}
	this.SecureJSON(200, Interface.Update())
}

//删除
func Delete(this *gin.Context) {
	var data = model.Menu{}
	this.ShouldBindUri(&data)
	Interface = &service.Menu{Model: data, Db: db.Master}
	this.SecureJSON(200, Interface.Delete())
}

//查询菜单
func LeftMenu(this *gin.Context) {
	user := middleware.GetAdmin(this)
	Interface = &service.Menu{Db: db.Master}
	result := Interface.LeftMenu(&user)
	this.SecureJSON(200, result)
}
