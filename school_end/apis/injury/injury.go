package injury

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get injury
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param searchText query string false "searchText"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/injury [get]
func GetInjury(c *gin.Context)  {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	searchText := ""
	if arg := c.Query("searchText"); arg != "" {
		searchText = arg
	}
	injuryParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"searchText": searchText,
	}
	err, info, total := models.GetAllInjury(injuryParam)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"OK")
}

// @Summary Get injurybyid
// @Produce  json
// @Param player_id path int true "player_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/injury/{player_id} [get]
func GetInjuryByID(c *gin.Context)  {
	player_id := c.Param("player_id")
	err, info := models.GetInjuryByID(player_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Add injury
// @Produce  json
// @Param data body models.injury_view true "injury_view"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/injury [post]
func AddInjury(c *gin.Context)  {
	err, info := models.AddInjury(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get doctor
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/doctor [get]
func GetDoctor(c *gin.Context)  {
	err, info := models.GetDoctor()
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Update injury
// @Produce  json
// @Param data body models.injury_view true "injury_view"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/injury [put]
func UpdateInjury(c *gin.Context)  {
	err, info := models.UpdateInjury(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Del injury
// @Produce  json
// @Param Id path int true "Id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/injury/{Id} [delete]
func DelInjury(c *gin.Context) {
	err := models.DelInjury(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,"injury deleted successfully","OK")
}