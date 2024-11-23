package rabbitmq

import "github.com/streadway/amqp"

// 定义 DataType 枚举常量
const (
	TextMessageType  = iota + 1 // 1: 文本消息
	ImageMessageType            // 2: 图像消息
)

// TextMessage 结构体
type TextMessage struct {
	Content string `json:"content"` // 文本消息内容
}

// ImageMessage 结构体
type ImageMessage struct {
	ImageURL string `json:"image_url"` // 图像 URL
	AltText  string `json:"alt_text"`  // 图像描述
}

// Data 结构体，包含 DataType 和 Body
type Data struct {
	DataType int    `json:"datatype"` // 消息类型
	Body     string `json:"body"`     // 消息体内容（JSON 字符串）
}

// RabbitMQ 结构体，用于封装 RabbitMQ 的连接和操作
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}
