package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Failure(code int, message string, ctx *gin.Context) {
	ctx.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func Success(message string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"token":   message,
		"message": "响应成功",
	})
}
