package delay

/**
环形队列 - 延迟重发
*/
import (
	"encoding/json"
	"errors"
	"fmt"
	"gin/model"
	"gin/package/curl"
	"gin/package/setting"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"strings"
	"time"
)

//队列结构元素
type Task struct {
	cycleNum int //执行任务的循环次数
	retryNum int //重试次数

	queueIndex string
	queueSlot  int

	notifyUrl    string    //请求通知地址
	planTime     time.Time //时间戳，指定要执行任务的计划时间
	methodName   string    //请求方法 post、get
	notifyParams string    //请求通知数据，请求查询字符串
}

//环形队列结构
type Queue struct {
	currentIndex int                //当前扫到的索引号
	slots        []map[string]*Task //队列中每个槽位的元素

	//信道
	closed        chan bool //关闭
	popTaskClose  chan bool //任务出队关闭
	pushTaskClose chan bool //任务入队关闭
	timeClose     chan bool //时间关闭
}

//创建队列
func New() (queue *Queue) {
	//初始化
	queue = &Queue{
		currentIndex: 0,
		closed:       make(chan bool),
		popTaskClose: make(chan bool),
		timeClose:    make(chan bool),
	}

	queue.slots = make([]map[string]*Task, setting.ConfigParam.Queue.Slot)

	//创建槽位
	for i := 0; i < setting.ConfigParam.Queue.Slot; i++ {
		queue.slots[i] = make(map[string]*Task)
	}

	//加载待处理的任务
	queue.loadTasks()

	return
}

//获取uuid
func GetUuid() string {
	return uuid.NewV4().String()
}

//构建任务数据
func (queue *Queue) buildTaskData(params string) (task *Task, err error) {
	task = &Task{
		cycleNum: 0,
		retryNum: 0,

		queueSlot: 0,

		notifyUrl:    "",
		planTime:     time.Time{},
		methodName:   "",
		notifyParams: "",
	}

	//json转换
	notifyData := RedisTask{}
	err = json.Unmarshal([]byte(params), &notifyData)
	if err != nil {
		return
	}

	//请求方法
	if notifyData.MethodName != "" {
		//转化为大写
		notifyData.MethodName = strings.ToUpper(notifyData.MethodName)
	} else {
		notifyData.MethodName = "GET"
	}

	//时间戳，指定要执行任务的计划时间
	if notifyData.PlanTime.Before(time.Now()) {
		notifyData.PlanTime = time.Now().Add(time.Second * time.Duration(setting.ConfigParam.Queue.TimeInterval))
	}

	task.notifyUrl = notifyData.NotifyUrl
	//转化为查询URL参数
	task.notifyParams = notifyData.NotifyParams
	task.methodName = notifyData.MethodName
	task.planTime = notifyData.PlanTime

	//队列索引
	task.queueIndex = GetUuid()

	//设置任务所在队列的槽位和循环次数
	queue.setTaskParam(task)

	//数据保存
	go func() {
		taskModel := model.Task{}
		taskModel.CycleNum = task.cycleNum
		taskModel.QueueIndex = task.queueIndex
		taskModel.QueueSlot = task.queueSlot

		taskModel.NotifyParam = params
		taskModel.NotifyUrl = notifyData.NotifyUrl
		taskModel.RequestMethod = notifyData.MethodName

		taskModel.Save()
	}()

	return task, nil

}

//设置任务所在队列的槽位和循环次数
func (queue *Queue) setTaskParam(task *Task) {
	//当前时间与指定时间相差秒数
	subSecond := task.planTime.Unix() - time.Now().Unix()

	//执行任务的循环次数
	cycleNum := int(subSecond / int64(setting.ConfigParam.Queue.Slot))
	task.cycleNum = cycleNum

	//计算任务所在的槽位
	task.queueSlot = int(subSecond % int64(setting.ConfigParam.Queue.Slot))
}

//添加任务 json字符串
func (queue *Queue) AddTask(params string) error {
	//任务数据
	task, err := queue.buildTaskData(params)
	if nil != err {
		return err
	}

	//加入任务队列
	err = queue.SaveQueueTask(task)
	if nil != err {
		return err
	}

	return nil
}

//加入任务队列
func (queue *Queue) SaveQueueTask(task *Task) error {
	key := task.queueIndex
	fmt.Println(fmt.Sprintf("key: %s  循环次数： %d 槽位： %d", key, task.cycleNum, task.queueSlot))

	//把任务加入队列
	tasks := queue.slots[task.queueSlot]
	if _, ok := tasks[key]; ok {
		return errors.New("该槽位中已存在key为" + key + "的任务")
	}

	tasks[key] = task
	return nil
}

//处理每1秒移动下标
func (queue *Queue) loopTime() {
	defer func() {
		fmt.Println("loopTime exit")
	}()
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-queue.timeClose: //收到关闭信号
			{
				return
			}
		case <-tick.C: //每秒执行移动下标
			{
				fmt.Println("loopTime:", time.Now().Format("2006-01-02 15:04:05"), "当前槽位:", queue.currentIndex)
				//判断当前下标 闭环
				if queue.currentIndex == (setting.ConfigParam.Queue.Slot - 1) {
					queue.currentIndex = 0
				} else {
					queue.currentIndex++
				}
			}
		}
	}
}

//处理每1秒的任务
func (queue *Queue) loopPopTask() {
	defer func() {
		fmt.Println("loopPopTask exit")
	}()
	//每隔1秒就向通道发送当前时间
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-queue.popTaskClose: //收到关闭信号
			{
				return
			}
		case <-tick.C:
			{
				//取出当前槽位的任务
				tasks := queue.slots[queue.currentIndex]
				if len(tasks) > 0 {
					//遍历任务，判断任务循环次数等于0，则运行任务,否则任务循环次数减1，等待下次执行
					for k, task := range tasks {
						if task.cycleNum == 0 {
							fmt.Println(k, "循环次数为0，执行任务", task.cycleNum)
							//异步协程处理
							go queue.execFunc(task)
							//删除任务
							delete(tasks, k)
						} else {
							fmt.Println(k, "当前循环次数", task.cycleNum)
							task.cycleNum--
							fmt.Println(k, "循环次数减一后", task.cycleNum)
						}
					}
				}
			}
		}
	}
}

//每秒加入任务到队列
func (queue *Queue) loopPushTask() {
	defer func() {
		fmt.Println("loopPushTask exit")
	}()
	//每隔1秒就向通道发送当前时间
	tick := time.NewTicker(time.Second)

	//redis任务
	rt := RedisTask{}

	for {
		select {
		case <-queue.pushTaskClose: //收到关闭信号
			{
				return
			}
		case <-tick.C:
			{
				//redis出队
				tasks, err := rt.Pop()

				if err != nil {
					log.Printf(err.Error())
				}

				if len(tasks) > 0 {
					for _, task := range tasks {
						//加入队列
						err = queue.AddTask(task)
						if err != nil {
							log.Printf(err.Error())
						}
					}
				}

			}
		}
	}

}

//启动队列
func (queue *Queue) Start() {
	//每秒从redis中加入任务到队列
	go queue.loopPushTask()

	//处理定时任务
	go queue.loopPopTask()

	//每秒滚动队列，改变
	go queue.loopTime()

	select {
	case <-queue.closed: //如果接受到关闭队列信号，那么其他任务需要接受到信号后终止
		{
			queue.popTaskClose <- true
			queue.pushTaskClose <- true
			queue.timeClose <- true
			break
		}
	}
}

//关闭队列
func (queue *Queue) Close() {
	queue.closed <- true
}

//异步发起请求第三方
func (queue *Queue) execFunc(task *Task) {
	c := curl.New()

	var data []byte
	var responseStatusCode int

	if task.methodName == "POST" {
		data, responseStatusCode = c.Post(task.notifyUrl, task.notifyParams)
	} else {
		data, responseStatusCode = c.Get(task.notifyUrl)
	}

	//记录curl命令
	curlCmd := c.BuildCurlCmd()

	//数据记录
	go func() {
		taskLog := model.TaskLog{}
		taskLog.QueueIndex = task.queueIndex
		taskLog.Curl = curlCmd
		taskLog.Response = fmt.Sprintf("状态码： %d 响应信息： %s", responseStatusCode, string(data))

		var response curl.ResponseData

		//请求状态是否成功
		var state = model.STATE_FAIL
		//循环次数
		var cycleNum = task.cycleNum + 1
		//重试次数
		var retryNum = task.retryNum + 1

		if responseStatusCode == http.StatusOK {
			err := json.Unmarshal(data, &response)

			if err != nil {
				log.Printf("json数据解析错误")
			}

			if response.Success {
				state = model.STATE_SUCCESS
				cycleNum = cycleNum - 1
				retryNum = retryNum - 1
			}
		}

		//新增记录
		taskLog.State = byte(state)
		taskLog.Save()

		//如果失败，可能需要重新执行任务
		if state == model.STATE_FAIL {
			//重试次数是否超出
			if retryNum > setting.ConfigParam.Queue.MaxRetryNum {
				//任务关闭
				taskLog.State = model.STATE_FAIL_FINISHED
			} else {
				//修改下次计划执行的时间
				task.planTime = time.Now().Add(time.Second * time.Duration(setting.DelayConfig[retryNum]))

				//设置任务所在队列的槽位和循环次数
				queue.setTaskParam(task)

				//任务重放
				queue.SaveQueueTask(task)
			}
		}

		//更新记录
		taskModel := model.Task{}
		taskModel.QueueIndex = task.queueIndex
		taskModel.Update(model.Task{
			State:    taskLog.State,
			CycleNum: cycleNum,
			RetryNum: retryNum,
		})

	}()

	log.Printf("%s %d %s", string(data), responseStatusCode, curlCmd)
}

//加载待处理的任务
func (queue *Queue) loadTasks() {
	//读取待处理的任务
	taskModel := model.Task{}

	data, err := taskModel.GetTasks()

	if err != nil {
		log.Printf(err.Error())
	}

	if len(data) > 0 {
		for _, item := range data {

			task := &Task{
				cycleNum:     item.CycleNum,
				retryNum:     item.RetryNum,
				queueIndex:   item.QueueIndex,
				queueSlot:    item.QueueSlot,
				notifyUrl:    item.NotifyUrl,
				methodName:   item.RequestMethod,
				notifyParams: item.NotifyParam,
			}
			//加入队列
			queue.SaveQueueTask(task)
		}
	}

}
