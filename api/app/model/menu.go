package model

import (
	"ekgo/boot/db"
	"github.com/small-ek/antgo/conv"
	"time"
)

type Menu struct {
	Id        int        `gorm:"primary_key" json:"id" uri:"id"` //标识
	Title     string     `json:"title"`                          //菜单名称
	Path      string     `json:"path"`                           //菜单跳转地址
	Status    string     `json:"status"`                         //是否隐藏
	ParentId  int        `json:"parent_id"`                      //父级标识
	Sort      int        `json:"sort"`                           //排序
	Icon      string     `json:"icon"`                           //图标
	CreatedAt *Time      `json:"created_at"`                     //创建时间
	UpdatedAt *Time      `json:"updated_at"`                     //修改时间
	DeletedAt *time.Time `json:"deleted_at"`                     //删除时间
}

func (Menu) TableName() string {
	return "s_menu"
}

//获取角色权限的菜单
func GetRoleMenu(id, parent_id Json) []Menu {
	var data []Menu

	var newParentId []int
	conv.Struct(&newParentId, parent_id)

	var newId []int
	conv.Struct(&newId, id)

	db.Master.Where("id IN (?)", newParentId).Or("id IN (?)", newId).Order("sort desc,id asc").Find(&data)
	return data
}

//获取所有菜单
func GetMenuByList() []Menu {
	var data []Menu
	db.Master.Order("sort desc,id asc").Find(&data)
	return data
}
