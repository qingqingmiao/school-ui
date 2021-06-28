package uploadexcel

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"chinasoccer/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @Summary Import Excel
// @Produce  json
// @Accept multipart/form-data
// @Param file formData file true "Excel File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/uploadexcel [post]
func UploadExcel(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		logger.Warn(err)
		app.Error(c, e.ERROR, err, err.Error())
		return
	}

	err2:= models.AddExcel(file)
	if err2 !=nil{
		app.Error(c,e.ERROR,err2,err2.Error())
		return
	}
	app.OK(c, nil,"OK")
}