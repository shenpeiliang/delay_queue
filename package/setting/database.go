package setting

type Database struct {
	Host               string `yaml:"host"`
	Port               string `yaml:"port"`
	DbName             string `yaml:"db_name"`
	DbUser             string `yaml:"db_user"`
	DbPwd              string `yaml:"db_pwd"`
	TablePrefix        string `yaml:"prefix"`
	DbCharset          string `yaml:"db_charset"`
	DbMaxOpenConns     string `yaml:"db_max_open_conns"`     // 连接池最大连接数
	DbMaxIdleCconns    string `yaml:"db_max_idle_conns"`     // 连接池最大空闲数
	DbMaxLifetimeConns string `yaml:"db_max_lifetime_conns"` // 连接池链接最长生命周期s
}
