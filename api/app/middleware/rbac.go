package middleware

import (
	"ekgo/api/boot/casbin"
	"ekgo/api/lib/conv"
	"github.com/gin-gonic/gin"
	"strings"
)

//casbin验证角色权限
func CasbinHandler() gin.HandlerFunc {
	return func(this *gin.Context) {
		var user = GetAdmin(this)
		if user.Super == "false" {
			//获取用户的角色标识
			sub := conv.Int(user.RoleId)
			//获取请求的URI
			var obj = strings.Split(this.Request.RequestURI, "?")[0]
			//获取请求方法
			act := this.Request.Method
			e := casbin.Register()
			//判断策略中是否存在
			ok, err := e.Enforce(sub, obj, act)
			if ok == false || err != nil {
				this.SecureJSON(200, gin.H{
					"code": 401,
					"msg":  "权限不足,请求失败",
				})
				this.Abort()
				return
			}
		}
		this.Next()
	}
}
