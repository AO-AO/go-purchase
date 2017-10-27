package receipt

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

// 用于单独将market解析出来，进入不同的逻辑
type reqMarket struct {
	Market string `form:"market" binding:"required"`
}

// appleReqParams 请求参数
type reqParams struct {
	Receipt       interface{} `form:"receipt" bingding:"required"`
	Market        string      `form:"market" binding:"required"`
	IAPConfig     IAPConfig   `form:"iapConfig" binding:"required"`
	UserID        string      `form:"userId"`
	Product       string      `form:"product"`
	Platform      string      `form:"platform"`
	Version       string      `form:"version"`
	TransactionID string      `form:"transactionId"`
	TestMode      bool        `form:"test"`
	OrderID       string      `form:"orderId"`
}

// googleeqParams 请求参数
type googleReqParams struct {
	Receipt       googleReceipt `form:"receipt" bingding:"required"`
	Market        string        `form:"market" binding:"required"`
	IAPConfig     IAPConfig     `form:"iapConfig" binding:"required"`
	UserID        string        `form:"userId"`
	Product       string        `form:"product"`
	Platform      string        `form:"platform"`
	Version       string        `form:"version"`
	TransactionID string        `form:"transactionID"`
	TestMode      bool          `form:"test"`
}

type googleReceipt struct {
	Data         string      `form:"data" binding:"required"`
	PurchaseData interface{} //`form:"purchaseData"`
	Signature    string      `form:"signature"`
}

// appleConfig ios校验参数
type appleConfig struct {
	ReceiptData string `json:"receipt-data"`
	Password    string `json:"password"`
}

// appleValidateRes 向服务器校验的返回结果
type appleValidateRes struct {
	Status            int              `json:"status"`
	ReceiptInfo       appleReceiptData `json:"receipt"`
	LatestReceiptInfo []inAppProduct   `json:"latest_receipt_info"`
	LatestReceipt     string           `json:"latest_receipt"` //auto-renewal订单有该数据
	IsSubscription    bool             `json:"is_subscription"`
	Receipt           string
}

type appleReceiptData struct {
	InApps []inAppProduct `json:"in_app"`
}

type googleTokenRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type googleReceiptReq struct {
	ProductID     string `json:"productId"`
	PackageName   string `json:"packageName"`
	AutoRenewing  bool   `json:"autoRenewing"`
	PurchaseToken string `json:"purchaseToken"`
}

// type googleReceiptDataReq struct {
// 	ProductID     string `json:"productId"`
// 	PackageName   string `json:"packageName"`
// 	AutoRenewing  bool   `json:"autoRenewing"`
// 	PurchaseToken string `json:"purchaseToken"`
// }

// googleValidateRes 向服务器校验的返回结果
type googleValidateRes struct {
	Status            int            `json:"purchaseState"`
	PurchaseDateMs    string         `json:"purchaseTimeMillis"`
	StartDateMs       string         `json:"startTimeMillis"`
	ExpireDateMs      string         `json:"expiryTimeMillis"`
	ConsumptionState  int            `json:"consumptionState"`
	Kind              string         `json:"kind"`
	DeveloperPayload  string         `json:"developerPayload"`
	OrderID           string         `json:"orderId"`
	PriceCurrencyCode string         `json:"priceCurrencyCode"`
	PriceAmountMicros string         `json:"priceAmountMicros"`
	CountryCode       string         `json:"countryCode"`
	CancelReason      int            `json:"cancelReason"`
	InApps            []inAppProduct // 手工注入
}

type googleReceiptData struct {
	InApps []inAppProduct `json:"in_app"`
}

type responseData struct {
	Status            int            `json:"status"`
	InApps            []inAppProduct `json:"in_app"`
	LatestReceiptInfo []inAppProduct `json:"latest_receipt_info,omitempty"`
	LatestReceipt     string         `json:"latest_receipt,omitempty"` //auto-renewal订单有该数据
	IsSubscription    bool           `json:"is_subscription"`
	Receipt           interface{}    `json:"receipt"`
}

type inAppProduct struct {
	Quantity               int    `json:"quantity,omitempty"`
	ProductID              string `json:"product_id,omitempty"`
	TransactionID          string `json:"transaction_id,omitempty"`
	OriginalTransactionID  string `json:"original_transaction_id,omitempty"`
	PurchaseDateMs         string `json:"purchase_date_ms,omitempty"`
	OriginalPurchaseDateMs string `json:"original_purchase_date_ms,omitempty"`
	ExpireDateMs           string `json:"expires_date_ms,omitempty"`
	//	TrialPeriod            bool   `json:"is_trial_period"`
	OrderID string `json:"order_id,omitempty"`
}
