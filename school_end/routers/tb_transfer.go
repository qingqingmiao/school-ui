package routers

import (
	"chinasoccer/apis/transfer"
	"github.com/gin-gonic/gin"
)

func TransferRouter(r *gin.RouterGroup)  {
	r.GET("/transfer", transfer.GetTransfer)
}
