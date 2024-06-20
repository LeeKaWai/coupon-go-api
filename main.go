package main

import (
	"coupon-go-api/app"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engin := gin.New()
	app.SetupInit(engin, os.Getenv("ENV"))
}
