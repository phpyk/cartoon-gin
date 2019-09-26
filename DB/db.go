package DB

import (
	"cartoon-gin/configs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func OpenBookDB() (*gorm.DB, error) {
	conn := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=%v&parseTime=%v",
		configs.DB_USERNAME,
		configs.DB_PASSWORD,
		configs.DB_HOST,
		configs.DB_PORT,
		"books",
		configs.DB_CHARSET,
		configs.DB_PARSE_TIME)
	db, err := gorm.Open(configs.DB_CONNECTION,conn)
	if err != nil {
		panic(err)
		defer db.Close()
	}
	return db,err
}

func OpenCartoon() (*gorm.DB, error) {
	conn := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=%v&parseTime=%v",
		configs.DB_USERNAME,
		configs.DB_PASSWORD,
		configs.DB_HOST,
		configs.DB_PORT,
		"cartoon",
		configs.DB_CHARSET,
		configs.DB_PARSE_TIME)

	db, err :=  gorm.Open(configs.DB_CONNECTION,conn)
	if err != nil {
		panic(err)
		defer db.Close()
	}
	return db,err
}