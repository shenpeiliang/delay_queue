package driver

import (
	"fmt"
	"gin/package/setting"
	"github.com/gomodule/redigo/redis"
	"time"
)

func initRedis() {
	var redisServer string

	//连接地址
	redisServer = fmt.Sprintf("%s:%s", setting.ConfigParam.Redis.Host, setting.ConfigParam.Redis.Port)

	RedisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			//数据库
			if setting.ConfigParam.Redis.DbName > 0 {
				redis.DialDatabase(setting.ConfigParam.Redis.DbName)
			}

			//密码验证
			if setting.ConfigParam.Redis.DbPwd != "" {
				redis.DialPassword(setting.ConfigParam.Redis.DbPwd)
			}

			//连接
			client, err := redis.Dial("tcp", redisServer)
			if err != nil {
				panic("redis error: " + err.Error())
			}

			return client, err
		},
		MaxIdle:     setting.ConfigParam.Redis.DbMaxOpenConns,                    //最初的连接数量
		MaxActive:   setting.ConfigParam.Redis.DbMaxIdleCconns,                   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: time.Duration(setting.ConfigParam.Redis.DbMaxLifetimeConns), //连接关闭时间 （>0不使用自动关闭）
	}
}
