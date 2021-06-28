package routers

import (
	"chinasoccer/handler"
	"github.com/gin-gonic/gin"
)

func SysBaseRouter(r *gin.RouterGroup) {
	r.GET("/ping", handler.Ping)
}

func SysBaseRouterWithMiddleWare(r *gin.RouterGroup, middle gin.HandlerFunc )  {
	r.Use(middle).GET("/ping", handler.Ping)
}