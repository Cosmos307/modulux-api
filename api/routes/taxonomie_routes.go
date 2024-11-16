package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func TaxonomieRoutes(router *gin.Engine) {
	taxonomieGroup := router.Group("/zielqualifikation")
	{
		taxonomieGroup.POST("/bewertungsampel", controllers.GetTaxonomieFeedback)
		taxonomieGroup.GET("/verben/:kategorie", controllers.GetVerbsByCategory)

	}
}
