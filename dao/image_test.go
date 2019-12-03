package dao

import (
	"fmt"
	"testing"
)

func TestFindImageByChapterId(t *testing.T) {
	cid := 231754
	list := GetImagesByChapterId(cid, 0)
	fmt.Printf("%+v", list)
}
