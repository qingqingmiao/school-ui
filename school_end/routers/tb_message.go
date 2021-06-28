package routers

import (
	"chinasoccer/apis/message"
	"github.com/gin-gonic/gin"
)

func MessageRouter(r *gin.RouterGroup)  {
	r.POST("/message", message.AddMessage)
	r.GET("/allmessage", message.GetAllMessage)
	r.DELETE("/message/:id", message.DelMessage)
	r.PUT("/message", message.UpdateMessage)
	r.GET("/fourmessage", message.GetFourMessage)
}
