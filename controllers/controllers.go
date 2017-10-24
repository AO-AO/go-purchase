package controllers

import (
	"github.com/gin-gonic/gin"
	"pincloud.purchase/controllers/lib"
)

// Controller 用于定义标准controller接口的interface
type Controller interface {
	PickIncomingParams(*gin.Context) (interface{}, lib.ERROR)
	DataManipulate(*gin.Context, interface{}) (interface{}, lib.ERROR)
	FormatResponse(*gin.Context, interface{}) (interface{}, lib.ERROR)
	SendResponse(*gin.Context, interface{}) lib.ERROR
	ErrorHandle(*gin.Context, lib.ERROR)
}

// NewExecuter 返回controller的执行器
func NewExecuter() func(context *gin.Context, controller Controller) {
	return func(context *gin.Context, controller Controller) {
		if params, err := controller.PickIncomingParams(context); err.Type == "" {
			if rawRata, err := controller.DataManipulate(context, params); err.Type == "" {
				if response, err := controller.FormatResponse(context, rawRata); err.Type == "" {
					if err := controller.SendResponse(context, response); err.Type != "" {
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
