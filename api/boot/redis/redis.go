package redis

import (
	"context"
	"ekgo/api/boot/config"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var ctx = context.Background()

func Client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Get.Redis.Address,
		Password: config.Get.Redis.Password, // no password set
		DB:       config.Get.Redis.Db,       // use default DB
	})
	_, err := client.Ping(ctx).Result()

	if err != nil {
		log.Println("Redis加载失败:" + err.Error())
	}
	return client
}

//获取值
func Get(key string) string {
	var result, err = Client().Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}

//修改值
func Set(key string, value interface{}, expiration ...int64) {
	var ex time.Duration

	if len(expiration) > 0 {
		ex = time.Duration(expiration[0])
	} else {
		ex = time.Duration(0)
	}

	err := Client().Set(ctx, key, value, ex).Err()
	if err != nil {
		log.Println("redis修改值错误或者配置不正确:" + err.Error())
	}

}
