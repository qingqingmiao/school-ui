package playermatch

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Summary Get Playermatch
// @Produce  json
// @Param player_id path int true "Player_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/playerdetailachieve/{player_id} [get]
func GetPlayermatchByPlayerID(c *gin.Context)  {
	player_id := c.Param("player_id")
	err, info := models.GetPlayermatchByPlayerID(player_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get Playermatchteam
// @Produce  json
// @Param match_id path int true "Match_id"
// @Param team_id path int true "Team_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/playermatchteam/{match_id}/{team_id} [get]
func GetPlayermatchteamByTeamID(c *gin.Context)  {
	match_id := c.Param("match_id")
	team_id := c.Param("team_id")
	err, info := models.GetPlayermatchteamByTeamID(match_id, team_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get Playermatchplayer
// @Produce  json
// @Param match_id path int true "Match_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/playermatchplayer/{match_id} [get]
func GetPlayermatchplayerByMatchID(c *gin.Context)  {
	match_id := c.Param("match_id")
	err, info := models.GetPlayermatchplayerByMatchID(match_id)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}