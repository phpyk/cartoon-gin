package main

import (
	"fmt"
	"os"
)

func main() {
	file,err := os.Open("")
	if err != nil {
		panic(err)
	}
	file.Close()

	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
