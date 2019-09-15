package controllers

import (
	"cartoon-gin/common"
	"cartoon-gin/dao"
	"github.com/gin-gonic/gin"
)

func GetHomeDataAction(c *gin.Context) {
	cg := common.Gin{C:c}
	//user := dao.FindUserByID(cg.C.Keys["uid"].(uint))
	homeData := make(map[string]interface{})
	homeData["scroll"] = dao.GetConfigRows(dao.MODULE_TYPE_SCROLL,5)
	cg.Success(homeData)
}