package service

import (
	"ekgo/app/model"
	"ekgo/app/repository"
	"github.com/small-ek/antgo/request"
	"github.com/small-ek/antgo/response"
)

type Menu struct {
	PageParam request.PageParam
	Model     model.Menu
	Menu      repository.MenuRepository
	RoleModel model.Role
	Role      repository.RoleRepository
}

//Index 分页
func (s *Menu) Index() (*response.Page, error) {
	list, total, err := s.Menu.Factory(s.Model).GetPage(s.PageParam)
	return &response.Page{List: list, Total: total}, err
}

//Show 查询
func (s *Menu) Show() (map[string]interface{}, error) {
	var result, err = s.Menu.Factory(s.Model).FindByMap(s.Model.Id)
	return result, err
}

//Store 添加
func (s *Menu) Store() (model.Menu, error) {
	var _, err = s.Menu.Factory(s.Model).Create()
	var result = s.Menu.GetModel()
	return result, err
}

//Update 修改
func (s *Menu) Update() error {
	var err = s.Menu.Factory(s.Model).Update()
	return err
}

//Delete 删除
func (s *Menu) Delete() error {
	var err = s.Menu.Factory(s.Model).Delete(s.Model.Id)
	return err
}

//获取菜单
func (s *Menu) LeftMenu(user *model.Admin) []model.Menu {
	var menu []model.Menu
	//获取对应的角色
	if user.Super == "true" {
		menu = s.Menu.New(s.Model).GetByList()
	} else {
		s.Role.Factory(s.RoleModel).FindByID(user.RoleId)
		var getRole = s.Role.GetModel()
		menu = s.Menu.New(s.Model).GetRoleMenu(getRole.MenuId, getRole.ParentMenuId)
	}
	return menu

}
