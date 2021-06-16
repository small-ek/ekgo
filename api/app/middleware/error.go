package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
)

// Recovery 捕获异常并且写入日志
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {

			if err := recover(); err != nil {
				request := c.Request
				var body []byte

				if request.Body != nil {
					body, _ = ioutil.ReadAll(request.Body)
					// 把刚刚读出来的再写进去其他地方使用没有
					c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				}
				// 请求URL
				path := request.URL.RequestURI()
				// 请求类型
				method := request.Method
				// 请求IP
				ip := c.ClientIP()
				logger.Write.Error("错误报错",
					zap.Any("ip", ip),
					zap.Any("path", path),
					zap.Any("request", body),
					zap.Any("method", method),
					zap.Any("header", request.Header),
					zap.Any("err", err),
					zap.Any("stack", debug.Stack()),
				)
				c.AbortWithStatus(http.StatusInternalServerError)
				log.Println(err)
				debug.PrintStack()
			}
		}()

		// 继续往下处理
		c.Next()
	}
}
