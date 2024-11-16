package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func ModulStudiengangRoutes(router *gin.Engine) {

	moduleSGroup := router.Group("/module_studiengaenge")
	{
		moduleSGroup.GET("/", controllers.GetAllModulStudiengang)
		moduleSGroup.GET("/:modul_kuerzel/modul_version/studiengang_id ", controllers.GetSpecificModulStudiengang)
		moduleSGroup.GET("/modul/:modul_kuerzel/:modul_version", controllers.GetModulStudiengangByModul)
		moduleSGroup.GET("/studiengang/:studiengang_id", controllers.GetModulStudiengangByStudiengangID)
		moduleSGroup.POST("/", controllers.AddModulStudiengang)
	}
}
