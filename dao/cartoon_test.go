package dao

import (
	"cartoon-gin/utils"
	"fmt"
	"testing"
)

func TestSearchCartoonByConditions(t *testing.T) {
	var request SearchRequest
	request.IsEnd = "1"
	request.SortType = 1
	request.Keywords = "战士!@#$%^&^%$%^&*("
	//keywords := strings.Trim(request.Keywords," ")
	keywords := utils.FilterSpecialChar(request.Keywords)
	fmt.Println("keywords: ", keywords)
	result := SearchCartoonByConditions(request)
	fmt.Printf("%+v", result)
}

func TestGetRecommendFromCache(t *testing.T) {
	userId := 133
	list := GetRecommendFromCache(userId)
	fmt.Printf("%v", list)
}
