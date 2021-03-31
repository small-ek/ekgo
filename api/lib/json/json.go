package json

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/i18n"
	"github.com/small-ek/antgo/response"
)

//Fail
func Fail(c *gin.Context, msg string, err ...error) {
	var msgs = i18n.Get(msg)
	if msgs == "" {
		msgs = msg
	}

	if len(err) > 0 {
		response.Json(response.Fail(msgs, err[0]), c)
	} else {
		response.Json(response.Fail(msgs), c)
	}
}

//Success
func Success(c *gin.Context, msg string, data ...interface{}) {
	var msgs = i18n.Get(msg)
	if msgs == "" {
		msgs = msg
	}

	if len(data) > 0 {
		response.Json(response.Success(msgs, data[0]), c)
	} else {
		response.Json(response.Success(msgs), c)
	}
}
