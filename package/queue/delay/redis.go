package delay

import (
	"encoding/json"
	"errors"
	"gin/driver"
	"gin/package/setting"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

//redis缓存键名
const DEFAULT_TASK_KEY = "delay_queue_task"

type RedisTask struct {
	NotifyUrl    string    `json:"notify_url"`
	PlanTime     time.Time `json:"plan_time"`
	MethodName   string    `json:"method_name"`
	NotifyParams string    `json:"notify_param"`
}

//获取队列键名
func (task RedisTask) GetKey() string {
	if setting.ConfigParam.Queue.Key == "" {
		return DEFAULT_TASK_KEY
	}
	return setting.ConfigParam.Queue.Key
}

//入队
func (task RedisTask) Push() (err error) {

	if task.NotifyUrl == "" {
		return errors.New("通知地址参数必须")
	}

	//如果任务计划执行时间小于当前时间
	if task.PlanTime.Before(time.Now()) {
		task.PlanTime = time.Now().Add(time.Second * time.Duration(setting.ConfigParam.Queue.TimeInterval))
	}

	//json格式话字符串
	params, err := json.Marshal(task)
	if err != nil {
		return
	}

	//redis客户端
	client, err := driver.RedisPool.Dial()
	if err != nil {
		return
	}

	//键名
	key := task.GetKey()

	//任务入队
	_, err = client.Do("rpush", key, params)

	return
}

//出队
func (task RedisTask) Pop() (t []string, err error) {
	//redis客户端
	client, err := driver.RedisPool.Dial()
	if err != nil {
		log.Printf(err.Error())
		return
	}

	//键名
	key := task.GetKey()

	//当前队列长度
	l, err := client.Do("llen", key)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	var taskLen int

	if v, ok := l.(int64); ok {
		taskLen = int(v)
	}

	//当前所有队列元素
	tasks, err := redis.Values(client.Do("lrange", key, 0, taskLen-1))

	if err != nil {
		log.Printf(err.Error())
		return
	}

	//字符串数组
	for _, item := range tasks {
		//追加结果集
		t = append(t, string(item.([]byte)))

		//删除redis元素，相当于弹出元素
		_, err = client.Do("lrem", key, 1, item)
		if err != nil {
			log.Printf(err.Error())
			return
		}
	}

	return
}
