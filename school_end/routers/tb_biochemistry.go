package routers

import (
	"chinasoccer/apis/biochemistry"
	"github.com/gin-gonic/gin"
)

func BiochemistryRouter(r *gin.RouterGroup)  {
	r.GET("/biochemistry/:trainplan_id", biochemistry.GetBiochemistryByID)
	r.POST("/biochemistry", biochemistry.AddBiochemistry)
	r.PUT("/biochemistry", biochemistry.UpdateBiochemistry)
}
