package utils

import (
	"fmt"
	"testing"
)

func TestGetNewCaptcha(t *testing.T) {
	cap, imgData := GetNewCaptcha()
	fmt.Println(cap)
	fmt.Println(imgData)
	//val := GetCaptchaVal(cap)
	//
	//res := CheckCaptcha(cap,val)
	//fmt.Println(res)
}

func TestCheckCaptcha(t *testing.T) {
	capid := "J0aVmCA0gnVtBR3eocId"
	value := []byte{7, 7, 9, 3}
	res := CheckCaptcha(capid, value)
	fmt.Println(res)
}
