package helper

import (
	"fmt"
	"gin/package/setting"
	"html/template"
	"strings"
	"time"
)

var FunctionHelper template.FuncMap = map[string]interface{}{
	"formatAsDate":   formatAsDate,
	"unixTimeToDate": unixTimeToDate,
	"loadJs":         loadJs,
	"loadCss":        loadCss,
}

//加载js
func loadJs(files ...string) template.HTML {
	var html string
	for _, file := range files {
		html += "<script type='text/javascript' src='" + assetVersion(file, "js") + "'></script>\n"
	}
	return template.HTML(html)
}

//加载css
func loadCss(files ...string) template.HTML {
	var html string
	for _, file := range files {
		html += "<link type='text/css' rel='stylesheet' href='" + assetVersion(file, "css") + "'/>\n"
	}
	return template.HTML(html)
}

//静态文件版本号
func assetVersion(fileName, fileType string) (version string) {
	fileType = strings.ToLower(fileType)
	if fileType == "js" {
		version = fileName + "?v=" + setting.ConfigParam.StaticVersion.Js
	} else {
		version = fileName + "?v=" + setting.ConfigParam.StaticVersion.Css
	}
	return
}

//格式化时间
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

//时间戳转日期
func unixTimeToDate(t int64, format string) string {
	tm := time.Unix(t, 0)
	return tm.Format(format)
}
