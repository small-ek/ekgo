package repository

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"gorm.io/gorm"
)

type AdminFactory struct {
	Model    model.Admin
	List     []model.Admin
	Db       *gorm.DB
	Factorys Factory
}

//Factory 继承增删改查公共工厂
func (r *AdminFactory) Factory(model model.Admin, Db ...*gorm.DB) *Factory {
	if len(Db) > 0 {
		r.Factorys.Db = Db[0]
	} else {
		r.Factorys.Db = db.Master
	}
	r.Factorys.New(model)
	return &r.Factorys
}

//New
func (r *AdminFactory) New(model model.Admin, Db ...*gorm.DB) *AdminFactory {
	if len(Db) > 0 {
		r.Db = Db[0]
	} else {
		r.Db = db.Master
	}
	r.Model = model
	return r
}

//Default
func (r *AdminFactory) Default() *AdminFactory {
	return r
}

//SetDb
func (r *AdminFactory) SetDb(Db *gorm.DB) *AdminFactory {
	r.Db = Db
	r.Factorys.Db = Db
	return r
}

//SetModel
func (r *AdminFactory) SetModel(model model.Admin) *AdminFactory {
	r.Factorys.Model = model
	r.Model = model
	return r
}

//GetModel
func (r *AdminFactory) GetModel() model.Admin {
	var row = r.Factorys.Model.(*model.Admin)
	r.Model = *row
	return r.Model
}

//GetModelList
func (r *AdminFactory) GetModelList() []model.Admin {
	return r.List
}

//FindByUserName 用户名查询
func (r *AdminFactory) FindByUserName() (model.Admin, error) {
	err := r.Db.Where("status='true' AND username = ?", r.Model.Username).First(&r.Model).Error
	return r.Model, err
}
