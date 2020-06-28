package router

import (
	"gin/controller/admin"
	"gin/middleware/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutersAdmin(router *gin.Engine) {
	//登录
	router.GET("/admin/login/index", admin.Login{}.Index)
	router.GET("/admin/login/captcha", admin.Login{}.Captcha)
	router.POST("/admin/login/check", admin.Login{}.Check)

	//分组路由
	adminModule := router.Group("/admin", auth.Authorize())
	{
		//退出登录
		adminModule.GET("/login/quit", admin.Login{}.Quit)
		//用户管理
		adminModule.GET("/user/index", admin.User{}.Index)
		adminModule.GET("/user/edit/:id", admin.User{}.Edit)
		adminModule.GET("/user/add", admin.User{}.Add)
		adminModule.GET("/user/delete/:id", admin.User{}.Delete)
		adminModule.POST("/user/save", admin.User{}.Save)

		adminModule.GET("/log/index", admin.Log{}.Index)
		adminModule.GET("/log/detail/:name", admin.Log{}.Detail)
	}

}
