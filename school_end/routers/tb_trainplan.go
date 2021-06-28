package routers

import (
	"chinasoccer/apis/trainplan"
	"github.com/gin-gonic/gin"
)

func TrainplanRouter(r *gin.RouterGroup)  {
	r.GET("/trainplan", trainplan.GetTrainplan)
	r.GET("/trainplan/:id", trainplan.GetTrainplanByID)
	r.POST("/trainplan", trainplan.AddTrainplan)
	r.PUT("/schedule/:trainplan_id", trainplan.AddSchedule)
	r.DELETE("/schedule/:trainplan_id", trainplan.DelSchedule)
}