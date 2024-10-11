package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func ModulVoraussetzungRoutes(router *gin.Engine) {
	modulVoraussetzung := router.Group("/modul_voraussetzungen")
	{
		modulVoraussetzung.GET("/", controllers.GetAllModulVoraussetzungen)
		modulVoraussetzung.GET("/:modul_kuerzel/:modul_version", controllers.GetModulVoraussetzungen)
		modulVoraussetzung.POST("/", controllers.CreateModulVoraussetzung)
		modulVoraussetzung.DELETE("/:modul_kuerzel/:modul_version/:vorausgesetztes_modul_kuerzel/:vorausgesetztes_modul_version", controllers.DeleteModulVoraussetzung)
	}
}
