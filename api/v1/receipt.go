package v1

import (
	"github.com/gin-gonic/gin"
	"pincloud.purchase/controllers"
	"pincloud.purchase/controllers/receipt"
)

// MountReceiptRouter 组织/receipt/valid路由
func MountReceiptRouter(r *gin.RouterGroup) {
	r.POST("/receipt/validate", func(context *gin.Context) {
		ctrl := receipt.ValidateController{}
		executer := controllers.NewExecuter()
		executer(context, &ctrl)
	})
}
