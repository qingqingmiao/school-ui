package message

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Add Message
// @Produce  json
// @Param data body models.tb_message true "message"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/message [post]
func AddMessage(c *gin.Context) {
	fmt.Print(c)
	err,info:=models.AddMessage(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get Allmessage
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param searchText query string false "searchText"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/allmessage [get]
func GetAllMessage(c *gin.Context)  {
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
	messageParam := map[string]interface{}{
		"page": page,
		"limit": limit,
		"searchText": searchText,
	}
	err,info, total:=models.GetAllMessage(messageParam)
	//fmt.Println(total)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total}, "OK")
}

// @Summary Del message
// @Produce  json
// @Param Id path int true "Id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/message/{Id} [delete]
func DelMessage(c *gin.Context) {
	err := models.DelMessage(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,"message deleted successfully","OK")
}

// @Summary Update message
// @Produce  json
// @Param data body models.tb_message true "tb_message"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/message [put]
func UpdateMessage(c *gin.Context)  {
	err, info := models.UpdateMessage(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get Fourmessage
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/fourmessage [get]
func GetFourMessage(c *gin.Context)  {
	//page := -1
	//if arg := c.Query("page"); arg != "" {
	//	page = com.StrTo(arg).MustInt()
	//}
	//limit := -1
	//if arg := c.Query("limit"); arg != "" {
	//	limit = com.StrTo(arg).MustInt()
	//}

	//messageParam := map[string]interface{}{
	//	"page": page,
	//	"limit": limit,
	//}
	err,info:=models.GetFourMessage()
	//fmt.Println(total)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info}, "OK")
}