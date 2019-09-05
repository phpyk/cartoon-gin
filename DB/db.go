package DB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func OpenBookDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql","root:hajgv8t24oA9@(123.206.107.76:3306)/books?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
		defer db.Close()
	}
	return db,err
}

func OpenCartoon() (*gorm.DB, error) {
	db, err :=  gorm.Open("mysql","root:hajgv8t24oA9@(123.206.107.76:3306)/cartoon?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
		defer db.Close()
	}
	return db,err
}