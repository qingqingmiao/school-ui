package match

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get match
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param searchText query string false "searchText"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/match [get]
func GetMatch(c *gin.Context)  {
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
	matchParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"searchText": searchText,
	}
	err, info, total := models.GetAllMatch(matchParam)
	//fmt.Println(info);
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"OK")
}

// @Summary Add match
// @Produce  json
// @Param data body models.match_view true "match_view"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/match [post]
func AddMatch(c *gin.Context)  {
	err, info := models.AddMatch(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get competition
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/competition [get]
func GetCompetition(c *gin.Context)  {
	err, info := models.GetCompetition()
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get footballteam
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/footballteam [get]
func GetFootballTeam(c *gin.Context)  {
	err, info := models.GetFootballTeam()
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Update match
// @Produce  json
// @Param data body models.match_view true "match_view"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/match [put]
func UpdateMatch(c *gin.Context)  {
	err, info := models.UpdateMatch(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Del match
// @Produce  json
// @Param Id path int true "Id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/match/{Id} [delete]
func DelMatch(c *gin.Context) {
	err := models.DelMatch(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,"match deleted successfully","OK")
}

// @Summary Get teamplayer
// @Produce  json
// @Param Id path int true "Team_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/teamplayer/{Id} [get]
func GetTeamPlayerByID(c *gin.Context)  {
	err, info := models.GetTeamPlayerByID(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Add teamformation_data
// @Produce  json
// @Param team_type query string true "team_type"
// @Param data body models.teamformation_data true "teamformation_data"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/teamformation_data [post]
func AddTeamFormation_data(c *gin.Context)  {
	err, info := models.AddTeamFormation_data(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get season
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/season [get]
func GetMatchSeason(c *gin.Context)  {
	err, info := models.GetMatchSeason()
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get playerdetailachieveinfo
// @Produce  json
// @Param player_id query int false "player_id"
// @Param season query string false "season"
// @Param leaguematch_type query int false "leaguematch_type"
// @Param is_hometeam query string false "is_hometeam"
// @Param is_victory query string false "is_victory"
// @Param com_number query int false "com_number"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/playerdetailachieveinfo [get]
func GetPlayerdetailachieveinfo(c *gin.Context)  {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	player_id := -1
	if arg := c.Query("player_id"); arg != "" {
		player_id = com.StrTo(arg).MustInt()
	}
	season := ""
	if arg := c.Query("season"); arg != "" {
		season = arg
	}
	leaguematch_type := -1
	if arg := c.Query("leaguematch_type"); arg != "" {
		leaguematch_type = com.StrTo(arg).MustInt()
	}
	is_hometeam := ""
	if arg := c.Query("is_hometeam"); arg != "" {
		is_hometeam = arg
	}
	is_victory := ""
	if arg := c.Query("is_victory"); arg != "" {
		is_victory = arg
	}
	com_number := -1
	if arg := c.Query("com_number"); arg != "" {
		com_number = com.StrTo(arg).MustInt()
	}
	searchParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"player_id": player_id,
		"season": season,
		"leaguematch_type": leaguematch_type,
		"is_hometeam": is_hometeam,
		"is_victory": is_victory,
		"com_number": com_number,
	}
	err, info := models.GetPlayerdetailachieveinfo(searchParam)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get playerdetailachieveinfo
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param player_id query int false "player_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/playerdetailtransferinfo [get]
func GetPlayerdetailtransferinfo(c *gin.Context)  {
	page := -1
	if arg := c.Query("page"); arg != "" {
		page = com.StrTo(arg).MustInt()
	}
	limit := -1
	if arg := c.Query("limit"); arg != "" {
		limit = com.StrTo(arg).MustInt()
	}
	player_id := -1
	if arg := c.Query("player_id"); arg != "" {
		player_id = com.StrTo(arg).MustInt()
	}
	transferinfoParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"player_id": player_id,
	}
	err, info, total := models.GetPlayerdetailtransferinfo(transferinfoParam)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"OK")
}