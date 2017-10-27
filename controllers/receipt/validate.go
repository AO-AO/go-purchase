package receipt

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"pincloud.purchase/controllers/lib"
	"pincloud.purchase/logger"
)

// ValidateController /receipt/validate接口
type ValidateController struct{}

// PickIncomingParams 为ValidateController实现的数据解析模块
func (controller *ValidateController) PickIncomingParams(context *gin.Context) (interface{}, error) {
	requestParams := reqParams{}
	var err error
	err = context.Bind(&requestParams)
	if err != nil {
		logger.Error(err.Error(), *context)
		err = errors.New("PARAMS_ERROR")
		return nil, err
	}

	if jsonRequestParams, err := json.Marshal(requestParams); err == nil {
		logger.Info(string(jsonRequestParams), *context)
	} else {
		logger.Warn(err.Error(), *context)
		err = nil
	}
	return requestParams, err
}

// DataManipulate 为ValidateController实现的数据操作模块
func (controller *ValidateController) DataManipulate(context *gin.Context, request interface{}) (interface{}, error) {
	var validateResult responseData
	var err error
	requestParams := request.(reqParams)
	if strings.ToLower(requestParams.Market) == "ios" {
		receipt := requestParams.Receipt.(string)
		var appleValidateResult appleValidateRes
		appleValidateResult, err = validateApple(receipt, requestParams.TestMode, requestParams.IAPConfig)
		if err != nil {
			logger.Error("Apple validate error: "+err.Error(), *context)
			err = errors.New("RECEIPT_VALIDATE_ERROR")
			return nil, err
		}
		//归一
		validateResult = responseData{
			Status:            appleValidateResult.Status,
			InApps:            appleValidateResult.ReceiptInfo.InApps,
			LatestReceiptInfo: appleValidateResult.LatestReceiptInfo,
			LatestReceipt:     appleValidateResult.LatestReceipt,
			Receipt:           requestParams.Receipt,
		}

		if validateResult.LatestReceipt != "" {
			validateResult.IsSubscription = true
		}

		// 为与本次请求transactionID匹配的inApp装入orderID
		if requestParams.TransactionID != "" && requestParams.OrderID != "" {
			for _, inApp := range validateResult.InApps {
				if inApp.TransactionID == requestParams.TransactionID {
					inApp.OrderID = requestParams.OrderID
				}
			}
		}
	}

	if strings.ToLower(requestParams.Market) == "android" {
		receiptData := requestParams.Receipt.(map[string]interface{})
		receiptStr := receiptData["data"].(string)
		var googleValidateResult googleValidateRes
		googleValidateResult, err = validateGoogle(receiptStr, requestParams.IAPConfig)
		if err != nil {
			logger.Error("Google validate error: "+err.Error(), *context)
			err = errors.New("RECEIPT_VALIDATE_ERROR")
			return nil, err
		}

		// 归一
		validateResult = responseData{
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
		logger.Info("ValidateResponse: "+string(message), *context)
	}
	return validateResult, err
}

// FormatResponse 为ValidateController实现的response组织模块
func (controller *ValidateController) FormatResponse(context *gin.Context, rawData interface{}) (interface{}, error) {
	response := rawData.(responseData)

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
func (controller *ValidateController) SendResponse(context *gin.Context, rawData interface{}) error {
	resData := rawData.(lib.StandardResponse)
	context.JSON(http.StatusOK, resData)
	return nil
}

// ErrorHandle 为ValidateController实现的错误处理模块
func (controller *ValidateController) ErrorHandle(context *gin.Context, err error) {
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
}
