package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ShowReted(c *gin.Context) bool {
	if IsUSAIp(c) {
		return false
	}
	return false
}

func IsUSAIp(c *gin.Context) bool {
	//TODO 添加IP归属地检测
	//ip := GetClientIP(c)
	return false;
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

func IsApplePayUser(c *gin.Context) bool {
	return false
}

func GetAppVersion(c *gin.Context) string {
	v := c.Request.Header.Get("h-v")
	return v
}
func GetDeviceType(c *gin.Context) string {
	deviceType := strings.ToLower(c.Request.Header.Get("h-device-type"))
	return deviceType
}
func GetAndroidChannel(c *gin.Context) string {
	return strings.ToLower(c.Request.Header.Get("h-channel"))
}
func IsIOS(c *gin.Context) bool {
	return GetDeviceType(c) == "ios"
}
func IsAndroid(c *gin.Context) bool {
	return GetDeviceType(c) == "android"
}
