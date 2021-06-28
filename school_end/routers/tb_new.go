package routers

import (
	"chinasoccer/apis/new"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.RouterGroup)  {
	r.POST("/new", new.AddNew)
	r.GET("/allnew", new.GetAllNew)
	r.DELETE("/new/:id", new.DelNew)
	r.PUT("/new", new.UpdateNew)
	r.GET("/fournew", new.GetFourNew)
}
