package main

import (
	. "cartoon-gin/controllers"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/")
	router.POST("/user/create", AddUserAction)
	router.GET("/user/get",QueryUserAction)

	router.GET("/category/all",GetAllAction)
	return router
}