package controllers

import (
	"fmt"
	"strconv"

	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func GetRecommend(c *gin.Context)  {
	cg := utils.Gin{C: c,}
	user := CurrentUser(c)
	totalCount,err := strconv.Atoi(c.Request.FormValue("count"))
	utils.CheckError(err)
	if totalCount <= 0 {
		totalCount = 1
	}

	responseData := make(map[string]interface{})

	if totalCount == 1 {
		if utils.InNight() || user.UserType == dao.UserTypeTarget || user.UserType == dao.UserTypeWilling {
			var list  = []string{}
			responseData["data"] = list
			cg.Success(responseData)
			return
		}
	}

	var counter utils.RecommendCalculator
	counter.SetUserType(int(user.UserType))
	counter.SetCanShowRated(ShowReted(c))
	counter.SetAppVerifyStatus(IsVerifying(c))

	ratedCount := counter.GetRatedCount(totalCount)
	fmt.Println("rated_count: ",ratedCount)

	list := dao.GetRecommend(totalCount,ratedCount)
	responseData["data"] = list
	cg.Success(responseData)
}
