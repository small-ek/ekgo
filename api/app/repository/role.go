package repository

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"gorm.io/gorm"
)

type RoleFactory struct {
	Model    model.Role
	List     []model.Role
	Db       *gorm.DB
	Factorys Factory
}

//Factory 继承增删改查公共工厂
func (r *RoleFactory) Factory(model model.Role, Db ...*gorm.DB) *Factory {
	if len(Db) > 0 {
		r.Factorys.Db = Db[0]
	} else {
		r.Factorys.Db = db.Master
	}
	r.Factorys.New(model)
	return &r.Factorys
}

//New
func (r *RoleFactory) New(model model.Role, Db ...*gorm.DB) *RoleFactory {
	if len(Db) > 0 {
		r.Db = Db[0]
	} else {
		r.Db = db.Master
	}
	r.Model = model
	return r
}

//Default
func (r *RoleFactory) Default() *RoleFactory {
	return r
}

//SetDb
func (r *RoleFactory) SetDb(Db *gorm.DB) *RoleFactory {
	r.Db = Db
	r.Factorys.Db = Db
	return r
}

//SetModel
func (r *RoleFactory) SetModel(model model.Role) *RoleFactory {
	r.Factorys.Model = model
	r.Model = model
	return r
}

//GetModel
func (r *RoleFactory) GetModel() model.Role {
	var row = r.Factorys.Model.(*model.Role)
	r.Model = *row
	return r.Model
}

//GetModelList
func (r *RoleFactory) GetModelList() []model.Role {
	return r.List
}