package models

import (
	"cartoon-gin/common"
	"log"
)

func Migrate() {
	db,err := OpenBookDB()
	common.CheckError(err)

	obj1 := Book{}
	if !db.HasTable(&obj1) {
		db.AutoMigrate(&obj1)
		log.Println("created -- Book")
	}

	obj2 := Chapter{}
	if !db.HasTable(&obj2) {
		db.AutoMigrate(&obj2)
		log.Println("created -- Chapter")
	}
}