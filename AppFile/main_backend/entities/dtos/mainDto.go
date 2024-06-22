package dtos

import "database/sql"

type TableNumber struct {
	Table_number int `json:"table_number" validate:"number,min=1"`
}

type MovieTable struct {
	Movie_name         string         `json:"movie_name"`
	Movie_english_name sql.NullString `json:"movie_english_name,omitempty"`
	Production_year    sql.NullInt64  `json:"production_year,omitempty"`
	Production_country sql.NullString `json:"production_country,omitempty"`
	Film_type          sql.NullString `json:"film_type,omitempty"`
	Genre              sql.NullString `json:"genre,omitempty"`
	Production_status  sql.NullString `json:"production_status,omitempty"`
	Director           sql.NullString `json:"director"`
	Production_company sql.NullString `json:"production_company"`
}
