package user

import (
	"github.com/gin-gonic/gin"
	"go-sample/api-gin/middleware"
)

func RouterGroup(router *gin.Engine) {
	userGroup := router.Group("/v1/user").Use(middleware.HeaderHandler()).Use(middleware.PermissionHandler())
	// my demo
	userGroup.GET("/say", Say)
}
