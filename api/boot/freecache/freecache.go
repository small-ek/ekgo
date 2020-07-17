package freecache

import (
	"github.com/coocood/freecache"
)

const (
	//设置开始内存大小
	cacheSize = 1024 * 1024
)

var cache = freecache.NewCache(cacheSize)

//设置值
func Set(key []byte, val []byte, expire int) {
	cache.Set(key, val, expire)
}

//根据key获取值
func Get(key []byte) ([]byte, error) {
	got, err := cache.Get(key)
	return got, err
}
//删除
func Del(key []byte) bool {
	result := cache.Del(key)
	return result

}
//清空
func Clear() bool {
	cache.Clear()
	return true

}
