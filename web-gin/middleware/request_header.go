package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HeaderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cip := c.ClientIP()
		fmt.Printf("Ip: %s, Path: %s", cip, c.FullPath())
	}
}
