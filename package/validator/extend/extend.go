package extend

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	splitParamsRegexString = `'[^']*'|\S+`
)

var (
	oneofValsCache       = map[string][]string{}
	oneofValsCacheRWLock = sync.RWMutex{}
	splitParamsRegex     = regexp.MustCompile(splitParamsRegexString)
)

//多参数
func parseOneOfParam2(s string) []string {
	oneofValsCacheRWLock.RLock()
	vals, ok := oneofValsCache[s]
	oneofValsCacheRWLock.RUnlock()
	if !ok {
		oneofValsCacheRWLock.Lock()
		vals = splitParamsRegex.FindAllString(s, -1)
		for i := 0; i < len(vals); i++ {
			vals[i] = strings.Replace(vals[i], "'", "", -1)
		}
		oneofValsCache[s] = vals
		oneofValsCacheRWLock.Unlock()
	}
	return vals
}

//验证是否合法手机号码
func (valid ValidatorExtend) mobile(f validator.FieldLevel) bool {
	return ValidatorExtend{}.IsMobile(f.Field().String())
}

//验证用户名
func (valid ValidatorExtend) userName(f validator.FieldLevel) bool {
	vals := parseOneOfParam2(f.Param())
	minLength := defaultMinLenthUserName
	maxLength := defaultMaxLenthUserName

	if len(vals) == 2 {
		min, err := strconv.Atoi(vals[0])
		if nil == err {
			minLength = min
		}
		max, err := strconv.Atoi(vals[1])
		if nil == err {
			maxLength = max
		}
	}
	return ValidatorExtend{}.IsUserName(f.Field().String(), minLength, maxLength)
}

//验证密码
func (valid ValidatorExtend) password(f validator.FieldLevel) bool {
	vals := parseOneOfParam2(f.Param())
	minLength := defaultMinLenthPassWord
	maxLength := defaultMaxLenthPassWord

	if len(vals) == 2 {
		min, err := strconv.Atoi(vals[0])
		if nil == err {
			minLength = min
		}
		max, err := strconv.Atoi(vals[1])
		if nil == err {
			maxLength = max
		}
	}

	return ValidatorExtend{}.IsPassword(f.Field().String(), minLength, maxLength)
}
