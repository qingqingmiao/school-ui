package coach

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Summary Get Coach
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/coach [get]
func GetCoach(c *gin.Context)  {
	err,info:=models.GetAllCoach()
	//fmt.Println(info);
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}

// @Summary Add Coach
// @Produce  json
// @Param data body models.coach true "coach"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/coach [post]
func AddCoach(c *gin.Context) {
	err,info:=models.AddCoach(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}

// @Summary Update Coach
// @Produce  json
// @Param Id path int true "Id"
// @Param data body models.coach true "coach"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/updatecoach/{Id} [put]
func UpdateCoach(c *gin.Context) {
	err,info:=models.UpdateCoach(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}

// @Summary Del Coach
// @Produce  json
// @Param Id path int true "Id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/delcoach/{Id} [delete]
func DelCoach(c *gin.Context) {
	err,info:=models.DelCoach(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,info,"OK")
}