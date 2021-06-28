package biochemistry

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Summary Get Biochemistry
// @Produce  json
// @Param trainplan_id path int true "Trainplan_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/biochemistry/{trainplan_id} [get]
func GetBiochemistryByID(c *gin.Context)  {
	trainplan_id := c.Param("trainplan_id")
	err, info := models.GetBiochemistryByID(trainplan_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Add Biochemistry
// @Produce  json
// @Param data body models.biochemistry true "Biochemistry"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/biochemistry [post]
func AddBiochemistry(c *gin.Context)  {
	err, info := models.AddBiochemistry(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Update Biochemistry
// @Produce  json
// @Param data body models.biochemistry true "Biochemistry"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/biochemistry [put]
func UpdateBiochemistry(c *gin.Context) {
	err,info:=models.UpdateBiochemistry(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}