package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/sebastiankennedy/gohub/app/http/controllers/api/v1"
	"github.com/sebastiankennedy/gohub/app/models/user"
	requests "github.com/sebastiankennedy/gohub/app/requests"
)

// 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// 检查手机号码是是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// 检查邮箱是否已被注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}

	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		// 打印错误信息
		fmt.Println(err.Error())

		return
	}

	// 表单验证
	errs := requests.ValidateSignupEmailExist(&request, c)

	// errs 返回长度等于零即通过，大于 0 即有错误发生
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
