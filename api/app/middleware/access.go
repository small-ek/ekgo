package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//跨域请求
func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")                                      //允许访问所有域
	c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE,OPTIONS") //允许请求类型
	c.Header("Access-Control-Allow-Credentials", "true")                              //服务器是否接受浏览器发送的Cookie
	c.Header("Connection", "keep-alive")                                              //可以使一次TCP连接为同意用户的多次请求服务,提高了响应速度。
	c.Header("Cache-Control", "no-store, no-cache, must-revalidate")                  //缓存请求指令
	c.Header("Pragma", "no-cache")                                                    //缓存请求指令
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, If-Match, If-Modified-Since, If-None-Match, If-Unmodified-Since, X-Requested-With")
	//放行所有OPTIONS方法
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	//处理请求
	c.Next()
}
