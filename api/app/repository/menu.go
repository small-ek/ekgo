package repository

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"github.com/small-ek/antgo/conv"
	"gorm.io/gorm"
)

type MenuRepository struct {
	Model    model.Menu
	List     []model.Menu
	Db       *gorm.DB
	Base Base
}

//Factory 继承增删改查公共工厂
func (r *MenuRepository) Factory(model model.Menu, Db ...*gorm.DB) *Base {
	if len(Db) > 0 {
		r.Base.Db = Db[0]
	} else {
		r.Base.Db = db.Master
	}
	r.Base.New(model)
	return &r.Base
}

//New
func (r *MenuRepository) New(model model.Menu, Db ...*gorm.DB) *MenuRepository {
	if len(Db) > 0 {
		r.Db = Db[0]
	} else {
		r.Db = db.Master
	}
	r.Model = model
	return r
}

//Default
func (r *MenuRepository) Default() *MenuRepository {
	return r
}

//SetDb
func (r *MenuRepository) SetDb(Db *gorm.DB) *MenuRepository {
	r.Db = Db
	r.Base.Db = Db
	return r
}

//SetModel
func (r *MenuRepository) SetModel(model model.Menu) *MenuRepository {
	r.Base.Model = model
	r.Model = model
	return r
}

//GetModel
func (r *MenuRepository) GetModel() model.Menu {
	var row = r.Base.Model.(*model.Menu)
	r.Model = *row
	return *row
}

//GetModelList
func (r *MenuRepository) GetModelList() []model.Menu {
	return r.List
}

//GetByList 获取所有菜单
func (r *MenuRepository) GetByList() []model.Menu {
	r.Db.Order("sort desc,id asc").Find(&r.List)
	return r.List
}

//GetRoleMenu 获取角色权限的菜单
func (r *MenuRepository) GetRoleMenu(id, parent_id model.JSON) []model.Menu {
	var newParentId []int
	conv.Struct(&newParentId, parent_id)

	var newId []int
	conv.Struct(&newId, id)

	r.Db.Where("id IN (?)", newParentId).Or("id IN (?)", newId).Order("sort desc,id asc").Find(&r.List)
	return r.List
}
