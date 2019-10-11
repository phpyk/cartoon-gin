package common

import (
	"fmt"
	"testing"
)

func TestGetNewCaptcha(t *testing.T) {
	cap := GetNewCaptcha()
	fmt.Println(cap)
}

func TestCheckCaptcha(t *testing.T) {
	capid := "Zmo7JWU5lrB6jkfgIB7f"
	value := []byte{6,2,9,6,9,9}
	res := CheckCaptcha(capid,value)
	fmt.Println(res)
}