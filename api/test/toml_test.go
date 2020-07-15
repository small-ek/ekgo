package test

import (
	"ekgo/api/boot/config"
	"log"
	"testing"
)

func TestToml(t *testing.T) {
	config.Load("../config/config.toml")
	log.Println(config.Get)
}
