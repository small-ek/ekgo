package test

import (
	"ekgo/api/boot/config"
	"ekgo/api/boot/redis"
	"log"
	"testing"
)

func TestRedis(t *testing.T) {
	config.Load("../config/config.toml")
	redis.Set("test", "name")
	log.Println(redis.Get("test"))
	log.Println(redis.Get("test2"))
}
