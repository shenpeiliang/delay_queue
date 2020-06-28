package auth

import (
	"encoding/gob"
	"gin/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Register记录value下层具体值的类型和其名称，后面跨路由才可以获取到user结构体数据
		gob.Register(model.User{})

		session := sessions.Default(ctx)
		//session数据
		user := session.Get("user")
		//fmt.Println("user信息:", user)
		if user == nil {
			ctx.Abort()
			ctx.Redirect(http.StatusMovedPermanently, "/admin/login/index")
			/*c.JSON(http.StatusUnauthorized, gin.H{
				"code": "ERROR-AUTH",
				"data": "访问未授权",
			})*/
			return
		}
		//设置变量
		ctx.Set("user", user)
		ctx.Next()
	}
}
