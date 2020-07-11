package freecache

import (
	"github.com/coocood/freecache"
)

const (
	//设置开始内存大小
	cacheSize = 1024 * 1024
)

var cache = freecache.NewCache(cacheSize)

//key    建
//value  值
//expire 过期时间毫秒
func Set(key []byte, val []byte, expire int) {
	cache.Set(key, val, expire)
}

//根据key获取参数
func Get(key []byte) ([]byte, error) {
	got, err := cache.Get(key)
	return got, err
}

func Del(key []byte) bool {
	result := cache.Del(key)
	return result

}

func Clear() bool {
	cache.Clear()
	return true

}
