package uploadimg

import (
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"chinasoccer/pkg/logger"
	"chinasoccer/pkg/upload"
	"github.com/gin-gonic/gin"
)

// @Summary Import Image
// @Produce  json
// @Accept multipart/form-data
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/uploadimg [post]
func UploadImage(c *gin.Context) {
	//appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logger.Warn(err)
		//appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		app.Error(c, e.ERROR, err, err.Error())
		return
	}

	if image == nil {
		//appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		app.Error(c, e.INVALID_PARAMS, err, err.Error())
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		//appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		app.Error(c, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, err, err.Error())
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logger.Warn(err)
		//appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		app.Error(c, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, err, err.Error())
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logger.Warn(err)
		//appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		app.Error(c, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, err, err.Error())
		return
	}

	//appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
	//	"image_url":      upload.GetImageFullUrl(imageName),
	//	"image_save_url": savePath + imageName,
	//})
	app.OK(c, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	}, "Ok")
}
