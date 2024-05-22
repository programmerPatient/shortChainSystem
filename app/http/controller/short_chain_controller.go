package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laravelGo/app/helper"
	"github.com/laravelGo/app/services"
)

type ShortChainController struct{}

var ShortChainC = ShortChainController{}

func (c *ShortChainController) Create(ctx *gin.Context) {
}

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
