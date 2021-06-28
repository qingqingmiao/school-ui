package user

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"chinasoccer/pkg/logger"
	"chinasoccer/pkg/upload"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get User
// @Produce  json
// @Param user_id path int true "User_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/userinfo/{user_id} [get]
func GetUserByID(c *gin.Context)  {
	user_id := c.Param("user_id")
	err, info := models.GetUserByID(user_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Update User
// @Produce  json
// @Param data body models.user true "User"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/userinfo [put]
func UpdateUser(c *gin.Context) {
	err,info:=models.UpdateUser(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}

// @Summary Update UserPortrait
// @Produce  json
// @Accept multipart/form-data
// @Param Id path int true "Id"
// @Param portrait formData file true "Portrait File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/userinfoPortrait/{Id} [put]
func UpdateUserPortrait(c *gin.Context) {
	user_portrait := ""
	portraitfile, portrait, portraiterr := c.Request.FormFile("portrait");
	if portraiterr != nil {
		logger.Warn(portraiterr)
		app.Error(c, e.ERROR, portraiterr, portraiterr.Error())
		return
	}
	if portrait == nil {
		app.Error(c, e.INVALID_PARAMS, portraiterr, portraiterr.Error())
		return
	}
	portraitName := upload.GetImageName(portrait.Filename)
	portraitFullPath := upload.GetImageFullPath()
	portraitSavePath := upload.GetImagePath()
	user_portrait = portraitSavePath + portraitName
	portraitsrc := portraitFullPath + portraitName
	if !upload.CheckImageExt(portraitName) || !upload.CheckImageSize(portraitfile) {
		app.Error(c, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, portraiterr, portraiterr.Error())
		return
	}

	portraiterr = upload.CheckImage(portraitFullPath)
	if portraiterr != nil {
		logger.Warn(portraiterr)
		app.Error(c, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, portraiterr, portraiterr.Error())
		return
	}

	if err := c.SaveUploadedFile(portrait, portraitsrc); err != nil {
		logger.Warn(err)
		app.Error(c, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, err, err.Error())
		return
	}
	err,info:=models.UpdateUserPortrait(c, user_portrait)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}

// @Summary Insert User
// @Produce  json
// @Param data body models.user true "User"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/userinsert [post]
func InsertUser(c *gin.Context) {

	err,info:=models.InsertUser(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}

// @Summary Delete User
// @Produce  json
// @Param Id path int true "Id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/userdelete/{Id} [delete]
func DeleteUser(c *gin.Context) {

	err :=models.DeleteUser(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,"user deleted successfully","OK")
}

// @Summary Get Alluser
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/alluser [get]
func GetAllUser(c *gin.Context)  {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}

	userParam := map[string]interface{}{
		"page": page,
		"limit": limit,
	}
	err,info, total:=models.GetAllUser(userParam)
	//fmt.Println(total)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total}, "OK")
}

// @Summary Get Focus
// @Produce  json
// @Param user_id path int true "User_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/focus/{user_id} [get]
func GetUserFocus(c *gin.Context)  {
	user_id := c.Param("user_id")
	err, info1, info2 := models.GetUserFocus(user_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"teamvalue": info1, "playervalue": info2},"OK")
}

// @Summary Update Focus
// @Produce  json
// @Param data body models.userfocus true "Userfocus"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/focus [put]
func UpdateUserFocus(c *gin.Context)  {
	err, info := models.UpdateUserFocus(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get dashboardNum
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/dashboardnum [get]
func GetDashboardnum(c *gin.Context)  {
	err, info := models.GetDashboardnum()
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

