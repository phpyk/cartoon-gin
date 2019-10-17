package utils

import (
	"fmt"
	"testing"
)

func TestIsPhone(t *testing.T) {
	phone := "1750818455"
	isPhone := IsPhone(phone)
	if isPhone {
		t.Log("ok")
	} else {
		t.Error("fail:", isPhone)
	}
}
func TestRandomString(t *testing.T) {
	l := 30
	str := RandomString(l,5)
	fmt.Println(str)
}

func TestGeneralInviteCode(t *testing.T) {
	for i := 0;i<10;i++ {
		fmt.Println(GeneralInviteCode())
	}
}
