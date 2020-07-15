package middleware

import (
	"bytes"
	"ekgo/api/boot/config"
	"ekgo/api/boot/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

func Logger() gin.HandlerFunc {
	//是否开启日志
	return func(c *gin.Context) {
		request := c.Request
		if request.Method != "OPTIONS" {

			if config.Get.Log.Close == false {

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
				//执行时间
				run_time := end.Sub(start)
				//请求URL
				path := c.Request.URL.RequestURI()
				//请求IP
				ip := c.ClientIP()
				//请求类型
				method := c.Request.Method
				//请求状态
				status := c.Writer.Status()

				logger.Log.WithFields(logrus.Fields{
					"run_time": run_time,
					"status":   status,
					"method":   method,
					"path":     path,
					"ip":       ip,
					"request":  string(body),
					"header":   request.Header,
				}).Info("请求日志")
				/*initLog.Log.Infof("| %3d | %13v | %15s | %s  %s |%s|",
					statusCode,
					latency,
					clientIP,
					method, path, requestParams,
				)*/
			}
		}

	}
}
