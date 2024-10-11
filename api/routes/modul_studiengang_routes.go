package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func ModulStudiengangRoutes(router *gin.Engine) {

	moduleGroup := router.Group("/module_studiengaenge")
	{
		moduleGroup.GET("/", controllers.GetAllModulStudiengang)
		moduleGroup.GET("/:modul_kuerzel/modul_version/studiengang_id ", controllers.GetSpecificModulStudiengang)
		moduleGroup.GET("/modul/:modul_kuerzel/:modul_version", controllers.GetModulStudiengangByModul)
		moduleGroup.GET("/studiengang/:studiengang_id", controllers.GetModulStudiengangByStudiengang)
		moduleGroup.POST("/", controllers.AddModulStudiengang)
	}
}
