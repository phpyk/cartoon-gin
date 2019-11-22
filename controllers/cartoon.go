package controllers

import (
	"encoding/json"
	"strconv"

	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func CartoonBaseInfoAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	cartoonId,_ := strconv.Atoi(cg.C.Request.FormValue("cartoon_id"))
	cartoon := dao.GetCartoonById(cartoonId)
	jsonData,err := json.Marshal(cartoon)
	utils.CheckError(err)

	var outData map[string]interface{}
	err = json.Unmarshal([]byte(jsonData),&outData)
	utils.CheckError(err)
	outData["tags"] = utils.GetTagsArray(outData["tags"].(string),2)
	cg.Success(outData)
}

func CartoonSearchAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	var searchRequest dao.SearchRequest
	if err := c.Bind(&searchRequest); err != nil {
		cg.Failed("bind request failed")
	}

	var outData []map[string]interface{}
	outData = dao.SearchCartoonByConditions(searchRequest)
	cg.Success(outData)
}

func CartoonChapterListAction(c *gin.Context)  {
	cg := utils.Gin{C: c,}
	cartoonId,err := strconv.Atoi(c.Request.FormValue("cartoon_id"))
	utils.CheckError(err)
	if dao.CartoonExists(cartoonId) {
		cg.Failed("漫画不存在")
		return
	}
	sortType,err := strconv.Atoi(c.Request.FormValue("sort_type"))
	if sortType == 0 || err != nil {
		sortType = 1
	}
	page, pageSize := GeneralPageInfo(c)

	chapterList := dao.GetChapterList(cartoonId,sortType,true,page,pageSize)
	totalCount := dao.GetChaptersCount(cartoonId)

	responseData := make(map[string]interface{})
	responseData["data"]  = chapterList
	responseData = utils.AppendPaginateData(responseData, totalCount, page, pageSize, c.Request.RequestURI)
	cg.Success(responseData)
}

func CartoonReadAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	user := CurrentUser(c)
	cartoonId,err := strconv.Atoi(c.Request.FormValue("cartoon_id"))
	chapterId,err := strconv.Atoi(c.Request.FormValue("chapter_id"))
	utils.CheckError(err)
	cartoon := dao.GetCartoonById(cartoonId)
	chapter := dao.GetChapterRow(chapterId)
	//vip 默认is_buy=1
	var hasBought bool
	if user.IsVip == 1 {
		hasBought = true
	}else {
		hasBought = dao.HasBoughtChapter(user.ID,chapterId)
	}
	var limit int
	//如果是H,且没有购买，且不是VIP的话只返回第一张图
	if cartoon.IsRated == 1 && user.IsVip == 0 && !hasBought && chapter.SalePrice > 0 {
		limit = 1
	}
	images := dao.GetImagesByChapterId(chapterId,limit)

	var imageUrlList []string
	for _,img := range images {
		imageUrlList  = append(imageUrlList,img.ImageAddr)
	}

	//前后章节信息
	lastChapter,nextChapter := dao.GetChapterNeighbors(chapter)

	responseData := make(map[string]interface{})
	responseData["images"] = imageUrlList
	responseData["last_chapter_id"] = lastChapter.ID
	responseData["next_chapter_id"] = nextChapter.ID
	if hasBought {
		chapter.IsBuy = 1
	} else {
		chapter.IsBuy = 0
	}
	responseData["chapter_info"] = chapter

	//DB记录阅读历史
	dao.
	//redis记录最近阅读章节id
	cg.Success(responseData)
}

