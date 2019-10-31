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
	//登录、退出、游客登录
	authorize := router.Group("/api/auth")
	{
		authorize.POST("/login", LoginAction)
		authorize.GET("/login-dev",LoginDevAccount)
		authorize.POST("/pwd-login", PasswordLoginAction)
		authorize.POST("/visitor-login", VisitorLoginAction)
		authorize.POST("/logout",auth.ValidateRedisToken(),LogoutAction)
		authorize.GET("/me", auth.ValidateRedisToken(), CurrentUserAction)
	}
	//首页tab
	home := router.Group("/api/home")
	{
		home.GET("/", GetHomeDataAction)
		home.GET("/more", GetMoreAction)
		home.GET("/ranking",GetRankAction)
		home.GET("/new-cartoons",GetNewCartoonsAction)
	}
	//通用功能（图片验证码、手机验证码）
	common := router.Group("/api/common")
	{
		common.POST("/captcha",CaptchaAction)
		common.POST("/captcha-check",CaptchaCheckAction)
		common.POST("/verify-code", SendSMSVerifyCodeAction)
	}
	//分类tab
	router.GET("/api/cat/labels",CategoryLabelsAction)
	//漫画相关接口
	cartoon := router.Group("/api/cartoon")
	{
		cartoon.GET("/base-info",CartoonBaseInfoAction)
	}
	return router
}
