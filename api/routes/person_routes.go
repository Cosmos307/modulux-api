package routes

import (
	"modulux/api/controllers"
	"modulux/api/middleware"

	"github.com/gin-gonic/gin"
)

func PersonRoutes(router *gin.Engine) {
	personGroup := router.Group("/personen")

	// GET endpoints without AuthMiddleware
	personGroup.GET("/", controllers.GetPersons)
	personGroup.GET("/:id", controllers.GetPerson)

	// PUT, POST, DELETE endpoints with AuthMiddleware
	auth := personGroup.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.PUT("/:id", controllers.UpdatePerson)
	auth.POST("/", controllers.CreatePerson)
	auth.DELETE("/:id", controllers.DeletePerson)
}
