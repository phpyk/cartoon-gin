package dao

import (
	"fmt"
	"testing"
)

func TestBuyChapter(t *testing.T) {
	user := UserFindByID(169)
	chapter := GetChapterRow(217984)
	//fmt.Printf("%+v \n", user)
	//fmt.Printf("%+v \n", chapter)
	//return
	err := BuyChapter(&user, &chapter)
	if err != nil {
		fmt.Println(err)
	}
}
