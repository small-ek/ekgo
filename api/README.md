#环境搭建
####Go安装包下载安装
https://studygolang.com/dl
####golang环境搭建设置项目目录
在计算机系统用户变量修改 SET GOPATH=G:\www 为项目运行目录

```
SET GO111MODULE=on
```
###被墙设置GOPROXY环境变量
```
SET GOPROXY=https://goproxy.io
```
###初始化目录
```
go mod init [project name]
```
###或者被墙添加依赖或修改依赖版本，这里支持模糊匹配版 可以是master或者latest
```
go mod edit -replace=golang.org/x/net@v0.0.0=github.com/x/net@master
```
###下载依赖
```
go mod download
```
###将依赖包复制到项目下的 vendor 目录。建议一些使用了被墙包的话可以这么处理，方便用户快速使用命令go build -mod=vendor编译。
```
go mod vendor
```

###清楚所有包
```
go clean -modcache
```

###文件自动处理依赖关系
```
go mod tidy
```
###显示依赖关系
```
go list -m all
```

#监视文件保存自动热编译
go get github.com/silenceper/gowatch
#####执行运行命令
```
gowatch
```


#编译部署
#####window编译
```
SET GOOS=windows
```

#####交叉编译Linux
```
SET GOOS=linux
```

##linux上后台运行的方法
#后台运行或者使用screen、Supervisor、systemctl等守护进程启动
```
nohup ./main &
```
##关闭
```
pkill main
```
