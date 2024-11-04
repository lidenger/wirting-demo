package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func result(ctx *gin.Context, httpCode int, code int, message string, data any) {
	ctx.JSON(httpCode, gin.H{
		"code": code,
		"msg":  message,
		"data": data,
	})
}

func success(ctx *gin.Context, data any) {
	result(ctx, http.StatusOK, 0, "success", data)
}

func Health(c *gin.Context) {
	data := make(map[string]string, 2)
	data["status"] = "ok"
	data["time"] = time.Now().Format(time.DateTime)
	success(c, data)
}
