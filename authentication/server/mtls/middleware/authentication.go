package middleware

import (
	"log"
	"mtls/service"

	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	// 获取client证书信息
	tlsState := c.Request.TLS
	if tlsState == nil {
		c.Next()
		return
	}
	clientCert := tlsState.PeerCertificates[0]
	if clientCert == nil {
		c.Next()
		return
	}
	serverSign := clientCert.Subject.CommonName
	// 获取服务信息
	server, err := service.ServerSvcIns.GetBySign(serverSign)
	if err != nil {
		log.Println(err)
		c.Next()
		return
	}
	c.Set("server`", server)
	c.Next()
}
