package service

import (
	"ekgo/api/ek/conv"
	"log"
)

type IUserService interface {
	GetName(id int)string
}
type UserService struct {

}
func (this *UserService)GetName(id int) string {
	log.Println(id)
	return conv.String(id)
}