package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccssWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK, // 我看你又用了这个200的code
		"message": msg,
		"config":  data,
	})
}

func FailWithDetailed(c *gin.Context, msg string, data any) {
	// 错误的时候要返回错误code和message
	c.JSON(http.StatusOK, gin.H{
		"code":    0, // 没有这个 又用了0
		"message": msg,
		"data":    data,
	})
}

// 前端登入需要data中存在code为0
func LoginSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
		"token":   "123456",
	})
}

func InfoSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"data":    data,
	})
}

func GomessageSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"result":  data,
		"error":   "null",
	})
}
