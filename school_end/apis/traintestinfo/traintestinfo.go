package traintestinfo

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Summary Get Traintestinfo
// @Produce  json
// @Param trainplan_id path int true "Trainplan_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/traintestinfo/{trainplan_id} [get]
func GetTraintestinfoByID(c *gin.Context)  {
	trainplan_id := c.Param("trainplan_id")
	err, info := models.GetTraintestinfoByID(trainplan_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Add Traintestinfo
// @Produce  json
// @Param data body models.traintestinfo true "Traintestinfo"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/traintestinfo [post]
func AddTraintestinfo(c *gin.Context)  {
	err, info := models.AddTraintestinfo(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Update Traintestinfo
// @Produce  json
// @Param data body models.traintestinfo true "Traintestinfo"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/traintestinfo [put]
func UpdateTraintestinfo(c *gin.Context) {
	err,info:=models.UpdateTraintestinfo(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}