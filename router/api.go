/*
 * @Description:
 * @Author: mali
 * @Date: 2022-09-05 17:41:24
 * @LastEditTime: 2022-11-08 10:58:03
 * @LastEditors: VSCode
 * @Reference:
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/laravelGo/app/http/controller"
)

func AddApiRouter(e *gin.Engine) {
	e.GET("/:shortCode", controller.ShortChainC.Access)
	e.POST("/shortUrlCreate", controller.ShortChainC.Create)
}
