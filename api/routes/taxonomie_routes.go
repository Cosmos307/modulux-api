package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func TaxonomieRoutes(router *gin.Engine) {
	taxonomieGroup := router.Group("/zielqualifikation")
	{
		taxonomieGroup.POST("/", controllers.GetTaxonomieFeedback)
		taxonomieGroup.GET("/verben/:category", controllers.GetVerbsByCategory)

	}
}
