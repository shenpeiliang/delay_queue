package main

import (
	"gin/driver"
	_ "gin/driver"
	"gin/server"
)

func main() {

	// 服务停止时清理数据库链接
	defer driver.SlaveDb.Close()
	defer driver.MasterDb.Close()

	defer driver.RedisPool.Close()

	// 启动服务
	server.Run()

}
