package user

import "github.com/gin-gonic/gin"

func RouterGroup(router *gin.Engine) {
	userGroup := router.Group("/v1/user")
	// my demo
	userGroup.GET("/say", UserSay)
}
