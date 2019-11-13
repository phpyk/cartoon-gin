package controllers

import (
	"encoding/json"
	"log"
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

	//err := cg.C.Request.ParseForm()
	//utils.CheckError(err)
	//catId := cg.C.Request.Form.Get("cat_id")
	log.Printf("search request:%+v",searchRequest)
}


