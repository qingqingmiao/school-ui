package catapult

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Summary Get Catapult
// @Produce  json
// @Param trainplan_id path int true "Trainplan_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/catapult/{trainplan_id} [get]
func GetCatapultByID(c *gin.Context)  {
	trainplan_id := c.Param("trainplan_id")
	err, info := models.GetCatapultByID(trainplan_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Add Catapult
// @Produce  json
// @Param data body models.catapult true "Catapult"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/catapult [post]
func AddCatapult(c *gin.Context)  {
	err, info := models.AddCatapult(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Update Catapult
// @Produce  json
// @Param data body models.catapult true "Catapult"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/catapult [put]
func UpdateCatapult(c *gin.Context) {
	err,info:=models.UpdateCatapult(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}