package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	success = iota + 1
	fail
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  "成功",
	})
}

func SuccssWithMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  msg,
	})
}

func SuccssWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  msg,
		"data": data,
	})
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  "成功",
	})
}

func FailWithMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  msg,
	})
}

func FailWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": fail,
		"msg":  msg,
		"data": data,
	})
}

func LoginSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  msg,
		"data": data,
	})
}

func InfoSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": success,
		"msg":  msg,
		"data": data,
	})
}

func GomessageSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":   success,
		"msg":    msg,
		"result": data,
		"error":  "null",
	})
}
