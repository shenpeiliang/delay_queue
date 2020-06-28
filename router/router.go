package router

import (
	"gin/controller/www"
	"github.com/gin-gonic/gin"
)

//初始化路由
func RegisterRoutes(router *gin.Engine) {

	//分组路由 - admin
	RegisterRoutersAdmin(router)

	//分组路由 - www
	RegisterRoutersHome(router)

	//分组路由 - user
	RegisterRoutersUser(router)

	//分组路由 - api
	RegisterRoutersApi(router)

	//没有找到路由
	router.NoRoute(www.Err{}.Show404)

}
