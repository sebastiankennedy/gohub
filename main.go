package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	r := gin.Default()

	// 注册一个路由
	r.GET("/", func(ctx *gin.Context) {

		// 以 JSON 格式响应
		ctx.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	// 运行服务
	r.Run()
}