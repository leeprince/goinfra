package services

import (
	"encoding/json"
	"gowebsocket/logger"
	"sync"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[string]*Client // // prince@TODO: 防止并发写，建议使用sync.Map 2024/4/15 18:04
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	Mux        sync.Mutex
}

// Message is return msg
type Message struct {
	WsToken string `json:"wstoken,omitempty"`
	Content string `json:"content,omitempty"`
}

// Manager define a ws server manager
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[string]*Client),
}

func InitWebSocket() {
	// 启动 websocket
	go Manager.Start()
}

// Start is  项目运行前, 协程开启start -> go Manager.Start()
func (manager *ClientManager) Start() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("ManagerStart异常退出: ", err)
		}
	}()
	for {
		select {
		case conn := <-Manager.Register:
			logger.Info("新用户加入:", conn.ID)
			Manager.Clients[conn.ID] = conn
			/*jsonMessage, _ := json.Marshal(&Message{Content: "Successful connection to socket service"})
			conn.Send <- jsonMessage*/
		case conn := <-Manager.Unregister:
			logger.Info("用户离开:", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				/*jsonMessage, _ := json.Marshal(&Message{Content: "A socket has disconnected"})
				conn.Send <- jsonMessage*/
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		case message := <-Manager.Broadcast:
			MessageStruct := Message{}
			logger.Info("广播message: ", string(message))
			json.Unmarshal(message, &MessageStruct)
			for id, conn := range Manager.Clients {
				if id != manager.CreatId(MessageStruct.WsToken) {
					continue
				}
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
		}
	}
}

func (manager *ClientManager) CreatId(token string) string {
	return token
}
