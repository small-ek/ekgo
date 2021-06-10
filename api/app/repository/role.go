package repository

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"gorm.io/gorm"
)

type RoleRepository struct {
	Model    model.Role
	List     []model.Role
	Db       *gorm.DB
	Base Base
}

//Factory 继承增删改查公共工厂
func (r *RoleRepository) Factory(model model.Role, Db ...*gorm.DB) *Base {
	if len(Db) > 0 {
		r.Base.Db = Db[0]
	} else {
		r.Base.Db = db.Master
	}
	r.Base.New(model)
	return &r.Base
}

//New
func (r *RoleRepository) New(model model.Role, Db ...*gorm.DB) *RoleRepository {
	if len(Db) > 0 {
		r.Db = Db[0]
	} else {
		r.Db = db.Master
	}
	r.Model = model
	return r
}

//Default
func (r *RoleRepository) Default() *RoleRepository {
	return r
}

//SetDb
func (r *RoleRepository) SetDb(Db *gorm.DB) *RoleRepository {
	r.Db = Db
	r.Base.Db = Db
	return r
}

//SetModel
func (r *RoleRepository) SetModel(model model.Role) *RoleRepository {
	r.Base.Model = model
	r.Model = model
	return r
}

//GetModel
func (r *RoleRepository) GetModel() model.Role {
	var row = r.Base.Model.(*model.Role)
	r.Model = *row
	return r.Model
}

//GetModelList
func (r *RoleRepository) GetModelList() []model.Role {
	return r.List
}