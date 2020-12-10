package casbin

import (
	"ekgo/boot/db"
	"github.com/casbin/casbin/util"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/small-ek/ginp/os/config"
	"log"
)

// 自定义规则函数
func ParamsMatch(key1, key2 string) bool {
	return util.KeyMatch2(key1, key2)

}

// 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	key1 := args[0].(string)
	key2 := args[1].(string)
	return ParamsMatch(key1, key2), nil
}

//持久化到数据库  引入自定义规则
func Register() *casbin.Enforcer {
	connection, userName, password, host, port, database, _ := db.GetDbConfig()

	var Master string
	switch connection {
	case "mysql":
		Master = "" + userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + ""
	case "postgres":
		Master = "host=" + host + " port=" + port + " user=" + userName + " dbname=" + database + " sslmode=disable" + " password=" + password
	}

	a, err := gormadapter.NewAdapter(connection, Master, true) // Your driver and data source.

	if err != nil {
		log.Println("数据库连接失败！" + err.Error())
	}
	var rbacPach = config.Decode().Get("casbin").Get("path").String()

	e, err2 := casbin.NewEnforcer(rbacPach, a)
	if err2 != nil {
		log.Println("配置读取失败！" + err.Error())
	}
	//添加函数必须在conf 里面定于执行模型否则不执行
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	e.LoadPolicy()
	return e
}
