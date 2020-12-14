package model

import (
	"time"
)

//内容
type Content struct {
	Id         int        `gorm:"PRIMARY_KEY" json:"id" uri:"id"`                                   //标识
	CategoryId int        `json:"category_id"`                                                      //分类标识
	Title      string     `json:"title"`                                                            //标题
	Body       string     `json:"body"`                                                             //内容
	CreatedAt  *Time      `json:"created_at"`                                                       //创建时间
	UpdatedAt  *Time      `json:"updated_at"`                                                       //修改时间
	DeletedAt  *time.Time `json:"deleted_at"`                                                       //删除时间
	Category   Category   `json:"category" gorm:"FOREIGNKEY:category_id;ASSOCIATION_FOREIGNKEY:id"` //
}

func (Content) TableName() string {
	return "b_content"
}
