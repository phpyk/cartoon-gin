package controllers

import (
	"cartoon-gin/common"
	"cartoon-gin/models"
	"github.com/gin-gonic/gin"
)

func GetAllAction(ctx *gin.Context) {
	cg := common.Gin{C:ctx}
	cats := models.GetAllCategories()
	cg.Success(cats)
}