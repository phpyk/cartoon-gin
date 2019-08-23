package controllers

import (
	"cartoon-gin/common"
	"cartoon-gin/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUserAction(c *gin.Context) {
	cg := common.Gin{C:c}
	phone := c.Request.FormValue("phone")
	nickName := c.Request.FormValue("nick_name")
	user,err := models.AddUser(phone,nickName)

	if err != nil {
		cg.Failed("新建用户失败",err,user)
	}else {
		cg.Success(user)
	}
}

func QueryUserAction(c *gin.Context) {
	cg := common.Gin{C:c}
	phone := c.Request.FormValue("phone")
	users,err := models.GetUsersByPhone(phone)

	if err != nil{
		cg.Failed("获取用户失败",err,users)
	}
	cg.Success(users)
}

func TestDB(c *gin.Context) {
	db,err := sql.Open("mysql","root:123_QWE_asd@(127.0.0.1:3306)/gotest?parseTime=True")
	if err != nil {
		panic(err)
	}
	res,err := db.Exec("insert into users(phone,nick_name) values (?,?)","18657173832","yuekai")
	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK,gin.H{
		"user_id": id,
	})
}