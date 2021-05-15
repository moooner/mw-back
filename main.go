package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moooner/mw-back/APIs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:pass@tcp(127.0.0.1:3306)/mw?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	router := gin.Default()
	APIs.Router(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
