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
	captchaVal := c.Request.FormValue("captcha_value")
	captchaKey := c.Request.FormValue("captcha_key")
	phone := c.Request.FormValue("phone")


}