package routes

import (
	"modulux/api/controllers"
	"modulux/api/middleware"

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
		moduleGroup.POST("/:kuerzel/:version/reset", controllers.ResetModuleToPreviousState)
		moduleGroup.DELETE("/:kuerzel/:version", controllers.DeleteModule)

		// endpoints with AuthMiddleware
		moduleGroup.PUT("/:kuerzel/:version", middleware.AuthMiddleware(), middleware.Authorize("modul_bearbeiten"), controllers.UpdateOrCreateModuleVersion)
		moduleGroup.GET("/roles", middleware.AuthMiddleware(), controllers.GetUserRoles)
	}
}
