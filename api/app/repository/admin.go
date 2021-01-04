package repository

import (
	"ekgo/app/model"
	"github.com/small-ek/antgo/orm"
	"github.com/small-ek/antgo/request"
	"github.com/small-ek/antgo/response"
	"gorm.io/gorm"
)

type AdminFactory interface {
	GetPage(request.PageParam) (*response.Page, error)
	FindByID() (*model.Admin, error)
	Create() (*model.Admin, error)
	Update() error
	Delete() error
}

type Admins struct {
	Model model.Admin
	List  []model.Admin
	Db    *gorm.DB
}

//Admin
func Admin(Model model.Admin, Db *gorm.DB) *Admins {
	return &Admins{
		Model: Model,
		Db:    Db,
	}
}

//GetPage 获取分页
func (m *Admins) GetPage(Page request.PageParam) (*response.Page, error) {
	err := m.Db.Model(&m.Model).Scopes(
		orm.Filters(Page.Filter),
		orm.Order(Page.Order),
		orm.Paginate(Page.PageSize, Page.CurrentPage),
	).Find(&m.List).Offset(0).Count(&Page.Total).Error
	return &response.Page{Total: Page.Total, List: m.List}, err
}

//FindByID 根据ID查询
func (m *Admins) FindByID() (*model.Admin, error) {
	err := m.Db.First(&m.Model, m.Model.Id).Error
	return &m.Model, err
}

//Create 创建数据
func (m *Admins) Create() (*model.Admin, error) {
	var err = m.Db.Create(&m.Db).Error
	return &m.Model, err
}

//Update 更新数据
func (m *Admins) Update() error {
	var err = m.Db.Model(&m.Model).Updates(m.Model).Error
	return err
}

//Delete 删除数据
func (m *Admins) Delete() error {
	var err = m.Db.Delete(&m.Model, m.Model.Id).Error
	return err
}
