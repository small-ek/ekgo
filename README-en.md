[English](./README-en.md) | [简体中文](./README.md)
#####Contact QQ Group:68667333

-Front-end UI framework: [element-ui](https://element.eleme.cn/#/zh-CN)
-Background framework: [gin](https://gin-gonic.com/zh-cn/docs)
-Orm framework: [gorm](http://gorm.book.jasperxu.com)
-Casbin authentication: [casbin] (https://gin-gonic.com/zh-cn/docs)
-Cache: [cache](https://github.com/coocood/freecache)
-Configuration: [go-ini](https://github.com/go-ini/ini)
-Go module proxy: [goproxy](https://goproxy.io/)
-Code demo: [https://go.smallek.com/]
-Account: admin
-Password: admin

## 1. Basic introduction

> Ekgo is a full-stack front-end and back-end management system developed based on vue and gin. It integrates jwt login authentication, dynamic routing, dynamic menu, casbin authority authentication, form generator, code generator, and highly decoupled code. Driven design in the DDD field, custom data validators, language packs, LRU algorithm cache libraries, coroutine-based queues, Hook event callback registration allows you to focus more time on business development.
## 2. Instructions

```
-node version> v8.6.0
-golang version >= v1.11
-IDE recommendation: Goland
```

### 2.1 web

```bash
# Enter the directory
cd web

# Installation package management
npm install

# Run
npm run serve

# Compile
npm run build
```

### 2.2 server

```bash
#Enter the directory
cd api

#Enable module support
SET GO111MODULE=on

#GOPROXY environment variable set by the wall
SET GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct

# Install go dependencies
go mod tidy

#Copy the dependencies to the vendor folder of the project path
go mod vendor

#Run (If you want to enjoy the joy of static language, please use gowatch hot compilation)
go run main

# Cross-compile the Linux environment or mac, the operating system of the target platform
SET GOOS=linux
SET GOOS=darwin
SET GOOS=windows

#Target platform architecture
GOARCH=amd64
GOARCH=386
GOARCH=arm

# Compile
go build main

# Download hot compilation
go get github.com/silenceper/gowatch

# Run hot compilation
gowatch

#Enter unit test
cd test

#Execute unit test
go test -v -cover
```

### 2.3 swagger automation API documentation
#### 2.3.1 Install swagger
``##
# Download
go get -u github.com/swaggo/swag/cmd/swag
# Initialize
swagger swag init
``''

### 2.4 Service Deployment
#### Or use screen, Supervisor, systemctl, tmux and other daemons to start
####The log output of the nohup deployment can be subcontracted using the log method
``##
#Switch to the project root directory and set executable permissions, use nohup demo here
nohup ./main &
#Hot update off
kill -l main
#Force close
pkill main
``''
## 3. EKGO directory structure

```
    ├─server (backend folder)
    │ ├─app (the directory contains the core code of the application)
    │ ├─├─admin (controller, almost all the logic to handle the request to enter the application is placed in this directory)
    │ ├─├─common (public interface)
    │ ├─├─middleware (middleware)
    │ ├─├─model (data model configuration)
    │ ├─├─queue (asynchronous queue task)
    │ ├─├─service (interface service layer, generally used to encapsulate data interface operations)
    │ ├─├─validate (data verification)
    │ ├─boot (the directory contains the boot frame)
    │ ├─├─cache (cache)
    │ ├─├─casbin (casbin authentication)
    │ ├─├─config (configuration)
    │ ├─├─db (database setting)
    │ ├─├─logger (log setting)
    │ ├─├─router (route registration)
    │ ├─├─serve (program start service, can be used to start multiple services)
    │ ├─├─validate (validator, can be used for data data parameter verification)
    │ ├─config (configuration file)
    │ ├─docs (swagger document directory)
    │ ├─lib (public function package, does not include the realization of business requirements.)
    │ ├─resource (resource)
    │ ├─router (the directory contains all route definitions of the application)
    └─web (front-end file)
        ├─public (post template)
        └─src (source package)
            ├─assets (static files)
            ├─common (tool library)
            ├─components (components)
            ├─lang (language pack)
            ├─layout (back-end layout)
            ├─menu (local loading menu configuration file)
            ├─router (front-end router)
            ├─store (vuex state management warehouse)
            └─view (front-end page)
```