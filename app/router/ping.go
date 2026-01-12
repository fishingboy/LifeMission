package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingBooking(r *gin.Engine) {
	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
