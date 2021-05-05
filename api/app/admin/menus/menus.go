package menus

import (
	"ekgo/app/middleware"
	"ekgo/app/model"
	"ekgo/app/service"
	"ekgo/lib/json"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/request"
)

//Index 查询分页
func Index(c *gin.Context) {
	var page = request.DefaultPage()
	c.ShouldBindQuery(&page)
	var service = service.Menu{PageParam: page}
	var result, err = service.Index()

	if err != nil {
		json.Fail(c, "failed", err)
	} else {
		json.Success(c, "success", result)
	}
}

//Show 显示对于id的内容
func Show(c *gin.Context) {
	var data = model.Menu{}
	c.ShouldBindUri(&data)
	var service = &service.Menu{Model: data}
	var result, err = service.Show()

	if err != nil {
		json.Fail(c, "failed", err)
	} else {
		json.Success(c, "success", result)
	}
}

//Store 创建
func Store(c *gin.Context) {
	var data = model.Menu{}
	c.ShouldBind(&data)

	var service = &service.Menu{Model: data}
	var result, err = service.Store()

	if err != nil {
		json.Fail(c, "failed", err)
	} else {
		json.Success(c, "success", result)
	}
}

//Update 修改
func Update(c *gin.Context) {
	var data = model.Menu{}
	c.ShouldBind(&data)
	var service = &service.Menu{Model: data}
	var err = service.Update()

	if err != nil {
		json.Fail(c, "failed", err)
	} else {
		json.Success(c, "success")
	}
}

//Delete 删除
func Delete(c *gin.Context) {
	var data = model.Menu{}
	c.ShouldBindUri(&data)
	var service = &service.Menu{Model: data}
	var err = service.Delete()

	if err != nil {
		json.Fail(c, "failed", err)
	} else {
		json.Success(c, "success")
	}
}

//LeftMenu 获取用户角色的菜单
func LeftMenu(c *gin.Context) {
	user := middleware.GetAdmin(c)
	service := &service.Menu{}
	result := service.LeftMenu(&user)
	json.Success(c, "success", result)
}
