package controllers

import (
	"cartoon-gin/common"
	"cartoon-gin/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllAction(ctx *gin.Context) {
	cg := common.Gin{C:ctx}
	cats := models.GetAllCategories()
	cg.Success(cats)
}

func AddCatAction(ctx *gin.Context) {
	cg := common.Gin{C: ctx}
	catName := ctx.Request.FormValue("cat_name")
	parentId,err := strconv.Atoi(ctx.Request.FormValue("parent_id"))
	common.CheckError(err)

	exists := models.GetCatByName(catName)
	if exists.CateName == catName {
		cg.Failed("分类已存在")
		return
	}

	res := models.AddCategory(catName,parentId)
	if res {
		cg.Success(res)
	}else {
		cg.Failed("添加分类失败")
	}
}

func UpdateCatAction(ctx *gin.Context) {
	cg := common.Gin{C: ctx}
	id,err := strconv.Atoi(ctx.Request.FormValue("cat_id"))
	common.CheckError(err)

}