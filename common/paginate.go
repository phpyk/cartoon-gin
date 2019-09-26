package common

import (
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
	fv := math.Ceil(float64(dataCount / perPage))
	return int(math.Ceil(fv))
}
