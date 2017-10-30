package filter

import "pincloud.purchase/purchaseApp/controllers/receipt/validate"

type reqParams struct {
	ValidateRes validate.ResponseData `form:"validate_result" binding:"required"`
}
