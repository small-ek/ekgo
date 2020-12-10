package role

import (
	"ekgo/app/model"
	"ekgo/app/service"
	"ekgo/boot/db"
	"ekgo/lib/response"
	"github.com/gin-gonic/gin"
)

//接口服务
var Interface service.RoleInterface

//查询
func Index(this *gin.Context) {
	var param = response.PageParam{CurrentPage: 1, PageSize: 10}
	this.ShouldBindQuery(&param)
	param.Filter = this.QueryArray("filter[]")
	Interface = &service.Role{PageParam: param, Db: db.Master}
	this.SecureJSON(200, Interface.Index())

}

//查询
func Read(this *gin.Context) {
	var data = model.Role{}
	this.ShouldBindUri(&data)
	Interface = &service.Role{Model: data, Db: db.Master}
	this.SecureJSON(200, Interface.Show())
}

//创建
func Store(this *gin.Context) {
	var data = model.Role{}
	this.ShouldBind(&data)
	Interface = &service.Role{Model: data, Db: db.Master}
	this.SecureJSON(200, Interface.Store())
}

//修改
func Update(this *gin.Context) {
	var data = model.Role{}
	this.ShouldBind(&data)
	Interface = &service.Role{Model: data, Db: db.Master}
	this.SecureJSON(200, Interface.Update())
}

//删除
func Delete(this *gin.Context) {
	var data = model.Role{}
	this.ShouldBindUri(&data)
	Interface = &service.Role{Model: data, Db: db.Master}
	this.SecureJSON(200, Interface.Delete())
}
