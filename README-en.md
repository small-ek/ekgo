w[English](./README-en.md) | [简体中文](./README.md)
#####Contact QQ Group:68667333

-Front-end UI framework: [element-ui](https://element.eleme.cn/#/zh-CN)
-Background framework: [gin](https://gin-gonic.com/zh-cn/docs)
-Orm framework: [gorm](http://gorm.book.jasperxu.com)
-Casbin authentication: [casbin] (https://gin-gonic.com/zh-cn/docs)
-Cache: [cache](https://github.com/coocood/freecache)
-Go module proxy: [goproxy](https://goproxy.io/)
-Code demo: [https://go.smallek.com/]
-Account: admin
-Password: admin

## 1. Basic introduction

>Ekgo made further improvements to the underlying architecture to reduce dependence. The main features of the back-end management system based on the separation of the front and back ends of the full stack developed by gin and vue include:
 
  + Adopt DDD field-driven design
  + Routing cross domain request support
  + Custom data validator
  + Code generator (model, api, service, route, vue form)
  + Form builder
  + Integrated jwt login authentication
  + LRU algorithm cache
  + Based on coroutine queue
  + Hook event (plug-in mode)
  + Perfect dependency injection
  + Middleware support
  + casbin authority authentication (RBAC)
  + Dynamic routing, dynamic menu
  + Data type conversion
  + Gorm, Scopes component paging, automatically build where and other common methods
  + Gorm time type conversion and Json format support
  + Common encryption and decryption encoding and decoding operations
  + Redis integration
  + Http request
  + Request log split by time
  + Common help methods

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

### 2.2 api side

```bash
# Enter the directory
cd api

# Turn on module support
SET GO111MODULE=on

# Set GOPROXY environment variable by the wall
SET GOPROXY=https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,https://goproxy.io,direct

# Install go dependencies
go mod tidy

# Copy the dependencies to the vendor folder of the project path
go mod vendor

# Run (if you want to enjoy the joy of static language, please use gowatch hot compilation)
go run main

# Cross-compile the Linux environment or mac, the operating system of the target platform
SET GOOS=linux
SET GOOS=darwin
SET GOOS=windows

# System architecture of target platform
GOARCH=amd64
GOARCH=386
GOARCH=arm

# Compile
go build main

# Download hot compilation
go get github.com/silenceper/gowatch

# Run hot compilation
gowatch

# Enter unit test
cd test

# Perform unit testing
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
    │ ├─├─hook (hook)
    │ ├─├─middleware (middleware)
    │ ├─├─model (data model configuration)
    │ ├─├─queue (asynchronous queue task)
    │ ├─├─service (interface service layer, generally used to encapsulate data interface operations)
    │ ├─├─validate (data verification)
    │ ├─boot (the directory contains the boot frame)
    │ ├─├─casbin (casbin authentication)
    │ ├─├─config (configuration)
    │ ├─├─db (database setting)
    │ ├─├─freecache (cache)
    │ ├─├─logger (log setting)
    │ ├─├─queue (coroutine queue)
    │ ├─├─redis (Redis)
    │ ├─├─router (route registration)
    │ ├─├─serve (program start service, can be used to start multiple services)
    │ ├─config (configuration file)
    │ ├─docs (swagger document directory)
    │ ├─lib (public function package, does not include the realization of business requirements.)
    │ ├─resource (resource)
    │ ├─router (the directory contains all route definitions of the application)
    └─web (front-end file)
        ├─public (post template)
        └─src (source package)
            ├─assets (static files)
            ├─common (tool library, configuration items, etc.)
            ├─components (components)
            ├─lang (language pack)
            ├─layout (back-end layout)
            ├─menu (local loading menu configuration file)
            ├─router (front-end router)
            ├─store (vuex state management warehouse)
            └─view (front-end page)
```