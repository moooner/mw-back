package v1

import (
	"github.com/gin-gonic/gin"
)

func Weather(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
		"message": "pong",
		"data": 26,
	})
}

func Weathers(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": 200,
		"message": "pong",
		"data": [3]int{26, 31, 25},
	})
}