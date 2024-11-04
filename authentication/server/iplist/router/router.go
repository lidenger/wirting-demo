package router

import (
	"github.com/gin-gonic/gin"
	"iplist/handler"
	"iplist/middleware"
)

func Initialize() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.Use(middleware.Authentication)
	g.Use(middleware.Authority)
	api(g)
	return g
}

func api(g *gin.Engine) {
	g.GET("/health", handler.Health)

}
