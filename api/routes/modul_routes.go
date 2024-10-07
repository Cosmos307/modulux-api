package routes

import (
	"modulux/api/controller"

	"github.com/gin-gonic/gin"
)

func ModulRoutes(router *gin.Engine) {

	moduleGroup := router.Group("/modules")
	{
		moduleGroup.GET("/", controller.GetModules)
	}
}
