package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moooner/mw-back/api"
	"github.com/moooner/mw-back/domain"
)

func main() {
	if domain.Err != nil {
		panic("DB connection failure.")
	}

	router := gin.Default()
	api.Router(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
