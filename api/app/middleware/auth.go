package middleware

import (
	"ekgo/app/model"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/conv"
	"github.com/small-ek/antgo/jwt"
	"github.com/small-ek/antgo/os/atime"
	"github.com/small-ek/antgo/os/config"
	"time"
)

//后台管理中间件授权认证登录
func AdminAuth() gin.HandlerFunc {
	return func(this *gin.Context) {
		token := this.Request.Header.Get("Authorization")
		if token == "" || token == "null" || token == "undefined" {
			this.JSON(200, gin.H{
				"code": 402,
				"msg":  "权限不足,请求失败",
			})
			this.Abort()
			return
		}
		//解析JWT
		var private_key = config.Decode().Get("jwt").Get("private_key").Bytes()
		var public_key = config.Decode().Get("jwt").Get("public_key").Bytes()
		result, err := jwt.Default(public_key, private_key).Decode(token)
		if err == nil {
			row := model.Admin{Id: conv.Int(result["id"])}
			//var service = &service.Admin{Model: model}
			//log.Println(service)
			//var row, err = service.Show()
			//log.Println(row)
			if row.Id == 0 || err != nil {
				this.JSON(200, gin.H{
					"code": 402,
					"msg":  "权限不足,请求失败",
				})
				this.Abort()
				return
			}

			if conv.Int64(result["exp"]) < atime.Timestamp(time.Now()) {
				this.JSON(200, gin.H{
					"code": 402,
					"msg":  "登陆已超时,重新登陆",
				})
				this.Abort()
				return
			}
			this.Set("user", row)
			this.Next()
		} else {
			this.JSON(200, gin.H{
				"code": 402,
				"msg":  "权限不足,请求失败",
			})
			this.Abort()
			return
		}
	}
}

//获取后台用户信息
func GetAdmin(this *gin.Context) model.Admin {
	user := this.MustGet("user").(model.Admin)
	return user
}
