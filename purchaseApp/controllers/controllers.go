package controllers

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"pincloud.purchase/purchaseApp/controllers/lib"
)

// Controller 用于定义标准controller接口的interface
type Controller interface {
	PickIncomingParams(*gin.Context) (interface{}, error)
	DataManipulate(interface{}) (interface{}, error)
	FormatResponse(interface{}) (interface{}, error)
	SendResponse(*gin.Context, interface{}) error
}

// NewExecuter 返回controller的执行器
func NewExecuter() func(context *gin.Context, controller Controller) {
	return func(context *gin.Context, controller Controller) {
		if params, err := controller.PickIncomingParams(context); err == nil {
			if rawRata, err := controller.DataManipulate(params); err == nil {
				if response, err := controller.FormatResponse(rawRata); err == nil {
					if err := controller.SendResponse(context, response); err != nil {
						ErrorHandle(context, err)
					}
				} else {
					ErrorHandle(context, err)
				}
			} else {
				ErrorHandle(context, err)
			}
		} else {
			ErrorHandle(context, err)
		}
	}
}

// ErrorHandle 通用错误处理模块
func ErrorHandle(context *gin.Context, err error) {
	errDetail := lib.ERRORS[err.Error()]
	errResponseJSON := lib.StandardResponse{
		Meta: lib.StandardResponseMeta{
			ErrorMessage: errDetail.Message,
			Code:         errDetail.Code,
			ErrorType:    errDetail.Type,
		},
		Salt: time.Now().UnixNano() / 1000000,
	}
	context.AbortWithStatusJSON(http.StatusBadRequest, errResponseJSON)
	debug.PrintStack()
}
