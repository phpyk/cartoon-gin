package common

import (
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
