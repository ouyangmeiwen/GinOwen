package rabbitmqextend

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

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

// Consume 开始消费消息
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
		handler(msg) // 调用用户定义的消息处理函数
	}
	return nil
}

// Close 关闭消费者的连接
func (c *RabbitMQConsumer) Close() {
	c.Channel.Close()
	c.Conn.Close()
}
func TestConsumer() {
	url := "amqp://guest:guest@localhost:5672/"
	exchangeName := "topic_exchange"
	exchangeType := "topic"
	queueName := "user_queue"
	routingKey := "user.create"

	// 初始化消费者
	consumer, err := NewRabbitMQConsumer(url, exchangeName, exchangeType, queueName, routingKey)
	if err != nil {
		log.Fatalf("Failed to initialize consumer: %v", err)
	}
	defer consumer.Close()

	// 开始消费消息
	err = consumer.Consume(func(msg amqp.Delivery) {
		log.Printf("Received message: %s", string(msg.Body))
	})
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}
}
