package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Load person routes
	PersonRoutes(r)
	ModulRoutes(r)

	// Ping route for testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
