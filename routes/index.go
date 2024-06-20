package routes

import (
	v1 "coupon-go-api/routes/v1"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1.SetupTaoBaoRoutes(router)
}
