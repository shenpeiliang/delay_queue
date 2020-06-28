package extend

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type ValidatorExtend struct {
}

//创建实例
func New() *ValidatorExtend {
	return &ValidatorExtend{}
}

var (
	defaultMinLenthUserName = 2
	defaultMaxLenthUserName = 20

	defaultMinLenthPassWord = 6
	defaultMaxLenthPassWord = 20
)

//注册自定义验证器
func (valid *ValidatorExtend) RegisterValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//手机号
		v.RegisterValidation("mobile", New().mobile)
		//用户名
		v.RegisterValidation("username", New().userName)
		//密码
		v.RegisterValidation("password", New().password)

		//是否为空
		v.RegisterValidation("notblank", validators.NotBlank)

	}
}
