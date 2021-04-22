package repository

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"gorm.io/gorm"
)

var Admin AdminInterface

func init() {
	Admin = &AdminFactory{}
}

type AdminInterface interface {
	Factory(model.Admin, ...*gorm.DB) *Factory
	New(model.Admin, ...*gorm.DB) *AdminFactory
	SetDb(Db *gorm.DB) *AdminFactory
	SetModel(model.Admin) *AdminFactory
}

type AdminFactory struct {
	Model    model.Admin
	List     []model.Admin
	Db       *gorm.DB
	Factorys Factory
}

//New
func (m *AdminFactory) New(model model.Admin, Db ...*gorm.DB) *AdminFactory {
	if len(Db) > 0 {
		m.Db = Db[0]
	} else {
		m.Db = db.Master
	}
	m.Factorys.Model = model
	return m
}

//Factory
func (m *AdminFactory) Factory(model model.Admin, Db ...*gorm.DB) *Factory {
	if len(Db) > 0 {
		m.Factorys.Db = Db[0]
	} else {
		m.Factorys.Db = db.Master
	}
	m.Factorys.Model = model
	return &m.Factorys
}

//SetDb
func (m *AdminFactory) SetDb(Db *gorm.DB) *AdminFactory {
	m.Db = Db
	return m
}

//SetModel
func (m *AdminFactory) SetModel(model model.Admin) *AdminFactory {
	m.Model = model
	return m
}

//GetModel
func (m *AdminFactory) GetModel() model.Admin {
	return m.Model
}

//GetModelList
func (m *AdminFactory) GetModelList() []model.Admin {
	return m.List
}

//FindByUserName 用户名查询
func (m *AdminFactory) FindByUserName() (model.Admin, error) {
	err := m.Db.Where("status='true' AND username = ?", m.Model.Username).First(&m.Model).Error
	return m.Model, err
}
