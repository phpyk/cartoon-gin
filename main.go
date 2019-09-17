package main

import (
	"cartoon-gin/configs"
	"fmt"
	"log"
)

func main() {
	r := initRouter()
	address := fmt.Sprintf(":%v",configs.PORT)
	err := r.Run(address)
	if err != nil {
		log.Fatal("failed to start gin",err)
	}
}
