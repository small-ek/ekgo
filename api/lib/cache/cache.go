package cache

import (
	"ekgo/api/boot/freecache"
	"ekgo/api/lib/conv"
	"ekgo/api/lib/secret"
)

type New struct {
	Key    string
	Group  string
	Expire int
}

//获取缓存数据(LUR算法)
func (this *New) Get() []byte {
	//判断是否有缓存
	var hash = secret.Sha256(this.Group + this.Key)
	getData, _ := freecache.Get([]byte(hash))
	return getData
}

//设置缓存数据(LUR算法)
func (this *New) Set(data interface{}) {
	//判断是否有缓存
	if this.Expire == 0 {
		this.Expire = 10000
	}
	var hash = secret.Sha256(this.Group + this.Key)
	go freecache.Set([]byte(hash), conv.StructToBytes(data), this.Expire)
}
