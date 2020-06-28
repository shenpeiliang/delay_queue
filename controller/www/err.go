package www

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Err struct {
}

//404错误
func (e Err) Show404(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "common/404.html", gin.H{
		"data": "抱歉！页面无法访问……",
	})
}

func (e Err) Com(ctx *gin.Context) {
	httpStatus, err := strconv.Atoi(ctx.Param("id"))
	if nil != err {
		log.Fatalf("参数错误: %s\n", err)
	}

	if httpStatus == http.StatusNotFound {
		ctx.HTML(http.StatusOK, "common/404.html", gin.H{
			"data": "抱歉！页面无法访问……",
		})
	} else if httpStatus == http.StatusInternalServerError {
		ctx.HTML(http.StatusOK, "common/404.html", gin.H{
			"data": "抱歉！系统出错了……",
		})
	}

}
