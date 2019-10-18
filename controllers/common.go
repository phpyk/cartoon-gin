package controllers

import (
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

func SendVerifyCodeAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	phone := c.Request.FormValue("phone")
	if !utils.IsPhone(phone) {
		cg.Failed("手机号格式不正确")
	}

	captchaVal := c.Request.FormValue("captcha_value")
	captchaKey := c.Request.FormValue("captcha_key")
	if !utils.CheckCaptcha(captchaKey, []byte(captchaVal)) {
		cg.Failed("图形验证码不正确")
	}


}