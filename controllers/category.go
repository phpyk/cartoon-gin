package controllers

import (
	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func CategoryLabelsAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	cats := dao.GetAllCategories()
	cg.Success(cats)
}
