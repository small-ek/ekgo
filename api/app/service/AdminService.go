package service

import (
	"ekgo/app/model"
	"ekgo/app/repository"
	"github.com/small-ek/ginp/request"
	"github.com/small-ek/ginp/response"
)

type Admin struct {
	PageParam request.PageParam
	Interface repository.AdminInterface
	Model     model.Admin
}



//Index 分页
func (get *Admin) Index() (*response.Page, error) {
	get.Interface = repository.Admin(get.Model)
	var list, err = get.Interface.GetPage(get.PageParam)
	return list, err
}

//Store 添加
func (get *Admin) Store() (*model.Admin, error) {
	get.Interface = repository.Admin(get.Model)
	var result, err = get.Interface.Create()
	return result, err
}

//Update 修改
func (get *Admin) Update() error {
	get.Interface = repository.Admin(get.Model)
	var err = get.Interface.Update()
	return err
}

//Delete 修改
func (get *Admin) Delete() error {
	get.Interface = repository.Admin(get.Model)
	var err = get.Interface.Delete()
	return err
}
