package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Authority(c *gin.Context) {
	server, isExists := c.Get("server")
	if isExists {
		fmt.Println(server)
	}
	c.Next()
}
