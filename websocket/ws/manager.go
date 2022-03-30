package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)

// Manager WebSocket 管理器
var Manager = ClientManager{
	Clients:    make(map[string]*Client),
	register:   make(chan *Client),
	unRegister: make(chan *Client),
}

// ClientManager websocket Client Manager struct
type ClientManager struct {
	//clientGroup map[string]map[string]*Clients
	Clients    map[string]*Client
	register   chan *Client
	unRegister chan *Client
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// RegisterClient 向 manage 中注册 client
func (manager *ClientManager) RegisterClient(c *gin.Context) {
	handlers := c.HandlerNames()
	log.Println(handlers)

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

	client := &Client{
		ID:     uuid.NewString(),
		Group:  c.Param("channel"),
		Socket: conn,
		Send:   make(chan *ResponseMessage),
	}
	manager.Clients[client.ID] = client
	//manager.register <- client
	go client.Read()
	go client.Write()
}

// 推送消息
func (manager *ClientManager) Send(message *ResponseMessage) {
	for _, client := range manager.Clients {
		if strings.Index(client.Tags, message.Symbol) != -1 {
			client.Send <- message
		}
	}
}

func (c *Client) Read() {
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			delete(Manager.Clients, c.ID) // 正常或异常退出、断开连接等
			log.Println(err)              // log.Fatal导致应用退出
			break
		}
		c.Tags = string(message)
		log.Printf("%s set Tag: %s", c.Socket.RemoteAddr(), c.Tags)
	}
}

func (c *Client) Write() {
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
			c.Socket.WriteJSON(message)
			//msg, _ := json.Marshal(message)
			//c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

//// Start 启动 websocket 管理器
//func (manager *clientManager) Start() {
//	log.Printf("Websocket manage start")
//	//for {
//	//	select {
//	//	case client := <-manager.register:
//	//		log.Printf("Websocket client %s connect", client.ID)
//	//		if manager.client[client.ID] == nil {
//	//			manager.client[client.ID] = client
//	//		}
//	//
//	//	case client := <-manager.unRegister:
//	//		log.Printf("Unregister websocket client %s", client.ID)
//	//
//	//		//case data := <-manager.broadcast:
//	//		//	if groupMap, ok := manager.clientGroup[data.GroupID]; ok {
//	//		//		for _, conn := range groupMap {
//	//		//			conn.Send <- data.Data
//	//		//		}
//	//		//	}
//	//	}
//	//}
//
//	consumer_start()
//}
