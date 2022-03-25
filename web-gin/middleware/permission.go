package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func PermissionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.GetHeader("Agent"))
	}
}
