package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

// SuccessJson响应成功的API请求，返回一个包含状态码、消息和数据的JSON响应。
// ctx: 代表HTTP请求上下文的Gin上下文。
// code: HTTP状态码，指示请求的处理状态。
// message: 与请求处理相关的信息消息。
// data: 要返回给客户端的数据。
func (c *BaseController) SuccessJson(ctx *gin.Context, code int, message string, data interface{}) {
	// 使用Gin的JSON方法将指定的数据作为JSON响应返回。
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code, // 使用固定值200表示操作成功。
		"message": message,
		"data":    data,
	})
}

// ErrorJson 是一个用于返回错误信息的JSON格式的方法。
// ctx：代表Gin框架的上下文，用于处理HTTP请求和响应。
// code：HTTP状态码，表明请求的处理状态。
// message：错误信息，对发生的问题进行描述。
// data：返回给客户端的额外数据。
func (c *BaseController) ErrorJson(ctx *gin.Context, code int, message string, data interface{}) {
	// 使用Gin框架的上下文，返回一个JSON格式的响应，包含错误代码、消息和数据。
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code, // 这里固定为500，可能需要根据实际需求进行调整
		"message": message,
		"data":    data,
	})
}
