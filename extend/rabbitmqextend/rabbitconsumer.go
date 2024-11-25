package rabbitmqextend

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var InstanceConsumer *RabbitMQConsumer

// RabbitMQConsumer 封装消费者的功能
type RabbitMQConsumer struct {
	Conn     *amqp.Connection
	Channel  *amqp.Channel
	Queue    string
	Exchange string
}

// NewRabbitMQConsumer 初始化 RabbitMQ 消费者
func NewRabbitMQConsumer(url, exchangeName, exchangeType, queueName, routingKey string) (*RabbitMQConsumer, error) {
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
		exchangeType, // 交换机类型
		true,         // durable
		false,        // autoDelete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	// 声明队列
	queue, err := ch.QueueDeclare(
		queueName, // 队列名称
		true,      // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	// 绑定队列到交换机
	err = ch.QueueBind(
		queue.Name,   // 队列名称
		routingKey,   // 路由键
		exchangeName, // 交换机名称
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &RabbitMQConsumer{
		Conn:     conn,
		Channel:  ch,
		Queue:    queue.Name,
		Exchange: exchangeName,
	}, nil
}

// Close 关闭消费者的连接
func (c *RabbitMQConsumer) Close() {
	c.Channel.Close()
	c.Conn.Close()
}
func RegisterMQConsumer(url, exchangeName, exchangeType, queueName, routingKey string) {
	var err error
	// 初始化消费者
	InstanceConsumer, err = NewRabbitMQConsumer(url, exchangeName, exchangeType, queueName, routingKey)
	if err != nil {
		log.Fatalf("Failed to initialize consumer: %v", err)
	}
	// 开始消费消息
	// 启动一个 goroutine 来异步消费消息
	go func() {
		err := InstanceConsumer.Consume(messageHandler)
		if err != nil {
			log.Fatalf("消费消息失败: %v", err)
		}
	}()
}

// 消费消息并处理不同类型的消息
func (c *RabbitMQConsumer) Consume(handler func(msg amqp.Delivery)) error {
	// 注册消费者
	messages, err := c.Channel.Consume(
		c.Queue, // 队列名称
		"",      // 消费者标签
		true,    // autoAck
		false,   // exclusive
		false,   // noLocal
		false,   // noWait
		nil,     // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	// 消费消息
	log.Printf("Waiting for messages on queue [%s]...", c.Queue)
	for msg := range messages {
		// 调用用户定义的消息处理函数
		handler(msg)
	}
	return nil
}

// 消息处理函数
func messageHandler(msg amqp.Delivery) {
	var data Data
	err := json.Unmarshal(msg.Body, &data)
	if err != nil {
		handleDefaultMessage(string(msg.Body))
		return
	}

	// 根据 DataType 解析不同类型的消息
	switch data.DataType {
	case "TextMessage":
		var textMsg TextMessage
		if err := json.Unmarshal(data.Body, &textMsg); err != nil {
			log.Printf("Error unmarshalling TextMessage: %v", err)
			return
		}
		handleTextMessage(textMsg)
	case "ImageMessage":
		var imageMsg ImageMessage
		if err := json.Unmarshal(data.Body, &imageMsg); err != nil {
			log.Printf("Error unmarshalling ImageMessage: %v", err)
			return
		}
		handleImageMessage(imageMsg)
	default:
		handleDefaultMessage(string(msg.Body))
	}
}

// 处理文本消息
func handleTextMessage(msg TextMessage) {
	log.Printf("Received TextMessage: %s", msg.Content)
}

// 处理图像消息
func handleImageMessage(msg ImageMessage) {
	log.Printf("Received ImageMessage: URL=%s, AltText=%s", msg.ImageURL, msg.AltText)
}
func handleDefaultMessage(msg string) {
	log.Printf("Received MSG:%s", msg)
}
