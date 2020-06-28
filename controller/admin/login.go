package admin

import (
	"encoding/gob"
	"gin/model"
	"gin/package/vcode"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Login struct {
	Controller
}

// 登录
func (login Login) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/login_index.html", gin.H{})
}

// 退出
func (login Login) Quit(ctx *gin.Context) {
	//清除session
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()

	//跳转
	ctx.Redirect(http.StatusMovedPermanently, "/admin/login/index")

}

// 验证码
func (login Login) Captcha(ctx *gin.Context) {
	option := vcode.Option{
		CaptchaName: "admin_captcha",
		Len:         4,
		Width:       90,
		Height:      42,
		FileExt:     ".png",
		Lang:        "zh",
	}
	vcode.Create(ctx, option)
}

// 检查
func (login Login) Check(ctx *gin.Context) {
	uname := strings.TrimSpace(ctx.DefaultPostForm("uname", ""))
	pwd := strings.TrimSpace(ctx.DefaultPostForm("pwd", ""))
	code := strings.TrimSpace(ctx.DefaultPostForm("code", ""))

	if uname == "" || pwd == "" || code == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": "请填写必要信息",
		})
		return
	}

	//验证码
	option := vcode.Option{
		CaptchaName: "admin_captcha",
	}
	if ret := vcode.Verify(ctx, option, code); !ret {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": "验证码错误",
		})
		return
	}

	userModel := model.User{Uname: uname}
	//用户信息
	user, err := userModel.GetByUser()
	if nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": "用户名或密码错误，请重新输入",
		})
		return
	}

	//验证密码
	userModel.Password = pwd
	ok, err := userModel.CheckLogin(user)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": err.Error(),
		})
		return
	}

	//Register记录value下层具体值的类型和其名称，后面跨路由才可以获取到user结构体数据
	gob.Register(model.User{})

	//登录成功，设置SESSION
	session := sessions.Default(ctx)
	session.Set("user", user)
	session.Save()

	//返回成功
	ctx.JSON(http.StatusOK, gin.H{
		"code": "ok",
		"data": "登录成功",
		"url":  "/admin/user/index",
	})
}
