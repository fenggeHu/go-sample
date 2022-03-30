package main

import (
	"github.com/gin-gonic/gin"
	"go-sample/api-gin/user"
)

func main() {
	router := gin.Default()
	// gin官方demo
	defaultRouter(router)
	// my demo
	user.RouterGroup(router)

	router.Run()
}
