package admin

import (
	"gin/model"
	"gin/package/pagination"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Log struct {
	Controller
}

//列表
func (log Log) Index(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "20")
	keyword := ctx.DefaultQuery("keyword", "")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	taskModel := model.Task{}

	if keyword != "" {
		taskModel.QueueIndex = strings.TrimSpace(keyword)
	}

	tasks, _ := taskModel.GetPage(pageInt, pageSizeInt)

	//分页统计
	total, _ := taskModel.GetPageCount()

	//分页标签
	pagination := pagination.New(ctx.Request, total, 10)

	//活动菜单
	Menu.DtMenu = "style_30"
	Menu.DdMenu = "style_30_1"

	ctx.HTML(http.StatusOK, "admin/log_index.html", gin.H{
		"data":   tasks,
		"config": log.getData(ctx),
		"page":   template.HTML(pagination.Page()),
	})

}

//查看详情
func (log Log) Detail(ctx *gin.Context) {
	//参数获取
	name := ctx.Param("name")

	if name == "" {
		ctx.Redirect(http.StatusMovedPermanently, "/admin/log/index")
		return
	}

	taskLogModel := model.TaskLog{QueueIndex: name}
	data, err := taskLogModel.GetTaskLogs()
	if nil != err {
		ctx.Redirect(http.StatusMovedPermanently, "/admin/log/index")
		return
	}

	ctx.HTML(http.StatusOK, "admin/log_detail.html", gin.H{
		"data":   data,
		"config": log.getData(ctx),
	})

}
