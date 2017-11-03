package validate

// IAPConfig In-app-purchase所需配置
// apple:
// google:
type IAPConfig struct {
	ApplePassword          string `form:"apple_password" json:"apple_password,omitempty"`
	GoogleClientID         string `form:"google_client_id" json:"google_client_id,omitempty"`
	GoogleClientSecret     string `form:"google_client_secret" json:"google_client_secret,omitempty"`
	GooglePublicKeyStrLive string `form:"google_public_key_str_live" json:"google_public_key_str_live,omitempty"`
	GoogleRefToken         string `form:"google_refresh_token" json:"google_refresh_token,omitempty"`
}

// 用于单独将market解析出来，进入不同的逻辑
type reqMarket struct {
	Market string `form:"market" binding:"required"`
}

type reqParams struct {
	Receipt     interface{} `form:"receipt" bingding:"required" json:"receipt"`
	Market      string      `form:"market" binding:"required" json:"market"`
	IAPConfig   IAPConfig   `form:"iap_config" json:"iap_config"`
	Product     string      `form:"product" json:"product"`           //日志用
	Platform    string      `form:"platform" json:"platform"`         // 日志用
	Version     string      `form:"version" json:"version"`           // 日志用
	SandboxMode bool        `form:"sandbox_mode" json:"sandbox_mode"` //apple测试需要指定为true
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
	LatestReceiptInfo []InAppProduct   `json:"latest_receipt_info"`
	LatestReceipt     string           `json:"latest_receipt"` //auto-renewal订单有该数据
	IsSubscription    bool             `json:"is_subscription"`
	Receipt           string
}

type appleReceiptData struct {
	InApps []InAppProduct `json:"in_app"`
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
	InApps            []InAppProduct // 手工注入
}

type googleReceiptData struct {
	InApps []InAppProduct `json:"in_app"`
}

// ResponseData validate返回数据类型
type ResponseData struct {
	Status            int            `json:"status"`
	InApps            []InAppProduct `json:"in_app,omitempty"`
	LatestReceiptInfo []InAppProduct `json:"latest_receipt_info,omitempty"`
	LatestReceipt     string         `json:"latest_receipt,omitempty"`
	IsSubscription    bool           `json:"is_subscription,omitempty"`
	Receipt           interface{}    `json:"receipt,omitempty"`
}

// InAppProduct validate返回中的in_app字段，表示本次receipt的购买内容
type InAppProduct struct {
	Quantity               int    `json:"quantity,omitempty"`
	ProductID              string `json:"product_id,omitempty"`
	TransactionID          string `json:"transaction_id,omitempty"`
	OriginalTransactionID  string `json:"original_transaction_id,omitempty"`
	PurchaseDateMs         string `json:"purchase_date_ms,omitempty"`
	OriginalPurchaseDateMs string `json:"original_purchase_date_ms,omitempty"`
	ExpireDateMs           string `json:"expires_date_ms,omitempty"`
	//	TrialPeriod            bool   `json:"is_trial_period"`
}
