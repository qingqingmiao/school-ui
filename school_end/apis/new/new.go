package new

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// @Summary Add New
// @Produce  json
// @Param data body models.tb_new true "new"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/new [post]
func AddNew(c *gin.Context) {
	fmt.Print(c)
	err,info:=models.AddNew(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get Allnew
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Param searchText query string false "searchText"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/allnew [get]
func GetAllNew(c *gin.Context)  {
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
	err,info, total:=models.GetAllNew(messageParam)
	//fmt.Println(total)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info, "total": total}, "OK")
}

// @Summary Del new
// @Produce  json
// @Param Id path int true "Id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/new/{Id} [delete]
func DelNew(c *gin.Context) {
	err := models.DelNew(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c,"new deleted successfully","OK")
}

// @Summary Update new
// @Produce  json
// @Param data body models.tb_new true "tb_new"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/new [put]
func UpdateNew(c *gin.Context)  {
	err, info := models.UpdateNew(c)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info,"OK")
}

// @Summary Get Fournew
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/fournew [get]
func GetFourNew(c *gin.Context)  {
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
	err,info:=models.GetFourNew()
	//fmt.Println(total)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, map[string]interface{}{"value": info}, "OK")
}