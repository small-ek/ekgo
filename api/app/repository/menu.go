package repository

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"github.com/small-ek/antgo/conv"
	"gorm.io/gorm"
)

type MenuFactory struct {
	Model    model.Menu
	List     []model.Menu
	Db       *gorm.DB
	Factorys Factory
}

//Factory 继承增删改查公共工厂
func (r *MenuFactory) Factory(model model.Menu, Db ...*gorm.DB) *Factory {
	if len(Db) > 0 {
		r.Factorys.Db = Db[0]
	} else {
		r.Factorys.Db = db.Master
	}
	r.Factorys.New(model)
	return &r.Factorys
}

//New
func (r *MenuFactory) New(model model.Menu, Db ...*gorm.DB) *MenuFactory {
	if len(Db) > 0 {
		r.Db = Db[0]
	} else {
		r.Db = db.Master
	}
	r.Model = model
	return r
}

//Default
func (r *MenuFactory) Default() *MenuFactory {
	return r
}

//SetDb
func (r *MenuFactory) SetDb(Db *gorm.DB) *MenuFactory {
	r.Db = Db
	r.Factorys.Db = Db
	return r
}

//SetModel
func (r *MenuFactory) SetModel(model model.Menu) *MenuFactory {
	r.Factorys.Model = model
	r.Model = model
	return r
}

//GetModel
func (r *MenuFactory) GetModel() model.Menu {
	var row = r.Factorys.Model.(*model.Menu)
	r.Model = *row
	return *row
}

//GetModelList
func (r *MenuFactory) GetModelList() []model.Menu {
	return r.List
}

//GetByList 获取所有菜单
func (r *MenuFactory) GetByList() []model.Menu {
	r.Db.Order("sort desc,id asc").Find(&r.List)
	return r.List
}

//GetRoleMenu 获取角色权限的菜单
func (r *MenuFactory) GetRoleMenu(id, parent_id model.JSON) []model.Menu {
	var newParentId []int
	conv.Struct(&newParentId, parent_id)

	var newId []int
	conv.Struct(&newId, id)

	r.Db.Where("id IN (?)", newParentId).Or("id IN (?)", newId).Order("sort desc,id asc").Find(&r.List)
	return r.List
}
