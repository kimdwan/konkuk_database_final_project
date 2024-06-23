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
		db              *sql.DB
		body            *dtos.TableNumber
		send_datas      []dtos.MovieTable
		command_comment string = `SELECT COUNT(*) AS "총갯수" FROM movies`
		total_numbers   int
		err             error
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

	err = services.CountDataBase(db, &command_comment, &total_numbers)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"send_datas":    send_datas,
		"total_numbers": total_numbers,
	})
}
