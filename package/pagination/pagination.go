package pagination

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//默认分页参数名
const PAGEPARAMNAME = "page"

type Pagination struct {
	Request       *http.Request
	PageParamName string //分页参数名
	Total         int    //总记录数
	PerNum        int    //每页显示数
}

//创建分页器
func New(request *http.Request, total, perNum int) *Pagination {
	return &Pagination{
		Request:       request,
		PageParamName: PAGEPARAMNAME,
		Total:         total,
		PerNum:        perNum,
	}
}

//设置分页参数名
func (p *Pagination) SetPageParamName(pageParamName string) {
	p.PageParamName = pageParamName
}

//获取分页链接
func (p *Pagination) getPageURL(page string) string {
	//基于当前url新建一个url对象
	u, _ := url.Parse(p.Request.URL.String())
	q := u.Query()
	q.Set("page", page)
	u.RawQuery = q.Encode()
	return u.String()
}

//生成html分页标签
func (p *Pagination) Page() string {
	//查询参数
	queryParams := p.Request.URL.Query()

	//获取分页参数
	page := queryParams.Get(p.PageParamName)

	if page == "" {
		page = "1"
	}

	var (
		totalDes string
		//当前分页
		currentPage int
		//分页总数
		totalPageNum int
		//首页链接
		firstLink string
		//上一页链接
		prevLink string
		//下一页链接
		nextLink string
		//末页链接
		lastLink string
		//中间页码链接
		pageLinks []string
	)

	//当前分页
	currentPage, _ = strconv.Atoi(page)

	//分页总数
	totalPageNum = p.Total / p.PerNum

	//总记录说明
	totalDes = fmt.Sprintf(`<span class="des-section">共<span class="total">%d</span>条记录，<span class="page">%d</span>页</span>`, p.Total, totalPageNum)

	//首页和上一页链接
	if currentPage > 1 {
		firstLink = fmt.Sprintf(`<a href="%s">首页</a>`, p.getPageURL("1"))
		prevLink = fmt.Sprintf(`<a href="%s">上一页</a>`, p.getPageURL(strconv.Itoa(currentPage-1)))
	} else {
		firstLink = `<a href="javascript:;" class="disabled">首页</a>`
		prevLink = `<a href="javascript:;" class="disabled">上一页</a>`
	}

	//末页和下一页
	if currentPage < totalPageNum {
		lastLink = fmt.Sprintf(`<a href="%s">末页</a>`, p.getPageURL(strconv.Itoa(totalPageNum)))
		nextLink = fmt.Sprintf(`<a href="%s">下一页</a>`, p.getPageURL(strconv.Itoa(currentPage+1)))
	} else {
		lastLink = `<a href="javascript:;" class="disabled">末页</a>`
		nextLink = `<a href="javascript:;" class="disabled">下一页</a>`
	}

	//生成中间页码链接
	pageLinks = make([]string, 0, 10)
	//中间页面只显示当前页的前N页和后N页
	startPos := currentPage - 5
	endPos := currentPage + 5
	if startPos < 1 {
		endPos = endPos + int(math.Abs(float64(startPos))) + 1
		startPos = 1
	}
	if endPos > totalPageNum {
		endPos = totalPageNum
	}
	for i := startPos; i <= endPos; i++ {
		var s string
		if i == currentPage {
			s = fmt.Sprintf(`<a href="%s" class="active">%d</a>`, p.getPageURL(strconv.Itoa(i)), i)
		} else {
			s = fmt.Sprintf(`<a href="%s">%d</a>`, p.getPageURL(strconv.Itoa(i)), i)
		}
		pageLinks = append(pageLinks, s)
	}

	var htmlDom string
	if totalPageNum > 1 {
		htmlDom = fmt.Sprintf(`<div class="pagination">%s%s%s%s%s%s</div>`, totalDes, firstLink, prevLink, strings.Join(pageLinks, ""), nextLink, lastLink)
	} else {
		htmlDom = fmt.Sprintf(`<div class="pagination">%s</div>`, totalDes)
	}
	return htmlDom
}
