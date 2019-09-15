package dao

import (
	"fmt"
	"testing"
)

func TestGetConfigRows(t *testing.T) {
	list := GetConfigRows(MODULE_TYPE_SCROLL,5)
	fmt.Printf("%+v",list)
}
