package dao

import (
	"fmt"
	"testing"
)

func TestGetUserReadingHistories(t *testing.T) {
	var req BookCaseSearchRequest
	req.UserId = 152
	req.Tab = "history"
	req.SortType = 3
	req.ShowRated =  true
	req.IsVerifying = false
	req.IsAndroid = false
	req.Page = 1
	req.PerPage = 18

	res := GetUserReadingHistories(req)
	for i,v := range res {
		fmt.Printf("%d -- %+v\n",(i+1),v)
	}
}
