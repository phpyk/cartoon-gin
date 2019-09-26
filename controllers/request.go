package controllers

import (
	"cartoon-gin/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GeneralPageInfo(c *gin.Context) (page, pageSize int) {
	page, err := strconv.Atoi(c.Request.FormValue("page"))
	common.CheckError(err)
	pageSize, err = strconv.Atoi(c.Request.FormValue("per_page"))
	common.CheckError(err)
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return page, pageSize
}
