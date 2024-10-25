package routes

import (
	"modulux/api/controllers"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
}
