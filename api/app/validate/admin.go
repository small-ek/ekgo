package validate

import (
	"ekgo/app/model"
)

var UserRule = map[string][]string{
	"username": {"require|请输入用户名"},
	"password": {"require|请输入原有密码"},
}

//用户名验证器
func CheckAdminRegister(request model.Admin) error {
	//需要验证的字段
	/*var err = validator.Default([]string{"username", "password"}, UserRule).Check(request)
	if err != nil {
		return err
	}

	var row, _ = repository.Admin.Default(request).FindByUserName()
	if row.Id > 0 {
		return errors.New("当前账号已存在")
	}*/
	return nil
}
