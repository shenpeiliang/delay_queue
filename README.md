# easy-gin
一套基于 Gin 框架的 MVC 脚手架
- 封装了 Gin Web 服务配置、路由配置、数据库/连接池配置、视图配置，方便快速体验及构建 Go Web 工程
- 自带一套用于体验及演示的 Restful Api 代码示例


## 安装步骤
### 安装govendor包管理工具
```
# go get -u -v github.com/kardianos/govendor
```
### 拉取源码
```sh
# cd $GOPATH/src && git clone git@github.com:sqrtcat/easy-gin.git && cd easy-gin
```
### 使用 govendor 安装依赖包
```sh
# govendor sync
```
### 服务配置
```go
package configs

// 服务配置 防止变量污染故用函数组织
func GetServerConfig() (serverConfig map[string]string) {
	serverConfig = make(map[string]string)

	serverConfig["HOST"] = "0.0.0.0"                     //监听地址
	serverConfig["PORT"] = "8080"                        //监听端口
	serverConfig["VIEWS_PATTERN"] = "easy-gin/views/*/*" //视图模板路径pattern
	serverConfig["ENV"] = "debug"                        //环境模式 release/debug/test
	return
}

```
### 数据库及连接池配置
因框架启动时会创建连接池服务，故需配置好数据库后运行
```go
package configs

// 数据库配置
func GetDbConfig() map[string]string {
	// 初始化数据库配置map
	dbConfig := make(map[string]string)

	dbConfig["DB_HOST"] = "127.0.0.1"          //主机
	dbConfig["DB_PORT"] = "3306"               //端口
	dbConfig["DB_NAME"] = "golang"             //数据库
	dbConfig["DB_USER"] = "root"               //用户名
	dbConfig["DB_PWD"] = ""                    //密码
	dbConfig["DB_CHARSET"] = "utf8"

	dbConfig["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	dbConfig["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	dbConfig["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return dbConfig
}
```
### 启动服务
```sh
# go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] Loaded HTML Templates (3):
        -
        - index.html
        - index/index.html

[GIN-debug] GET    /                         --> easy-gin/controllers.IndexHome (3 handlers)
[GIN-debug] GET    /index                    --> easy-gin/controllers.IndexHome (3 handlers)
[GIN-debug] GET    /users/:id                --> easy-gin/controllers.UserGet (3 handlers)
[GIN-debug] GET    /users                    --> easy-gin/controllers.UserGetList (3 handlers)
[GIN-debug] POST   /users                    --> easy-gin/controllers.UserPost (3 handlers)
[GIN-debug] PUT    /users/:id                --> easy-gin/controllers.UserPut (3 handlers)
[GIN-debug] PATCH  /users/:id                --> easy-gin/controllers.UserPut (3 handlers)
[GIN-debug] DELETE /users/:id                --> easy-gin/controllers.UserDelete (3 handlers)
[GIN-debug] Listening and serving HTTP on 0.0.0.0:8080
```
### 访问服务
http://yourhost:8080/

## 快速体验
### 导入框架示例 Sql
```sql
CREATE DATABASE `golang` DEFAULT CHARSET uft8mb4 DEFAULT utf8mb4_general_ci;
USE `golang`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(25) NOT NULL,
  `age` tinyint unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
```
### 定义控制器
controllers<br>
controllers/IndexController.go<br>
```go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHome(ctx *gin.Context) {

	//// query string
	//queryVal1 := ctx.Query("val1")
	//queryVal2 := ctx.DefaultQuery("val2", "val2_default")
	//
	//// post form data
	//formVal3 := ctx.PostForm("val3")
	//formVal4 := ctx.DefaultPostForm("val4", "val4_default")
	//
	//// path info
	//pathVal5 := ctx.Param("val5")

	//ctx.String(http.StatusOK, "hello %s %s %s %s %s", queryVal1, queryVal2, formVal3, formVal4, pathVal5)
	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"msg": "easy gin",
	})
}
```
### 定义模型
models<br>
models/User.go<br>
```go
package models

import (
	"log"
)

type User struct {
	Model
	Id   int    `json:"id" form:"id" primaryKey:"true"`
	Name string `json:"name" form:"name" binding:"required"`
	Age  int    `json:"age" form:"age" binding:"required"`
}

// get one
func (model *User) UserGet(id int) (user User, err error) {
  ....
}

// get list
func (model *User) UserGetList(page int, pageSize int) (users []User, err error) {
	....
}

// create
func (model *User) UserAdd() (id int64, err error) {
	....
}

// update
func (model *User) UserUpdate(id int) (afr int64, err error) {
	....
}

// delete
func (model *User) UserDelete(id int) (afr int64, err error) {
	....
}
```
### 定义视图
views<br>
views/index<br>
views/index/index.html<br>

### 定义路由 Restful
routes<br>
routes/router.go<br>
```go
package routes

import (
	"easy-gin/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controllers.IndexHome)
	router.GET("/index", controllers.IndexHome)

	router.GET("/users/:id", controllers.UserGet)
	router.GET("/users", controllers.UserGetList)
	router.POST("/users", controllers.UserPost)
	router.PUT("/users/:id", controllers.UserPut)
	router.PATCH("/users/:id", controllers.UserPut)
	router.DELETE("/users/:id", controllers.UserDelete)
}
```

### 热更新开发模式借用 rizla 插件
```sh
# go get -u github.com/kataras/rizla
# rizla main.go
```
### 注意
因 golang 的包载入机制问题，项目名如需改为其他，需修改框架内的部分包的载入路径 `easy-gin` to `your-projectName`
```
main.go
server/server.go
routes/router.go
server/server.go
models/*
```
go 1.11 后大家可以改用 gomod 模式
