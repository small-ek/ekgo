package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//分类
type Category struct {
	Id        int        `gorm:"primary_key" json:"id" uri:"id"` //标识
	Ids       int        `json:"ids" uri:"ids"`                  //标识
	Name      string     `json:"name"`                           //分类名称
	CreatedAt *Time      `json:"created_at"`                     //创建时间
	UpdatedAt *Time      `json:"updated_at"`                     //修改时间
	DeletedAt *time.Time `json:"deleted_at`                      //删除时间
}

func (Category) TableName() string {
	return "b_category"
}

//查询之后
func (this *Category) AfterFind(scope *gorm.Scope) error {
	return nil
}

//创建之前
func (this *Category) BeforeCreate(scope *gorm.Scope) error {
	return nil
}

//创建之后
func (this *Category) AfterCreate(scope *gorm.Scope) error {
	return nil
}

//更新之前
func (this *Category) BeforeUpdate(scope *gorm.Scope) error {
	return nil
}

//更新之后
func (this *Category) AfterUpdate(scope *gorm.Scope) error {
	return nil
}

//删除之前
func (this *Category) BeforeDelete(scope *gorm.Scope) error {
	return nil
}

//删除之后
func (this *Category) AfterDelete(scope *gorm.Scope) error {
	return nil
}
