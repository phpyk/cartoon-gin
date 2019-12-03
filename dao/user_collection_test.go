package dao

import (
	"fmt"
	"testing"
)

func TestCancelCollectCartoon(t *testing.T) {
	uId := 63
	cid := 10000
	res := CancelCollectCartoon(uId, cid)
	fmt.Println(res)
}

func TestGetUserCollection(t *testing.T) {
	var req BookCaseSearchRequest
	req.UserId = 152
	req.Tab = "collect"
	req.SortType = 3
	req.ShowRated = true
	req.IsVerifying = false
	req.IsAndroid = false
	req.Page = 1
	req.PerPage = 18

	res := GetUserCollections(req)
	for i, v := range res {
		fmt.Printf("%d -- %+v\n", (i + 1), v)
	}
}
