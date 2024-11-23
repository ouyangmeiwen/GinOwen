package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// 定义全局的 RabbitMQ 实例
var Instance *RabbitMQ

// 创建一个新的 RabbitMQ 实例
func NewRabbitMQ(url string, queueName string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	// QueueDeclare 用于声明一个新的队列，或确认一个已经存在的队列。
	// 参数如下：
	// queueName: 队列的名称。
	// durable: 队列是否持久化，如果 RabbitMQ 重启，队列和其中的消息会保留下来。
	//         如果为 true，队列会在 RabbitMQ 重启后依然存在。
	// autoDelete: 是否在队列没有消费者时自动删除队列。如果为 true，
	//             当队列没有消费者连接时，它会被自动删除。
	// exclusive: 如果为 true，队列会是独占的。也就是说，只能被创建它的连接使用。
	//         当连接关闭时，队列会自动删除。
	// noWait: 如果为 true，RabbitMQ 不会等待队列声明的结果。
	//         如果为 false，RabbitMQ 会在声明队列时等待确认。
	// arguments: 其他队列参数，用于传递队列的附加设置，如消息 TTL、最大队列长度等。

	queue, err := ch.QueueDeclare(
		queueName, // 队列的名称
		true,      // durable: 队列是否持久化。设置为 true，队列在 RabbitMQ 重启后依然存在。
		false,     // autoDelete: 是否在没有消费者时自动删除队列。设置为 false，队列不会在没有消费者时自动删除。
		false,     // exclusive: 队列是否独占，只能被创建它的连接使用。设置为 false，队列对所有连接可见。
		false,     // noWait: 如果设置为 true，RabbitMQ 不会等待确认队列声明是否成功。通常设置为 false，表示等待确认。
		nil,       // arguments: 队列的其他参数，通常为空。
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %w", err)
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
		Queue:   queue,
	}, nil
}

// InitRabbitMQ 初始化全局 RabbitMQ 实例
func InitRabbitMQ(url string, queueName string) error {
	var err error
	Instance, err = NewRabbitMQ(url, queueName)
	if err != nil {
		return fmt.Errorf("failed to initialize RabbitMQ: %w", err)
	}
	return nil
}

// Close 关闭 RabbitMQ 的连接和通道
func (r *RabbitMQ) Close() error {
	// 关闭 Channel 和 Connection
	if err := r.Channel.Close(); err != nil {
		return fmt.Errorf("failed to close channel: %w", err)
	}
	if err := r.Conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}
	return nil
}
