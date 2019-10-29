package main

import (
	"cartoon-gin/auth"
	"cartoon-gin/configs"
	. "cartoon-gin/controllers"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	gin.SetMode(configs.GIN_MODE)
	router := gin.Default()

	router.GET("/")
	authorize := router.Group("/api/auth")
	{
		authorize.POST("/login", LoginAction)
		authorize.GET("/login-dev",LoginDevAccount)
		authorize.POST("/pwd-login", PasswordLoginAction)
		authorize.POST("/visitor-login", VisitorLoginAction)
		authorize.POST("/logout",auth.ValidateRedisToken(),LogoutAction)
		authorize.GET("/me", auth.ValidateRedisToken(), CurrentUserAction)
	}

	home := router.Group("/api/home")
	{
		home.GET("/", GetHomeDataAction)
		home.GET("/more", GetMoreAction)
		home.GET("/ranking",GetRankAction)
		home.GET("/new-cartoons",GetNewCartoonsAction)
	}
	common := router.Group("/api/common")
	{
		common.POST("/captcha",CaptchaAction)
		common.POST("/captcha-check",CaptchaCheckAction)
		common.POST("/verify-code", SendSMSVerifyCodeAction)
	}
	return router
}
