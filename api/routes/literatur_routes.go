package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func LiteratureRoutes(router *gin.Engine) {
	literatureGroup := router.Group("/literatur")
	{
		literatureGroup.GET("/test", controllers.TestCrossRefConnection)
		literatureGroup.GET("/vorschläge", controllers.GetLiteratureSuggestions)
		literatureGroup.POST("/bestätigen", controllers.ConfirmLiteratureSelection)
	}
}
