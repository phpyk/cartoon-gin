package controllers

import (
	"cartoon-gin/configs"
	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func ConfigApiUrlAction(c *gin.Context) {
	cg := utils.Gin{C: c,}

	pushDuration := []int{2, 5, 15, 30,}
	ip1 := make(map[string]string)
	ip1["ip"] = "picaacgi.com"
	ip1["port"] = "80"
	ipList := []map[string]string{ip1,}
	outData := make(map[string]interface{})
	outData["data"] = ipList
	outData["duration"] = pushDuration
	cg.Success(outData)
	return
}

func ConfigReportAction(c *gin.Context) {
	var reasonTypes = [5]int{
		dao.ReportReasonTypeQueshi,
		dao.ReportReasonTypeDisu,
		dao.ReportReasonTypeQinquan,
		dao.ReportReasonTypeHuazhicha,
		dao.ReportReasonTypeOthers,
	}

	var rows []map[string]interface{}
	var report dao.UserReport
	for _, t := range reasonTypes {
		r := make(map[string]interface{})
		r["type"] = t
		r["text"] = report.GetReportReason(t)
		rows = append(rows, r)
	}

	cg := utils.Gin{C: c}
	responseData := make(map[string]interface{})
	responseData["list"] = rows
	cg.Success(responseData)
}

func ConfigVipAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	//TODO 安卓的channel和ios的channel含义不一样，需重新确定
	packageType := GetChannel(c)
	if packageType == "" {
		packageType = "beecar"
	}
	rows := dao.GetChargeVipConfigs(packageType)

	responseData := make(map[string]interface{})
	responseData["charge_info"] = getChargeExplain("vip")
	responseData["pack_list"] = rows
	cg.Success(responseData)
}

func ConfigCoinAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	//TODO 安卓的channel和ios的channel含义不一样，需重新确定
	packageType := GetChannel(c)
	if packageType == "" {
		packageType = "beecar"
	}
	rows := dao.GetChargeCoinConfigs(packageType)
	user := CurrentUser(c)
	for i, item := range rows {
		if item.IsDouble == 1 && dao.HasBought(user.ID, dao.OrderTypeCoin, item.ID) {
			item.IsDouble = 0
			rows[i] = item
		}
	}

	responseData := make(map[string]interface{})
	responseData["charge_info"] = getChargeExplain("coin")
	responseData["pack_list"] = rows
	cg.Success(responseData)
}

func ConfigPayChannelsAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	var list = []int{configs.PayChannelAlipay}
	responseData := make(map[string]interface{})
	responseData["pay_channels"] = list
	cg.Success(responseData)
}

//充值页面文案
func getChargeExplain(chargeType string) map[string]interface{} {
	res := make(map[string]interface{})
	res["currency"] = configs.HipopayCurrency
	res["symbol"] = configs.HipopayCurrencySymbol
	if chargeType == "coin" {
		var explainText = []string{
			"1、1美元≈700咔币。",
			"2、iOS端充值的咔币仅限iOS端使用。",
			"3、充值未完成时，请不要退出或卸载App，否则可能出现充值不到账的情况。",
			"4、首充奖励仅对第一笔充值有效，与充值档位无关。",
			"5、如果充值遇到问题，请联系客服QQ:3473807828",
			"6、对于使用游客身份登录的用户，强烈建议您绑定手机，绑定手机后可以在左右iOS设备上以手机号登录。",
		}
		res["explain"] = explainText
	}
	return res
}
