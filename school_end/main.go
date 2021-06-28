package main

import (
	_ "chinasoccer/docs"
	"chinasoccer/models"
	"chinasoccer/pkg/gredis"
	"chinasoccer/pkg/logger"
	"chinasoccer/pkg/setting"
	"chinasoccer/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	setting.Setup()
	models.Setup()
	gredis.Setup()
	logger.Setup()
}
// @title chinasoccer-backendServer
// @version 1.0
// @description base on gin
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	routersInit := routers.InitRouter()
	routersInit.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routersInit.Run(endPoint)
}