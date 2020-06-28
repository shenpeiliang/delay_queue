package www

import (
	"gin/model"
	"gin/package/pagination"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Controller
}

//列表
func (user User) Index(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "20")
	keyword := ctx.DefaultQuery("keyword", "")
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	userModel := model.User{}

	if keyword != "" {
		userModel.Uname = strings.TrimSpace(keyword)
	}

	users, _ := userModel.GetPage(pageInt, pageSizeInt)

	//分页统计
	total, _ := userModel.GetPageCount()

	//分页标签
	pagination := pagination.New(ctx.Request, total, 10)

	ctx.HTML(http.StatusOK, "admin/user_index.html", gin.H{
		"data":   users,
		"config": user.getData(ctx),
		"page":   template.HTML(pagination.Page()),
	})

}

//保存
func (user User) Save(ctx *gin.Context) {
	userModel := model.User{}
	var err error
	//表单验证
	if err = ctx.ShouldBind(&userModel); nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": userModel.GetError(err.(validator.ValidationErrors)),
		})

		return
	}

	//如果密码不为空
	if userModel.Password != "" {
		if ok := user.getValidateExtend().IsPassword(userModel.Password, 6, 20); !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"code": "no",
				"data": "请输入密码：6-20个字符",
			})

			return
		}
	}

	//检查用户名
	err = userModel.CheckUserName()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": err.Error(),
		})

		return
	}

	//数据保存
	if userModel.Id > 0 {
		_, err = userModel.Update()
	} else {
		_, err = userModel.Add()
	}

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

//新增
func (user User) Add(ctx *gin.Context) {
	data := model.User{
		Id:       0,
		Uname:    "",
		Email:    "",
		RealName: "",
		Tel:      "",
		AddTime:  0,
	}

	ctx.HTML(http.StatusOK, "admin/user_form.html", gin.H{
		"data":   data,
		"config": user.getData(ctx),
	})
}

//编辑
func (user User) Edit(ctx *gin.Context) {
	//参数获取
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)

	if nil != err || 0 == idInt {
		ctx.Redirect(http.StatusMovedPermanently, "/admin/user/index")
		return
	}

	userModel := model.User{Id: idInt}
	data, err := userModel.GetOne()
	if nil != err {
		ctx.Redirect(http.StatusMovedPermanently, "/admin/user/index")
		return
	}

	ctx.HTML(http.StatusOK, "admin/user_form.html", gin.H{
		"data":   data,
		"config": user.getData(ctx),
	})

}

//删除
func (user User) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)

	if nil != err || 0 == idInt {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "no",
			"data": "resource identifier not found",
		})
		return
	}

	userModel := model.User{Id: idInt}

	//是否是超级管理员
	userData, err := userModel.GetOne()
	if nil != err {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": "删除失败：" + err.Error(),
		})
		return
	}

	if userData.IsSuper == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": "超级管理员不允许被删除",
		})
		return
	}

	afr, e := userModel.Delete()

	if nil != e {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": "删除失败：" + e.Error(),
		})
		return
	}

	if afr == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "no",
			"data": "删除失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": "ok",
		"data": "删除成功",
	})
}
