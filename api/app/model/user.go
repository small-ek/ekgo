package model

import (
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
	Details   JSON       `json:"details"`                        //其他信息
	CreatedAt *Time      `json:"created_at"`                     //
	UpdatedAt *Time      `json:"updated_at"`                     //
	DeletedAt *time.Time `json:"deleted_at`                      //
}

//TableName
func (User) TableName() string {
	return "b_user"
}
