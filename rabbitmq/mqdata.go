package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// TextMessage 结构体
type TextMessage struct {
	Content string `json:"Content"` // 文本消息内容
}

// ImageMessage 结构体
type ImageMessage struct {
	ImageURL string `json:"ImageURL"` // 图像 URL
	AltText  string `json:"AltText"`  // 图像描述
}

// Data 结构体，包含 DataType 和 Body
type Data struct {
	DataType string          // 消息类型
	Body     json.RawMessage // 消息体内容（json对象）
}

// RabbitMQ 结构体，用于封装 RabbitMQ 的连接和操作
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}
