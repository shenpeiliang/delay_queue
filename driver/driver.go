package driver

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

var (
	MasterDb *gorm.DB // db pool instance
	SlaveDb  *gorm.DB // db pool instance
	DbErr    error    // db err instance

	RedisPool *redis.Pool
)

func init() {
	//主库初始化
	initMasterDb()

	//从库初始化
	initSlaveDb()

	//Redis初始化
	initRedis()
}
