package router

import (
	"gin/controller/www"
	"github.com/gin-gonic/gin"
)

func RegisterRoutersHome(router *gin.Engine) {
	//分组路由
	homeModule := router.Group("/")
	{
		homeModule.GET("/", www.Home{}.Index)
		homeModule.GET("/index", www.Home{}.Index)
	}
}
