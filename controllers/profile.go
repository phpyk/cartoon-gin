package controllers

import (
	"errors"
	"fmt"
	"strings"

	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func ProfileBaseInfoAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	user := CurrentUser(c)

	responseData := make(map[string]interface{})
	responseData["base_info"] = user
	cg.Success(responseData)
}

func ProfileUpdateAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	user := CurrentUser(c)

	nickName := utils.FilterSpecialChar(strings.Trim(c.Request.FormValue("nick_name")," "))
	if nickName == "" {
		cg.Failed("昵称不能为空")
		return
	}
	if len(nickName) < 2 || len(nickName) > 10 {
		cg.Failed("昵称长度为2到10个字符")
		return
	}
	existUser := dao.UserFindByNickName(nickName)
	if existUser.ID > 0 && existUser.ID != user.ID {
		cg.Failed("昵称已被占用，请换一个吧~")
		return
	}
	updateData := make(map[string]string)
	updateData["nick_name"] = nickName
	dao.UpdateUser(user.ID,updateData)
	cg.Success(nil)
}

func UserBindPhoneAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	err := checkPhoneCodeAndCaptcha(c)
	if err != nil {
		cg.Failed(err.Error())
		return
	}

	phone := utils.FilterSpecialChar(strings.Trim(c.Request.FormValue("phone"), " "))
	currentUser := CurrentUser(c)
	existUser := dao.UserFindByPhone(phone)
	if existUser.ID > 0 && existUser.ID != currentUser.ID {
		cg.Failed("此手机号已被占用，换一个吧")
		return
	}
	updateData := make(map[string]string)
	updateData["phone"] = phone
	updateData["user_type"] = fmt.Sprintf("%v",dao.UserTypeNormal)
	dao.UpdateUser(currentUser.ID,updateData)
	cg.Success(nil)
}
//修改手机号--校验旧号码
func UserChangePhoneVerify(c *gin.Context) {
	cg := utils.Gin{C: c,}
	err := checkPhoneCodeAndCaptcha(c)
	if err != nil {
		cg.Failed(err.Error())
		return
	}
	cg.Success(nil)
}
//修改手机号--保存新号码
func UserChangePhoneSubmit(c *gin.Context) {
	cg := utils.Gin{C: c,}
	err := checkPhoneCodeAndCaptcha(c)
	if err != nil {
		cg.Failed(err.Error())
		return
	}

	phone := utils.FilterSpecialChar(strings.Trim(c.Request.FormValue("phone"), " "))
	currentUser := CurrentUser(c)
	updateData := make(map[string]string)
	updateData["phone"] = phone
	dao.UpdateUser(currentUser.ID,updateData)
	cg.Success(nil)
}

func checkPhoneCodeAndCaptcha(c *gin.Context) (err error) {
	phone := utils.FilterSpecialChar(strings.Trim(c.Request.FormValue("phone"), " "))
	if utils.IsPhone(phone) {
		return errors.New("手机号格式不正确")
	}
	captchaKey := utils.FilterSpecialChar(strings.Trim(c.Request.FormValue("captcha_key"), " "))
	captchaValue := utils.FilterSpecialChar(strings.Trim(c.Request.FormValue("captcha_value"), " "))
	if !utils.CheckCaptcha(captchaKey, []byte(captchaValue)) {
		return errors.New("图形验证码不正确")
	}
	phoneCode := utils.FilterSpecialChar(strings.Trim(c.Request.FormValue("verify_code"), " "))
	if !checkSmsCode(phone, phoneCode) {
		return errors.New("手机验证码不正确")
	}
	return nil
}