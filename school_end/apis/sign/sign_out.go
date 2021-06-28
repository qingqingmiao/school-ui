package sign

import (
	"chinasoccer/pkg/app"
	"github.com/gin-gonic/gin"
)

// @Summary Post LogOutUser
// @Product json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /chinasoccer/user/logout [post]
func LogOut(c *gin.Context) {
	app.OK(c, "success", "登出成功")
}