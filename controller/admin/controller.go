package admin

import (
	"gin/package/validator/extend"
	"github.com/gin-gonic/gin"
)

//公共结构体
type Controller struct{}

//活动菜单
type ActiveMenu struct {
	DtMenu string
	DdMenu string
}

var Menu ActiveMenu

//获取自定义验证器实例
func (c Controller) getValidateExtend() *extend.ValidatorExtend {
	return extend.New()
}

//获取通用数据
func (c Controller) getData(ctx *gin.Context) interface{} {
	return map[string]interface{}{
		"user":       ctx.MustGet("user"), //授权登录用户信息
		"activeMenu": Menu,                //活动菜单
	}
}
