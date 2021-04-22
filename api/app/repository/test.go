package repository

import (
	"ekgo/app/model"
	"gorm.io/gorm"
)

type AdminFactorys struct {
	Model model.Admin
	List  []model.Admin
	Data  map[string]interface{}
	Db    *gorm.DB
	Crud  Factorys
}

func (m *AdminFactorys) FindByID() (interface{}, error) {
	result, err := m.Crud.New(m.Model).FindByID(m.Model.Id)
	return result, err
}
