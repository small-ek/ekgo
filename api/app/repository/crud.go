package repository

import (
	"ekgo/boot/db"
	"github.com/small-ek/antgo/orm"
	"github.com/small-ek/antgo/request"
	"gorm.io/gorm"
)

//Factory 所有repository继承
type Base struct {
	Model interface{}              //模型
	List  []map[string]interface{} //返回列表
	Data  map[string]interface{}   //返回单个数据
	Db    *gorm.DB                 //Db连接
}

//Get 获取模型
func (b *Base) Get() *Base {
	return b
}

//New 设置模型
func (b *Base) New(model interface{}, Db ...*gorm.DB) *Base {
	if len(Db) > 0 {
		b.Db = Db[0]
	} else {
		b.Db = db.Master
	}
	b.Model = model
	return b
}

//New 设置数据库
func (b *Base) SetDb(db *gorm.DB) *Base {
	b.Db = db
	return b
}

//SetModel 设置模型
func (b *Base) SetModel(model interface{}) *Base {
	b.Model = model
	return b
}

//GetPage 获取分页
func (b *Base) GetPage(Page request.PageParam) ([]map[string]interface{}, int64, error) {
	var err = b.Db.Model(b.Model).Scopes(
		orm.Filters(Page.Filter),
		orm.Order(Page.Order),
		orm.Paginate(Page.PageSize, Page.CurrentPage),
	).Find(&b.List).Offset(-1).Limit(-1).Count(&Page.Total).Error
	return b.List, Page.Total, err
}

//FindByID 根据ID查询
func (b *Base) FindByID(id int) (interface{}, error) {
	result := db.Master.First(b.Model, id)
	return b.Model, result.Error
}

//FindByMap 根据ID查询
func (b *Base) FindByMap(id int) (map[string]interface{}, error) {
	var row = map[string]interface{}{}
	result := db.Master.Model(b.Model).First(&row, id)
	return row, result.Error
}

//Create 创建数据
func (b *Base) Create() (interface{}, error) {
	var err = b.Db.Create(b.Model).Error
	return b.Model, err
}

//CreateMap 批量创建
func (b *Base) CreateMap(data map[string]interface{}) error {
	var err = b.Db.Create(&data).Error
	return err
}

//Update 更新数据
func (b *Base) Update() error {
	var err = b.Db.Model(b.Model).Updates(b.Model).Error
	return err
}

//UpdateMap Map更新数据
func (b *Base) UpdateMap(data map[string]interface{}) error {
	var err = b.Db.Model(b.Model).Updates(data).Error
	return err
}

//Delete 软删除数据
func (b *Base) Delete(id int) error {
	var err = b.Db.Delete(b.Model, id).Error
	return err
}

//DeleteUnscoped 永久删除匹配的记录
func (b *Base) DeleteUnscoped() error {
	var err = b.Db.Unscoped().Delete(b.Model).Error
	return err
}

//FindByUnscoped 查找被软删除的记录
func (b *Base) FindByUnscoped(id int) (interface{}, error) {
	result := db.Master.Unscoped().First(b.Model, id)
	return b.Model, result.Error
}

//FindByInIds 根据id查找多个
func (b *Base) FindByInIds(ids []int) (interface{}, error) {
	result := db.Master.Find(b.List, ids)
	return b.List, result.Error
}
