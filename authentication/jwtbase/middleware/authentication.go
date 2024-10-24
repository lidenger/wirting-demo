package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jwtbase/result"
	"jwtbase/service"
	"log"
)

func Authentication(c *gin.Context) {
	jwt := c.GetHeader("Authorization")
	if jwt == "" {
		c.Next()
		return
	}
	// 验证jwt
	sign, err := service.VerifyJwt(jwt)
	if err != nil {
		result.R(c, errors.New("无效的JWT"), "")
		c.Abort()
	}
	// 获取服务信息
	server, err := service.ServerSvcIns.GetBySign(sign)
	if err != nil {
		log.Printf("+%v", err)
		c.Next()
		return
	}
	c.Set("server", server)
	c.Next()
}
