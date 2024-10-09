package routes

import (
	"modulux/api/controller"

	"github.com/gin-gonic/gin"
)

func StudiengangRoutes(router *gin.Engine) {

	studiengangGroup := router.Group("/studiengaenge")
	{
		studiengangGroup.GET("/", controller.GetStudiengaenge)
		studiengangGroup.GET("/:id", controller.GetStudiengang)
		studiengangGroup.PUT("/:id/", controller.UpdateStudiengang)
		studiengangGroup.POST("/", controller.CreateStudiengang)
		studiengangGroup.DELETE("/:id", controller.DeleteStudiengang)
	}
}
