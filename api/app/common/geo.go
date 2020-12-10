package common

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"ekgo/lib/response"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/ginp/orm"
)

//查询所有地理
func GeoAll(c *gin.Context) {
	var data = []model.Geo{}
	err := db.Master.Scopes(
		orm.Where("pid", "=", c.Param("pid")),
	).Order("id asc").Find(&data).Order("id desc").Error
	if err == nil {
		c.SecureJSON(200, response.Write{
			Code: 200,
			Data: data,
		})
	} else {
		c.SecureJSON(403, response.ErrorResponse(err))
	}
}
