package ws

import (
	"github.com/gin-gonic/gin"
	"go-sample/api-gin/middleware"
)

func RouterGroup(router *gin.Engine) {
	// Use Handler作用于Group
	wsGroup := router.Group("/ws").Use(middleware.HeaderHandler()).Use(middleware.PermissionHandler())
	wsGroup.GET("/quote/list", quoteHandler)
}
