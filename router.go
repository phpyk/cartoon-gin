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
		authorize.POST("/logout",auth.ValidateJWTToken(),LogoutAction)
		authorize.GET("/me", auth.ValidateJWTToken(), CurrentUserAction)
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
		cartoon.POST("/search",CartoonSearchAction)
		cartoon.GET("/chapter-list",CartoonChapterListAction)
		cartoon.GET("/reading",auth.ValidateJWTToken(),CartoonReadAction)
		cartoon.POST("/collect",auth.ValidateJWTToken(),CartoonCollectAction)
		cartoon.POST("/buy-chapter",auth.ValidateJWTToken(),CartoonBuyAction)
	}
	//书架tab
	bookcase := router.Group("/api/user-bookcase")
	{
		bookcase.GET("/books",auth.ValidateJWTToken(),BookcaseTabsAction)
		bookcase.POST("/delete",auth.ValidateJWTToken(),BookcaseDeleteAction)
	}
	//推荐
	router.GET("/api/recommend/get",auth.ValidateJWTToken(),GetRecommend)

	return router
}
