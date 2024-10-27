package middleware

import (
	"github.com/gin-gonic/gin"
	"jwtenhance/result"
	"jwtenhance/service"
	"log"
)

func Authentication(c *gin.Context) {
	jwt := c.GetHeader("Authorization")
	if jwt == "" {
		c.Next()
		return
	}
	// 验证jwt
	serverSign, err := service.VerifyJwt(jwt)
	if err != nil {
		result.AuthErr(c, "无效的JWT")
		c.Abort()
	}
	// 获取服务信息
	server, err := service.ServerSvcIns.GetBySign(serverSign)
	if err != nil {
		log.Printf("+%v", err)
		c.Next()
		return
	}
	c.Set("server", server)
	c.Next()
}
