package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moooner/mw-back/APIs"
)

func main() {
	router := gin.Default()
	APIs.Router(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
