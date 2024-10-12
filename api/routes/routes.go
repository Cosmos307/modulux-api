package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Routes
	PersonRoutes(r)
	ModulRoutes(r)
	StudiengangRoutes(r)
	ModulStudiengangRoutes(r)
	ModulVoraussetzungRoutes(r)
	TaxonomieRoutes(r)

	// Ping route for testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
