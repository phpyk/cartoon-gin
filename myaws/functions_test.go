package myaws

import (
	"fmt"
	"testing"
)

func TestReadSrcAndLocalSave(t *testing.T) {
	imgurl := "http://cartoon1.qiniu.tblinker.com/30c8e1/e52e41/45c48cce.data"
	fullName := ReadSrcAndLocalSave(imgurl)
	fmt.Println(fullName)
}
