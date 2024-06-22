package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kimdwan/konkuk_database_final_project/AppFile/main_backend/entities/dtos"
	"github.com/kimdwan/konkuk_database_final_project/AppFile/main_backend/pkgs/services"
)

func FindAllMovieDatas(ctx *gin.Context) {
	var (
		db         *sql.DB
		body       *dtos.TableNumber
		send_datas []dtos.MovieTable
		err        error
	)

	body, err = services.ParseAndCheckBody[dtos.TableNumber](ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db, err = services.ConnectDb()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer db.Close()

	err = services.GetDatas(db, body, &send_datas)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, send_datas)
}
