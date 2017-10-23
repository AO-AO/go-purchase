package controllers

import (
	"github.com/gin-gonic/gin"
)

// Controller 用于定义标准controller接口的interface
type Controller interface {
	PickIncomingParams(*gin.Context) (interface{}, error)
	DataManipulate(interface{}) (interface{}, error)
	FormatResponse(interface{}) (interface{}, error)
	SendResponse(*gin.Context, interface{}) error
	ErrorHandle(*gin.Context, error)
}

// NewExecuter 返回controller的执行器
func NewExecuter() func(context *gin.Context, controller Controller) {
	return func(context *gin.Context, controller Controller) {
		if params, err := controller.PickIncomingParams(context); err == nil {
			if rawRata, err := controller.DataManipulate(params); err == nil {
				if response, err := controller.FormatResponse(rawRata); err == nil {
					if err := controller.SendResponse(context, response); err != nil {
						controller.ErrorHandle(context, err)
					}
				} else {
					controller.ErrorHandle(context, err)
				}
			} else {
				controller.ErrorHandle(context, err)
			}
		} else {
			controller.ErrorHandle(context, err)
		}
	}
}
