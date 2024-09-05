package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccssWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"config":  data,
	})
}

func FailWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"config":  data,
	})
}

func LoginSuccessDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": msg,
		"data":    data,
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
