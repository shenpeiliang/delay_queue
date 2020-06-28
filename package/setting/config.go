package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	MasterDB      *Database `yaml:"master_db"`
	SlaveDB       *Database `yaml:"slave_db"`
	Redis         *Redis    `yaml:"redis"`
	Server        *Server   `yaml:"server"`
	Session       *Session  `yaml:"session"`
	StaticVersion *Static   `yaml:"static_version"`
	Queue         *Queue    `yaml:"queue"`
}

var ConfigParam *Config // 配置信息

func init() {
	config, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("File Not Found: %v\n", err) //配置文件缺少，服务终止
	}
	ConfigParam = new(Config)
	err = yaml.Unmarshal(config, ConfigParam)
	if err != nil {
		log.Fatalf("错误Unmarsha1: %v\n", err)
	}

	//延迟队列配置
	ConfigParam.Queue.buildDelayConfig()

}
