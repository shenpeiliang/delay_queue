package extend

import (
	"regexp"
	"strconv"
)

//验证是否合法手机号码
func (valid ValidatorExtend) IsMobile(mobile string) bool {
	pattern := "^[(86)|0]?(1[2-9]\\d{9})$"
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(mobile)
}

//验证用户名 2-20
func (valid ValidatorExtend) IsUserName(uname string, minLength, maxLength int) bool {
	pattern := "^[\\x{4e00}-\\x{9fa5}A-Za-z]{1}[\\x{4e00}-\\x{9fa5}A-Za-z0-9_]{" + strconv.Itoa(minLength-1) + "," + strconv.Itoa(maxLength-1) + "}$"
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(uname)
}

//验证密码 6-20
func (valid ValidatorExtend) IsPassword(password string, minLength, maxLength int) bool {
	pattern := "^\\S{" + strconv.Itoa(minLength) + "," + strconv.Itoa(maxLength) + "}$"
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(password)
}

//是否全为指定字符
func (valid ValidatorExtend) IsAllCharacter(field string, charName string) bool {
	pattern := "^\\" + charName + "+$"
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(field)
}

//是否全为指定字符
func (valid ValidatorExtend) IsAllNumber(field string) bool {
	pattern := "^\\d+$"
	reg, _ := regexp.Compile(pattern)
	return reg.MatchString(field)
}
