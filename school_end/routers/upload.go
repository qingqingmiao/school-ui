package routers

import (
	"chinasoccer/apis/uploadimg"
	"chinasoccer/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadimgRouter(r *gin.RouterGroup) {
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.POST("/uploadimg", uploadimg.UploadImage)
}