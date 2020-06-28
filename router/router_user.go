package router

import (
	"gin/controller/admin"
	"gin/middleware/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutersUser(router *gin.Engine) {
	userModule := router.Group("/user", auth.Authorize())
	{
		//用户管理
		userModule.GET("/user/index", admin.User{}.Index)
		userModule.GET("/user/edit/:id", admin.User{}.Edit)
		userModule.GET("/user/add", admin.User{}.Add)
		userModule.GET("/user/delete/:id", admin.User{}.Delete)
		userModule.POST("/user/save", admin.User{}.Save)
	}

}
