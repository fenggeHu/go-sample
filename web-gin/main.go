package main

import (
	"github.com/gin-gonic/gin"
	"go-sample/web-gin/user"
)

func main() {
	router := gin.Default()
	// gin官方demo
	defaultRouter(router)
	// my demo
	user.RouterGroup(router)

	router.Run()
}
