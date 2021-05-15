package main

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/moooner/mw-back/main/v1"
)

func Router(router *gin.Engine) {
	routerV1 := router.Group("/v1")
	{
		routerV1.POST("signUp", v1.SignUp)
		routerV1.GET("signIn", v1.SignIn)
	}
}
