package controllers

import (
	"cartoon-gin/common"
	"cartoon-gin/dao"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetHomeDataAction(c *gin.Context) {
	cg := common.Gin{C:c}
	//user := dao.FindUserByID(cg.C.Keys["uid"].(uint))
	homeData := make(map[string]interface{})

	homeData["scroll_list"] = dao.GetHomeConfigRows(dao.MODULE_TYPE_SCROLL,5)
	//TODO 按照用户喜好推荐
	homeData["recommend_list"] = dao.GetHomeConfigRows(dao.MODULE_TYPE_RECOMMEND,6)
	homeData["elite_list"] = dao.GetHomeConfigRows(dao.MODULE_TYPE_ELITE,4)
	homeData["hot_list"] = dao.GetHomeConfigRows(dao.MODULE_TYPE_HOT,6)
	homeData["ended_list"] = dao.GetHomeConfigRows(dao.MODULE_TYPE_ENDED,6)

	cg.Success(homeData)
}

func GetMoreAction(c *gin.Context){
	cg := common.Gin{C:c}
	log.Println(c.Request.Host)
	log.Println(c.Request.RequestURI)
	log.Println(c.Request.URL)
	log.Printf("%+v",c.Request.Form)

	page,_ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize,_ := strconv.Atoi(c.Request.FormValue("per_page"))
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10;
	}

	moduleName := c.Request.FormValue("module_name")
	totalCount := dao.GetHomeConfigRowCount(moduleName)
	list := dao.GetMoreHomeConfigRows(moduleName,page,pageSize)

	responseData := make(map[string]interface{})
	responseData["data"] = list
	responseData = common.AppendPaginateData(responseData,totalCount ,page,pageSize,c.Request.RequestURI)
	cg.Success(responseData)
}