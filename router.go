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
		authorize.POST("login", LoginAction)
		authorize.GET("login-dev",LoginDevAccount)
		authorize.POST("pwd-login", PasswordLoginAction)
		authorize.POST("visitor-login", VisitorLoginAction)
		authorize.POST("logout",auth.ValidateJWTToken(),LogoutAction)
		authorize.GET("me", auth.ValidateJWTToken(), CurrentUserAction)
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
		common.POST("captcha",CaptchaAction)
		common.POST("captcha-check",CaptchaCheckAction)
		common.POST("verify-code", SendSMSVerifyCodeAction)
	}
	//分类tab
	router.GET("/api/cat/labels",CategoryLabelsAction)
	//漫画相关接口
	cartoon := router.Group("/api/cartoon/")
	{
		cartoon.GET("base-info",CartoonBaseInfoAction)
		cartoon.POST("search",CartoonSearchAction)
		cartoon.GET("chapter-list",CartoonChapterListAction)
		cartoon.GET("reading",auth.ValidateJWTToken(),CartoonReadAction)
		cartoon.POST("collect",auth.ValidateJWTToken(),CartoonCollectAction)
		cartoon.POST("buy-chapter",auth.ValidateJWTToken(),CartoonBuyAction)
	}
	//书架tab
	bookcase := router.Group("/api/user-bookcase/").Use(auth.ValidateJWTToken())
	{
		bookcase.GET("books",BookcaseTabsAction)
		bookcase.POST("delete",BookcaseDeleteAction)
	}
	//推荐
	router.GET("/api/recommend/get",auth.ValidateJWTToken(),GetRecommend)
	//修改资料
	profile := router.Group("/api/profile/").Use(auth.ValidateJWTToken())
	{
		profile.GET("base-info",ProfileBaseInfoAction)
		profile.POST("update",ProfileUpdateAction)
		profile.POST("bind-phone",UserBindPhoneAction)
		profile.POST("change-phone-verify",UserChangePhoneVerify)
		profile.POST("change-phone-submit",UserChangePhoneSubmit)
	}
	//获取App-config
	appConfig := router.Group("/api/config/")
	{
		appConfig.GET("reports",ConfigReportAction)
		appConfig.GET("vip-config",ConfigVipAction)
		appConfig.GET("coin-config",auth.ValidateJWTToken(),ConfigCoinAction)
		appConfig.GET("pay-channels",ConfigPayChannelsAction)
	}
	//举报
	router.POST("/api/report",auth.ValidateJWTToken(),ReportStoreAction)
	//用户反馈
	feedback := router.Group("/api/feedback/").Use(auth.ValidateJWTToken())
	{
		feedback.POST("save",FeedbackStoreAction)
		feedback.GET("list",FeedbackListAction)
	}
	//检查版本升级
	router.GET("/api/version/check",CheckAppUpdateAction)
	//获取最新接口
	router.GET("/api/creative",ConfigApiUrlAction)
	//创建订单、查询订单
	order := router.Group("/api/order/").Use(auth.ValidateJWTToken())
	{
		order.POST("create", OrderCreateAction)
		//order.GET("query", OrderQueryAction)
	}
	return router
}
