package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jwtenhance/param"
	"jwtenhance/result"
	"jwtenhance/service"
	"time"
)

func success(ctx *gin.Context, data any) {
	result.R(ctx, nil, data)
}

func Health(c *gin.Context) {
	data := make(map[string]string, 2)
	data["status"] = "ok"
	data["time"] = time.Now().Format(time.DateTime)
	success(c, data)
}

func GenToken(c *gin.Context) {
	p := &param.AkSk{}
	if err := c.ShouldBindJSON(&p); err != nil {
		result.R(c, err, nil)
		return
	}
	sk := service.AkSkSvcIns.GetSk(p.Ak)
	if sk == nil || sk.Sk != p.Sk {
		result.R(c, errors.New("无效的aksk"), nil)
		return
	}
	jwt, err := service.GenJwt(sk.ServerSign, c.ClientIP())
	if err != nil {
		result.R(c, err, nil)
		return
	}
	success(c, jwt)
}

func FetchOrderByNo(c *gin.Context) {
	orderNo := c.Param("orderNo")
	order, err := service.OrderSvcIns.FetchOrderByNo(orderNo)
	result.R(c, err, order)
}
