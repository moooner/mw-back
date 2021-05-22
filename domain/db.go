package domain

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "moooner:dev123@tcp(127.0.0.1:3306)/mw?charset=utf8mb4&parseTime=True&loc=Local"
var Db, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
