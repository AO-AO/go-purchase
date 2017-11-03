package validate

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"pincloud.purchase/purchaseApp/controllers/lib"
	"pincloud.purchase/purchaseApp/logger"
)

// Controller /receipt/validate接口
type Controller struct {
	Context gin.Context
}

// PickIncomingParams 为ValidateController实现的数据解析模块
func (controller *Controller) PickIncomingParams(context *gin.Context) (interface{}, error) {
	requestParams := reqParams{}
	var err error
	err = context.Bind(&requestParams)
	if err != nil {
		logger.Error(err.Error(), controller.Context)
		err = errors.New("PARAMS_ERROR")
		return nil, err
	}

	if jsonRequestParams, err := json.Marshal(requestParams); err == nil {
		logger.Info(string(jsonRequestParams), controller.Context)
	} else {
		logger.Warn(err.Error(), controller.Context)
		err = nil
	}
	return requestParams, err
}

// DataManipulate 为ValidateController实现的数据操作模块
func (controller *Controller) DataManipulate(request interface{}) (interface{}, error) {
	var validateResult ResponseData
	var err error
	requestParams := request.(reqParams)
	if strings.ToLower(requestParams.Market) == "ios" {
		receipt := requestParams.Receipt.(string)
		var appleValidateResult appleValidateRes
		appleValidateResult, err = validateApple(receipt, requestParams.SandboxMode, requestParams.IAPConfig)
		if err != nil {
			logger.Error("Apple validate error: "+err.Error(), controller.Context)
			err = errors.New("RECEIPT_VALIDATE_ERROR")
			return nil, err
		}
		//归一
		validateResult = ResponseData{
			Status:            appleValidateResult.Status,
			InApps:            appleValidateResult.ReceiptInfo.InApps,
			LatestReceiptInfo: appleValidateResult.LatestReceiptInfo,
			LatestReceipt:     appleValidateResult.LatestReceipt,
			Receipt:           requestParams.Receipt,
		}

		if validateResult.LatestReceipt != "" {
			validateResult.IsSubscription = true
		}

	}

	if strings.ToLower(requestParams.Market) == "android" {
		receiptData := requestParams.Receipt.(map[string]interface{})
		receiptStr := receiptData["data"].(string)
		var googleValidateResult googleValidateRes
		googleValidateResult, err = validateGoogle(receiptStr, requestParams.IAPConfig)
		if err != nil {
			logger.Error("Google validate error: "+err.Error(), controller.Context)
			err = errors.New("RECEIPT_VALIDATE_ERROR")
			return nil, err
		}

		// 归一
		validateResult = ResponseData{
			Status:  googleValidateResult.Status,
			InApps:  googleValidateResult.InApps,
			Receipt: requestParams.Receipt,
		}
		if validateResult.InApps[0].ExpireDateMs != "" {
			validateResult.IsSubscription = true
		}
	}

	message, marshalErr := json.Marshal(validateResult)
	if marshalErr == nil {
		logger.Info("ValidateResponse: "+string(message), controller.Context)
	}
	return validateResult, err
}

// FormatResponse 为ValidateController实现的response组织模块
func (controller *Controller) FormatResponse(rawData interface{}) (interface{}, error) {
	response := rawData.(ResponseData)

	responseJSON := lib.StandardResponse{
		Meta: lib.StandardResponseMeta{
			ErrorMessage: "success",
			Code:         200,
			ErrorType:    "",
		},
		Data: response,
		Salt: time.Now().UnixNano() / 1000000,
	}

	return responseJSON, nil
}

// SendResponse 为ValidateController实现的response发送模块
func (controller *Controller) SendResponse(context *gin.Context, rawData interface{}) error {
	resData := rawData.(lib.StandardResponse)
	context.JSON(http.StatusOK, resData)
	return nil
}
