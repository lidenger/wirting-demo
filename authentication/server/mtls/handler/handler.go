package handler

import (
	"github.com/gin-gonic/gin"
	"mtls/result"
	"mtls/service"
	"time"
)

func Health(c *gin.Context) {
	data := make(map[string]string, 2)
	data["status"] = "ok"
	data["time"] = time.Now().Format(time.DateTime)
	result.Success(c, data)
}

func FetchOrderByNo(c *gin.Context) {
	orderNo := c.Param("orderNo")
	order, err := service.OrderSvcIns.FetchOrderByNo(orderNo)
	result.R(c, err, order)
}
