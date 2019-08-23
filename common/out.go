package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
const (
	SUCCESS_CODE = 1
	FAILD_CODE = 0
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Success(data interface{}) {
	g.C.JSON(http.StatusOK,gin.H{
		"state":   SUCCESS_CODE,
		"message": "success",
		"result":data,
	})
}

func (g *Gin) Failed(errmsg string,err error,data interface{})  {
	g.C.JSON(http.StatusInternalServerError,gin.H{
		"state":   FAILD_CODE,
		"message": errmsg,
		"err_msg":err.Error(),
	})
}