package main

import (
	"cartoon-gin/auth"
	. "cartoon-gin/controllers"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/")
	router.POST("/auth/login",LoginAction)
	router.GET("/auth/me",CurrentUserAction, auth.ValidateTokenV2())
	return router
}