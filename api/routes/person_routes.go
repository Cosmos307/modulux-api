package routes

import (
	"modulux/api/controller"

	"github.com/gin-gonic/gin"
)

func PersonRoutes(router *gin.Engine) {

	personGroup := router.Group("/persons")
	{
		personGroup.GET("/", controller.GetPersons)
	}
}
