package router

import (
	"gin/controller/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutersApi(router *gin.Engine) {
	//分组路由
	adminModule := router.Group("/api")
	{
		//退出登录
		adminModule.POST("/task/save", api.Task{}.Save)
	}

}
