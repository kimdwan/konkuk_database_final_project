package services

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kimdwan/konkuk_database_final_project/AppFile/main_backend/entities/dtos"
)

func ParseAndCheckBody[T dtos.TableNumber](ctx *gin.Context) (*T, error) {
	var (
		body T
		err  error
	)

	if err = ctx.ShouldBindJSON(&body); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return nil, errors.New("(json) 클라이언트에서 보낸 폼을 파싱하는데 오류가 발생했습니다")
	}

	validate := validator.New()
	if err = validate.Struct(body); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return nil, errors.New("(validate) 클라이언트에서 보낸 폼을 파싱하는데 오류가 발생했습니다")
	}

	return &body, nil
}

func ConnectDb() (*sql.DB, error) {
	var (
		dsn string = os.Getenv("DATABASE_DSN")
		err error
	)

	if dsn == "" {
		return nil, errors.New("환경변수에 데이터 베이스 기본 설정을 해놓지 않았습니다")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return nil, errors.New("데이터 베이스에 연결하는데 오류가 발생했습니다")
	}

	return db, nil
}

func GetDatas(db *sql.DB, body *dtos.TableNumber, send_datas *[]dtos.MovieTable) (err error) {
	var (
		query_text string = "SELECT movie_name, movie_english_name, production_year, production_country, film_type, genre, production_status, director,production_company FROM movies\n"
		nums_first int    = (body.Table_number - 1) * 10
		nums_last  int    = body.Table_number * 10
	)

	add_text := "WHERE id <= " + strconv.Itoa(nums_last) + " AND id > " + strconv.Itoa(nums_first) + ";"
	query_text += add_text

	rows, err := db.Query(query_text)
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("데이터베이스에서 열을 가져오는데 오류가 발생했습니다")
	}

	defer rows.Close()

	for rows.Next() {
		var (
			send_data dtos.MovieTable
		)
		err = rows.Scan(&send_data.Movie_name, &send_data.Movie_english_name, &send_data.Production_year, &send_data.Production_country, &send_data.Film_type, &send_data.Genre, &send_data.Production_status, &send_data.Director, &send_data.Production_company)
		if err != nil {
			fmt.Println("시스템 오류: ", err.Error())
			return errors.New("행에서 데이터를 스캔하는데 오류가 발생했습니다")
		}

		*send_datas = append(*send_datas, send_data)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("데이터 베이스에서 해당 값을 찾는데 오류가 발생했습니다")
	}

	return nil
}

func CountDataBase(db *sql.DB, command_comment *string, total_numbers *int) (err error) {
	rows := db.QueryRow(*command_comment)
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("특정 조건에 맞는 쿼리를 찾는데 오류가 발생했습니다")
	}

	if err = rows.Scan(total_numbers); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("총갯수를 읽는데 오류가 발생했습니다")
	}

	return nil

}
