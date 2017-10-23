package receipt

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "pincloud.purchase/controllers"
	. "pincloud.purchase/"
)

// ValidateController /receipt/validate接口返回数据
type ValidateController struct {
	ValidateRes       string
	TransactionIDList []string
}

type requestParams struct {
	receipt       string `form:"receipt" bingding:"required"`
	market        string `form:"market" binding:"required"`
	userID        string `form:"userId"`
	product       string `form:"product"`
	platform      string `form:"platform"`
	version       string `form:"version"`
	transactionID string `form:"transactionID"`
}

// PickIncomingParams 为ValidateController实现的数据解析模块
func (controller *ValidateController) PickIncomingParams(context *gin.Context) (interface{}, error) {
	var requestParams requestParams
	err := context.Bind(&requestParams)
	return requestParams, err
}

// DataManipulate 为ValidateController实现的数据操作模块
func (controller *ValidateController) DataManipulate(requestParams interface{}) (interface{}, error) {
	var resultData = requestParams

	if !ok {
      errors.New('') 
	}
	return resData, err
}

// FormatResponse 为ValidateController实现的response组织模块
func (controller *ValidateController) FormatResponse(rawData interface{}) (interface{}, error) {
	return nil, nil
}

// SendResponse 为ValidateController实现的response发送模块
func (controller *ValidateController) SendResponse(context *gin.Context, rawData interface{}) error {
	return nil
}

// ErrorHandle 为ValidateController实现的错误处理模块
func (controller *ValidateController) ErrorHandle(*gin.Context, error) {

}
