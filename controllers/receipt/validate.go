package receipt

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"pincloud.purchase/controllers/lib"
	"pincloud.purchase/logger"
)

// 沙盒模式校验地址
var appleSandBoxHost = "sandbox.itunes.apple.com"
var prodHost = "buy.itunes.apple.com"
var applePath = "/verifyReceipt"

// ValidateController /receipt/validate接口
type ValidateController struct{}

// IAPConfig In-app-purchase所需配置
// apple:
// google:
type IAPConfig struct {
	ApplePassword          string `form:"applePassword"`
	GoogleClientID         string `form:"googleClientID"`
	GoogleClientSecret     string `form:"googleClientSecret"`
	GooglePublicKeyStrLive string `form:"googlePublicKeyStrLive"`
	GoogleRefToken         string `form:"googleRefToken"`
}

// reqParams 请求参数
type reqParams struct {
	Receipt       string    `form:"receipt" bingding:"required"`
	Market        string    `form:"market" binding:"required"`
	IAPConfig     IAPConfig `form:"iapConfig" binding:"required"`
	UserID        string    `form:"userId"`
	Product       string    `form:"product"`
	Platform      string    `form:"platform"`
	Version       string    `form:"version"`
	TransactionID string    `form:"transactionID"`
	TestMode      bool      `form:"test"`
}

type appleRequst struct {
	ReceiptData string `json:"receipt-data"`
	Password    string `json:"password"`
}

type googleRequest struct {
}

// validateRes 向服务器校验的返回结果
type validateRes struct {
	Status            int            `json:"status"`
	InApps            []inAppProduct `json:"in_app"`
	LatestReceiptInfo []inAppProduct `json:"latest_receipt_info"`
	LatestReceipt     string         `json:"latest_receipt"` //auto-renewal订单有该数据
	IsSubscription    bool           `json:"is_subscription"`
}

type inAppProduct struct {
	Quantity               int    `json:"quantity"`
	ProductID              string `json:"product_id"`
	TransactionID          string `json:"transaction_id"`
	OriginalTransactionID  string `json:"original_transaction_id"`
	PurchaseDateMs         string `json:"purchase_date_ms"`
	OriginalPurchaseDateMs string `json:"original_purchase_date_ms"`
	ExpireDateMs           string `json:"expires_date_ms"`
	TrialPeriod            bool   `json:"is_trial_period"`
}

// PickIncomingParams 为ValidateController实现的数据解析模块
func (controller *ValidateController) PickIncomingParams(context *gin.Context) (interface{}, error) {
	requestParams := reqParams{}
	err := context.Bind(&requestParams)
	if err != nil {
		logger.Error(err.Error(), *context)
		err = errors.New("PARAMS_ERROR")
		return nil, err
	}
	if jsonRequestParams, err := json.Marshal(requestParams); err == nil {
		logger.Info(string(jsonRequestParams), *context)
	} else {
		logger.Warn(err.Error(), *context)
	}
	return requestParams, err
}

// DataManipulate 为ValidateController实现的数据操作模块
func (controller *ValidateController) DataManipulate(context *gin.Context, request interface{}) (interface{}, error) {
	requestParams := request.(reqParams)
	var validateResult validateRes
	var err error
	if strings.ToLower(requestParams.Market) == "ios" {
		validateResult, err = validateApple(requestParams.Receipt, requestParams.TestMode, requestParams.IAPConfig)
		if err != nil {
			logger.Error("Validate error: "+err.Error(), *context)
			err = errors.New("RECEIPT_VALIDATE_ERROR")
			return nil, err
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
	resData := rawData.(validateRes)
	if resData.LatestReceipt != "" {
		resData.IsSubscription = true
	}
	responseJSON := lib.StandardResponse{
		Meta: lib.StandardResponseMeta{
			ErrorMessage: "success",
			Code:         200,
			ErrorType:    "",
		},
		Data: resData,
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

func validateApple(receipt string, testMode bool, iapConfig IAPConfig) (validateRes, error) {
	var checkURL string
	if testMode {
		checkURL = appleSandBoxHost + applePath
	} else {
		checkURL = prodHost + applePath
	}
	checkURL = "https://" + checkURL
	applePassword := iapConfig.ApplePassword
	config := appleRequst{
		ReceiptData: receipt,
		Password:    applePassword,
	}
	result := validateRes{}

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return result, err
	}

	ioBody := bytes.NewReader([]byte(jsonConfig))
	response, err := http.Post(checkURL, "application/json", ioBody)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	resultBuf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	json.Unmarshal(resultBuf, &result)
	return result, err
}
