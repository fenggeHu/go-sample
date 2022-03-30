package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func PermissionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("PermissionHandler: %s", c.Request.Header)
	}
}
