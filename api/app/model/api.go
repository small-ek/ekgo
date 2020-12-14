package model

import (
	"ekgo/boot/db"
	"time"
)

//api接口
type Api struct {
	Id        int        `gorm:"primary_key" json:"id" uri:"id"` //标识
	Title     string     `json:"title"`                          //api名称
	Path      string     `json:"path"`                           //api路径
	Group     string     `json:"group"`                          //api分组
	Method    string     `json:"method"`                         //请求类型
	CreatedAt *Time      `json:"created_at"`                     //
	UpdatedAt *Time      `json:"updated_at"`                     //
	DeletedAt *time.Time `json:"deleted_at`                      //
}

func (Api) TableName() string {
	return "s_api"
}

//获取所有API
func GetApiList(id []int) ([]Api, error) {
	var data []Api
	err := db.Master.Where("id IN (?)", id).Find(&data).Error
	return data, err
}
