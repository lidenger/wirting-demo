package handler

import (
	"accesstoken/result"
	"accesstoken/service"
	"github.com/gin-gonic/gin"
	"time"
)

func Health(c *gin.Context) {
	data := make(map[string]string, 2)
	data["status"] = "ok"
	data["time"] = time.Now().Format(time.DateTime)
	result.R(c, nil, data)
}

func GenAccessToken(c *gin.Context) {
	p := &GenAccessTokenParam{}
	if err := c.ShouldBindJSON(&p); err != nil {
		result.R(c, err, "")
		return
	}
	token, err := service.AccessTokenSvcIns.GenAccessToken(p.Sign, p.Ak, p.Sk)
	result.R(c, err, token)
}

func FetchOrderByNo(c *gin.Context) {
	orderNo := c.Param("orderNo")
	order, err := service.OrderSvcIns.FetchOrderByNo(orderNo)
	result.R(c, err, order)
}
