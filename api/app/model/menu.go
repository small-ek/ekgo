package model

import (
	"gorm.io/gorm"
	"time"
)

type Menu struct {
	Id        int            `gorm:"primary_key" json:"id" uri:"id"` //标识
	Title     string         `json:"title"`                          //菜单名称
	Path      string         `json:"path"`                           //菜单跳转地址
	Status    string         `json:"status"`                         //是否隐藏
	ParentId  int            `json:"parent_id"`                      //父级标识
	Sort      int            `json:"sort"`                           //排序
	Icon      string         `json:"icon"`                           //图标
	CreatedAt time.Time      `json:"created_at"`                     //创建时间
	UpdatedAt time.Time      `json:"updated_at"`                     //修改时间
	DeletedAt gorm.DeletedAt `json:"deleted_at"`                     //删除时间
}

func (Menu) TableName() string {
	return "s_menu"
}
