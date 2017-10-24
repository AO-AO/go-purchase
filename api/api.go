package api

import (
	"github.com/gin-gonic/gin"
	"pincloud.purchase/api/v1"
)

// MountRouters mounts api routers
func MountRouters(r *gin.Engine) {
	groupV1 := r.Group("/api/v1")
	{
		v1.MountReceiptRouter(groupV1)
	}
}
