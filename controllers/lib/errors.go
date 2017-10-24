package lib

import "errors"

// ERRORS 定义错误输出
var ERRORS = map[string]ERROR{
	"RESPONSE_FORMAT_ERROR": ERROR{
		Code:  400,
		Error: errors.New("Error response format"),
		Type:  "RESPONSE_FORMAT_ERROR",
	},
	"RECEIPT_VALIDATE_ERROR": ERROR{
		Code:  400,
		Error: errors.New("Error validate the receipt"),
		Type:  "RECEIPT_VALIDATE_ERROR",
	},
	"PARAMS_ERROR": ERROR{
		Code:  400,
		Error: errors.New("Params error"),
		Type:  "PARAMS_ERROR",
	},
}

// ERROR Error详情
type ERROR struct {
	Code  int
	Error error
	Type  string
}
