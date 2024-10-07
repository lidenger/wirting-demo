package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"requestsign/result"
)

func Authority(c *gin.Context) {
	server, isExists := c.Get("server")
	if !isExists || !checkAuth(server) {
		result.R(c, errors.New("无权访问"), nil)
		c.Abort()
		return
	}
	log.Printf("%+v", server)
	c.Next()
}

func checkAuth(server any) bool {
	// TODO 鉴权逻辑
	return true
}
