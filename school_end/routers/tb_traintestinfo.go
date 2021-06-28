package routers

import (
"chinasoccer/apis/traintestinfo"
"github.com/gin-gonic/gin"
)

func TraintestinfoRouter(r *gin.RouterGroup)  {
	r.GET("/traintestinfo/:trainplan_id", traintestinfo.GetTraintestinfoByID)
	r.POST("/traintestinfo", traintestinfo.AddTraintestinfo)
	r.PUT("/traintestinfo", traintestinfo.UpdateTraintestinfo)
}