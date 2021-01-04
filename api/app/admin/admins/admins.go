package admins

import (
	"ekgo/app/model"
	"ekgo/app/service"
	"ekgo/lib/response"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/request"
	"log"
)

//Index 查询分页
func Index(c *gin.Context) {
	var param = request.DefaultPage()
	c.ShouldBindQuery(&param)
	param.Filter = c.QueryArray("filters[]")
	log.Println(param.Filter)

	var service = service.Admin{PageParam: param}
	var result, err = service.Index()
	if err != nil {
		response.Fail(c, "failed", err)
	} else {
		response.Success(c, "success", result)
	}
}

//Store 创建
func Store(c *gin.Context) {
	var data = model.Admin{}
	c.ShouldBind(&data)
	var service = &service.Admin{Model: data}
	var result, err = service.Store()
	if err != nil {
		response.Fail(c, "failed", err)
	} else {
		response.Success(c, "success", result)
	}
}

//Update 修改
func Update(c *gin.Context) {
	var data = model.Admin{}
	c.ShouldBind(&data)
	var service = &service.Admin{Model: data}
	var err = service.Update()
	if err != nil {
		response.Fail(c, "failed", err)
	} else {
		response.Success(c, "success")
	}
}

//Delete 删除
func Delete(c *gin.Context) {
	var data = model.Admin{}
	c.ShouldBindUri(&data)
	var service = &service.Admin{Model: data}
	var err = service.Delete()
	if err != nil {
		response.Fail(c, "failed", err)
	} else {
		response.Success(c, "success")
	}
}
