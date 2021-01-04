package response

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/i18n"
	"github.com/small-ek/antgo/response"
)

//Fail
func Fail(c *gin.Context, msg string, err error) {
	response.Json(response.Fail(i18n.Get(msg), err), c)
}

//Success
func Success(c *gin.Context, msg string, data ...interface{}) {
	response.Json(response.Success(i18n.Get(msg), data), c)
}
