package test

import (
	"github.com/small-ek/antgo/conv"
	"log"
	"testing"
)

func TestOrm(t *testing.T) {
	FA(Factorys{Name: "test"})
}

type Factorys struct {
	Name string
}

//Default 获取分页
func Default(model interface{}) {
	log.Println(model)
	log.Println(&model)
	log.Println(conv.InterfaceToStruct(model))
}

func FA(model Factorys) () {
	log.Println(model)
	log.Println(&model)
	Default(model)
}
