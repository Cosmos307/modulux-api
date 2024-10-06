package routes

import (
	"modulux/api/controller"

	"github.com/gin-gonic/gin"
)

func PersonRoutes(router *gin.Engine) {

	personGroup := router.Group("/persons")
	{
		personGroup.GET("/", controller.GetPersons)
		personGroup.GET("/:id", controller.GetPerson)
		personGroup.PUT("/:id", controller.UpdatePerson)

	}
}
