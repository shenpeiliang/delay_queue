package setting

type Server struct {
	Host           string `yaml:"host"`             //监听地址
	Port           string `yaml:"port"`             //监听端口
	ReadTimeout    uint16 `yaml:"read_timeout"`     //读超时 s
	WriteTimeout   uint16 `yaml:"write_timeout"`    //写超时 s
	MaxHeaderBytes int    `yaml:"max_header_bytes"` //最大请求头字节数
	ViewsPattern   string `yaml:"views_pattern"`    //模板路径
	Env            string `yaml:"env"`              //环境模式 release/debug/test
	LeftDelims     string `yaml:"left_delims"`      //模板渲染分隔符
	RightDelims    string `yaml:"right_delims"`     //模板渲染分隔符
}
