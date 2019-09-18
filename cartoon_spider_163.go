package main

import (
	"cartoon-gin/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const GET_DATA_API string = "https://163.bilibili.com/category/getData.json?csrfToken=6d86e40f170ac908f33cda0b74c9e084&sort=2&pageSize=72&page=1&_=1568729330794"

func main() {
	resp,err := http.Get(GET_DATA_API)
	common.CheckError(err)
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	common.CheckError(err)
	fmt.Println(string(body))

	result := make(map[string]interface{})
	er := json.Unmarshal(body,&result)
	common.CheckError(er)
	fmt.Printf("%+v",result)
}