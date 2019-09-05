package main

import (
	"cartoon-gin/common"
	"fmt"
)

func main() {
	pwd := "123456"
	enPwd := common.Md5Str(common.Md5Str(pwd))
	enPwd2 := common.EncryptPwd(pwd)
	fmt.Println(pwd)
	fmt.Println(enPwd)
	fmt.Println(enPwd2)
}
