package test

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestOrm(t *testing.T) {
	//var user model.Admin
	//db.Master.First(&user)
	//log.Println(user)
	var model=AdminFactory{
		Model:model.Admin{Id: 2},
	}
	var result,err=model.FindByID()
	log.Println(result)
	log.Println(err)
}

type Factorys struct {
	Model interface{}
	List  interface{}
	Data  map[string]interface{}
	Db    *gorm.DB
	Id 	  interface{}
}
//Default 获取分页
func (m *Factorys) Default(model interface{}) *Factorys {
	m.Db = db.Master
	m.Model=model
	return m
}
//FindByID 根据ID查询
func (m *Factorys) FindByID(id interface{}) (interface{}, error) {
	result := m.Db.First(m.Model, id)
	return m.Model, result.Error
}

type AdminFactory struct {
	Model model.Admin
	List  []model.Admin
	Data  map[string]interface{}
	Db    *gorm.DB
	Crud   Factorys
}

func (m *AdminFactory) FindByID() (interface{}, error) {
	result,err := m.Crud.Default(&m.Model).FindByID(m.Model.Id)
	return result, err
}