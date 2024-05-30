package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laravelGo/app/helper"
	requeststruct "github.com/laravelGo/app/http/request_struct"
	"github.com/laravelGo/app/services"
	"github.com/laravelGo/core/config"
)

type ShortChainController struct {
	BaseController
}

var ShortChainC = ShortChainController{}

func (c *ShortChainController) Access(ctx *gin.Context) {
	shortCode := ctx.Param("shortCode")
	original_url := services.GetShortChainService().GetOriginalUrlByShortUrl(shortCode)
	if helper.Empty(original_url) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": fmt.Sprintf("短链接 %s 不存在", shortCode),
		})
		return
	}
	//301是永久重定向，302是临时重定向。短地址一经生成就不会变化，所以用301是符合http语义的。
	//但是如果用了301， Google，百度等搜索引擎，搜索的时候会直接展示真实地址，那我们就无法统计到短地址被点击的次数了，
	//也无法收集用户的Cookie, User Agent 等信息，这些信息可以用来做很多有意思的大数据分析，也是短网址服务商的主要盈利来源
	ctx.Redirect(http.StatusFound, original_url)
}

// Create处理创建短链的请求。
// ctx: 代表HTTP请求的上下文。
func (c *ShortChainController) Create(ctx *gin.Context) {
	// 从请求中获取原始URL参数
	param := requeststruct.ShortChainCreate{}
	ctx.ShouldBindJSON(&param)
	// 检查原始URL是否为空
	if helper.Empty(param.OriginalUrl) {
		c.ErrorJson(ctx, 400, "原始URL不能为空", nil)
		return
	}
	// 调用服务创建短链
	link, err := services.GetShortChainService().CreateShortUrl(param.OriginalUrl)
	// 检查短链创建过程中是否出错
	if err != nil {
		// 如果出错，返回500错误
		c.ErrorJson(ctx, 500, "短链接生成失败", err.Error())
		return
	}
	// 构造完整短链地址
	url := fmt.Sprintf("%s/%s", config.GetString("app.foreign_host"), link.ShortUrl)
	// 返回200成功响应，包含短链地址
	c.SuccessJson(ctx, 200, "短链接生成成功", url)

}
