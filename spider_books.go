package main

import (
	"cartoon-gin/common"
	"cartoon-gin/models"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"

	//"strconv"
)

var bookListUrl = "https://www.qidian.com/free"

func main() {
	ParseBookListPage()
}

var score,bookId,imageUrl,bookDetailUrl,descLong,bookName string
var isEnd int
var book models.Book


func ParseBookListPage() {
	//db,_ := models.OpenBookDB()

	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting:", request.URL.String())
	})
	c.OnError(func(response *colly.Response, e error) {
		fmt.Printf("error_code:%v,error_msg:%v \n", response.StatusCode, e.Error())
	})

	freebooksSelector := "#limit-list > div > ul > li"
	c.OnHTML(freebooksSelector, func(e *colly.HTMLElement) {
		exists := true
		//评分
		score = e.DOM.Find("div.score").First().Text()
		//书详情链接
		aDom := e.DOM.Find("a").First()
		bookId,exists = aDom.Attr("data-bid")
		bookDetailUrl,exists = aDom.Attr("href")
		imageUrl,exists = aDom.Find("img").First().Attr("src")
		if !exists {
			log.Println("bookId bookDetailUrl imageUrl not found")
		}

		//访问书本详情页
		c.Visit(bookDetailUrl)

		scoreFloat,err := strconv.ParseFloat(score,32)
		common.CheckError(err)
		book.Score = scoreFloat
		bookidInt,err := strconv.Atoi(bookId)
		common.CheckError(err)
		book.OriginBookId = bookidInt
		book.BookDetailUrl = bookDetailUrl
		book.CoverUrl = imageUrl
		//fmt.Printf("bookdetail:%+v \n",book)
	})

	descLongSelector := "body > div > div.book-detail-wrap.center990 > div.book-content-wrap.cf.hidden > div.left-wrap.fl > div.book-info-detail > div.book-intro"
	c.OnHTML(descLongSelector, func(e *colly.HTMLElement) {
		descLong = e.Text
		book.DescLong = descLong
	})
	infoSelector := "body > div > div.book-detail-wrap.center990 > div.book-information.cf > div.book-info"
	c.OnHTML(infoSelector, func(e *colly.HTMLElement) {
		bookName = e.DOM.Find("em").First().Text()
		isEnd = 0
		bookStatusText := e.DOM.Find("p.tag > span:nth-child(1)").Text()
		if bookStatusText == "完本" {
			isEnd = 1
		}
		book.Name = bookName
		book.IsEnd = isEnd
		//catDom := e.DOM.Find("p.tag > a")
		//fmt.Printf("bookdetail:%+v \n",book)
	})

	fmt.Printf("bookdetail:%+v \n",book)

	err := c.Visit(bookListUrl)
	common.CheckError(err)
}