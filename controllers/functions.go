package controllers

import (
	"log"
	"strings"

	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) (user *dao.User) {
	if value, exists := c.Get("user"); exists && value != nil {
		user, ok := value.(*dao.User)
		if !ok {
			panic("context user type error")
		}
		return user
	}
	return
}

func FlushCurrentUser(c *gin.Context, user *dao.User) {
	c.Set("user", user)
}

func ShowReted(c *gin.Context) bool {
	if IsUSAIp(c) {
		return false
	}
	row := getAppVersionRow(c)
	return row.ShowRated == 1
}

func IsVerifying(c *gin.Context) bool {
	row := getAppVersionRow(c)
	log.Printf("version row:%+v\n", row)
	return row.IsVerifying == 1
}

func getAppVersionRow(c *gin.Context) dao.AppVersion {
	dtype := GetDeviceType(c)
	version := GetAppVersion(c)
	channel := GetChannel(c)
	row := dao.GetAppVersionRow(version, dtype, channel)
	return row
}

func IsUSAIp(c *gin.Context) bool {
	//TODO 添加IP归属地检测
	//ip := GetClientIP(c)
	return false
}

func IsApplePayUser(userId int) bool {
	redisClient := utils.NewRedisClient()
	isMember, err := redisClient.SIsMember(utils.RDS_KEY_APPLE_PAY_USERS, userId).Result()
	utils.CheckError(err)
	return isMember
}

func GetClientIP(c *gin.Context) string {
	r := c.Request
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

func GetAppVersion(c *gin.Context) string {
	v := c.Request.Header.Get("h-v")
	return v
}
func GetDeviceType(c *gin.Context) string {
	deviceType := strings.ToLower(c.Request.Header.Get("h-device-type"))
	return deviceType
}
func GetChannel(c *gin.Context) string {
	return strings.ToLower(c.Request.Header.Get("h-channel"))
}
func IsIOS(c *gin.Context) bool {
	return GetDeviceType(c) == "ios"
}
func IsAndroid(c *gin.Context) bool {
	return GetDeviceType(c) == "android"
}
