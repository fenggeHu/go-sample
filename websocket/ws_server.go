package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

// WsManager WebSocket 管理器
var WsManager = clientManager{
	client:     make(map[string]*wsClient),
	register:   make(chan *wsClient),
	unRegister: make(chan *wsClient),
}

// ClientManager websocket client Manager struct
type clientManager struct {
	//clientGroup map[string]map[string]*wsClient
	client     map[string]*wsClient
	register   chan *wsClient
	unRegister chan *wsClient
}

type ResponseMessage struct {
	Time   int32     `json:"time"`
	Symbol string    `json:"symbol"`
	Data   []float64 `json:"data"`
}

// wsClient Websocket 客户端
type wsClient struct {
	ID     string
	Group  string
	Socket *websocket.Conn
	Send   chan *ResponseMessage
	Tags   string
}

func (c *wsClient) Read() {
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			log.Fatalln(err)
			break
		}
		c.Tags = string(message)
	}
}

func (c *wsClient) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			msg, _ := json.Marshal(message)
			c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

// Start 启动 websocket 管理器
func (manager *clientManager) Start() {
	log.Printf("Websocket manage start")
	//for {
	//	select {
	//	case client := <-manager.register:
	//		log.Printf("Websocket client %s connect", client.ID)
	//		if manager.client[client.ID] == nil {
	//			manager.client[client.ID] = client
	//		}
	//
	//	case client := <-manager.unRegister:
	//		log.Printf("Unregister websocket client %s", client.ID)
	//
	//		//case data := <-manager.broadcast:
	//		//	if groupMap, ok := manager.clientGroup[data.GroupID]; ok {
	//		//		for _, conn := range groupMap {
	//		//			conn.Send <- data.Data
	//		//		}
	//		//	}
	//	}
	//}

	consumer_start()
}

// RegisterClient 向 manage 中注册 client
func (manager *clientManager) RegisterClient(c *gin.Context) {
	upgrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("websocket client connect %v error", c.Param("channel"))
		return
	}

	client := &wsClient{
		ID:     uuid.NewString(),
		Group:  c.Param("channel"),
		Socket: conn,
		Send:   make(chan *ResponseMessage),
	}
	manager.client[client.ID] = client
	//manager.register <- client
	go client.Read()
	go client.Write()
}

func quoteHandler(c *gin.Context) {
	WsManager.RegisterClient(c)
}

// 推送消息
func (manager *clientManager) send(message *ResponseMessage) {
	for _, client := range manager.client {
		if strings.Index(client.Tags, message.Symbol) != -1 {
			client.Send <- message
		}
	}
}
