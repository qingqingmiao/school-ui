package routers

import (
	"chinasoccer/apis/player"
	"github.com/gin-gonic/gin"
)

func PlayerRouter(r *gin.RouterGroup)  {
	//r.Use(middleWare.JWT())
	//{
	//	r.GET("/player", player.GetPlayer)
	//}
	r.GET("/player", player.GetPlayer)
	r.POST("/player", player.AddPlayer)
	r.PUT("/player", player.UpdatePlayer)
	r.DELETE("/player/:id", player.DelPlayer)
	r.GET("/allplayer", player.GetAllPlayerKeyData)
}
