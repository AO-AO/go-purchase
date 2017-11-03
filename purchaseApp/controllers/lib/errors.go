package lib

// ERRORS 错误输列表
var ERRORS = map[string]ERROR{
	"RESPONSE_FORMAT_ERROR": ERROR{
		Code:    400,
		Message: "Error response format",
		Type:    "RESPONSE_FORMAT_ERROR",
	},
	"RECEIPT_VALIDATE_ERROR": ERROR{
		Code:    400,
		Message: "Error validate the receipt",
		Type:    "RECEIPT_VALIDATE_ERROR",
	},
	"PARAMS_ERROR": ERROR{
		Code:    400,
		Message: "Params error",
		Type:    "PARAMS_ERROR",
	},
	"STATUS_ERROR": ERROR{
		Code:    400,
		Message: "The receipt status is invalid",
		Type:    "STATUS_ERROR",
	},
}

// ERROR Error详情
type ERROR struct {
	Code    int
	Message string
	Type    string
}
