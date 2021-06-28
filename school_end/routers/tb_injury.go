package routers

import (
	"chinasoccer/apis/injury"
	"github.com/gin-gonic/gin"
)

func InjuryRouter(r *gin.RouterGroup)  {
	r.GET("/injury", injury.GetInjury)
	r.POST("/injury", injury.AddInjury)
	r.GET("/doctor", injury.GetDoctor)
	r.PUT("/injury", injury.UpdateInjury)
	r.DELETE("/injury/:id", injury.DelInjury)
	r.GET("/injury/:player_id", injury.GetInjuryByID)
}
