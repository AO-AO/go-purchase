package lib

// StandardResponse 用来组织通用接口返回的数据
type StandardResponse struct {
	Meta StandardResponseMeta `json:"meta"`
	Data interface{}          `json:"data,omitempty"`
	Salt int64                `json:"salt"`
}

// StandardResponseMeta 用来组织api返回json中的meta字段
type StandardResponseMeta struct {
	ErrorType    string `json:"error_type"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
}

// StandardErrorResponse 用来组织api返回的错误数据
type StandardErrorResponse struct {
	Meta StandardResponseMeta `json:"meta"`
	Salt int64                `json:"salt"`
}
