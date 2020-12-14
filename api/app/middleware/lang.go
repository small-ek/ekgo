package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/small-ek/ginp/i18n"
	"strings"
)

//Lang 设置语言
func Lang() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request = c.Request
		var language = request.Header["Accept-Language"][0]
		var firstLanguage = strings.Split(language, ",")
		i18n.SetLanguage(firstLanguage[0])
	}
}
