package setting

import (
	"strconv"
	"strings"
)

type Queue struct {
	Key              string   `yaml:"key"`                    //队列键名
	Slot             int      `yaml:"slot"`                   //队列槽位数
	TimeInterval     int      `yaml:"time_interval"`          // 如果加入队列时的计划时间小于当前时间，设置任务计划时间为当前时间之后的配置秒数
	MaxRetryNum      int      `yaml:"max_retry_num"`          // 最大重试次数
	RetryDefaultTime int      `yaml:"retry_default_time"`     // 重试次数默认间隔时间
	RetryTimeConfig  []string `yaml:"retry_time_config,flow"` // 重试时间配置
}

//延迟队列
var DelayConfig map[int]int

//延迟队列配置
func (queue *Queue) buildDelayConfig() {
	delay := make(map[int]int)

	//重试时间配置字符串拆分
	retryTimeConfig := queue.RetryTimeConfig
	for i := 0; i < len(retryTimeConfig); i++ {
		numTime := strings.Split(retryTimeConfig[i], "_")
		index, _ := strconv.Atoi(numTime[0])
		val, _ := strconv.Atoi(numTime[1])
		delay[index] = val
	}

	//填充默认值
	for i := 0; i < queue.MaxRetryNum; i++ {
		if _, exist := delay[i]; !exist {
			delay[i] = queue.TimeInterval
		}
	}

	DelayConfig = delay
}
