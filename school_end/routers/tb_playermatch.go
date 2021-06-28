package routers

import (
	"chinasoccer/apis/playermatch"
	"github.com/gin-gonic/gin"
)

func PlayermatchRouter(r *gin.RouterGroup)  {
	r.GET("/playerdetailachieve/:player_id", playermatch.GetPlayermatchByPlayerID)
	r.GET("/playermatchteam/:match_id/:team_id", playermatch.GetPlayermatchteamByTeamID)
	r.GET("/playermatchplayer/:match_id", playermatch.GetPlayermatchplayerByMatchID)
}
