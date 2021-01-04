package model

import (
	"time"
)

//管理员
type Admin struct {
	Id        int        `gorm:"primaryKey;autoIncrement" json:"id" uri:"id"` //标识
	Name      string     `json:"name"`                                        //名称
	Username  string     `json:"username"`                                    //用户名
	Salt      string     `gorm:"->:false" json:"salt,omitempty"`              //盐
	Password  string     `gorm:"->:false" json:"password,omitempty"`          //密码
	Status    string     `json:"status"`                                      //状态:启用=true,禁用=false
	Super     string     `json:"super"`                                       //超级管理员:是=true,否=false
	RoleId    int        `json:"role_id"`                                     //角色标识
	CreatedAt time.Time  `json:"created_at"`                                  //创建时间
	UpdatedAt time.Time  `json:"updated_at"`                                  //修改时间
	DeletedAt *time.Time `json:"deleted_at"`                                  //删除时间
}

//TableName
func (Admin) TableName() string {
	return "s_admin"
}
