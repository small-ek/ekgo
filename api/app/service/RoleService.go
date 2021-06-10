package service

import (
	"ekgo/app/model"
	"ekgo/app/repository"
	"github.com/small-ek/antgo/request"
	"github.com/small-ek/antgo/response"
)

type Role struct {
	PageParam request.PageParam
	Model     model.Role
	repository.RoleRepository
}

//Index 分页
func (s *Role) Index() (*response.Page, error) {
	list, total, err := s.Factory(s.Model).GetPage(s.PageParam)
	return &response.Page{List: list, Total: total}, err
}

//Show 查询
func (s *Role) Show() (map[string]interface{}, error) {
	var result, err = s.Factory(s.Model).FindByMap(s.Model.Id)
	return result, err
}

//Store 添加
func (s *Role) Store() (model.Role, error) {
	var _, err = s.Factory(s.Model).Create()
	var result = s.GetModel()
	return result, err
}

//Update 修改
func (s *Role) Update() error {
	var err = s.Factory(s.Model).Update()
	return err
}

//Delete 删除
func (s *Role) Delete() error {
	var err = s.Factory(s.Model).Delete(s.Model.Id)
	return err
}
