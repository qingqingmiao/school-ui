package routers

import (
	"chinasoccer/apis/uploadexcel"
	"github.com/gin-gonic/gin"
)

func UploadexcelRouter(r *gin.RouterGroup) {
	r.POST("/uploadexcel", uploadexcel.UploadExcel)
}