package dao

import (
	"fmt"
	"strings"
	"testing"
)

func TestSearchCartoonByConditions(t *testing.T) {
	var request SearchRequest
	request.IsEnd = "1"
	request.SortType = 1
	request.Keywords = "');select * from users ; -- "
	keywords := strings.Trim(request.Keywords," ")
	fmt.Println("keywords: ",keywords)
	//result := SearchCartoonByConditions(request)
	//fmt.Printf("%+v",result)
}