package rabbitmqextend

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var InstancePublisher *RabbitMQPublisher

// RabbitMQPublisher 封装发布者的功能
type RabbitMQPublisher struct {
	Conn     *amqp.Connection
	Channel  *amqp.Channel
	Exchange string
}

// NewRabbitMQPublisher 初始化 RabbitMQ 发布者
func NewRabbitMQPublisher(url, exchangeName, exchangeType string) (*RabbitMQPublisher, error) {
	// 建立连接
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	// 声明交换机
	err = ch.ExchangeDeclare(
		exchangeName, // 交换机名称
		exchangeType, // 交换机类型 (direct, topic, fanout)
		true,         // durable: 是否持久化
		false,        // autoDelete: 自动删除
		false,        // internal: 是否为内部交换机
		false,        // noWait: 不等待服务器确认
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	return &RabbitMQPublisher{
		Conn:     conn,
		Channel:  ch,
		Exchange: exchangeName,
	}, nil
}

// PublishMessage 发送消息到交换机
func (p *RabbitMQPublisher) PublishData(routingKey string, data Data) error {
	message, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return p.PublishMessage(routingKey, message)
}

// PublishMessage 发送消息到交换机
func (p *RabbitMQPublisher) PublishMessage(routingKey string, message []byte) error {
	err := p.Channel.Publish(
		p.Exchange, // 交换机名称
		routingKey, // 路由键
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	log.Printf("Message published: %s", message)
	return nil
}

// Close 关闭发布者的连接
func (p *RabbitMQPublisher) Close() {
	p.Channel.Close()
	p.Conn.Close()
}

func RegisterMQPublisher(url string, exchangeName string, exchangeType string) {
	// 初始化发布者
	var err error
	InstancePublisher, err = NewRabbitMQPublisher(url, exchangeName, exchangeType)
	if err != nil {
		log.Fatalf("Failed to initialize publisher: %v", err)
	}
}
