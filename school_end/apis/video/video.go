package video

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"chinasoccer/pkg/file"
	"chinasoccer/pkg/logger"
	"chinasoccer/pkg/upload"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get Video
// @Produce  json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param video_keyword query string false "Video_keyword"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/video [get]
func GetVideo(c *gin.Context)  {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	video_keyword := ""
	if arg := c.Query("video_keyword"); arg != "" {
		video_keyword = arg
	}
	videoParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"video_keyword": video_keyword,
	}
	err, info, total := models.GetAllVideo(videoParam)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"OK")
}

// @Summary Add Video
// @Produce  json
// @Accept multipart/form-data
// @Param video_title query string true "Video_title"
// @Param video_content query string true "Video_content"
// @Param video_keyword query string true "Video_keyword"
// @Param video_time query string true "Video_time"
// @Param storage_path formData file true "Video File"
// @Param video_thumbnail formData file true "Thumbnail File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/video [post]
func AddVideo(c *gin.Context) {
	video_title := ""
	if arg := c.Query("video_title"); arg != "" {
		video_title = arg
	}
	video_content := ""
	if arg := c.Query("video_content"); arg != "" {
		video_content = arg
	}
	video_keyword := ""
	if arg := c.Query("video_keyword"); arg != "" {
		video_keyword = arg
	}
	video_time := ""
	if arg := c.Query("video_time"); arg != "" {
		video_time = arg
	}

	storage_path := ""
	videofile, video, videoerr := c.Request.FormFile("storage_path");
	video_thumbnail := ""
	thumbnailfile, thumbnail, thumbnailerr := c.Request.FormFile("video_thumbnail");
	if videoerr != nil {
		logger.Warn(videoerr)
		app.Error(c, e.ERROR, videoerr, videoerr.Error())
		return
	}
	if video == nil {
		app.Error(c, e.INVALID_PARAMS, videoerr, videoerr.Error())
		return
	}
	if thumbnailerr != nil {
		logger.Warn(thumbnailerr)
		app.Error(c, e.ERROR, thumbnailerr, thumbnailerr.Error())
		return
	}
	if thumbnail == nil {
		app.Error(c, e.INVALID_PARAMS, thumbnailerr, thumbnailerr.Error())
		return
	}

	//视频
	videoName := upload.GetVideoName(video.Filename)
	videoFullPath := upload.GetVideoFullPath()
	videoSavePath := upload.GetVideoPath()
	storage_path = videoSavePath + videoName
	viodesrc := videoFullPath + videoName

	videoerr = upload.CheckVideo(videoFullPath)

	if videoerr != nil {
		logger.Warn(videoerr)
		app.Error(c, e.ERROR_UPLOAD_CHECK_VIDEO_FAIL, videoerr, videoerr.Error())
		return
	}

	if err := c.SaveUploadedFile(video, viodesrc); err != nil {
		logger.Warn(err)
		app.Error(c, e.ERROR_UPLOAD_SAVE_VIDEO_FAIL, err, err.Error())
		return
	}

	//视频首页图
	thumbnailName := upload.GetImageName(thumbnail.Filename)
	thumbnailFullPath := upload.GetImageFullPath()
	thumbnailSavePath := upload.GetImagePath()
	video_thumbnail = thumbnailSavePath + thumbnailName
	thumbnailsrc := thumbnailFullPath + thumbnailName

	if !upload.CheckImageExt(thumbnailName) || !upload.CheckImageSize(thumbnailfile) {
		app.Error(c, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, thumbnailerr, thumbnailerr.Error())
		return
	}

	thumbnailerr = upload.CheckImage(thumbnailFullPath)
	if thumbnailerr != nil {
		logger.Warn(thumbnailerr)
		app.Error(c, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, thumbnailerr, thumbnailerr.Error())
		return
	}

	if err := c.SaveUploadedFile(thumbnail, thumbnailsrc); err != nil {
		logger.Warn(err)
		app.Error(c, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, err, err.Error())
		return
	}

	// 获取视频时间长度
	video_length, err := file.GetMP4Duration(videofile)
	h := video_length / 3600
	video_length = video_length - h * 3600
	m := video_length / 60
	video_length = video_length - m * 60
	s := video_length
	video_duration := ""
	if(h != 0) {
		video_duration = fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	} else {
		video_duration = fmt.Sprintf( "%02d:%02d", m, s)
	}
	videodata := map[string]interface{}{
		"video_title": video_title,
		"video_content": video_content,
		"video_keyword": video_keyword,
		"video_time": video_time,
		"storage_path": storage_path,
		"video_thumbnail": video_thumbnail,
		"video_duration": video_duration,
	}

	err,info := models.AddVideo(videodata)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}

// @Summary Update Video
// @Produce  json
// @Accept multipart/form-data
// @Param video_id query int true "Video_id"
// @Param video_title query string true "Video_title"
// @Param video_content query string true "Video_content"
// @Param video_keyword query string true "Video_keyword"
// @Param video_time query string true "Video_time"
// @Param storage_path formData file true "Video File"
// @Param video_thumbnail formData file true "Thumbnail File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/video [put]
func UpdateVideo(c *gin.Context) {
	video_id := -1
	if arg := c.Query("video_id"); arg != "" {
		video_id = com.StrTo(arg).MustInt()
	}
	video_title := ""
	if arg := c.Query("video_title"); arg != "" {
		video_title = arg
	}
	video_content := ""
	if arg := c.Query("video_content"); arg != "" {
		video_content = arg
	}
	video_keyword := ""
	if arg := c.Query("video_keyword"); arg != "" {
		video_keyword = arg
	}
	video_time := ""
	if arg := c.Query("video_time"); arg != "" {
		video_time = arg
	}

	storage_path := ""
	videofile, video, videoerr := c.Request.FormFile("storage_path");
	video_thumbnail := ""
	thumbnailfile, thumbnail, thumbnailerr := c.Request.FormFile("video_thumbnail");
	if videoerr != nil {
		logger.Warn(videoerr)
		app.Error(c, e.ERROR, videoerr, videoerr.Error())
		return
	}
	if video == nil {
		app.Error(c, e.INVALID_PARAMS, videoerr, videoerr.Error())
		return
	}
	if thumbnailerr != nil {
		logger.Warn(thumbnailerr)
		app.Error(c, e.ERROR, thumbnailerr, thumbnailerr.Error())
		return
	}
	if thumbnail == nil {
		app.Error(c, e.INVALID_PARAMS, thumbnailerr, thumbnailerr.Error())
		return
	}

	//视频
	videoName := upload.GetVideoName(video.Filename)
	videoFullPath := upload.GetVideoFullPath()
	videoSavePath := upload.GetVideoPath()
	storage_path = videoSavePath + videoName
	viodesrc := videoFullPath + videoName

	videoerr = upload.CheckVideo(videoFullPath)

	if videoerr != nil {
		logger.Warn(videoerr)
		app.Error(c, e.ERROR_UPLOAD_CHECK_VIDEO_FAIL, videoerr, videoerr.Error())
		return
	}

	if err := c.SaveUploadedFile(video, viodesrc); err != nil {
		logger.Warn(err)
		app.Error(c, e.ERROR_UPLOAD_SAVE_VIDEO_FAIL, err, err.Error())
		return
	}

	//视频首页图
	thumbnailName := upload.GetImageName(thumbnail.Filename)
	thumbnailFullPath := upload.GetImageFullPath()
	thumbnailSavePath := upload.GetImagePath()
	video_thumbnail = thumbnailSavePath + thumbnailName
	thumbnailsrc := thumbnailFullPath + thumbnailName

	if !upload.CheckImageExt(thumbnailName) || !upload.CheckImageSize(thumbnailfile) {
		app.Error(c, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, thumbnailerr, thumbnailerr.Error())
		return
	}

	thumbnailerr = upload.CheckImage(thumbnailFullPath)
	if thumbnailerr != nil {
		logger.Warn(thumbnailerr)
		app.Error(c, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, thumbnailerr, thumbnailerr.Error())
		return
	}

	if err := c.SaveUploadedFile(thumbnail, thumbnailsrc); err != nil {
		logger.Warn(err)
		app.Error(c, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, err, err.Error())
		return
	}

	// 获取视频时间长度
	video_length, err := file.GetMP4Duration(videofile)
	h := video_length / 3600
	video_length = video_length - h * 3600
	m := video_length / 60
	video_length = video_length - m * 60
	s := video_length
	video_duration := ""
	if(h != 0) {
		video_duration = fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	} else {
		video_duration = fmt.Sprintf( "%02d:%02d", m, s)
	}
	videodata := map[string]interface{}{
		"video_id": video_id,
		"video_title": video_title,
		"video_content": video_content,
		"video_keyword": video_keyword,
		"video_time": video_time,
		"storage_path": storage_path,
		"video_thumbnail": video_thumbnail,
		"video_duration": video_duration,
	}

	err,info := models.UpdateVideo(videodata)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}

// @Summary Del Video
// @Produce  json
// @Param Id path int true "Id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/video/{Id} [delete]
func DelVideo(c *gin.Context) {
	err := models.DelVideo(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,"video deleted successfully","OK")
}