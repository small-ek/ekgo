package model

import (
	"ekgo/api/boot/db"
	"time"
)

//角色表
type Role struct {
	ID           int        `gorm:"primary_key" json:"id" uri:"id"` //标识
	Name         string     `json:"name"`                           //角色名称
	ApiId        Json       `json:"api_id"`                         //角色标识
	MenuId       Json       `json:"menu_id"`                        //菜单标识
	ParentMenuId Json       `json:"parent_menu_id"`                 //选中的一级菜单标识
	CreatedAt    *Time      `json:"created_at"`                     //创建时间
	UpdatedAt    *Time      `json:"updated_at"`                     //修改时间
	DeletedAt    *time.Time `json:"deleted_at"`                     //删除时间
}

func (Role) TableName() string {
	return "s_role"
}
func GetRole(id int) Role {
	var data Role
	db.Master.First(&data, id)
	return data
}
