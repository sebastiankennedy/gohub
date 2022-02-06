package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// new 一个 Gin Engine 实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(ctx *gin.Context) {
		// 以 JSON 格式响应
		ctx.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	// 运行服务，指定端口为 8000
	r.Run(":8000")
}
