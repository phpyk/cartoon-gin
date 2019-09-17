package DB

import (
	"cartoon-gin/configs"
	"fmt"
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
	connectionUrl := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=%v&parseTime=%v",
		configs.DB_USERNAME,
		configs.DB_PASSWORD,
		configs.DB_HOST,
		configs.DB_PORT,
		configs.DB_DATABASE,
		configs.DB_CHARSET,
		configs.DB_PARSE_TIME)

	db, err :=  gorm.Open(configs.DB_CONNECTION,connectionUrl)
	if err != nil {
		panic(err)
		defer db.Close()
	}
	return db,err
}