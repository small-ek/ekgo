package service

import (
	"ekgo/app/model"
	"ekgo/lib/response"
	"github.com/jinzhu/gorm"
)

type MenuInterface interface {
	Index() *response.Write                      //分页
	Store() *response.Write                      //添加
	Update() *response.Write                     //修改
	Delete() *response.Write                     //删除
	LeftMenu(admin *model.Admin) *response.Write //左侧菜单权限
}

type Menu struct {
	ID        int
	PageParam response.PageParam
	Model     model.Menu
	Db        *gorm.DB
}

//获取菜单
func (this *Menu) Index() *response.Write {
	var list = []model.Menu{}
	this.Db.Order("sort desc,id asc").Find(&list)

	return &response.Write{
		Code: 200,
		Data: list,
	}

}

//添加
func (this *Menu) Store() *response.Write {
	err := this.Db.Create(&this.Model).Error

	if err == nil {
		return response.Success("保存成功")
	}

	return response.Success("保存失败", err.Error())
}

//修改
func (this *Menu) Update() *response.Write {
	err := this.Db.Model(&this.Model).Update(this.Model).Error

	if err == nil {
		return response.Success("修改成功")
	}

	return response.Success("修改失败", err.Error())
}

//删除
func (this *Menu) Delete() *response.Write {
	err := this.Db.First(&this.Model, this.Model.Id).Error

	if err == nil {
		this.Db.Where("id=?", this.Model.Id).Delete(&this.Model)
		return response.Success("删除成功")
	}

	return response.Success("删除失败", err.Error())
}

//获取菜单
func (this *Menu) LeftMenu(user *model.Admin) *response.Write {
	var menu []model.Menu
	//获取对应的角色
	if user.Super == "true" {
		menu = model.GetMenuByList()
	} else {
		getRole := model.GetRole(user.RoleId)
		menu = model.GetRoleMenu(getRole.MenuId, getRole.ParentMenuId)
	}

	return &response.Write{
		Code: 200,
		Data: menu,
	}

}
