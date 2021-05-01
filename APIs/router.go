package APIs

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/moooner/mw-back/APIs/v1"
)

func Router(router *gin.Engine) {
	routerV1 := router.Group("/v1")
	{
		routerV1.GET("weather", v1.Weather)
		routerV1.GET("weathers", v1.Weathers)
	}
}
