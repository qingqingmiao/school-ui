package routers

import (
	middlewares "chinasoccer/middleWare"
	_ "chinasoccer/models"
	"github.com/gin-gonic/gin"
)

//无需认证
func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/chinasoccer")
	SysBaseRouter(r)
	AuthRouter(r)
	StudentRouter(r)
	CoachRouter(r)
	UserRouter(r)
	InjuryRouter(r)
	MatchRouter(r)
	MessageRouter(r)
	NewRouter(r)
	PlayerRouter(r)
	VideoRouter(r)
	TrainplanRouter(r)
	TraintestinfoRouter(r)
	TransferRouter(r)
	PlayermatchRouter(r)
	BiochemistryRouter(r)
	CatapultRouter(r)
	UploadimgRouter(r)
	UploadexcelRouter(r)
}

func sysCheckRoleRouter(r *gin.RouterGroup, authMiddle gin.HandlerFunc) {
	SysBaseRouterWithMiddleWare(r, authMiddle)
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())

	g := r.Group("")
	//jwt := middleWare.JWT()

	sysNoCheckRoleRouter(g)
	//sysCheckRoleRouter(g,jwt)
	return r
}
