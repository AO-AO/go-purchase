package main

import (
	"github.com/gin-gonic/gin"
	"pincloud.purchase/purchaseApp/api"
	"pincloud.purchase/purchaseApp/middlewares"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Serve static files
	//	r.Static("/public", "./public")

	// TODO add recursive load template file support
	// Load template files
	//	r.LoadHTMLGlob("./views/*.html")
	// r.LoadHTMLGlob("./views/**/*.html")

	// Mount routers
	//	r.Use(middlewares.GetSession)
	// r.Use(middlewares.CheckUserLogin)

	//健康检查接口
	r.GET("/healthy", func(context *gin.Context) {
		context.String(200, "I'm healthy~")
	})

	r.Use(middlewares.SetRequestID)
	api.MountRouters(r)

	// 初始化各种cache
	//	initCaches()
	// Listen and Server in 0.0.0.0:9527
	r.Run(":9401")
}
