package repository

import (
	"ekgo/boot/db"
	"github.com/small-ek/antgo/conv"
	"github.com/small-ek/antgo/orm"
	"github.com/small-ek/antgo/request"
	"gorm.io/gorm"
)

//Factory 工厂模式 repository都去继承这个
type Factory struct {
	Model interface{}
	Value interface{}
	List  []map[string]interface{}
	Data  map[string]interface{}
	Db    *gorm.DB
}

//Default
func (m *Factory) Default() *Factory {
	return m
}

//New 设置模型
func (m *Factory) New(model interface{}, Db ...*gorm.DB) *Factory {
	if len(Db) > 0 {
		m.Db = Db[0]
	} else {
		m.Db = db.Master
	}
	m.Value = model
	m.Model = conv.InterfaceToStruct(model)
	return m
}

//New 设置模型
func (m *Factory) SetDb(db *gorm.DB) *Factory {
	m.Db = db
	return m
}

//GetPage 获取分页
func (m *Factory) GetPage(Page request.PageParam) ([]map[string]interface{}, int64, error) {
	var err = m.Db.Model(m.Model).Scopes(
		orm.Filters(Page.Filter),
		orm.Order(Page.Order),
		orm.Paginate(Page.PageSize, Page.CurrentPage),
	).Find(&m.List).Offset(0).Count(&Page.Total).Error
	return m.List, Page.Total, err
}

//FindByID 根据ID查询
func (m *Factory) FindByID(id int) (interface{}, error) {
	result := db.Master.First(m.Model, id)
	return m.Model, result.Error
}

//FindByMap 根据ID查询
func (m *Factory) FindByMap(id int) (map[string]interface{}, error) {
	var row = map[string]interface{}{}
	result := db.Master.Model(m.Model).First(&row, id)
	return row, result.Error
}

//Create 创建数据
func (m *Factory) Create() (interface{}, error) {
	var err = m.Db.Create(m.Model).Error
	return m.Model, err
}

//CreateMap 批量创建
func (m *Factory) CreateMap(data map[string]interface{}) error {
	var err = m.Db.Create(&data).Error
	return err
}

//Update 更新数据
func (m *Factory) Update() error {
	var err = m.Db.Model(m.Model).Updates(m.Model).Error
	return err
}

//UpdateMap Map更新数据
func (m *Factory) UpdateMap(data map[string]interface{}) error {
	var err = m.Db.Model(m.Model).Updates(data).Error
	return err
}

//Delete 软删除数据
func (m *Factory) Delete(id int) error {
	var err = m.Db.Delete(m.Model, id).Error
	return err
}

//DeleteUnscoped 永久删除匹配的记录
func (m *Factory) DeleteUnscoped() error {
	var err = m.Db.Unscoped().Delete(m.Model).Error
	return err
}

//FindByUnscoped 查找被软删除的记录
func (m *Factory) FindByUnscoped(id int) (interface{}, error) {
	result := db.Master.Unscoped().First(m.Model, id)
	return m.Model, result.Error
}
