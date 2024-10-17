package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func ModulRoutes(router *gin.Engine) {

	moduleGroup := router.Group("/modul")
	{
		moduleGroup.GET("/", controllers.GetModules)
		moduleGroup.GET("/:kuerzel/:version", controllers.GetModule)
		moduleGroup.GET("opal-links", controllers.GetOpalLinks)
		moduleGroup.GET("/:kuerzel/:version/opal-link", controllers.GetOpalLink)
		moduleGroup.POST("/", controllers.CreateModule)
		moduleGroup.PUT("/:kuerzel/:version", controllers.UpdateOrCreateModuleVersion)
		moduleGroup.DELETE("/:kuerzel/:version", controllers.DeleteModule)

	}
}
