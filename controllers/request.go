package controllers

import (
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GeneralPageInfo(c *gin.Context) (page, pageSize int) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	utils.CheckError(err)
	pageSize, err = strconv.Atoi(c.DefaultQuery("per_page", "20"))
	utils.CheckError(err)

	return page, pageSize
}
