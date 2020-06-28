package main

import (
	"gin/driver"
	"gin/package/queue/delay"
)

func main() {
	queue := delay.New()
	/*
		//添加任务
		planTime := time.Now().Add(time.Second * 10)

		task := delay.RedisTask{
			NotifyUrl: "http://sksystem.sk0.com/queue/vip",
			PlanTime:  planTime,
		}

		err := task.Push()
		if nil != err {
			fmt.Println("错误信息：", err.Error())
		}

		planTime = time.Now().Add(time.Second * 130)
		task.PlanTime = planTime

		err = task.Push()
		if nil != err {
			fmt.Println("错误信息：", err.Error())
		}*/

	//1小时候后关闭
	/*time.AfterFunc(time.Second*3600, func() {
		queue.Close()
	})*/

	// 服务停止时清理数据库链接
	defer driver.SlaveDb.Close()
	defer driver.MasterDb.Close()

	defer driver.RedisPool.Close()

	//启动队列服务
	queue.Start()
}
