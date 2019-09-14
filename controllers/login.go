package controllers

import (
	"cartoon-gin/auth"
	"cartoon-gin/common"
	"cartoon-gin/dao"
	"github.com/gin-gonic/gin"
)

func LoginAction(c *gin.Context) {
	cg := common.Gin{C:c}
	phone := c.Request.FormValue("phone")
	password := c.Request.FormValue("password")

	if !common.IsPhone(phone) {
		cg.Failed("手机号格式不正确")
	}
	user := dao.FindUserByPhone(phone)
	if !(user.ID > 0) {
		cg.Failed("用户不存在")
	}
	encryptPwd := common.EncryptPwd(password)
	if encryptPwd != user.Password {
		cg.Failed("密码不正确")
	}

	token,err := auth.GenerateToken(&user)
	if err != nil {
		cg.Failed("Login failed:"+err.Error())
	}

	response := make(map[string]interface{})
	response["token"] = token
	response["token_type"] = "Bearer"
	response["user_info"] = user
	cg.Success(response)
}

func CurrentUserAction(c *gin.Context) {
	cg := common.Gin{C:c}
	//interface 转 uint类型
	//cg.C.Keys["uid"].(uint)
	me := dao.FindUserByID(cg.C.Keys["uid"].(uint))
	cg.Success(me)
}