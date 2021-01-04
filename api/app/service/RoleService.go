package service

import (
	"ekgo/app/model"
	"github.com/small-ek/antgo/orm"
	"github.com/small-ek/antgo/request"
	"github.com/small-ek/antgo/response"
	"gorm.io/gorm"
)

type RoleInterface interface {
	Index() *response.Write  //分页
	Store() *response.Write  //添加
	Update() *response.Write //修改
	Delete() *response.Write //删除
	Show() *response.Write   //查询
}

type Role struct {
	PageParam request.PageParam
	Model     model.Role
	Db        *gorm.DB
}

//分页
func (this *Role) Index() *response.Write {
	var list = []model.Role{}

	err := this.Db.Scopes(
		orm.Filters(this.PageParam.Filter),
		orm.Order(this.PageParam.Order),
		orm.Paginate(this.PageParam.PageSize, this.PageParam.CurrentPage),
	).Find(&list).Offset(0).Count(&this.PageParam.Total).Error

	if err != nil {
		return response.Fail("获取失败")
	}

	return &response.Write{
		Code: 200,
		Data: response.Page{
			Total: this.PageParam.Total,
			List:  list,
		},
	}
}

//查询单个
func (this *Role) Show() *response.Write {
	err := this.Db.First(&this.Model, this.Model.ID).Error

	if err == nil {
		return response.Success("获取成功", this.Model)
	}

	return response.Success("获取失败", err.Error())
}

//添加
func (this *Role) Store() *response.Write {
	err := this.Db.Create(&this.Model).Error

	if err == nil {
		return response.Success("保存成功")
	}

	return response.Success("保存失败", err.Error())
}

//修改
func (this *Role) Update() *response.Write {
	err := this.Db.Model(&this.Model).Updates(this.Model).Error

	if err == nil {
		return response.Success("修改成功")
	}

	return response.Success("保存失败", err.Error())
}

//删除
func (this *Role) Delete() *response.Write {
	err := this.Db.First(&this.Model, this.Model.ID).Error

	if err == nil {
		this.Db.Where("id=?", this.Model.ID).Delete(&this.Model)
		return response.Success("删除成功")
	}

	return response.Success("删除失败", err.Error())
}
