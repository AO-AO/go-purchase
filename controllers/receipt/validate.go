package receipt

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"pincloud.purchase/logger"
)

// 沙盒模式校验地址
var appleSandBoxHost = "sandbox.itunes.apple.com"
var prodHost = "buy.itunes.apple.com"
var applePath = "/verifyReceipt"

// ValidateController /receipt/validate接口返回数据
type ValidateController struct {
	ValidateRes       string
	TransactionIDList []string
}

// IAPConfig In-app-purchase所需配置
// apple:
// google:
type IAPConfig struct {
	test                   bool
	applePassword          string
	googleClientID         string
	googleClientSecret     string
	googlePublicKeyStrLive string
	googleRefToken         string
}

// RequestParams 请求参数
type RequestParams struct {
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

// ValidateResult 向服务器校验的返回结果
type ValidateResult struct {
	TransactionID string `json:"transaction_id"`
}

// PickIncomingParams 为ValidateController实现的数据解析模块
func (controller *ValidateController) PickIncomingParams(context *gin.Context) (RequestParams, error) {
	var requestParams RequestParams
	err := context.Bind(&requestParams)
	if err != nil {
		logger.Error(err.Error(), *context)
	} else {
		if jsonRequestParams, err := json.Marshal(requestParams); err == nil {
			logger.Info(string(jsonRequestParams), *context)
		} else {
			logger.Warn(err.Error(), *context)
		}
	}
	return requestParams, err
}

// DataManipulate 为ValidateController实现的数据操作模块
func (controller *ValidateController) DataManipulate(requestParams RequestParams) (ValidateResult, error) {
	var validateResult ValidateResult
	var err error
	if strings.ToLower(requestParams.Market) == "ios" {
		validateResult, err = validateApple(requestParams.Receipt, requestParams.TestMode, requestParams.IAPConfig)
	}
	return validateResult, err
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
func (controller *ValidateController) ErrorHandle(context *gin.Context, err error) {
	logger.Error(err.Error(), *context)
	panic(err)
}

func validateApple(receipt string, testMode bool, iapConfig IAPConfig) (ValidateResult, error) {
	var checkURL string
	if testMode {
		checkURL = appleSandBoxHost + applePath
	} else {
		checkURL = prodHost + applePath
	}
	applePassword := iapConfig.applePassword
	config := appleRequst{
		ReceiptData: receipt,
		Password:    applePassword,
	}
	result := ValidateResult{}

	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return result, err
	}

	ioBody := bytes.NewBufferString(string(jsonConfig))
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
