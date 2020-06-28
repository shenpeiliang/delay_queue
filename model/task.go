package model

import (
	"errors"
	"gin/package/setting"
	"time"
)

const (
	STATE_WAIT          = 0 //待处理
	STATE_SUCCESS       = 1 //处理成功
	STATE_FAIL          = 2 //处理失败 - 待处理
	STATE_FAIL_FINISHED = 3 //处理失败 - 结束处理
)

//任务
type Task struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"`
	QueueIndex string `json:"queue_index" form:"queue_index"`
	QueueSlot  int    `json:"queue_slot" form:"queue_slot"`
	CycleNum   int    `json:"cycle_num" form:"cycle_num"`
	RetryNum   int    `json:"retry_num" form:"retry_num"`
	Dateline   int64  `json:"dateline" form:"dateline"`
	AddTime    int64  `json:"add_time" form:"add_time"`
	State      byte   `json:"state" form:"state"`

	NotifyUrl     string `json:"notify_url" form:"notify_url"`
	RequestMethod string `json:"request_method" form:"request_method"`
	NotifyParam   string `json:"notify_param" form:"notify_param"`
}

//任务记录返回
type TaskResult struct {
	Task
	StateDes string
}

//任务处理记录
type TaskLog struct {
	Id         int    `json:"id" form:"id" gorm:"primary_key"`
	QueueIndex string `json:"queue_index" form:"queue_index"`
	Response   string `json:"response" form:"response"`
	Curl       string `json:"curl" form:"curl"`
	Dateline   int64  `json:"dateline" form:"dateline"`
	State      byte   `json:"state" form:"state"`
}

//任务处理记录返回
type TaskLogResult struct {
	TaskLog
	StateDes string
}

//任务新增
func (t *Task) Save() (id int, err error) {
	timeNow := time.Now().Unix()
	t.AddTime = timeNow
	t.Dateline = timeNow

	db := masterDb.Create(t)
	if err = db.Error; err != nil {
		return
	}

	if task, ok := db.Value.(*Task); ok {
		return task.Id, nil
	} else {
		return 0, errors.New("新增记录出错")
	}
}

//任务更新
func (t *Task) Update(task Task) (afr int64, err error) {
	timeNow := time.Now().Unix()
	task.Dateline = timeNow

	db := masterDb.Model(&t).Where("queue_index = ?", t.QueueIndex).Updates(task)
	if err = db.Error; err != nil {
		return
	}

	afr = db.RowsAffected

	return
}

//任务操作记录新增
func (log *TaskLog) Save() (id int, err error) {
	timeNow := time.Now().Unix()
	log.Dateline = timeNow

	db := masterDb.Create(log)
	if err = db.Error; err != nil {
		return
	}

	if taskLog, ok := db.Value.(*TaskLog); ok {
		return taskLog.Id, nil
	} else {
		return 0, errors.New("新增记录出错")
	}
}

//获取待处理的任务
func (t *Task) GetTasks() (tasks []Task, err error) {
	tasks = make([]Task, 0)

	db := slaveDb.Where("state in (?)", []int{STATE_WAIT, STATE_FAIL}).Find(&tasks)

	if err = db.Error; err != nil {
		return
	}

	return
}

//分页数据
func (t *Task) GetPage(page int, pageSize int) (taskResult []TaskResult, err error) {
	taskResult = make([]TaskResult, 0)

	if page < 1 {
		page = 1
	}
	offset := pageSize * (page - 1)
	limit := pageSize

	db := slaveDb.Table(setting.ConfigParam.MasterDB.TablePrefix + "task")

	if t.QueueIndex != "" {
		db = db.Where("queue_index = ?", t.QueueIndex)
	}

	db.Limit(limit).Offset(offset).Order("id desc").Find(&taskResult)

	if err = db.Error; err != nil {
		return
	}

	//获取状态标题
	if len(taskResult) > 0 {
		for key, item := range taskResult {
			taskResult[key].StateDes = GetStateDes(item.State)
		}
	}

	return
}

//获取状态标题
func GetStateDes(state byte) (des string) {
	states := map[byte]string{
		STATE_WAIT:          "待处理",
		STATE_SUCCESS:       "处理成功",
		STATE_FAIL:          "处理失败 - 待处理",
		STATE_FAIL_FINISHED: "处理失败 - 结束处理",
	}
	if val, ok := states[state]; ok {
		des = val
	}
	return
}

//分页数
func (t *Task) GetPageCount() (count int, err error) {
	db := slaveDb.Model(&User{})

	if t.QueueIndex != "" {
		db = db.Where("queue_index = ?", t.QueueIndex)
	}

	db.Table(setting.ConfigParam.MasterDB.TablePrefix + "task").Count(&count)

	if err = db.Error; err != nil {
		return
	}

	return
}

//任务执行记录
func (log *TaskLog) GetTaskLogs() (taskLogResult []TaskLogResult, err error) {
	taskLogResult = make([]TaskLogResult, 0)

	slaveDb.Table(setting.ConfigParam.MasterDB.TablePrefix + "task_log").Where(log).Order("id desc").Find(&taskLogResult)

	if err = masterDb.Error; err != nil {
		return
	}

	//获取状态标题
	if len(taskLogResult) > 0 {
		for key, item := range taskLogResult {
			taskLogResult[key].StateDes = GetStateDes(item.State)
		}
	}

	return
}
