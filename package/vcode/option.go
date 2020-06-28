package vcode

type Option struct {
	CaptchaName string //保存的缓存名称
	FileExt     string //文件后缀
	Lang        string //验证码语言
	Len         int    //字符长度
	Width       int    //图像宽度
	Height      int    //图像高度
}
