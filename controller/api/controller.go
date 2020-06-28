package api

import (
	"gin/package/validator/extend"
	"github.com/gin-gonic/gin"
)

//公共结构体
type Controller struct{}

//获取自定义验证器实例
func (c Controller) getValidateExtend() *extend.ValidatorExtend {
	return extend.New()
}

//获取通用数据
func (c Controller) getData(ctx *gin.Context) interface{} {
	return map[string]interface{}{
		"user": ctx.MustGet("user"), //授权登录用户信息
	}
}
