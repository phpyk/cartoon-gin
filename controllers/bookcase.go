package controllers

import (
	"fmt"

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
	//我的收藏
	if searchRequest.Tab == "collect" {
		responseData["data"] = dao.GetUserCollections(searchRequest)
		searchRequest.GetTotalCount = true
		totalCount := dao.GetUserCollectionCount(searchRequest)
		responseData = utils.AppendPaginateData(responseData,totalCount,searchRequest.Page,searchRequest.PerPage,c.Request.RequestURI)
	}
	//阅读历史
	if searchRequest.Tab == "history" {

	}
	//已购买
	if searchRequest.Tab == "pay" {

	}
	cg.Success(responseData)
}