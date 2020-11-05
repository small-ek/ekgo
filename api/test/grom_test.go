package test

import (
	"ekgo/api/boot/db"
	"testing"
)

type AgentCode struct {
	ID              uint
	AgentID         uint
	AgentName       string
	SendMsgFlatRate float32
	SendMsgPrice    float32
	Code            string `gorm:"index:idx_name"`
	CreatedAt       int64  `gorm:"autoCreateTime"`
}

func (AgentCode) TableName() string {
	return "newshopco.s_user2"
}
func TestGorm(t *testing.T) {
	db.RegisterMaster()
	/*var user=&User2{
		Name2:"223333",
		Name4:"66666",
	}*/
	db.Master.AutoMigrate(&AgentCode{})
}
