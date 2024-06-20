package app

import (
	"coupon-go-api/config"
	"coupon-go-api/db"
	"coupon-go-api/routes"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func SetupInit(eng *gin.Engine, env string) {
	if env == "" {
		env = "dev"
	}

	fmt.Println("env:", env)

	config.LoadConfig(env)

	routes.SetupRoutes(eng)

	db.InitDB()

	var wg sync.WaitGroup
	wg.Add(1)

	server := &http.Server{
		Addr:    ":" + config.GLOBAL_CONF.System.Port,
		Handler: eng,
	}
	go func() {
		defer wg.Done()
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Println("Server start err", err)
		}
	}()

	wg.Wait()
	fmt.Println("Server shutdown successfully")
}
