package model

import (
	"gorm.io/gorm"
	"time"
)

//角色表
type Role struct {
	Id           int            `gorm:"primary_key" json:"id" uri:"id"` //标识
	Name         string         `json:"name"`                           //角色名称
	ApiId        JSON           `json:"api_id"`                         //角色标识
	MenuId       JSON           `json:"menu_id"`                        //菜单标识
	ParentMenuId JSON           `json:"parent_menu_id"`                 //选中的一级菜单标识
	CreatedAt    time.Time      `json:"created_at"`                     //创建时间
	UpdatedAt    time.Time      `json:"updated_at"`                     //修改时间
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`                     //删除时间
}

func (Role) TableName() string {
	return "s_role"
}
