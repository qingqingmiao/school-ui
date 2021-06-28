package routers

import (
	"chinasoccer/apis/video"
	"chinasoccer/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VideoRouter(r *gin.RouterGroup) {
	r.StaticFS("/upload/videos", http.Dir(upload.GetVideoFullPath()))
	r.GET("/video", video.GetVideo)
	r.POST("/video", video.AddVideo)
	r.PUT("/video", video.UpdateVideo)
	r.DELETE("/video/:id", video.DelVideo)
}