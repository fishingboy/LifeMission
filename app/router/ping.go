package router

import (
	"github.com/fishingboy/LifeMission/app/controller"
	"github.com/gin-gonic/gin"
)

func PingBooking(r *gin.Engine) {
	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		controller.Ping(c)
	})
}
