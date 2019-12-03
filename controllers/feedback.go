package controllers

import (
	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

func FeedbackStoreAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	content := utils.FilterSpecialChar(c.Request.FormValue("content"))
	if content == "" {
		cg.Failed("请输入内容")
		return
	}
	user := CurrentUser(c)
	var feedback = dao.UserFeedback{CommitUserId: user.ID, Content: content, Type: dao.FeedbackTypeCommit}
	feedback.Save()
	cg.Success(nil)
}

func FeedbackListAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	page,perPage := GeneralPageInfo(c)
	user := CurrentUser(c)
	list := dao.GetFeedbacks(user.ID,page, perPage)
	totalCount := dao.GetFeedbackCount(user.ID)
	responseData := make(map[string]interface{})
	responseData["data"] = list
	responseData = utils.AppendPaginateData(responseData,totalCount,page,perPage,c.Request.RequestURI)
	cg.Success(responseData)
}
