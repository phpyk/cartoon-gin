package controllers

import (
	"cartoon-gin/common"
	"cartoon-gin/models"
	"github.com/gin-gonic/gin"
)

func AddUserAction(c *gin.Context) {
	cg := common.Gin{C:c}
	phone := c.Request.FormValue("phone")
	nickName := c.Request.FormValue("nick_name")
	user,err := models.AddUser(phone,nickName)

	if err != nil {
		cg.Failed("新建用户失败")
	}else {
		cg.Success(user)
	}
}

func QueryUserAction(c *gin.Context) {
	cg := common.Gin{C:c}
	phone := c.Request.FormValue("phone")
	users,err := models.GetUsersByPhone(phone)

	if err != nil{
		cg.Failed("获取用户失败")
	}
	cg.Success(users)
}
