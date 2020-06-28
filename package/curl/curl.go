package curl

/**
自动为请求添加cookie
可以获取curl命令请求记录
*/
import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Curl struct {
	cookies []*http.Cookie
	client  *http.Client

	url     string
	params  string
	headers map[string]string
	method  string
}

//响应信息
type ResponseData struct {
	Success bool `json:"success"`
}

//初始化
func New() (curl *Curl) {
	curl = &Curl{
		client: &http.Client{},
	}

	//为所有重定向的请求增加cookie
	curl.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > 0 {
			for _, v := range curl.GetCookie() {
				req.AddCookie(v)
			}
		}
		return nil
	}

	return
}

//GET请求
func (curl *Curl) Get(requestUrl string) ([]byte, int) {
	request, _ := http.NewRequest("GET", requestUrl, nil)

	//用于打印curl
	curl.method = "GET"
	curl.url = requestUrl

	curl.setRequestCookie(request)

	response, _ := curl.client.Do(request)

	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)

	return data, response.StatusCode

}

//POST请求
func (curl *Curl) Post(requestUrl string, params string) ([]byte, int) {
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(params))

	//用于打印curl
	curl.method = "POST"
	curl.url = requestUrl
	curl.setRequestHeader(request, "Content-Type", "application/x-www-form-urlencoded")

	//设置cookie
	curl.setRequestCookie(request)

	response, _ := curl.client.Do(request)

	defer response.Body.Close()

	//保存响应的 cookie
	curl.AddCookie(response.Cookies())

	data, _ := ioutil.ReadAll(response.Body)

	return data, response.StatusCode

}

//设置请求cookie
func (curl *Curl) setRequestCookie(request *http.Request) {
	for _, v := range curl.cookies {
		request.AddCookie(v)
	}
}

//设置请求cookie
func (curl *Curl) setRequestHeader(request *http.Request, key, value string) {
	curl.headers[key] = value
	request.Header.Set(key, value)
}

//设置请求cookie
func (curl *Curl) AddCookie(cookie []*http.Cookie) {
	curl.cookies = append(curl.cookies, cookie...)
}

//获取所有的cookie
func (curl *Curl) GetCookie() []*http.Cookie {
	return curl.cookies
}

//构建curl命令
func (curl *Curl) BuildCurlCmd() (cmd string) {
	cmd = "curl"
	if curl.method == "GET" {
		cmd += " -G"
	}

	header := ""
	for k, v := range curl.headers {
		header += " -H '" + k + ":" + v + "'"
	}

	var data string
	if curl.params == "" {
		data = " -d ''"
	} else {
		data = " -d " + curl.params
	}

	return fmt.Sprintf("%s %s %s %s", cmd, header, data, curl.url)
}

//URL参数串转换为集合
func Uri2Map(uri string) (m map[string]string, err error) {
	m = make(map[string]string)

	if len(uri) < 1 {
		return m, errors.New("空字符串")
	}

	//是否有包含？
	if uri[0:1] == "?" {
		uri = uri[1:]
	}

	//分解
	pars := strings.Split(uri, "&")

	for _, par := range pars {
		//分解
		parkv := strings.Split(par, "=")
		if len(parkv) == 2 {
			m[parkv[0]] = parkv[1]
		}
	}
	return
}

//map转url查询参数 encode
func EncodeParams(params map[string]string) string {
	paramsData := url.Values{}
	for k, v := range params {
		paramsData.Set(k, v)
	}
	//转码&排序
	return paramsData.Encode()
}
