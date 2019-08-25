package models

import (
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func OpenDB() (*sql.DB,error) {
	db, err := sql.Open("mysql","root:hajgv8t24oA9@(123.206.107.76:3306)/news?parseTime=True&charset=utf8mb4&loc=Local")
	return db,err
}

func OpenNewsDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql","root:hajgv8t24oA9@(123.206.107.76:3306)/news?charset=utf8mb4&parseTime=True")
	return db,err
}