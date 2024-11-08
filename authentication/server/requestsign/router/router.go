package router

import (
	"github.com/gin-gonic/gin"
	"requestsign/handler"
	"requestsign/middleware"
)

func Initialize() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.Use(middleware.Authentication)
	api(g)
	return g
}

func api(g *gin.Engine) {
	g.GET("/health", handler.Health)

	v1 := g.Group("/v1")
	v1.Use(middleware.Authority)
	{
		order := v1.Group("/order")
		order.GET(":orderNo", handler.FetchOrderByNo)
	}

}
