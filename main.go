package main

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mw-back/v1"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	routerV1 := router.Group("/v1")
	{
		routerV1.GET("weather", v1.Weather)
		routerV1.GET("weathers", v1.Weathers)
	}
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
