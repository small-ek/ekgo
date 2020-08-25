package middleware

import (
	"bytes"
	"ekgo/api/boot/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

//中间件记录日志模块
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.Request
		//请求状态不是OPTIONS和日志开启
		if request.Method != "OPTIONS" && config.Get.Log.Close == false {
			var body []byte

			if c.Request.Body != nil {
				body, _ = ioutil.ReadAll(c.Request.Body)
				// 把刚刚读出来的再写进去其他地方使用没有
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}

			// 开始时间
			/*start := time.Now()*/
			// 处理请求
			c.Next()
			// 结束时间
			/*end := time.Now()
			// 执行时间
			run_time := end.Sub(start)
			// 请求URL
			path := c.Request.URL.RequestURI()
			// 请求IP
			ip := c.ClientIP()
			// 请求类型
			method := c.Request.Method
			// 请求状态
			status := c.Writer.Status()*/
			// 写入日志
			/*logger.Log.WithFields(logrus.Fields{
				"ip":       ip,
				"path":     path,
				"request":  string(body),
				"run_time": run_time,
				"status":   status,
				"method":   method,
				"header":   request.Header,
			}).Info()*/
			/*logger.Log.Debug("test",
				zap.String("string", "string"),
				zap.Int("int", 3),
				zap.Duration("time", time.Second),
			)*/
		}
	}
}
