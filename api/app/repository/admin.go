package repository

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"github.com/small-ek/antgo/orm"
	"github.com/small-ek/antgo/request"
	"gorm.io/gorm"
)

var Admin AdminInterface

func init() {
	Admin = &AdminFactory{}
}

type AdminInterface interface {
	Default(model.Admin) *AdminFactory
	SetDb(Db *gorm.DB) *AdminFactory
	SetModel(model.Admin) *AdminFactory
	GetPage(request.PageParam) ([]model.Admin, int64, error)
}
type AdminFactory struct {
	Model model.Admin
	List  []model.Admin
	Data  map[string]interface{}
	Db    *gorm.DB
}

//Default
func (m *AdminFactory) Default(model model.Admin) *AdminFactory {
	m.Model = model
	m.Db = db.Master
	return m
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

//GetPage 获取分页
func (m *AdminFactory) GetPage(Page request.PageParam) ([]model.Admin, int64, error) {
	var err = m.Db.Model(&m.Model).Scopes(
		orm.Filters(Page.Filter),
		orm.Order(Page.Order),
		orm.Paginate(Page.PageSize, Page.CurrentPage),
	).Find(&m.List).Offset(0).Count(&Page.Total).Error
	return m.List, Page.Total, err
}

//FindByID 根据ID查询
func (m *AdminFactory) FindByID() (model.Admin, error) {
	result := m.Db.First(&m.Model, m.Model.Id)
	return m.Model, result.Error
}

//Create 创建数据
func (m *AdminFactory) Create() (model.Admin, error) {
	var err = m.Db.Create(&m.Model).Error
	return m.Model, err
}

//CreateMap 批量创建
func (m *AdminFactory) CreateMap(data map[string]interface{}) error {
	var err = m.Db.Create(&data).Error
	return err
}

//Update 更新数据
func (m *AdminFactory) Update() error {
	var err = m.Db.Model(&m.Model).Updates(m.Model).Error
	return err
}

//UpdateMap Map更新数据
func (m *AdminFactory) UpdateMap(data map[string]interface{}) error {
	var err = m.Db.Model(&m.Model).Updates(data).Error
	return err
}

//Delete 删除数据
func (m *AdminFactory) Delete() error {
	var err = m.Db.Delete(&m.Model, m.Model.Id).Error
	return err
}

//DeleteAll 批量删除
func (m *AdminFactory) DeleteAll(id []int) error {
	var err = m.Db.Delete(&m.Model, id).Error
	return err
}

//DeleteUnscoped 永久删除
func (m *AdminFactory) DeleteUnscoped() error {
	var err = m.Db.Delete(&m.Model).Error
	return err
}

//FindByUserName 用户名查询
func (m *AdminFactory) FindByUserName() (model.Admin, error) {
	err := m.Db.Where("status='true' AND username = ?", m.Model.Username).First(&m.Model).Error
	return m.Model, err
}
