package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func StudiengangRoutes(router *gin.Engine) {

	studiengangGroup := router.Group("/studiengaenge")
	{
		studiengangGroup.GET("/", controllers.GetStudiengaenge)
		studiengangGroup.GET("/:id", controllers.GetStudiengang)
		studiengangGroup.GET("/:id/modulverantwortliche", controllers.GetModulverantwortlicheByStudiengang)
		studiengangGroup.GET("/:id/opal-links", controllers.GetOpalLinks)
		studiengangGroup.GET("/:id/modul/zielqualifikationen", controllers.GetModuleGoalsByStudiengangID)
		studiengangGroup.PUT("/:id", controllers.UpdateStudiengang)
		studiengangGroup.POST("/", controllers.CreateStudiengang)
		studiengangGroup.DELETE("/:id", controllers.DeleteStudiengang)
	}
}
