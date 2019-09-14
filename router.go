package main

import (
	"cartoon-gin/auth"
	. "cartoon-gin/controllers"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/")

	router.POST("/api/auth/login",LoginAction)
	router.GET("/api/auth/me",auth.ValidateTokenV2(),CurrentUserAction);
	return router
}