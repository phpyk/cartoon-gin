package controllers

import (
	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func CheckAppUpdateAction(c *gin.Context)  {
	cg := utils.Gin{C: c,}
	deviceType := GetDeviceType(c)
	currentVersion := getAppVersionRow(c)
	newVersion := dao.GetNewestAppVersionRow(deviceType,"")
	if currentVersion.Version == newVersion.Version {
		cg.Success(nil)
		return
	}

	outData := make(map[string]interface{})
	outData["current_version"] = formatOutData(currentVersion)
	outData["new_version"] = formatOutData(newVersion)
	cg.Success(outData)
	return
}

func formatOutData(row dao.AppVersion) map[string]interface{} {
	data := make(map[string]interface{})
	data["is_verifying"] = row.IsVerifying
	data["is_forced_update"] = row.IsForcedUpdate
	data["content"] = row.Content
	data["new_version"] = row.Version
	data["device_type"] = row.DeviceType
	data["download_url"] = row.DownloadUrl
	return data
}