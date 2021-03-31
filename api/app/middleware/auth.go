package middleware

import (
	"ekgo/app/model"
	"ekgo/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/conv"
	"github.com/small-ek/antgo/jwt"
	"github.com/small-ek/antgo/os/config"
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
			model := model.Admin{Id: conv.Int(result["id"])}
			var row, _ = repository.Admin.Default(model).FindByID()

			if row.Id == 0 || model.Status == "false" {
				this.JSON(200, gin.H{
					"code": 402,
					"msg":  "权限不足,请求失败",
				})
				this.Abort()
				return
			}

			if result["exp"].(float64) < conv.GetTimestamp() {
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
