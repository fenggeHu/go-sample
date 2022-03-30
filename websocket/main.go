package main

import (
	"github.com/gin-gonic/gin"
	"go-sample/api-gin/user"
	"go-sample/kafka/quote"
	"go-sample/websocket/mock"
	"go-sample/websocket/ws"
)

// 使用gin -- 把gin-web升级为gin-websocket
func main() {
	go quote.TradeConsumer()
	go quote.SnapshotConsumer()

	r := gin.Default()
	//监听 get请求  /test路径
	r.GET("/mock", mock.QuoteSendHandler) // mock数据能正常推送到订阅方

	// ws demo
	ws.RouterGroup(r)
	// user demo
	user.RouterGroup(r)
	r.Run(":11888")

}
