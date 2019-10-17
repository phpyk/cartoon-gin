package controllers

import (
	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func GetHomeDataAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	//user := dao.UserFindByID(cg.C.Keys["uid"].(uint))
	homeData := make(map[string]interface{})

	homeData["scroll_list"] = dao.GetHomeConfigRows(dao.MODULE_TYPE_SCROLL, 5)
	//TODO 按照用户喜好推荐
	homeData["recommend_list"] = dao.GetHomeConfigRows(dao.MODULE_TYPE_RECOMMEND, 6)
	homeData["elite_list"] = dao.GetHomeConfigRows(dao.MODULE_TYPE_ELITE, 4)
	homeData["hot_list"] = dao.GetHomeConfigRows(dao.MODULE_TYPE_HOT, 6)
	homeData["ended_list"] = dao.GetHomeConfigRows(dao.MODULE_TYPE_ENDED, 6)

	cg.Success(homeData)
}

func GetMoreAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	page, pageSize := GeneralPageInfo(c)

	moduleName := c.Request.FormValue("module_name")
	totalCount := dao.GetHomeConfigRowCount(moduleName)
	list := dao.GetMoreHomeConfigRows(moduleName, page, pageSize)

	responseData := make(map[string]interface{})
	responseData["data"] = list
	responseData = utils.AppendPaginateData(responseData, totalCount, page, pageSize, c.Request.RequestURI)
	cg.Success(responseData)
}

func GetRankAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	page, pageSize := GeneralPageInfo(c)

	sortBy := "read_count"
	list := dao.GetCartoonRank(page, pageSize, sortBy, "DESC")

	totalCount := dao.GetCartoonCount()
	responseData := make(map[string]interface{})
	responseData["data"] = list
	responseData = utils.AppendPaginateData(responseData, totalCount, page, pageSize, c.Request.RequestURI)

	cg.Success(responseData)
}

func GetNewCartoonsAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	page, pageSize := GeneralPageInfo(c)
	sortBy := "created_at"
	list := dao.GetCartoonRank(page, pageSize, sortBy, "DESC")
	totalCount := dao.GetCartoonCount()
	responseData := make(map[string]interface{})
	responseData["data"] = list
	responseData = utils.AppendPaginateData(responseData, totalCount, page, pageSize, c.Request.RequestURI)

	cg.Success(responseData)
}
