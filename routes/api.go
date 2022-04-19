package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebastiankennedy/gohub/app/http/controllers/api/v1/auth"
)

func RegisterApiRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(ctx *gin.Context) {
			// 以 JSON 格式响应
			ctx.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})

		authGroup := v1.Group(("/auth"))
		{
			suc := new(auth.SignupController)
			// 判断手机是否已被注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)

			// 判断邮箱是否已被注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
		}
	}
}
