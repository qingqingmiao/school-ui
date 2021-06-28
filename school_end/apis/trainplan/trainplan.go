package trainplan

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get Trainplan
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/trainplan [get]
func GetTrainplan(c *gin.Context)  {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	trainplanParam := map[string]interface{}{
		"page": page,
		"limit": limit,
	}
	err, info, total:=models.GetAllTrainplan(trainplanParam)
	//fmt.Println(info);
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"OK")
}

// @Summary Get Trainplan
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/trainplan/{id} [get]
func GetTrainplanByID(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	err, info := models.GetTrainplanByID(id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Add Trainplan
// @Produce  json
// @Param data body models.trainplan true "trainplan"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/trainplan [post]
func AddTrainplan(c *gin.Context) {
	err,info:=models.AddTrainplan(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Update Schedule
// @Produce  json
// @Param trainplan_id path int false "trainplan_id"
// @Param data body models.traincontent true "traincontent"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/schedule/{trainplan_id} [put]
func AddSchedule(c *gin.Context) {
	//trainplan_id := com.StrTo(c.Param("trainplan_id")).MustInt()
	trainplan_id := com.StrTo(c.Params.ByName("trainplan_id")).MustInt()
	err,info:=models.AddSchedule(c,trainplan_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Delete Schedule
// @Produce  json
// @Param trainplan_id path int false "trainplan_id"
// @Param data body models.traincontent true "traincontent"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/schedule/{trainplan_id} [delete]
func DelSchedule(c *gin.Context) {
	//trainplan_id := com.StrTo(c.Param("trainplan_id")).MustInt()
	trainplan_id := com.StrTo(c.Params.ByName("trainplan_id")).MustInt()
	err,info:=models.DelSchedule(c,trainplan_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}