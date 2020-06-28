package setting

type Redis struct {
	Host               string `yaml:"host"`
	Port               string `yaml:"port"`
	DbName             int    `yaml:"db_name"`
	DbPwd              string `yaml:"db_pwd"`
	DbMaxOpenConns     int    `yaml:"db_max_open_conns"`     // 连接池初始化最大连接数
	DbMaxIdleCconns    int    `yaml:"db_max_idle_conns"`     // 连接池最大空闲数
	DbMaxLifetimeConns int    `yaml:"db_max_lifetime_conns"` // 连接池链接最长生命周期s
}
