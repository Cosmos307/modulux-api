package routes

import (
	"modulux/api/controller"

	"github.com/gin-gonic/gin"
)

func ModulRoutes(router *gin.Engine) {

	moduleGroup := router.Group("/modules")
	{
		moduleGroup.GET("/", controller.GetModules)
		moduleGroup.GET("/:kuerzel/:version", controller.GetModule)
		moduleGroup.GET("opal-links", controller.GetOpalLinks)
		moduleGroup.GET("/:kuerzel/:version/opal-link", controller.GetOpalLink)
		moduleGroup.PUT("/:kuerzel/:version", controller.UpdateModule)
	}
}
