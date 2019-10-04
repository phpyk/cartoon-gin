package main

import (
	"flag"
	"fmt"
)

func main() {
	var count,offset int
	flag.IntVar(&count,"count",0,"count")
	flag.IntVar(&offset,"offset",0,"offset")
	flag.Parse()
	fmt.Println("count:",count)
	fmt.Println("offset:",offset)
}
