package main

import (
	"fmt"
	"time"
	"unsafe"

	"cartoon-gin/utils"
)

const (
	Sunday = iota	//0
	Monday			//1
	Tuesday			//2
	Wednesday		//3
	Thursday		//4
	Friday			//5
	Saturday		//6
)

func init() {
	fmt.Println("test init")
}

func main() {
	smsClient := utils.NewSmsClient()

	fmt.Printf("client:%+v \n",smsClient)

	to := "17505818455"
	data := []string{"8810","5"}
	templateId := 424738

	result ,err := smsClient.SendTemplateSMS(to,data,templateId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result : %+v",result)
}

func test() {
	fmt.Println(time.Now().Format("20060102150405"))

	fmt.Println("Sunday:",Sunday)
	fmt.Println("Monday:",Monday)
	fmt.Println("Tuesday:",Tuesday)
	fmt.Println("Wednesday:",Wednesday)
	fmt.Println("Thursday:",Thursday)
	fmt.Println("Friday:",Friday)
	fmt.Println("Saturday:",Saturday)

	fmt.Printf("sunday : %T\n",Sunday)

	var a int = 1000
	var b float64
	b = float64(a)+0.001

	fmt.Println("b is :",b)
	fmt.Printf("size of a is :%d \n",unsafe.Sizeof(a))

	switch num := number(); { // num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
	}

	numbers := []int{1,2,3,4,5}
	fmt.Printf("numbers type:%T\n",numbers)

	numa := [3]int{78, 79 ,80}
	nums1 := numa[:] // creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	s1 := numa[1:]
	fmt.Printf("s1 %v, len %d, cap:%d, addr:%x array_addr:%x\n",s1,len(s1),cap(s1),&s1,&numa)
	s1  = append(s1,81)
	fmt.Printf("s1 %v, len %d, cap:%d, addr:%x array_addr:%x\n",s1,len(s1),cap(s1),&s1,&numa)

}



func number() int {
	return 15*5
}
