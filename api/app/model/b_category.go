package model

import (
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
