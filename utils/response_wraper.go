package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ResponseSuccess(c *gin.Context, data interface{}) {
	if data == nil {
		data = make(map[string]string)
	}
	result := gin.H{
		"code":      200,
		"msg":       "success",
		"data":      data,
		"timestamp": time.Now().Unix(),
	}
	c.JSON(http.StatusOK, result)
}

func ResponseError(c *gin.Context, code int, msg string) {
	ret := make(map[string]string)
	result := gin.H{
		"code":      code,
		"msg":       msg,
		"data":      ret,
		"timestamp": time.Now().Unix(),
	}
	c.JSON(code, result)
}
