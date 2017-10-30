package filter

import (
	"github.com/gin-gonic/gin"
)

// Controller /receipt/validate接口返回数据
type Controller struct {
	Context gin.Context
}

// PickIncomingParams 为Controller实现的数据解析模块
func (controller *Controller) PickIncomingParams(context *gin.Context) (interface{}, error) {
	filterParams := reqParams{}
	err := context.Bind(&filterParams)
	return filterParams, err
}

// DataManipulate 为Controller实现的数据操作模块
func (controller *Controller) DataManipulate(interface{}) (interface{}, error) {
	return nil, nil
}

// FormatResponse 为Controller实现的response组织模块
func (controller *Controller) FormatResponse(interface{}) (interface{}, error) {
	return nil, nil
}

// SendResponse 为Controller实现的response发送模块
func (controller *Controller) SendResponse(*gin.Context, interface{}) error {
	return nil
}

// ErrorHandle 为Controller实现的错误处理模块
func (controller *Controller) ErrorHandle(*gin.Context, error) {

}
