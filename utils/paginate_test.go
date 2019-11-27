package utils

import (
	"fmt"
	"testing"
)

func TestCalcuLastPage(t *testing.T) {
	pageCount := calcuLastPage(10, 3)
	fmt.Println(pageCount)
}
