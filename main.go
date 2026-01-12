package main

import (
	"github.com/fishingboy/LifeMission/app/config"
	"github.com/fishingboy/LifeMission/app/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// 載入 .env
	config.LoadEnv()

	// 載入 router
	router.Booking(r)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	listenUrl := config.GetConfig("LISTEN_URL")
	r.Run(listenUrl)
}
