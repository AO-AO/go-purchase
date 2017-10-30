package v1

import (
	"github.com/gin-gonic/gin"
	"pincloud.purchase/purchaseApp/controllers"
	"pincloud.purchase/purchaseApp/controllers/receipt/filter"
	"pincloud.purchase/purchaseApp/controllers/receipt/validate"
)

// MountReceiptRouter 组织/receipt/valid路由
func MountReceiptRouter(r *gin.RouterGroup) {
	r.POST("/receipt/validate", func(context *gin.Context) {
		ctrl := validate.Controller{
			Context: *context,
		}
		executer := controllers.NewExecuter()
		executer(context, &ctrl)
	})
	r.POST("/receipt/filter", func(context *gin.Context) {
		ctrl := filter.Controller{
			Context: *context,
		}
		executer := controllers.NewExecuter()
		executer(context, &ctrl)
	})
}
