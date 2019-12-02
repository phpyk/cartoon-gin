package controllers

import (
	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func ConfigReportAction(c *gin.Context) {
	var reasonTypes = [5]int{
		dao.ReportReasonTypeQueshi,
		dao.ReportReasonTypeDisu,
		dao.ReportReasonTypeQinquan,
		dao.ReportReasonTypeHuazhicha,
		dao.ReportReasonTypeOthers,
	}

	rows := make(map[int]string)
	var report dao.UserReport
	for _, t := range reasonTypes {
		rows[t] = report.GetReportReason(t)
	}

	cg := utils.Gin{C: c,}
	responseData := make(map[string]interface{}) 
	responseData["list"] = rows
	cg.Success(responseData)
}

func ConfigVipAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	//TODO 安卓的channel和ios的channel含义不一样，需重新确定
	packageType := GetChannel(c)
	rows := dao.GetChargeVipConfigs(packageType)
	cg.Success()
}