package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//用户
type User struct {
	Id        int        `gorm:"primary_key" json:"id" uri:"id"` //标识
	Username  string     `json:"username"`                       //用户名
	Salt      string     `json:"salt,omitempty"`                 //盐
	Password  string     `json:"password,omitempty"`             //密码
	Nickname  string     `json:"nickname"`                       //昵称
	Head      string     `json:"head"`                           //头像
	FullName  string     `json:"full_name"`                      //姓名
	Sex       string     `json:"sex"`                            //性别
	Email     string     `json:"email"`                          //邮箱
	Phone     string     `json:"phone"`                          //手机
	Status    string     `json:"status"`                         //状态
	Details   Json       `json:"details"`                        //其他信息
	CreatedAt *Time      `json:"created_at"`                     //
	UpdatedAt *Time      `json:"updated_at"`                     //
	DeletedAt *time.Time `json:"deleted_at`                      //
}

func (User) TableName() string {
	return "b_user"
}

//查询之后
func (this *User) AfterFind(scope *gorm.Scope) error {
	return nil
}

//创建之前
func (this *User) BeforeCreate(scope *gorm.Scope) error {
	return nil
}

//创建之后
func (this *User) AfterCreate(scope *gorm.Scope) error {
	return nil
}

//更新之前
func (this *User) BeforeUpdate(scope *gorm.Scope) error {
	return nil
}

//更新之后
func (this *User) AfterUpdate(scope *gorm.Scope) error {
	return nil
}

//删除之前
func (this *User) BeforeDelete(scope *gorm.Scope) error {
	return nil
}

//删除之后
func (this *User) AfterDelete(scope *gorm.Scope) error {
	return nil
}
