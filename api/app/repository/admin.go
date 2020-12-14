package repository

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"github.com/small-ek/ginp/orm"
	"github.com/small-ek/ginp/request"
	"github.com/small-ek/ginp/response"
	"gorm.io/gorm"
)

type AdminInterface interface {
	GetPage(request.PageParam) (*response.Page, error)
	FindByID() (*model.Admin, error)
	Create() (*model.Admin, error)
	Update() error
	Delete() error
}

type Admins struct {
	Model model.Admin
	Db    *gorm.DB
}

//Admin 默认
func Admin(model model.Admin) *Admins {
	return &Admins{
		Model: model,
		Db:    db.Master,
	}
}

//GetPage 获取分页
func (get *Admins) GetPage(Page request.PageParam) (*response.Page, error) {
	result := []model.Admin{}
	err := get.Db.Model(&get.Model).Scopes(
		orm.WhereBuildQuery(Page.Filter),
		orm.Order(Page.Order),
		orm.Paginate(Page.PageSize, Page.CurrentPage),
	).Find(&result).Offset(0).Count(&Page.Total).Error
	return &response.Page{Total: Page.Total, List: result}, err
}

//FindByID 根据ID查询
func (get *Admins) FindByID() (*model.Admin, error) {
	err := get.Db.First(&get.Model, get.Model.Id).Error
	return &get.Model, err
}

//Create 创建数据
func (get *Admins) Create() (*model.Admin, error) {
	var err = get.Db.Create(&get.Db).Error
	return &get.Model, err
}

//Update 更新数据
func (get *Admins) Update() error {
	var err = get.Db.Model(&get.Model).Updates(get.Model).Error
	return err
}

//Delete 删除数据
func (get *Admins) Delete() error {
	var err = get.Db.Delete(&get.Model, get.Model.Id).Error
	return err
}
