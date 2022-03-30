package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-sample/kafka/quote"
	"go-sample/web-gin/user"
	"go-sample/websocket/mock"
	"go-sample/websocket/ws"
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

// 使用gin -- 把gin-web升级为gin-websocket
func main() {
	go quote.TradeConsumer()
	go quote.SnapshotConsumer()

	r := gin.Default()
	//监听 get请求  /test路径
	r.GET("/mock", mock.MockSendHandler) // mock数据能正常推送到订阅方

	// ws demo
	ws.RouterGroup(r)
	// user demo
	user.RouterGroup(r)
	r.Run(":11888")

}
