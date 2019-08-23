package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (*sql.DB,error) {
	db, err := sql.Open("mysql","root:123_QWE_asd@(127.0.0.1:3306)/gotest?parseTime=True")
	return db,err
}
