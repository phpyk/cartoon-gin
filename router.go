package main

import (
	"cartoon-gin/auth"
	. "cartoon-gin/controllers"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/")

	authorize := router.Group("/api/auth")
	{
		authorize.POST("/login",LoginAction)
		authorize.GET("/me",auth.ValidateTokenV2(),CurrentUserAction);
	}
	home := router.Group("/api/home")
	{
		home.GET("/",GetHomeDataAction)
		home.GET("/more",GetMoreAction)
	}
	return router
}