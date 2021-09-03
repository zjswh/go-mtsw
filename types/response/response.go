package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code         int         `json:"code"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
	ErrorCode    int         `json:"errorCode"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Success(data interface{}, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		200,
		data,
		"",
		0,
	})
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		200,
		data,
		msg,
		code,
	})
}

func SystemError(msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusInternalServerError, Response{
		http.StatusInternalServerError,
		"",
		msg,
		0,
	})
	c.Abort()
}

func ParamError(message string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		200,
		"",
		message,
		2,
	})
}

func DbError(message string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		200,
		"",
		message,
		3,
	})
}
