package ws

import "github.com/gin-gonic/gin"

func quoteHandler(c *gin.Context) {
	Manager.RegisterClient(c)
}
