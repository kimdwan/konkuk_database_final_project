package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kimdwan/konkuk_database_final_project/AppFile/main_backend/middlewares"
	"github.com/kimdwan/konkuk_database_final_project/AppFile/main_backend/pkgs/routes"
	"github.com/kimdwan/konkuk_database_final_project/AppFile/main_backend/settings"
)

func init() {
	settings.LoadDotenv()
}

func main() {
	var (
		port string = os.Getenv("GO_PORT")
	)

	if port == "" {
		panic("환경변수에 port 번호를 입력하지 않았습니다.")
	}

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(middlewares.CorsMiddleware())
	routes.MainRouter(router)
	router.Run(port)
}
