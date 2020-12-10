package service

import (
	"ekgo/app/model"
	"ekgo/boot/db"
	"ekgo/lib/orm"
	"ekgo/lib/response"
	"github.com/jinzhu/gorm"
)

type ApiInterface interface {
	Index() *response.Write  //分页
	Store() *response.Write  //添加
	Update() *response.Write //修改
	Delete() *response.Write //删除
	Show() *response.Write   //查询
}

type Api struct {
	PageParam response.PageParam
	Model     model.Api
	Db        *gorm.DB
}

//分页
func (this *Api) Index() *response.Write {
	var list = []model.Api{}

	err := this.Db.Scopes(
		orm.WhereQueryBuild(this.PageParam.Filter),
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
func (this *Api) Show() *response.Write {
	err := db.Master.First(&this.Model, this.Model.Id).Error
	if err == nil {
		return response.Success("获取成功", this.Model)
	}

	return response.Fail("获取失败")
}

//添加
func (this *Api) Store() *response.Write {
	err := this.Db.Create(&this.Model).Error

	if err == nil {
		return response.Success("保存成功")
	}

	return response.Fail("保存失败", err.Error())
}

//修改
func (this *Api) Update() *response.Write {
	err := this.Db.Model(&this.Model).Update(this.Model).Error

	if err == nil {
		return response.Success("修改成功")
	}

	return response.Fail("保存失败", err.Error())
}

//删除
func (this *Api) Delete() *response.Write {
	err := this.Db.First(&this.Model, this.Model.Id).Error

	if err == nil {
		this.Db.Where("id=?", this.Model.Id).Delete(&this.Model)
		return response.Success("删除成功")
	}

	return response.Fail("删除失败", err.Error())
}
