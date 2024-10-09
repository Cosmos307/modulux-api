package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func ModulRoutes(router *gin.Engine) {

	moduleGroup := router.Group("/modules")
	{
		moduleGroup.GET("/", controllers.GetModules)
		moduleGroup.GET("/:kuerzel/:version", controllers.GetModule)
		moduleGroup.GET("opal-links", controllers.GetOpalLinks)
		moduleGroup.GET("/:kuerzel/:version/opal-link", controllers.GetOpalLink)
		moduleGroup.PUT("/:kuerzel/:version", controllers.UpdateModule)
		moduleGroup.POST("/", controllers.CreateModule)
		moduleGroup.DELETE("/:kuerzel/:version", controllers.DeleteModule)
	}
}
