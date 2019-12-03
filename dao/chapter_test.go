package dao

import (
	"fmt"
	"testing"
)

func TestGetChapterRow(t *testing.T) {
	id := 218702
	GetChapterRow(id)
}

func TestGetChapterList(t *testing.T) {
	list := GetChapterList(4166, 1, true, 1, 20)
	for _, row := range list {
		fmt.Printf("%+v \n", row)
	}
}
