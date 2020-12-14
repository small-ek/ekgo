package model

import (
	"ekgo/boot/db"
	"github.com/small-ek/ginp/conv"
)

//角色表
type CasbinRule struct {
	PType string `json:"p_type"` //访问类型
	V0    string `json:"v0"`     //访问的组
	V1    string `json:"v1"`     //访问地址
	V2    string `json:"v2"`     //请求类型
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}

//创建casbin列表
//role_id 角色标识
//api 标识
func CreateCasbinList(role_id int, api_id []int) {
	//获取所有api列表
	api_list, err := GetApiList(api_id)
	if err != nil {

	} else {
		casbin := CasbinRule{PType: conv.String(role_id)}

		go func() {
			casbin.DeleteCasbin()
			for _, v := range api_list {
				casbin_rule := CasbinRule{
					PType: "p",
					V0:    conv.String(role_id),
					V1:    v.Path,
					V2:    v.Method,
				}
				casbin_rule.CreateCasbin()
			}
		}()

	}
}

//清除Casbin
func (data *CasbinRule) DeleteCasbin() {
	db.Master.Where("v0 = ?", data.PType).Delete(&data)
}

//创建casebin
func (data *CasbinRule) CreateCasbin() error {
	err := db.Master.Create(&data).Error
	return err
}
