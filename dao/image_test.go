package dao

import (
	"fmt"
	"testing"
)

func TestFindImageByChapterId(t *testing.T) {
	cid := 231754
	list := FindImageByChapterId(cid)
	fmt.Printf("%+v",list)
}
