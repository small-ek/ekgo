package admin

import (
	"ekgo/api/app/middleware"
	"ekgo/api/app/model"
	"ekgo/api/app/service"
	"ekgo/api/app/validate"
	"ekgo/api/boot/db"
	"ekgo/api/lib/request"
	"ekgo/api/lib/response"
	"github.com/gin-gonic/gin"
)

//接口服务
var Interface service.AdminInterface

//登陆
func Login(this *gin.Context) {
	var data = model.Admin{}
	this.ShouldBind(&data)

	Interface = &service.Admin{Model: data, Db: db.Master}
	this.SecureJSON(200, Interface.Login())
}

//查询分页
func Index(this *gin.Context) {
	var param = response.PageParam{CurrentPage: 1, PageSize: 10}
	this.ShouldBindQuery(&param)
	param.Filter = this.QueryArray("filter[]")

	Interface = &service.Admin{
		PageParam: param,
		Db:        db.Master,
	}
	this.SecureJSON(200, Interface.Index())
}

//插入
func Store(this *gin.Context) {
	var data = model.Admin{}
	this.ShouldBind(&data)

	Interface = &service.Admin{
		Model: data,
		Db:    db.Master,
	}
	this.SecureJSON(200, Interface.Store())
}

//修改
func Update(this *gin.Context) {
	var data = model.Admin{}
	this.ShouldBind(&data)

	Interface = &service.Admin{
		Model: data,
		Db:    db.Master,
	}
	this.SecureJSON(200, Interface.Update())
}

//删除
func Delete(this *gin.Context) {
	var data = model.Admin{}
	this.ShouldBindUri(&data)

	Interface = &service.Admin{
		Model: data,
		Db:    db.Master,
	}
	this.SecureJSON(200, Interface.Delete())
}

//修改当前用户密码
func UpdatePassword(this *gin.Context) {
	var Request = request.Get(this)
	var getAdmin = middleware.GetAdmin(this)
	var data = model.Admin{}
	this.ShouldBind(&data)
	data.Id = getAdmin.Id
	Request["id"] = getAdmin.Id
	var validate = &validate.Admin{Request: Request}

	if err := validate.CheckPassword(); err != nil {
		this.SecureJSON(200, err)
	} else {
		Interface = &service.Admin{Model: data, Db: db.Master}
		this.SecureJSON(200, Interface.Update())
	}
}

//检查是否登录
func CheckLogin(c *gin.Context) {
	var Model = model.Admin{Id: middleware.GetAdmin(c).Id}
	Interface = &service.Admin{Model: Model, Db: db.Master}

	if Interface.Show().Code == 200 {
		c.SecureJSON(200, &response.Write{
			Code: 200,
			Msg:  "登录成功",
		})
	} else {
		c.SecureJSON(200, &response.Write{
			Code: 403,
			Msg:  "登录失败",
		})
	}
}
