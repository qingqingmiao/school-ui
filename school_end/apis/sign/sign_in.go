package sign

import (
	"chinasoccer/models"
	"chinasoccer/pkg/app"
	"chinasoccer/pkg/e"
	"chinasoccer/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Summary Post User_Account
// @Product json
// @Param data body models.User_Account true "User_Account"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/user/login [post]
func GetAuth(c *gin.Context) {
	userDataJson := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&userDataJson)
	var userdata models.User_Account
	userdata.Account = userDataJson["username"].(string)
	userdata.Password = userDataJson["password"].(string)
	userid := userdata.Account
	passwd := userdata.Password
	success, err := userdata.CheckPasswd()
	if err != nil {
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	if success {
		token, err := utils.GenerateToken(userid, passwd)
		fmt.Println(err)
		data := new(app.Data)
		data.Token = token
		app.OK(c, data, "OK")
	} else {
		app.PswError(c, err, "Password error!!!")
	}
}
// @Summary Get UserInfo
// @Product json
// @Param account query string true "Account"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/user/info [get]
func CheckAuth(c *gin.Context) {
	account := c.Query("account")
	err, info := models.GetUserAuthByAccount(account)
	if err !=nil{
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	responseData := new(app.ResponseData)
	responseData.Roles = append(responseData.Roles, "admin")
	responseData.Introduction = "I am a super administrator"
	responseData.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	responseData.Name = "Super Admin"
	app.OK(c, map[string]interface{}{"info": info, "responseData": responseData},"OK")
}

//func CheckAuth(c *gin.Context) {
//	responseData := new(app.ResponseData)
//	responseData.Roles = append(responseData.Roles, "admin")
//	responseData.Introduction = "I am a super administrator"
//	responseData.Avatar = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
//	responseData.Name = "Super Admin"
//	app.OK(c, responseData, "有权限")
//}