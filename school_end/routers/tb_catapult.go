package routers

import (
	"chinasoccer/apis/catapult"
	"github.com/gin-gonic/gin"
)

func CatapultRouter(r *gin.RouterGroup)  {
	r.GET("/catapult/:trainplan_id", catapult.GetCatapultByID)
	r.POST("/catapult", catapult.AddCatapult)
	r.PUT("/catapult", catapult.UpdateCatapult)
}
