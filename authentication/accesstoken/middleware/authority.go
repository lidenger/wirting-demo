package middleware

import (
	"accesstoken/result"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

func Authority(c *gin.Context) {
	server, isExists := c.Get("server")
	if !isExists || !checkAuth(server) {
		result.R(c, errors.New("无权访问"), "")
		c.Abort()
		return
	}
	log.Printf("%+v", server)
	c.Next()
}

func checkAuth(server any) bool {
	// TODO 检测权限逻辑
	return true
}
