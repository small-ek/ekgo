import {config} from "../config";
import {convTable} from "../template";

export const serviceTemplate = (table_name, data) => {
    var field = ""
    data.forEach(function (value) {
        var associativeTable = convTable(value.table_name)
        console.log(associativeTable)
        if (value.association_table == true && value.relation == "has_one") {//一对一关联表
            field = `Preload("` + associativeTable + `").`
        } else if (value.association_table == true && value.relation == "has_many") {//一对多关联表
            field = `Preload("` + associativeTable + `").`
        }
    })

    return `package service

import (
	"` + config.projectName + `/api/app/model"
	"` + config.projectName + `/api/boot/db"
	"` + config.projectName + `/api/lib/response"
	"` + config.projectName + `/api/lib/orm"
	"github.com/jinzhu/gorm"
	"log"
)

type ` + table_name + `Interface interface {
	Index() *response.Write  //分页
	Store() *response.Write  //添加
	Update() *response.Write //修改
	Delete() *response.Write //删除
	Show() *response.Write   //查询
}

type ` + table_name + ` struct {
	PageParam response.PageParam
	Model     model.` + table_name + `
	Db        *gorm.DB
}

//分页
func (this *` + table_name + `) Index() *response.Write {
	var list = []model.` + table_name + `{}

	err := this.Db.` + field + `Scopes(
		orm.WhereQueryBuild(this.PageParam.Filter),
		orm.Order(this.PageParam.Order),
		orm.Paginate(this.PageParam.PageSize, this.PageParam.CurrentPage),
	).Find(&list).Offset(0).Count(&this.PageParam.Total).Error

	if err != nil {
		log.Print(err.Error())
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
func (this *` + table_name + `) Show() *response.Write {
	err := db.Master.` + field + `First(&this.Model, this.Model.Id).Error
	if err == nil {
		return response.Success("获取成功", this.Model)
	}

	return response.Fail("获取失败")
}

//添加
func (this *` + table_name + `) Store() *response.Write {
	err := this.Db.Create(&this.Model).Error

	if err == nil {
		return response.Success("保存成功")
	}

	return response.Fail("保存失败", err.Error())
}

//修改
func (this *` + table_name + `) Update() *response.Write {
	err := this.Db.Model(&this.Model).Update(this.Model).Error

	if err == nil {
		return response.Success("修改成功")
	}

	return response.Fail("保存失败", err.Error())
}

//删除
func (this *` + table_name + `) Delete() *response.Write {
	err := this.Db.First(&this.Model, this.Model.Id).Error

	if err == nil {
		this.Db.Where("id=?", this.Model.Id).Delete(&this.Model)
		return response.Success("删除成功")
	}

	return response.Fail("删除失败", err.Error())
}
`
}