package routers

import (
	"chinasoccer/apis/sign"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup)  {
	r.POST("/auth",sign.GetAuth)
}


