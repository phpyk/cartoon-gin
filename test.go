package main

import (
	"cartoon-gin/common"
	"fmt"
	"math"
)

func test() {
	a := math.Ceil(11/float64(3))
	b := 1
	fmt.Println(a)
	fmt.Println(float64(b))

	arr := []int{1,2,3,4,5}
	//bb := []
	fmt.Println("len:",len(arr))

	pwd := "123456"
	enPwd := common.Md5Str(common.Md5Str(pwd))
	enPwd2 := common.EncryptPwd(pwd)
	fmt.Println(pwd)
	fmt.Println(enPwd)
	fmt.Println(enPwd2)
}
