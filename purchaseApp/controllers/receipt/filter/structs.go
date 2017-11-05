package filter

import "pincloud.purchase/purchaseApp/controllers/receipt/validate"

type reqParams struct {
	ValidateRes   validate.ResponseData `form:"validate_result" json:"validate_result" binding:"required"`
	Products      []dbProduct           `form:"product_list" json:"product_list" binding:"required"`
	TransactionID string                `form:"transaction_id" json:"transaction_id"`
	OfferID       string                `form:"offer_id" json:"offer_id"`
}

type dbProduct struct {
	BestDeal     bool              `form:"best_deal" json:"best_deal"`
	Discount     int               `form:"discount" json:"discount"`
	Effect       int               `form:"effect" json:"effect"`
	Iap          map[string]string `form:"iap" json:"iap"`
	IsHot        bool              `form:"is_hot" json:"is_hot"`
	Kind         string            `form:"kind" json:"kind"`
	OfferID      string            `form:"offer_id" json:"offer_id"`
	Way          string            `form:"way" json:"way,omitempty"`
	Subscription map[string]int    `form:"subscription" json:"subscription"`
}

// ResponseData filter的返回结果,是validate的结果，增加一个字段用来存放filter的数据库结果
type ResponseData struct {
	ValidateRes    validate.ResponseData   `json:"validate_result"`
	TransactionID  string                  `json:"transaction_id,omitempty"`
	OfferID        string                  `json:"offer_id,omitempty"`
	ValideIAPs     []validate.InAppProduct `json:"valide_iaps"`
	ValideProducts []dbProduct             `json:"valide_products"`
}
