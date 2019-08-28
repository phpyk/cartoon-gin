package main

import (
	"cartoon-gin/common"
	"cartoon-gin/models"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"

	//"strconv"
)

var bookListUrl = "https://www.qidian.com/free"

func main() {
	ParseBookListPage()
}

var isEnd int
var books []models.Book

func ParseBookListPage() {
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting:", request.URL.String())
	})
	c.OnError(func(response *colly.Response, e error) {
		fmt.Printf("error_code:%v,error_msg:%v \n", response.StatusCode, e.Error())
	})


	freebooksSelector := "#limit-list > div > ul > li"
	c.OnHTML(freebooksSelector, func(e *colly.HTMLElement) {
		aDom := e.DOM.Find("a").First()
		bookDetailUrl,_ := aDom.Attr("href")

		//访问书本详情页
		c.Visit(bookDetailUrl)
	})

	infoSelector := "body > div > div.book-detail-wrap.center990 > div.book-information.cf"
	c.OnHTML(infoSelector, func(e *colly.HTMLElement) {

		bookIdStr,_ := e.DOM.Find("#bookImg").Attr("data-bid")
		bookId,_ := strconv.Atoi(bookIdStr)
		coverUrl,_ := e.DOM.Find("#bookImg > img").First().Attr("src")
		bookName := e.DOM.Find("div.book-info > h1 > em").First().Text()// e.DOM.Find("em").First().Text()
		author := e.DOM.Find("div.book-info > h1 > span > a").First().Text()
		isEnd = 0
		bookStatusText := e.DOM.Find("div.book-info > p.tag > span:nth-child(1)").Text()
		if bookStatusText == "完本" {
			isEnd = 1
		}
		//分数获取不到，TODO
		score1 := e.DOM.Find("#score1").First().Text()
		//fmt.Println("score1",score1)
		score2 := e.DOM.Find("#score2").First().Text()
		//fmt.Println("score2",score2)
		score,err := strconv.ParseFloat(score1+"."+score2,32)
		common.CheckError(err)
		descShort := e.DOM.Find("p.intro").First().Text()
		descLong := e.DOM.Parent().Find("div.left-wrap.fl > div.book-info-detail > div.book-intro > p").First().Text()

		//分类需要用符号分隔开 TODO
		category := e.DOM.Find("div.book-info > p.tag > a").Text()
		bookDetailUrl,_ := e.DOM.Parent().Parent().Find("div.crumbs-nav.center990.top-op > span > a:nth-child(8)").First().Attr("href")

		book := models.Book{
			Name:bookName,
			Author:author,
			Score:score,
			DescShort:descShort,
			DescLong:descLong,
			Kinds:category,
			CoverUrl:coverUrl,
			BookDetailUrl:bookDetailUrl,
			IsEnd:isEnd,
			OriginBookId:bookId,
		}
		books = append(books, book)
		fmt.Printf("bookdetail:%+v \n",book)
	})

	err := c.Visit(bookListUrl)
	common.CheckError(err)
	saveBooks(books)
}

func saveBooks(books []models.Book) {
	db,err := models.OpenBookDB()
	common.CheckError(err)
	for b := range books {
		db.NewRecord(b)
		db.Create(&b)
	}
}