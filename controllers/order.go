package controllers

import (
	"strconv"

	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func OrderCreateAction(c *gin.Context) {
	cg := utils.Gin{C: c,}

	//1购买VIP，2金币充值
	orderType, err := strconv.Atoi(c.Request.FormValue("order_type"))
	packageId, err := strconv.Atoi(c.Request.FormValue("package_id"))
	utils.CheckError(err)
	if orderType == 0 || packageId == 0 {
		cg.Failed("params error:order_type or package_id ")
		return
	}
	user := CurrentUser(c)
	var order dao.Order
	dao.crea


}
