package handler

import (
	"github.com/gin-gonic/gin"
	"jwtbase/result"
	"jwtbase/service"
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

func FetchOrderByNo(c *gin.Context) {
	orderNo := c.Param("orderNo")
	order, err := service.OrderSvcIns.FetchOrderByNo(orderNo)
	result.R(c, err, order)
}

func GenToken(c *gin.Context) {

}
