package transfer

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Get transfer
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param player_id query int false "player_id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/transfer [get]
func GetTransfer(c *gin.Context)  {
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
	transferParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"player_id": player_id,
	}
	err, info, total := models.GetTransfer(transferParam)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total},"OK")
}