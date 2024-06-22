package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kimdwan/konkuk_database_final_project/AppFile/main_backend/pkgs/controllers"
)

func MainRouter(router *gin.Engine) {
	mainrouter := router.Group("main")
	mainrouter.POST("findalls", controllers.FindAllMovieDatas)
}
