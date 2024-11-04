package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jwtenhance/result"
	"log"
)

func Authority(c *gin.Context) {
	server, isExists := c.Get("server")
	if !isExists {
		result.R(c, errors.New("无效服务信息，无权访问"), "")
		c.Abort()
	}
	log.Printf("server:%+v\n", server)
	c.Next()
}
