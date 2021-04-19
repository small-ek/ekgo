package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/i18n"
	"strings"
)

//Lang 设置语言
func Lang() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request = c.Request
		var language = request.Header["Accept-Language"][0]
		var firstLanguage = strings.Split(language, ",")
		if len(firstLanguage) > 0 && firstLanguage[0] == "zh-Hans-CN" {
			firstLanguage[0] = "zh-CN"
		}
		i18n.SetLanguage(firstLanguage[0])
	}
}
