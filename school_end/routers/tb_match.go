package routers

import (
	"chinasoccer/apis/match"
	"github.com/gin-gonic/gin"
)

func MatchRouter(r *gin.RouterGroup)  {
	r.GET("/match", match.GetMatch)
	r.GET("/competition", match.GetCompetition)
	r.GET("/footballteam", match.GetFootballTeam)
	r.POST("/match", match.AddMatch)
	r.PUT("/match", match.UpdateMatch)
	r.DELETE("/match/:id", match.DelMatch)
	r.GET("/teamplayer/:id", match.GetTeamPlayerByID)
	r.POST("/teamformation_data", match.AddTeamFormation_data)
	r.GET("/season", match.GetMatchSeason)
	r.GET("/playerdetailachieveinfo", match.GetPlayerdetailachieveinfo)
	r.GET("/playerdetailtransferinfo", match.GetPlayerdetailtransferinfo)
}