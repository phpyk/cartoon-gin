package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func CartoonBaseInfoAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	cartoonId, _ := strconv.Atoi(cg.C.Request.FormValue("cartoon_id"))
	cartoon := dao.GetCartoonById(cartoonId)
	jsonData, err := json.Marshal(cartoon)
	utils.CheckError(err)

	var outData map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &outData)
	utils.CheckError(err)
	outData["tags"] = utils.GetTagsArray(fmt.Sprint(outData["tags"]), 2)
	cg.Success(outData)
}

func CartoonSearchAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	var searchRequest dao.SearchRequest
	if err := c.Bind(&searchRequest); err != nil {
		cg.Failed("bind request failed")
		return
	}

	var outData []map[string]interface{}
	outData = dao.SearchCartoonByConditions(searchRequest)
	cg.Success(outData)
}

func CartoonChapterListAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	cartoonId, err := strconv.Atoi(c.Request.FormValue("cartoon_id"))
	utils.CheckError(err)
	if dao.CartoonExists(cartoonId) {
		cg.Failed("漫画不存在")
		return
	}
	sortType, err := strconv.Atoi(c.Request.FormValue("sort_type"))
	if sortType == 0 || err != nil {
		sortType = 1
	}
	page, pageSize := GeneralPageInfo(c)

	chapterList := dao.GetChapterList(cartoonId, sortType, true, page, pageSize)
	totalCount := dao.GetChaptersCount(cartoonId)

	responseData := make(map[string]interface{})
	responseData["data"] = chapterList
	responseData = utils.AppendPaginateData(responseData, totalCount, page, pageSize, c.Request.RequestURI)
	cg.Success(responseData)
}

func CartoonReadAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	user := CurrentUser(c)
	cartoonId, err := strconv.Atoi(c.Request.FormValue("cartoon_id"))
	chapterId, err := strconv.Atoi(c.Request.FormValue("chapter_id"))
	utils.CheckError(err)
	cartoon := dao.GetCartoonById(cartoonId)
	chapter := dao.GetChapterRow(chapterId)
	//vip 默认is_buy=1
	var hasBought bool
	if user.IsVip == 1 {
		hasBought = true
	} else {
		hasBought = dao.HasBoughtChapter(user.ID, chapterId)
	}
	var limit int
	//如果是H,且没有购买，且不是VIP的话只返回第一张图
	if cartoon.IsRated == 1 && user.IsVip == 0 && !hasBought && chapter.SalePrice > 0 {
		limit = 1
	}
	images := dao.GetImagesByChapterId(chapterId, limit)

	var imageUrlList []string
	for _, img := range images {
		imageUrlList = append(imageUrlList, img.ImageAddr)
	}

	//前后章节信息
	lastChapter, nextChapter := dao.GetChapterNeighbors(chapter)

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
	dao.AddToReadingHistory(user.ID, cartoonId, chapterId)
	//redis记录最近阅读章节id
	saveLastReadChapterInfoToRedis(user.ID, cartoonId, chapterId, chapter.ChapterName)
	cg.Success(responseData)
}

func CartoonCollectAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	cartoonId, err := strconv.Atoi(c.Request.FormValue("cartoon_id"))
	utils.CheckError(err)
	actType := c.Request.FormValue("act_type")
	if actType == "" {
		actType = "add"
	}
	if dao.CartoonExists(cartoonId) {
		cg.Failed("漫画不存在")
		return
	}

	user := CurrentUser(c)
	//收藏
	if actType == "add" {
		if dao.CartoonHasBeenCollected(user.ID, cartoonId) || dao.CollectCartoon(user.ID, cartoonId) {
			cg.Success(nil)
			return
		} else {
			cg.Failed("加入收藏失败，请稍后再试~")
			return
		}
	} else {
		//取消收藏
		dao.CancelCollectCartoon(user.ID, cartoonId)
		cg.Success(nil)
		return
	}
}

func CartoonBuyAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	chapterId, err := strconv.Atoi(c.Request.FormValue("chapter_id"))
	utils.CheckError(err)

	user := CurrentUser(c)
	if dao.HasBoughtChapter(user.ID, chapterId) {
		cg.Failed("您已经购买过了，无需重复购买~")
		return
	}
	chapter := dao.GetChapterRow(chapterId)
	if chapter.ID == 0 {
		cg.Failed("章节不存在")
		return
	}
	if user.ValidCoin < uint(chapter.SalePrice) {
		cg.Failed("可用金币余额不足")
		return
	}

	err = dao.BuyChapter(user, &chapter)
	if err != nil {
		log.Printf("购买失败：user_id:%v,chapter_id:%v,error:%v", user.ID, chapterId, err.Error())
		cg.Failed("购买失败，请稍后再试")
		return
	}

	outData := make(map[string]interface{})
	outData["user_valid_coin"] = int(user.ValidCoin)
	cg.Success(outData)
	return
}

func saveLastReadChapterInfoToRedis(userId, cartoonId, chapterId int, chapterName string) {
	parms := make(map[string]string)
	parms["uid"] = strconv.Itoa(userId)
	parms["cid"] = strconv.Itoa(cartoonId)
	key := utils.GetRedisKey(utils.RDS_KEY_USER_LAST_READ_CHAPTER_INFO, parms)

	readingInfo := make(map[string]interface{})
	readingInfo["last_read_chapter_name"] = chapterName
	readingInfo["last_read_chapter_id"] = chapterId
	readingInfo["last_read_time"] = time.Now().Unix()

	if val, err := json.Marshal(readingInfo); err == nil {
		utils.RedisSave(key, string(val))
	}
}
