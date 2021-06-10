package middleware

import (
	"ekgo/app/model"
	"ekgo/app/service"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/conv"
	"github.com/small-ek/antgo/jwt"
	"github.com/small-ek/antgo/os/atime"
	"github.com/small-ek/antgo/os/config"
)

//后台管理中间件授权认证登录
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" || token == "null" || token == "undefined" {
			c.JSON(200, gin.H{
				"code": 402,
				"msg":  "权限不足,请求失败",
			})
			c.Abort()
			return
		}
		//解析JWT
		var private_key = config.Decode().Get("jwt").Get("private_key").Bytes()
		var public_key = config.Decode().Get("jwt").Get("public_key").Bytes()
		result, err := jwt.Default(public_key, private_key).Decode(token)
		if err == nil {
			models := model.Admin{Id: conv.Int(result["id"])}
			var service = &service.Admin{Model: models}
			var row, err = service.Show()
			if row.Id == 0 || err != nil {
				c.JSON(200, gin.H{
					"code": 402,
					"msg":  "权限不足,请求失败",
				})
				c.Abort()
				return
			}

			if conv.Int64(result["exp"]) < atime.Now().Timestamp() {
				c.JSON(200, gin.H{
					"code": 402,
					"msg":  "登陆已超时,重新登陆",
				})
				c.Abort()
				return
			}
			c.Set("user", row)
			c.Next()
		} else {
			c.JSON(200, gin.H{
				"code": 402,
				"msg":  "权限不足,请求失败",
			})
			c.Abort()
			return
		}
	}
}

//获取后台用户信息
func GetAdmin(c *gin.Context) model.Admin {
	user := c.MustGet("user").(model.Admin)
	return user
}
