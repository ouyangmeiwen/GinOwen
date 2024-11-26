package websocketextend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 定义 WebSocket 升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有跨域连接，生产环境请限制来源
		return true
	},
}

var Instance *ClientManager

// 定义客户端结构体
type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
}

// 定义客户端管理器
type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	mu         sync.Mutex
}
type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"` // targetID:message
}

// 初始化客户端管理器
func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// 启动客户端管理器
func (manager *ClientManager) Run() {
	for {
		select {
		case client := <-manager.Register:
			manager.mu.Lock()
			manager.Clients[client.ID] = client
			manager.mu.Unlock()
			fmt.Printf("Client connected: %s\n", client.ID)

		case client := <-manager.Unregister:
			manager.mu.Lock()
			if _, ok := manager.Clients[client.ID]; ok {
				close(client.Send)
				delete(manager.Clients, client.ID)
				fmt.Printf("Client disconnected: %s\n", client.ID)
			}
			manager.mu.Unlock()

		case message := <-manager.Broadcast:
			manager.mu.Lock()
			for _, client := range manager.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(manager.Clients, client.ID)
				}
			}
			manager.mu.Unlock()
		}
	}
}

func (client *Client) Read(manager *ClientManager) {
	defer func() {
		manager.Unregister <- client
		client.Conn.Close()
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			break
		}

		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			fmt.Printf("Invalid message format: %v\n", err)
			continue
		}

		switch msg.Type {
		case "broadcast":
			manager.Broadcast <- []byte(msg.Content)
		case "sendtoprivate":
			// 假设内容格式是 "targetID:message"
			parts := strings.SplitN(msg.Content, ":", 2) //限制分割为2个字符串
			if len(parts) == 2 {
				targetID := parts[0]
				message := parts[1]
				manager.SendToClient(targetID, []byte(message))
			}
		case "self":
			fmt.Printf("self: %s\n", msg.Content)
		default:
			fmt.Printf("Unknown message type: %s\n", msg.Type)
		}
	}
}
func (manager *ClientManager) SendToClient(clientID string, message []byte) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	client, exists := manager.Clients[clientID]
	if exists {
		client.Send <- message
	} else {
		fmt.Printf("Client %s not found\n", clientID)
	}
}

// 向客户端写入消息
func (client *Client) Write() {
	defer client.Conn.Close()
	for message := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Printf("Error writing message: %v\n", err)
			break
		}
	}
}

// WebSocket 路由处理
func serveWS(manager *ClientManager, c *gin.Context, clientID string) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Error upgrading connection: %v\n", err)
		return
	}

	client := &Client{
		ID:   clientID, // 使用从 URL 中获取的 clientID
		Conn: conn,
		Send: make(chan []byte),
	}

	manager.Register <- client

	go client.Read(manager)
	go client.Write()
}

func RegisterWebsocket(router *gin.Engine) {
	Instance = NewClientManager()
	go Instance.Run()

	// WebSocket 路由
	router.GET("/websocket/:clientID", func(c *gin.Context) {
		clientID := c.Param("clientID") // 获取路径参数 clientID
		serveWS(Instance, c, clientID)
	})
	// 测试主动发送消息接口
	router.POST("/websocket_send", func(c *gin.Context) {
		clientID := c.PostForm("clientID")
		message := c.PostForm("message")

		if clientID == "" {
			// 广播消息
			Instance.Broadcast <- []byte(message)
		} else {
			// 发送给特定客户端
			Instance.SendToClient(clientID, []byte(message))
		}
		c.JSON(http.StatusOK, gin.H{"status": "message sent"})
	})
}
