package main

import (
	"cartoon-gin/common"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"strconv"
	"time"
)

func main() {
	c ,err := ioutil.ReadFile("max_id.log")
	if err != nil {
		fmt.Println(err)
	}
	i,err := strconv.Atoi(string(c))
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))
	//maxId := []byte(strconv.Itoa(31415926))
	//ioutil.WriteFile("last_max_id.log",maxId,0644)

	//t := time.Now().Format("2006-01-02 15:04:05")
	count := flag.Int("count",1000,"count")
	fmt.Println("count:",*count)
	offset := flag.Int("offset",1000,"offset")
	fmt.Println("offset:",*offset)
	t1 := time.Now()

	fmt.Println(t1)

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

	escaped := time.Since(t1)
	fmt.Print(escaped)
}
