package main

import (
	"cartoon-gin/common"
	"fmt"
	"math"
	"reflect"
	"time"
)

func main() {
	//t := time.Now().Format("2006-01-02 15:04:05")
	t2 := time.Now().Unix()
	fmt.Println(reflect.TypeOf(t2))
	fmt.Println(t2)

	a := math.Ceil(11 / float64(3))
	b := 1
	fmt.Println(a)
	fmt.Println(float64(b))

	arr := []int{1, 2, 3, 4, 5}
	//bb := []
	fmt.Println("len:", len(arr))

	pwd := "123456"
	enPwd := common.Md5Str(common.Md5Str(pwd))
	enPwd2 := common.EncryptPwd(pwd)
	fmt.Println(pwd)
	fmt.Println(enPwd)
	fmt.Println(enPwd2)
}
