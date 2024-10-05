package middleware

import (
	"accesstoken/service"
	"github.com/gin-gonic/gin"
	"log"
)

func Authentication(c *gin.Context) {
	token := c.GetHeader("accessToken")
	if token == "" {
		c.Next()
		return
	}
	sign, err := service.AccessTokenSvcIns.VerifyAccessToken(token)
	if err != nil {
		log.Printf("+%v", err)
		c.Next()
		return
	}
	server, err := service.ServerSvcIns.GetBySign(sign)
	if err != nil {
		log.Printf("+%v", err)
		c.Next()
		return
	}
	c.Set("server", server)
	c.Next()
}
