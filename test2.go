package main

import (
	"fmt"
	"strings"
)

func main() {
	//url := "http://cartoon.qiniu.tblinker.com/2_d6d210_cover.jpg"
	//url := "http://cartoon1.qiniu.tblinker.com/cd0dce/41d0e299.data"
	url := "http://cartoon1.qiniu.tblinker.com/e034fb/703686/c9f0f895.data"
	idx := strings.Index(url,".com/")

	substr := url[idx+5:]
	arr := strings.Split(substr, "/")

	var name string
	for _,v := range arr {
		name += "/"+v
	}
	name = name[1:]
	fmt.Println(name)
}
