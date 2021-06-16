package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/os/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

//中间件记录日志模块
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.Request
		//请求状态不是OPTIONS和日志开启

		if request.Method != "OPTIONS" && config.Decode().Get("log").Get("close").Bool() == false {
			var body []byte

			if c.Request.Body != nil {
				body, _ = ioutil.ReadAll(c.Request.Body)
				// 把刚刚读出来的再写进去其他地方使用没有
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}

			// 开始时间
			start := time.Now()
			// 处理请求
			c.Next()
			// 结束时间
			end := time.Now()
			// 执行时间
			run_time := end.Sub(start)
			// 请求URL
			path := c.Request.URL.RequestURI()
			// 请求IP
			ip := c.ClientIP()
			// 请求类型
			method := c.Request.Method
			// 请求状态
			status := c.Writer.Status()

			// 写入日志
			logger.Write.Info("接口请求",
				zap.Any("ip", ip),
				zap.Any("path", path),
				zap.Any("request", body),
				zap.Any("run_time", run_time),
				zap.Any("status", status),
				zap.Any("method", method),
				zap.Any("header", request.Header),
			)
		}
	}
}
