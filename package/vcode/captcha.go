package vcode

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//创建图像验证码
func Create(c *gin.Context, option Option) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")

	if option.Len == 0 {
		option.Len = captcha.DefaultLen
	}

	//设置长度
	captchaId := captcha.NewLen(option.Len)

	//缓存
	session := sessions.Default(c)
	session.Set(option.CaptchaName, captchaId)
	session.Save()

	var content bytes.Buffer

	switch option.FileExt {
	case ".png":
		c.Writer.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, captchaId, option.Width, option.Height)
	case ".wav":
		c.Writer.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, captchaId, option.Lang)
	default:
		log.Println(captcha.ErrNotFound)
	}
	http.ServeContent(c.Writer, c.Request, captchaId+option.FileExt, time.Time{}, bytes.NewReader(content.Bytes()))
}

//验证图像验证码
func Verify(c *gin.Context, option Option, code string) bool {
	//缓存
	session := sessions.Default(c)
	if captchaId := session.Get(option.CaptchaName); captchaId != nil {
		//删除缓存
		session.Delete(option.CaptchaName)
		session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		}
	}

	return false
}
