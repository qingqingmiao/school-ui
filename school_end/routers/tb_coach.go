package routers

import (
	"chinasoccer/apis/coach"
	"github.com/gin-gonic/gin"
)

func CoachRouter(r *gin.RouterGroup)  {
	r.GET("/coach", coach.GetCoach)
	r.POST("/coach", coach.AddCoach)
	r.PUT("/updatecoach/:id", coach.UpdateCoach)
	r.DELETE("/delcoach/:id", coach.DelCoach)
}

