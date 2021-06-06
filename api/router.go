package api

import (
	"github.com/gin-gonic/gin"
	"github.com/moooner/mw-back/api/v1/user"
	"github.com/moooner/mw-back/api/v1/users"
)

func Router(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("users", users.Post)
		v1.GET("users", users.Get)
		v1.PUT("users/:idx", user.Put)
		v1.DELETE("users/:idx", user.Delete)
	}
}
