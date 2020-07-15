package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// 用https把这个中间件在router里面use一下就好
func LoadTls() gin.HandlerFunc {
	return func(this *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(this.Writer, this.Request)
		if err != nil {
			//如果出现错误，请不要继续。
			fmt.Println(err)
			return
		}
		// 继续往下处理
		this.Next()
	}
}
