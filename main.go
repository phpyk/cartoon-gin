package main

import (
	"cartoon-gin/models"
)


func main() {
	//r := initRouter()
	//err := r.Run(":8080")
	//if err != nil {
	//	log.Fatal("failed to start gin",err)
	//}
	models.Migrate()

}
