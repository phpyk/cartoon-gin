package controllers

import (
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
	if password !=
}