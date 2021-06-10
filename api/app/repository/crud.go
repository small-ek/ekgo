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
	Model interface{}              //模型
	List  []map[string]interface{} //返回列表
	Data  map[string]interface{}   //返回单个数据
	Db    *gorm.DB                 //Db连接
}

//Get
func (f *Factory) Get() *Factory {
	return f
}

//New 设置模型
func (f *Factory) New(model interface{}, Db ...*gorm.DB) *Factory {
	if len(Db) > 0 {
		f.Db = Db[0]
	} else {
		f.Db = db.Master
	}
	f.Model = conv.InterfaceToStruct(model)
	return f
}

//New 设置数据库
func (f *Factory) SetDb(db *gorm.DB) *Factory {
	f.Db = db
	return f
}

//SetModel 设置模型
func (f *Factory) SetModel(model interface{}) *Factory {
	f.Model = model
	return f
}

//GetPage 获取分页
func (f *Factory) GetPage(Page request.PageParam) ([]map[string]interface{}, int64, error) {
	var err = f.Db.Model(f.Model).Scopes(
		orm.Filters(Page.Filter),
		orm.Order(Page.Order),
		orm.Paginate(Page.PageSize, Page.CurrentPage),
	).Find(&f.List).Offset(0).Count(&Page.Total).Error
	return f.List, Page.Total, err
}

//FindByID 根据ID查询
func (f *Factory) FindByID(id int) (interface{}, error) {
	result := db.Master.First(f.Model, id)
	return f.Model, result.Error
}

//FindByMap 根据ID查询
func (f *Factory) FindByMap(id int) (map[string]interface{}, error) {
	var row = map[string]interface{}{}
	result := db.Master.Model(f.Model).First(&row, id)
	return row, result.Error
}

//Create 创建数据
func (f *Factory) Create() (interface{}, error) {
	var err = f.Db.Create(f.Model).Error
	return f.Model, err
}

//CreateMap 批量创建
func (f *Factory) CreateMap(data map[string]interface{}) error {
	var err = f.Db.Create(&data).Error
	return err
}

//Update 更新数据
func (f *Factory) Update() error {
	var err = f.Db.Model(f.Model).Updates(f.Model).Error
	return err
}

//UpdateMap Map更新数据
func (f *Factory) UpdateMap(data map[string]interface{}) error {
	var err = f.Db.Model(f.Model).Updates(data).Error
	return err
}

//Delete 软删除数据
func (f *Factory) Delete(id int) error {
	var err = f.Db.Delete(f.Model, id).Error
	return err
}

//DeleteUnscoped 永久删除匹配的记录
func (f *Factory) DeleteUnscoped() error {
	var err = f.Db.Unscoped().Delete(f.Model).Error
	return err
}

//FindByUnscoped 查找被软删除的记录
func (f *Factory) FindByUnscoped(id int) (interface{}, error) {
	result := db.Master.Unscoped().First(f.Model, id)
	return f.Model, result.Error
}

//FindByInIds 根据id查找多个
func (f *Factory) FindByInIds(ids []int) (interface{}, error) {
	result := db.Master.Find(f.List, ids)
	return f.List, result.Error
}
