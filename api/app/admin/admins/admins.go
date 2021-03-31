package admins

import (
	"ekgo/app/model"
	"ekgo/app/service"
	"ekgo/app/validate"
	"ekgo/lib/json"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/request"
)

//Index 查询分页
func Index(c *gin.Context) {
	var page = request.DefaultPage()
	c.ShouldBindQuery(&page)
	var service = service.Admin{PageParam: page}
	var result, err = service.Index()

	if err != nil {
		json.Fail(c, "failed", err)
	} else {
		json.Success(c, "success", result)
	}
}

//Show 显示对于id的内容
func Show(c *gin.Context) {
	var data = model.Admin{}
	c.ShouldBindUri(&data)
	var service = &service.Admin{Model: data}
	var result, err = service.Show()

	if err != nil {
		json.Fail(c, "failed", err)
	} else {
		json.Success(c, "success", result)
	}
}

//Store 创建
func Store(c *gin.Context) {

	var data = model.Admin{}
	c.ShouldBind(&data)

	var validate = validate.CheckAdminRegister(data)
	if validate != nil {
		json.Fail(c, validate.Error())
		return
	}

	var service = &service.Admin{Model: data}
	var result, err = service.Store()

	if err != nil {
		json.Fail(c, "failed", err)
	} else {
		json.Success(c, "success", result)
	}
}

//Update 修改
func Update(c *gin.Context) {
	var data = model.Admin{}
	c.ShouldBind(&data)
	var service = &service.Admin{Model: data}
	var err = service.Update()

	if err != nil {
		json.Fail(c, "failed", err)
	} else {
		json.Success(c, "success")
	}
}

//Delete 删除
func Delete(c *gin.Context) {
	var data = model.Admin{}
	c.ShouldBindUri(&data)
	var service = &service.Admin{Model: data}
	var err = service.Delete()

	if err != nil {
		json.Fail(c, "failed", err)
	} else {
		json.Success(c, "success")
	}
}

//Login 登录
func Login(c *gin.Context) {
	var data = model.Admin{}
	c.ShouldBind(&data)
	var service = &service.Admin{Model: data}
	var result, err = service.Login()

	if err != nil {
		json.Fail(c, "login_failure", err)
	} else if result == nil {
		json.Fail(c, "login_incorrect")
	} else {
		json.Success(c, "login_success", result)
	}
}
