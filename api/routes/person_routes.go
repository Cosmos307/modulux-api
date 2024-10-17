package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func PersonRoutes(router *gin.Engine) {

	personGroup := router.Group("/personen")
	{
		personGroup.GET("/", controllers.GetPersons)
		personGroup.GET("/:id", controllers.GetPerson)
		personGroup.PUT("/:id", controllers.UpdatePerson)
		personGroup.POST("/", controllers.CreatePerson)
		personGroup.DELETE("/:id", controllers.DeletePerson)
	}
}
