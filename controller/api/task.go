package api

import (
	"gin/package/queue/delay"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	Controller
}

//保存
func (task Task) Save(ctx *gin.Context) {
	//post表单
	notifyUrl := ctx.DefaultPostForm("notify_url", "")
	planTime := ctx.PostForm("plan_time")
	pt, err := strconv.ParseInt(planTime, 10, 64)

	if nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": "参数错误：" + err.Error(),
		})
		return
	}

	//把任务放大redis
	rt := delay.RedisTask{
		NotifyUrl: notifyUrl,
	}

	if pt > 0 {
		rt.PlanTime = time.Unix(pt, 0)
	}

	//入队
	err = rt.Push()
	if nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": err.Error(),
		})
		return
	}

	//成功返回
	ctx.JSON(http.StatusOK, gin.H{
		"code": "ok",
		"data": "操作成功",
	})
}
