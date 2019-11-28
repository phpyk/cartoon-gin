package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func BookcaseTabsAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	user := CurrentUser(c)

	var searchRequest dao.BookCaseSearchRequest
	if err := c.Bind(&searchRequest); err != nil {
		cg.Failed("bind request failed")
		return
	}
	if searchRequest.PerPage == 0 {
		searchRequest.PerPage = 18
	}
	searchRequest.UserId = user.ID
	searchRequest.IsAndroid = IsAndroid(c)
	searchRequest.IsVerifying = IsVerifying(c)
	searchRequest.ShowRated = ShowReted(c)

	fmt.Printf("%+v\n",searchRequest)

	responseData := make(map[string]interface{})
	var totalCount = 0
	//我的收藏
	if searchRequest.Tab == "collect" {
		responseData["data"] = dao.GetUserCollections(searchRequest)
		totalCount = dao.GetUserCollectionCount(searchRequest)
	}
	//阅读历史
	if searchRequest.Tab == "history" {
		responseData["data"] = dao.GetUserReadingHistories(searchRequest)
		totalCount = dao.GetUserReadingHistoryCount(searchRequest)
	}
	//已购买
	if searchRequest.Tab == "pay" {
		responseData["data"] = dao.GetUserBoughtCartoons(searchRequest)
		totalCount = dao.GetUserBoughtCartoonsCount(searchRequest)
	}
	responseData = utils.AppendPaginateData(responseData,totalCount,searchRequest.Page,searchRequest.PerPage,c.Request.RequestURI)

	cg.Success(responseData)
}

func BookcaseDeleteAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	cartoonIdStr := c.Request.FormValue("cartoon_id_str")
	tab := c.Request.FormValue("tab")

	if cartoonIdStr == "" || tab == "" {
		cg.Failed("params required")
		return
	}
	cartoonIdSlice := strings.Split(cartoonIdStr,",")
	var deletedIdSlice []int
	for _, stringId := range cartoonIdSlice {
		if stringId == "" {
			continue
		}
		intId,err := strconv.Atoi(stringId)
		if err != nil || intId == 0 {
			continue
		}
		deletedIdSlice = append(deletedIdSlice,intId)
	}

	user := CurrentUser(c)
	if tab == "collect" {
		dao.DeleteUserCollection(user.ID,deletedIdSlice)
	}else if tab == "history" {
		dao.DeleteUserReadingHistory(user.ID,deletedIdSlice)
	}
	cg.Success(nil)
}