package controllers

import (
	"time"

	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func CaptchaAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	captchaId,captchaData := utils.GetNewCaptcha()
	outData := make(map[string]string)
	outData["key"] = captchaId
	outData["img"] = captchaData
	cg.Success(outData)
}

func CaptchaCheckAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	key := c.Request.FormValue("key")
	code := c.Request.FormValue("code")
	isRight := utils.CheckCaptcha(key,[]byte(code))
	if isRight {
		cg.Success(nil)
	} else {
		cg.Failed("图形验证码不正确")
		return
	}
}

func SendSMSVerifyCodeAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	phone := c.Request.FormValue("phone")
	if !utils.IsPhone(phone) {
		cg.Failed("手机号格式不正确")
		return
	}

	captchaVal := c.Request.FormValue("captcha_value")
	captchaKey := c.Request.FormValue("captcha_key")
	if !utils.CheckCaptcha(captchaKey, []byte(captchaVal)) {
		cg.Failed("图形验证码不正确")
		return
	}

	randomStr := utils.RandomString(4, 1)
	smsClient := utils.NewSmsClient()
	sendResult,err := smsClient.SendTemplateSMS(phone,[]string{randomStr,"5分钟"},1)
	utils.CheckError(err)

	if err == nil && sendResult["statusCode"] == "000000" {
		saveCodeToRedis(randomStr)
		cg.Success(nil)
	}else {
		cg.Failed("发送验证码失败")
		return
	}
}

func saveCodeToRedis(code string) bool {
	redisClient := utils.NewRedisClient()
	par := make(map[string]string)
	par["phone"] = code
	key := utils.GetRedisKey(utils.RDS_KEY_SMS_CODE,par)
	expireTime := 5 * time.Minute
	ok,err := redisClient.Set(key,code,expireTime).Result()
	utils.CheckError(err)
	if ok == "OK" {
		return true
	}

	return false
}