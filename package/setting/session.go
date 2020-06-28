package setting

type Session struct {
	KeyPairs string `yaml:"key_pairs"` //秘钥
	Name     string `yaml:"name"`      //session名称
}
