package main

import (
	"fmt"
	"log"

	"cartoon-gin/configs"
)

func main() {
	//禁用控制台颜色，写入日志时开启此选项
	//gin.DisableConsoleColor()
	// 写入日志的文件
	//f,_ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果你需要同时写入日志文件和控制台上显示，使用下面代码
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := initRouter()
	address := fmt.Sprintf(":%v", configs.PORT)
	err := r.Run(address)
	if err != nil {
		log.Fatal("failed to start gin", err)
	}
}
