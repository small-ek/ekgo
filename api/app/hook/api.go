package hook

import (
	"log"
)

type Test struct{}

func (p *Test) Before() interface{} {
	log.Println("在之前执行")
	return nil
}
func (p *Test) After() {
	log.Println("在之后执行")
}
