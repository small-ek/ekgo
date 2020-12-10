package validate

import (
	"ekgo/app/model"
	"ekgo/app/service"
	"ekgo/boot/db"
	"ekgo/lib/conv"
	"ekgo/lib/response"
	"ekgo/lib/secret"
	"ekgo/lib/validator"
)

var Interface service.AdminInterface

var Rule = map[string][]string{
	"password":     {"require|请输入密码"},
	"old_password": {"require|请输入原有密码"},
}

type Admin struct {
	Request interface{}
}

//密码验证器
func (this *Admin) CheckPassword() *response.Write {
	//需要验证的字段
	var Scene = []string{"old_password", "password"}
	var validator = &validator.New{
		Request: this.Request,
		Rule:    Rule,
		Scene:   Scene,
	}
	if err := validator.CheckRule(); err != nil {
		return response.Fail(err.Error())
	}
	if this.CheckOldPassowrd() == false {
		return response.Fail("原有密码不正确")
	}
	return nil
}

//验证旧的密码
func (this *Admin) CheckOldPassowrd() bool {
	var request map[string]interface{}
	conv.Struct(&request, this.Request)
	var Model = model.Admin{Id: conv.Int(request["id"])}
	Interface = &service.Admin{Model: Model, Db: db.Master}
	conv.Struct(&Model, Interface.Show().Data)
	var oldPassword = secret.Sha256(Model.Salt + conv.String(request["old_password"]))

	if oldPassword == Model.Password {
		return true
	}
	return false
}
