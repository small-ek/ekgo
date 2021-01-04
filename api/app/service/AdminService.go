package service

import (
	"ekgo/app/model"
	"ekgo/app/repository"
	"ekgo/boot/db"
	"github.com/small-ek/antgo/request"
	"github.com/small-ek/antgo/response"
)

type Admin struct {
	PageParam request.PageParam
	Factory   repository.AdminFactory
	Model     model.Admin
}

//Index 分页
func (s *Admin) Index() (*response.Page, error) {
	s.Factory = repository.Admin(s.Model, db.Master)
	var result, err = s.Factory.GetPage(s.PageParam)
	return result, err
}

//Store 添加
func (s *Admin) Store() (*model.Admin, error) {
	s.Factory = repository.Admin(s.Model, db.Master)
	var result, err = s.Factory.Create()
	return result, err
}

//Update 修改
func (s *Admin) Update() error {
	s.Factory = repository.Admin(s.Model, db.Master)
	var err = s.Factory.Update()
	return err
}

//Delete 修改
func (s *Admin) Delete() error {
	s.Factory = repository.Admin(s.Model, db.Master)
	var err = s.Factory.Delete()
	return err
}
