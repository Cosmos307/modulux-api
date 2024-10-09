package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func ModulStudiengangRoutes(router *gin.Engine) {

	moduleGroup := router.Group("/module_studiengaenge")
	{
		moduleGroup.GET("/", controllers.GetAllModulStudiengang)
	}
}
