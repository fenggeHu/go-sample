package ws

import "github.com/gorilla/websocket"

//  Websocket 客户端
type Client struct {
	ID     string
	Group  string
	Socket *websocket.Conn
	Send   chan *ResponseMessage
	Tags   string
}

type ResponseMessage struct {
	Time   int64     `json:"time"`
	Symbol string    `json:"symbol"`
	Data   []float64 `json:"data"`
}
