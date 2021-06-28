package routers

import (
	"chinasoccer/apis/sign"
	"chinasoccer/apis/user"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup)  {
	r.POST("/user/login", sign.GetAuth)
	//r.Use(middleWare.JWT())
	//{
	//	r.GET("/user/info", sign.CheckAuth)
	//}
	r.GET("/user/info", sign.CheckAuth)
	r.POST("/user/logout", sign.LogOut)
	r.GET("/userinfo/:user_id", user.GetUserByID)
	r.PUT("/userinfo", user.UpdateUser)
	r.PUT("/userinfoPortrait/:user_id", user.UpdateUserPortrait)
	r.GET("/alluser", user.GetAllUser)
	r.POST("/userinsert", user.InsertUser)
	r.DELETE("/userdelete/:id", user.DeleteUser)
	r.GET("/focus/:user_id", user.GetUserFocus)
	r.PUT("/focus", user.UpdateUserFocus)
	r.GET("/dashboardnum", user.GetDashboardnum)
}