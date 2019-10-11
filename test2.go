package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "http://cartoon1.qiniu.tblinker.com/8b0d26/41d0e299.data"
	arr := strings.Split(url, "/")
	l := len(arr)

	name := arr[l-3] + "/" + arr[l-2] + "/" + arr[l-1]
	fmt.Println(name)
}
