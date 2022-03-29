package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// TODO 开启了kafka消费后，导致ws断开
// 使用gin -- 把gin-web升级为gin-websocket
func main() {
	go consumer_start()

	r := gin.Default()
	//监听 get请求  /test路径
	r.GET("/quote/list", quoteHandler)
	r.GET("/mock", mockSendHandler) // mock数据能正常推送到订阅方
	r.Run(":11888")

}
