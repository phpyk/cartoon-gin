package controllers

import (
	"strconv"
	"strings"

	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func ReportStoreAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	cartoonId, _ := strconv.Atoi(c.Request.FormValue("cartoon_id"))
	if cartoonId == 0 {
		cg.Failed("cartoon_id 不能为空")
		return
	}
	typeString := strings.Trim(c.Request.FormValue("type"), " ")
	typeList := strings.Split(typeString, ",")
	user := CurrentUser(c)
	for _, v := range typeList {
		iv, _ := strconv.Atoi(v)
		var row dao.UserReport
		row.UserId = user.ID
		row.CartoonId = cartoonId
		row.ReasonType = iv
		row.Reason = row.GetReportReason(row.ReasonType)
		row.Save()
	}
	cg.Success(nil)
}
