package middlewares

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// SetRequestID 为请求生成一个唯一requestID
func SetRequestID(context *gin.Context) {
	requestID := uuid.NewV4()
	context.Set("RequestID", requestID)
}
