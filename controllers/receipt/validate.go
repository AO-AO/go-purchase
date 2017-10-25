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

	if strings.ToLower(requestParams.Market) == "android" {
		googleRequestParams := googleReqParams{}
		err := context.Bind(&googleRequestParams)
		if err != nil {
			logger.Error(err.Error(), *context)
			err = errors.New("PARAMS_ERROR")
			return nil, err
		}
		if jsonRequestParams, err := json.Marshal(googleRequestParams); err == nil {
			logger.Info(string(jsonRequestParams), *context)
		} else {
			logger.Warn(err.Error(), *context)
		}
		requestParams = reqParams{
			Receipt:       googleRequestParams.Receipt.Data,
			Market:        googleRequestParams.Market,
			IAPConfig:     googleRequestParams.IAPConfig,
			UserID:        googleRequestParams.UserID,
			Product:       googleRequestParams.Product,
			Platform:      googleRequestParams.Platform,
			Version:       googleRequestParams.Version,
			TransactionID: googleRequestParams.TransactionID,
			TestMode:      googleRequestParams.TestMode,
		}
	}
	if strings.ToLower(requestParams.Market) == "ios" {
		appleRequestParams := appleReqParams{}
		errBind := context.Bind(&appleRequestParams)
		if errBind != nil {
			logger.Error(errBind.Error(), *context)
			err = errors.New("PARAMS_ERROR")
			return nil, errBind
		}
		if jsonRequestParams, err := json.Marshal(appleRequestParams); err == nil {
			logger.Info(string(jsonRequestParams), *context)
		} else {
			logger.Warn(err.Error(), *context)
		}
		requestParams = reqParams{
			Receipt:       appleRequestParams.Receipt,
			Market:        appleRequestParams.Market,
			IAPConfig:     appleRequestParams.IAPConfig,
			UserID:        appleRequestParams.UserID,
			Product:       appleRequestParams.Product,
			Platform:      appleRequestParams.Platform,
			Version:       appleRequestParams.Version,
			TransactionID: appleRequestParams.TransactionID,
			TestMode:      appleRequestParams.TestMode,
		}
	}
	return requestParams, err
}

// DataManipulate 为ValidateController实现的数据操作模块
func (controller *ValidateController) DataManipulate(context *gin.Context, request interface{}) (interface{}, error) {
	var validateResult responseData
	var err error
	requestParams := request.(reqParams)
	if strings.ToLower(requestParams.Market) == "ios" {
		requestParams := request.(appleReqParams)
		var appleValidateResult appleValidateRes
		appleValidateResult, err = validateApple(requestParams.Receipt, requestParams.TestMode, requestParams.IAPConfig)
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
			Receipt:           appleValidateResult.Receipt,
		}
	}

	if strings.ToLower(requestParams.Market) == "android" {
		requestParams := request.(googleReqParams)
		var googleValidateResult googleValidateRes
		googleValidateResult, err = validateGoogle(requestParams.Receipt.Data, requestParams.IAPConfig)
		if err != nil {
			logger.Error("Google validate error: "+err.Error(), *context)
			err = errors.New("RECEIPT_VALIDATE_ERROR")
			return nil, err
		}
		//归一
		validateResult = responseData{
			Status:            googleValidateResult.Status,
			InApps:            googleValidateResult.ReceiptInfo.InApps,
			LatestReceiptInfo: googleValidateResult.LatestReceiptInfo,
			LatestReceipt:     googleValidateResult.LatestReceipt,
			Receipt:           googleValidateResult.Receipt,
		}
	}

	message, marshalErr := json.Marshal(validateResult)
	if marshalErr == nil {
		logger.Info("ValidateResponse: "+string(message), *context)
	}
	validateResult.Receipt = requestParams.Receipt
	return validateResult, err
}

// FormatResponse 为ValidateController实现的response组织模块
func (controller *ValidateController) FormatResponse(context *gin.Context, rawData interface{}) (interface{}, error) {
	resData := rawData.(appleValidateRes)
	response := responseData{
		Status:            resData.Status,
		InApps:            resData.ReceiptInfo.InApps,
		LatestReceiptInfo: resData.LatestReceiptInfo,
		LatestReceipt:     resData.LatestReceipt,
		Receipt:           resData.Receipt,
	}
	if resData.LatestReceipt != "" {
		response.IsSubscription = true
	}
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
