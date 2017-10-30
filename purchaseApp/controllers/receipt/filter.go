package receipt

import "github.com/gin-gonic/gin"

// FilterController /receipt/validate接口返回数据
type FilterController struct {
	validateResult    string
	transactionIDList []string
}

type filterParams struct {
	receipt       string `form:"receipt" bingding:"required"`
	market        string `form:"market" binding:"required"`
	userID        string `form:"userId"`
	product       string `form:"product"`
	platform      string `form:"platform"`
	version       string `form:"version"`
	transactionID string `form:"transactionID"`
}

// PickIncomingParams 为FilterController实现的数据解析模块
func (controller *FilterController) PickIncomingParams(context *gin.Context) (interface{}, error) {
	var filterParams filterParams
	err := context.Bind(&filterParams)
	return filterParams, err
}

// DataManipulate 为FilterController实现的数据操作模块
func (controller *FilterController) DataManipulate(interface{}) (interface{}, error) {
	return nil, nil
}

// FormatResponse 为FilterController实现的response组织模块
func (controller *FilterController) FormatResponse(interface{}) (interface{}, error) {
	return nil, nil
}

// SendResponse 为FilterController实现的response发送模块
func (controller *FilterController) SendResponse(*gin.Context, interface{}) error {
	return nil
}

// ErrorHandle 为FilterController实现的错误处理模块
func (controller *FilterController) ErrorHandle(*gin.Context, error) {

}
