package routers

import (
	"chinasoccer/apis/student"
	"github.com/gin-gonic/gin"
)

func StudentRouter(r *gin.RouterGroup)  {
	//fmt.Println(123)
	r.GET("/student",student.GetStudent)
}