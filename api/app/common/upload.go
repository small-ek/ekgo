package common

import (
	"ekgo/app/service"
	"github.com/gin-gonic/gin"
)

//阿里云oss
func Oss(c *gin.Context) {
	var service service.UpLoadOSS
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err.Error())
	} else {
		result := service.GetToken()
		c.SecureJSON(200, result)
	}
}
