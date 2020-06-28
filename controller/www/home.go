package www

import (
	"github.com/gin-contrib/sessions"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Home struct {
}

func (home Home) Index(ctx *gin.Context) {

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

	//基于cookie的session使用
	session := sessions.Default(ctx)
	session.Set("hello", "world")
	session.Save()

	//设置cookie
	ctx.SetCookie("gin", "hello", 3600, "/", "127.0.0.1", false, true)

	//文件上传

	//密码加密

	//ctx.String(http.StatusOK, "hello %s %s %s %s %s", queryVal1, queryVal2, formVal3, formVal4, pathVal5)
	ctx.HTML(http.StatusOK, "www/home_index.html", gin.H{
		"msg": "Hello Go Developer",
		"now": time.Now().Format("2006-01-02 15:04:05"),
	})
}
