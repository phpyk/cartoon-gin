package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS_CODE = 1
	FAILD_CODE   = 0
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	State   int         `json:"state"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func (g *Gin) Success(data interface{}) {
	g.C.JSON(http.StatusOK, gin.H{
		"state":   SUCCESS_CODE,
		"message": "success",
		"result":  data,
	})
}

func (g *Gin) Failed(errmsg string) {
	g.C.JSON(http.StatusInternalServerError, gin.H{
		"state":   FAILD_CODE,
		"message": errmsg,
	})
}

func (g *Gin) Response(httpCode int, resp Response) {
	g.C.JSON(httpCode, resp)
}

func ResponseWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
