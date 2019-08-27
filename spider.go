package main

import (
	"cartoon-gin/common"
	"cartoon-gin/models"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
)

var startUrl = "https://book.qidian.com/info/1015444718#Catalog"
var imageBaseDir = "images2/";

func main() {
	ParsePage()
}
func ParsePage() {
	db,_ := models.OpenBookDB()

	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting:", request.URL.String())
	})
	c.OnError(func(response *colly.Response, e error) {
		fmt.Printf("error_code:%v,error_msg:%v \n", response.StatusCode, e.Error())
	})
	freeChapters := "#j-catalogWrap > div.volume-wrap > div:nth-child(1) > ul > li > a"
	c.OnHTML(freeChapters, func(e *colly.HTMLElement) {
		chapterSeq,exists := e.DOM.Parent().First().Attr("data-rid")
		if !exists {
			log.Println("chapter sequence not found")
		}
		seq,_ := strconv.Atoi(chapterSeq)

		chapterName := e.Text
		chapterUrl := e.Attr("href")
		fmt.Printf("seq:%v -- name:%v -- url:%v \n",seq,chapterName,chapterUrl)
		chapter := models.Chapter{Sequence:seq,Name:chapterName,Url:chapterUrl}
		db.Create(&chapter)
	})

	err := c.Visit(startUrl)
	common.CheckError(err)
}
