package utils

import (
	"fmt"
	"math"
)

func AppendPaginateData(result map[string]interface{}, count, currentPage, perPage int, url string) map[string]interface{} {
	result["current_page"] = currentPage
	result["from"] = 1
	lastPage := calcuLastPage(count, perPage)
	result["last_page"] = lastPage
	result["path"] = url
	result["total"] = count
	result["per_page"] = perPage
	return result
}

func calcuLastPage(dataCount, perPage int) int {
	var pcount float64
	pcount = float64(dataCount)/float64(perPage)
	fmt.Println(pcount)
	fv := math.Ceil(pcount)
	return int(math.Ceil(fv))
}
