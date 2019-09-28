package controllers

import (
	"cartoon-gin/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GeneralPageInfo(c *gin.Context) (page, pageSize int) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	common.CheckError(err)
	pageSize, err = strconv.Atoi(c.DefaultQuery("per_page", "20"))
	common.CheckError(err)

	return page, pageSize
}
