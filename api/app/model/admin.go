package model

import (
	"ekgo/api/boot/db"
	"github.com/jinzhu/gorm"
	"time"
)

//管理员
type Admin struct {
	Id        int        `gorm:"primary_key" json:"id" uri:"id"` //标识
	Name      string     `json:"name"`                           //名称
	Username  string     `json:"username"`                       //用户名
	Salt      string     `json:"salt,omitempty"`                 //盐
	Password  string     `json:"password,omitempty"`             //密码
	Status    string     `json:"status"`                         //状态:启用=true,禁用=false
	Super     string     `json:"super"`                          //超级管理员:是=true,否=false
	RoleId    int        `json:"role_id"`                        //角色标识
	CreatedAt *Time      `json:"created_at"`                     //创建时间
	UpdatedAt *Time      `json:"updated_at"`                     //修改时间
	DeletedAt *time.Time `json:"deleted_at"`                     //删除时间
}

func (Admin) TableName() string {
	return "s_admin"
}

//查询之后
func (this *Admin) AfterFind(scope *gorm.Scope) error {
	return nil
}

//创建之前
func (this *Admin) BeforeCreate(scope *gorm.Scope) error {
	return nil
}

//创建之后
func (this *Admin) AfterCreate(scope *gorm.Scope) error {
	return nil
}

//更新之前
func (this *Admin) BeforeUpdate(scope *gorm.Scope) error {
	return nil
}

//更新之后
func (this *Admin) AfterUpdate(scope *gorm.Scope) error {
	return nil
}

//删除之前
func (this *Admin) BeforeDelete(scope *gorm.Scope) error {
	return nil
}

//删除之后
func (this *Admin) AfterDelete(scope *gorm.Scope) error {
	return nil
}

//查询用户名
func CountAdminByUsername(username string) int {
	var count = 0
	db.Master.Model(&Admin{}).Where("username=?", username).Count(&count)
	return count
}
