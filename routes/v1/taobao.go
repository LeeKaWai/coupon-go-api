package v1

import (
	"coupon-go-api/src/api"

	"github.com/gin-gonic/gin"
)

func SetupTaoBaoRoutes(router *gin.Engine) {
	taobaoGroup := router.Group("/v1/tb")
	{
		taobaoGroup.POST("/search", api.TaobaoSearch)
	}
}
