[English](./README-en.md) | [简体中文](./README.md)
#####联系方式 QQ群:68667333

- 前端UI框架：[element-ui](https://element.eleme.cn/#/zh-CN) 
- 后台框架：[gin](https://gin-gonic.com/zh-cn/docs)
- Orm框架：[gorm](http://gorm.book.jasperxu.com)
- Casbin鉴权：[casbin](https://gin-gonic.com/zh-cn/docs)    
- Cache缓存：[cache](https://github.com/coocood/freecache)
- 配置：[go-ini](https://github.com/go-ini/ini)
- Go模块代理：[goproxy](https://goproxy.io/)
- 代码演示:[https://go.smallek.com/]
- 账号:admin
- 密码:admin

## 1. 基本介绍

> Ekgo是一个基于vue和gin开发的全栈前后端分离的后台管理系统，表单生成器、代码生成器、集成jwt登录鉴权、动态路由、动态菜单、casbin权限认证、高度代码解耦完全建立在DDD领域驱动设计、自定义数据验证器、语言包、LRU算法缓存库、基于协程队列、Hook事件回调注册让您把更多时间专注在业务开发上。
## 2. 使用说明

```
- node版本 > v8.6.0
- golang版本 >= v1.11
- IDE推荐：Goland
```

### 2.1 web端

```bash
# 进入目录
cd web

# 安装包管理
npm install

# 运行
npm run serve

# 编译
npm run build
```

### 2.2 server端

```bash
#进入目录
cd api

#开启模块支持
SET GO111MODULE=on

#被墙设置GOPROXY环境变量
SET GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct

# 安装go依赖包
go mod tidy

#将依赖复制到项目路径的vendor文件夹中
go mod vendor

#运行(如果想享受静态语言的快乐请使用gowatch热编译)
go run main

# 交叉编译Linux环境或者mac,目标平台的操作系统
SET GOOS=linux
SET GOOS=darwin
SET GOOS=windows

#目标平台的体系架构
GOARCH=amd64
GOARCH=386
GOARCH=arm

# 编译
go build main

# 下载热编译
go get github.com/silenceper/gowatch

# 运行热编译
gowatch

#进入单元测试
cd test

#执行单元测试
go test -v -cover
```

### 2.3 swagger自动化API文档
#### 2.3.1 安装 swagger
````#
# 下载
go get -u github.com/swaggo/swag/cmd/swag
# 初始化
swagger swag init
````

### 2.4 服务部署
#### 或者使用screen、Supervisor、systemctl，tmux等守护进程启动
####nohup部署的日志输出可以分包采用日志的方式
````#
#切换项目根目录并设置可执行权限,这里使用nohup演示
nohup ./main &
#热更新关闭
kill -l main
#强制关闭
pkill main
````
## 3. EKGO目录结构

```
    ├─server  	     （后端文件夹）
    │  ├─app            （目录包含应用程序的核心代码）
    │  ├─├─admin        （控制器,处理进入应用程序请求的所有逻辑几乎都放置在此目录）
    │  ├─├─common        (公共接口）
    │  ├─├─middleware    (中间件）
    │  ├─├─model         (数据模型配置）
    │  ├─├─queue         (异步队列任务）
    │  ├─├─service       (接口服务层,一般用于封装数据接口的操作）
    │  ├─├─validate      (数据验证）
    │  ├─boot           （目录包含引导框架的）
    │  ├─├─cache        （缓存）
    │  ├─├─casbin        (casbin鉴权）
    │  ├─├─config        (配置）
    │  ├─├─db            (数据库设置）
    │  ├─├─logger        (日志设置）
    │  ├─├─router        (路由注册）
    │  ├─├─serve         (程序启动服务,可以用于多个服务启动）
    │  ├─├─validate      (验证器,可以用于数据数据参数验证）
    │  ├─config         （配置文件）
    │  ├─docs  	        （swagger文档目录）
    │  ├─lib            （公共的功能封装包，不包含业务需求实现。）
    │  ├─resource       （资源）
    │  ├─router         （目录包含了应用的所有路由定义）
    └─web            （前端文件）
        ├─public        （发布模板）
        └─src           （源码包）
            ├─assets	（静态文件）
            ├─common    （工具库）
            ├─components（组件）
            ├─lang      （语言包）
            ├─layout    （后端布局）
            ├─menu      （本地加载菜单配置文件）
            ├─router	（前端路由）
            ├─store     （vuex 状态管理仓）
            └─view      （前端页面）
```