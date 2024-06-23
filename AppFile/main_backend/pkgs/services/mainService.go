package services

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kimdwan/konkuk_database_final_project/AppFile/main_backend/entities/dtos"
)

func ParseAndCheckBody[T dtos.TableNumber | dtos.FindMovieTableDto](ctx *gin.Context) (*T, error) {
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
	if rows.Err() != nil {
		fmt.Println("시스템 오류: ", rows.Err())
		return errors.New("특정 조건에 맞는 쿼리를 찾는데 오류가 발생했습니다")
	}

	if err = rows.Scan(total_numbers); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("총갯수를 읽는데 오류가 발생했습니다")
	}

	return nil

}

func FindWantMovieDatasService(db *sql.DB, body *dtos.FindMovieTableDto, send_datas *[]dtos.MovieTable, total_numbers *int) (err error) {
	var (
		find_text          string = "SELECT movie_name, movie_english_name, production_year, production_country, film_type, genre, production_status, director, production_company FROM movies"
		where_text         string
		find_total_numbers string = "SELECT COUNT(*) FROM movies"
		args               []interface{}
	)

	var where_clauses []string

	if body.Movie_name != "" {
		where_clauses = append(where_clauses, "movie_name LIKE ?")
		args = append(args, body.Movie_name+"%")
	}
	if body.Create_movie_year != 0 {
		where_clauses = append(where_clauses, "production_year = ?")
		args = append(args, body.Create_movie_year)
	}
	if body.Director != "" {
		where_clauses = append(where_clauses, "director LIKE ?")
		args = append(args, body.Director+"%")
	}

	if len(where_clauses) > 0 {
		where_text = " WHERE " + strings.Join(where_clauses, " AND ")
	}

	find_text += where_text + " LIMIT 10;"
	find_total_numbers += where_text

	// Prepare the query for fetching movies
	stmt, err := db.Prepare(find_text)
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("쿼리 준비 중 오류가 발생했습니다")
	}
	defer stmt.Close()

	// Execute the query
	rows, err := stmt.Query(args...)
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("해당 조건에 맞는 데이터를 찾는 중 오류가 발생했습니다")
	}
	defer rows.Close()

	for rows.Next() {
		var send_data dtos.MovieTable
		if err = rows.Scan(&send_data.Movie_name, &send_data.Movie_english_name, &send_data.Production_year, &send_data.Production_country, &send_data.Film_type, &send_data.Genre, &send_data.Production_status, &send_data.Director, &send_data.Production_company); err != nil {
			fmt.Println("시스템 오류: ", err.Error())
			return errors.New("쿼리 결과를 스캔하는 중 오류가 발생했습니다")
		}
		*send_datas = append(*send_datas, send_data)
	}

	if rows.Err() != nil {
		fmt.Println("시스템 오류: ", rows.Err())
		return errors.New("쿼리 결과 처리 중 오류가 발생했습니다")
	}

	// Query for total number of matching movies
	stmtTotal, err := db.Prepare(find_total_numbers)
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("총 갯수 쿼리 준비 중 오류가 발생했습니다")
	}
	defer stmtTotal.Close()

	if err := stmtTotal.QueryRow(args...).Scan(total_numbers); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("총 갯수를 가져오는 중 오류가 발생했습니다")
	}

	return nil
}
