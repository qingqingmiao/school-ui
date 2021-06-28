package player

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get player
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param ageGroup query []int false "ageGroup"
// @Param heightGroup query []int false "heightGroup"
// @Param weightGroup query []int false "weightGroup"
// @Param initials query string false "initials"
// @Param leaguematch_type query []int false "leaguematch_type"
// @Param team_name query string false "team_name"
// @Param position query string false "position"
// @Param state query string false "state"
// @Param technical_level query string false "technical_level"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/player [get]
func GetPlayer(c *gin.Context)  {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	ageGroup := []string{}
	if arg, _ := c.GetQueryArray("ageGroup[]"); len(arg) != 0 {
		ageGroup = arg
	}
	heightGroup := []string{}
	if arg, _ := c.GetQueryArray("heightGroup[]"); len(arg) != 0 {
		heightGroup = arg
	}
	weightGroup := []string{}
	if arg, _ := c.GetQueryArray("weightGroup[]"); len(arg) != 0 {
		weightGroup = arg
	}
	initials := ""
	if arg := c.Query("initials"); arg != "" {
		initials = arg
	}
	leaguematch_type := []string{}
	if arg, _ := c.GetQueryArray("leaguematch_type[]"); len(arg) != 0 {
		leaguematch_type = arg
	}
	team_name := ""
	if arg := c.Query("team_name"); arg != "" {
		team_name = arg
	}
	position := ""
	if arg := c.Query("position"); arg != "" {
		position = arg
	}
	state := ""
	if arg := c.Query("state"); arg != "" {
		state = arg
	}
	technical_level := ""
	if arg := c.Query("technical_level"); arg != "" {
		technical_level = arg
	}
	searchText := ""
	if arg := c.Query("searchText"); arg != "" {
		searchText = arg
	}
	playerParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"ageGroup[]": ageGroup,
		"heightGroup[]": heightGroup,
		"weightGroup[]": weightGroup,
		"initials": initials,
		"leaguematch_type[]": leaguematch_type,
		"team_name": team_name,
		"position": position,
		"state": state,
		"technical_level": technical_level,
		"searchText": searchText,
	}
	err,info, total:=models.GetAllPlayer(playerParam)
	//fmt.Println(total);
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total}, "OK")
}

// @Summary Add player
// @Produce  json
// @Param data body models.player true "Player"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/player [post]
func AddPlayer(c *gin.Context)  {
	err, info := models.AddPlayer(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Update player
// @Produce  json
// @Param data body models.player true "Player"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/player [put]
func UpdatePlayer(c *gin.Context)  {
	err, info := models.UpdatePlayer(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Del player
// @Produce  json
// @Param Id path int true "Id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/player/{Id} [delete]
func DelPlayer(c *gin.Context) {
	err := models.DelPlayer(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,"player deleted successfully","OK")
}

// @Summary Get allPlayer
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/allplayer [get]
func GetAllPlayerKeyData(c *gin.Context) {
	err, info := models.GetAllPlayerKeyData()
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}
