package model

import (
	"gin/driver"
	"github.com/jinzhu/gorm"
)

//Db数据库实例
var (
	masterDb *gorm.DB = driver.MasterDb
	slaveDb  *gorm.DB = driver.SlaveDb
)
