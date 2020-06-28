package model

import (
	"errors"
	"gin/package/password"
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	Id       int64  `json:"id" form:"id" gorm:"primary_key"`
	IsSuper  byte   `json:"is_super"`
	Uname    string `json:"uname" form:"uname" binding:"required,notblank,username=2 20"`
	Email    string `json:"email" form:"email" binding:"required,email,notblank"`
	Password string `json:"password" form:"password" binding:"-"`
	RealName string `json:"real_name" form:"real_name" binding:"required,notblank"`
	Tel      string `json:"tel" form:"tel" binding:"required,mobile,notblank"`
	AddTime  int64  `json:"add_time" form:"add_time" binding:"-"`
}

//自定义错误
func (u *User) GetError(errs validator.ValidationErrors) string {
	for _, v := range errs {
		if "User.Uname" == v.Namespace() {
			switch v.Tag() {
			case "required":
				return "请输入用户名"
			case "username":
				return "请输入用户名：2-20个字符"
			}
		}

		if "User.Email" == v.Namespace() {
			switch v.Tag() {
			case "required":
				return "请输入邮箱地址"
			case "email":
				return "请输入正确的邮箱地址"
			}
		}

		if "User.RealName" == v.Namespace() {
			switch v.Tag() {
			case "required":
				return "请输入真实姓名"
			}
		}

		if "User.Tel" == v.Namespace() {
			switch v.Tag() {
			case "required":
				return "请输入手机号码"
			case "mobile":
				return "请输入正确的手机号码"
			}
		}

		if "User.Password" == v.Namespace() {
			switch v.Tag() {
			case "password":
				return "请输入合法的密码： 6-20个字符"
			}
		}
	}

	return "参数错误"
}

// 获取一条用户信息
func (u *User) GetOne() (user User, err error) {
	slaveDb.Where(u).Order("id desc").Limit(1).Find(&user)
	if err = slaveDb.Error; err != nil {
		return
	}

	return
}

//通过用户信息查询
func (u *User) GetByUser() (user User, err error) {
	slaveDb.Where(u).Order("id desc").Limit(1).Find(&user)
	if err = slaveDb.Error; err != nil {
		return
	}
	return
}

//验证用户登录
func (u *User) CheckLogin(user User) (ok bool, err error) {
	pwdBcrypt := password.New()
	ok, err = pwdBcrypt.ValidatePassword(user.Password, u.Password)
	return
}

//分页数据
func (u *User) GetPage(page int, pageSize int) (users []User, err error) {
	users = make([]User, 0)

	if page < 1 {
		page = 1
	}
	offset := pageSize * (page - 1)
	limit := pageSize

	db := slaveDb.Model(&User{})

	if u.Uname != "" {
		db = db.Where("uname LIKE ?", u.Uname+"%")
	}

	db.Limit(limit).Offset(offset).Find(&users)

	if err = db.Error; err != nil {
		return
	}

	return
}

//分页数
func (u *User) GetPageCount() (count int, err error) {
	db := slaveDb.Model(&User{})
	if u.Uname != "" {
		db = db.Where("uname LIKE ?", u.Uname+"%")
	}
	db.Model(&User{}).Count(&count)

	if err = db.Error; err != nil {
		return
	}

	return
}

//检查用户名是否存在
func (u *User) CheckUserName() (err error) {
	var user User

	db := slaveDb.Where("uname = ?", u.Uname)

	if u.Id > 0 {
		db = db.Where("id != ?", u.Id)
	}

	db.Order("id desc").Limit(1).Find(&user)

	if user.Id > 0 {
		return errors.New("用户已经存在，请重新填写")
	}

	return
}

// 添加
func (u *User) Add() (id int64, err error) {
	//密码
	bcrypt := password.Bcrypt{}
	var pwd []byte

	pwd, err = bcrypt.GeneratePassword(u.Password)
	if nil != err {
		return
	}
	u.Password = string(pwd)

	//添加时间
	u.AddTime = time.Now().Unix()

	db := masterDb.Create(u)
	if err = db.Error; err != nil {
		return
	}

	if user, ok := db.Value.(*User); ok {
		return user.Id, nil
	} else {
		return 0, errors.New("新增记录出错")
	}
}

// 更新
func (u *User) Update() (afr int64, err error) {
	var pwd []byte
	//密码
	if u.Password != "" {
		bcrypt := password.Bcrypt{}
		pwd, err = bcrypt.GeneratePassword(u.Password)
		if nil != err {
			return
		}
		u.Password = string(pwd)
	}

	db := masterDb.Save(u)
	if err = db.Error; err != nil {
		return
	}

	afr = db.RowsAffected

	return
}

// 删除
func (u *User) Delete() (afr int64, err error) {

	db := masterDb.Delete(u)
	if err = db.Error; err != nil {
		return
	}

	afr = db.RowsAffected

	return
}
