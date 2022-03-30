package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HeaderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cip := c.ClientIP()
		log.Printf("HeaderHandler: Ip: %s, Path: %s", cip, c.FullPath())
	}
}
