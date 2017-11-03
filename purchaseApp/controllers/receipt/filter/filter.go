package filter

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"pincloud.purchase/purchaseApp/controllers/lib"
	"pincloud.purchase/purchaseApp/logger"
)

// Controller /receipt/validate接口返回数据
type Controller struct {
	Context gin.Context
}

// PickIncomingParams 为Controller实现的数据解析模块
func (controller *Controller) PickIncomingParams(context *gin.Context) (interface{}, error) {
	filterParams := reqParams{}
	err := context.Bind(&filterParams)
	if err != nil {
		logger.Error(err.Error(), controller.Context)
		err = errors.New("PARAMS_ERROR")
		return nil, err
	}
	if jsonFilterParams, marshalErr := json.Marshal(filterParams); marshalErr == nil {
		logger.Info(string(jsonFilterParams), controller.Context)
	} else {
		logger.Warn(err.Error(), controller.Context)
	}
	return filterParams, nil
}

// DataManipulate 为Controller实现的数据操作模块
func (controller *Controller) DataManipulate(request interface{}) (interface{}, error) {
	requestParams := request.(reqParams)
	resData := ResponseData{}
	resData.ValidateRes = requestParams.ValidateRes
	// status = 0才是正常receipt
	validateStatus := requestParams.ValidateRes.Status
	if validateStatus != 0 {
		logger.Error("Receipt is in error status: "+strconv.Itoa(validateStatus), controller.Context)
		statErr := errors.New("INVALIDE_STATUS")
		return resData, statErr
	}

	// 去掉过期的iap信息的列表
	for _, iap := range resData.ValidateRes.InApps {
		iapExpireDateMs, convErr := strconv.Atoi(iap.ExpireDateMs)
		timeNowMs := time.Now().UnixNano() / 1000000
		if convErr != nil {
			logger.Warn("Convert expireDateMs to int error: "+convErr.Error(), controller.Context)
			resData.ValideIAPs = append(resData.ValideIAPs, iap)
		} else if iap.ExpireDateMs == "" || timeNowMs < int64(iapExpireDateMs) {
			resData.ValideIAPs = append(resData.ValideIAPs, iap)
		}
	}

	// 如果指定了transactionID，则要求transactionID一致
	if requestParams.TransactionID != "" {
		for index, iap := range resData.ValidateRes.InApps {
			if iap.TransactionID != requestParams.TransactionID {
				resData.ValidateRes.InApps = append(resData.ValidateRes.InApps[:index], resData.ValidateRes.InApps[index+1:]...)
			}
		}
	}

	// 通过过滤过的iap，组装products；至少有1个valideIAPs
	if len(resData.ValideIAPs) > 0 {
		// 超过1个，打印一个warn日志
		if len(resData.ValideIAPs) > 1 {
			logger.Warn("More than one validated IAPs, "+strconv.Itoa(len(resData.ValideIAPs))+" exist.", controller.Context)
		}

		// 筛选匹配的product
		for _, product := range requestParams.Products {
			// 指定了OfferID，则通过OfferID匹配product；没有指定OfferID，则通过productID匹配product
			if requestParams.OfferID != "" {
				if requestParams.OfferID == product.OfferID {
					resData.ValideProducts = append(resData.ValideProducts, product)
				}
			} else {
				for _, iap := range resData.ValideIAPs {
					if dbProductID, ok := product.Iap["product_id"]; ok {
						if iap.ProductID == dbProductID {
							resData.ValideProducts = append(resData.ValideProducts, product)
						}
					}
				}
			}
		}
	}

	resData.OfferID = requestParams.OfferID
	return resData, nil
}

// FormatResponse 为Controller实现的response组织模块
func (controller *Controller) FormatResponse(resData interface{}) (interface{}, error) {
	response := resData.(ResponseData)
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

// SendResponse 为Controller实现的response发送模块
func (controller *Controller) SendResponse(context *gin.Context, rawData interface{}) error {
	resData := rawData.(lib.StandardResponse)
	context.JSON(http.StatusOK, resData)
	return nil
}
