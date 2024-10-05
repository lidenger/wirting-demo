package router

import (
	"accesstoken/handler"
	"accesstoken/middleware"
	"github.com/gin-gonic/gin"
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

	g.POST("/access_token", handler.GenAccessToken)

	v1 := g.Group("/v1")
	v1.Use(middleware.Authority)
	{
		order := v1.Group("/order")
		order.GET(":orderNo", handler.FetchOrderByNo)
	}

}
