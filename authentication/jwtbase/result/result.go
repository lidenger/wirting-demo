package result

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func result(ctx *gin.Context, httpCode int, code int, message string, data any) {
	ctx.JSON(httpCode, gin.H{
		"code": code,
		"msg":  message,
		"data": data,
	})
}

func R(ctx *gin.Context, err error, data any) {
	if err == nil {
		result(ctx, http.StatusOK, 0, "success", data)
		return
	}
	log.Printf("%+v", err)
	result(ctx, http.StatusOK, -1, err.Error(), data)
}
