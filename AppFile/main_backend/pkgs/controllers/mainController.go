package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllMovieDatas(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "모든 영화 데이터 여기 있습니다.",
	})
}
