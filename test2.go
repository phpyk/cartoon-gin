package main

import "fmt"

func main() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key,val := range slice {
		m[key] = &val
	}
	ls := len(slice)
	i:=0
	for i= 0;i<ls;i++{
		fmt.Println(slice[i])
	}


	for k,v := range m {
		fmt.Println(k,"->",*v)
	}
}
