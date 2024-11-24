package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/streadway/amqp"
)

// 消费消息并解析不同类型的消息
// ConsumeData 用于从 RabbitMQ 队列中消费消息。
// 它会返回一个消息通道（msgs），消费者可以从该通道接收消息。
// 如果有错误发生，方法会返回 nil 和错误信息。
func (r *RabbitMQ) ConsumeData() (<-chan amqp.Delivery, error) {
	// 使用 Channel.Consume 方法从指定的队列中消费消息
	msgs, err := r.Channel.Consume(
		r.Queue.Name, // 队列名称，用于指定要消费的队列
		"",           // 消费者的标识（consumerTag），如果为空字符串，RabbitMQ 会为消费者自动生成一个唯一的标识
		true,         // autoAck：是否自动确认消息。如果为 true，RabbitMQ 在消息传递给消费者后立即认为消息已成功处理并确认。// 如果为 false，消费者需要手动确认消息。
		false,        // exclusive：如果为 true，队列会变为独占，意味着只能由该消费者访问此队列。// 如果为 false，其他消费者也可以消费该队列。
		false,        // noLocal：如果为 true，消费者将不会接收自己发布到队列的消息。通常为 false。
		false,        // noWait：如果为 true，RabbitMQ 不会等待队列声明的结果（即不会返回确认信息）。
		nil,          // arguments：可选的附加参数，用于传递队列或消费者的额外配置，例如最大消息大小等。
	)

	// 如果发生错误，返回 nil 和错误信息
	if err != nil {
		return nil, fmt.Errorf("failed to register a consumer: %w", err)
	}

	// 如果成功，返回消息通道（msgs）和 nil 错误
	return msgs, nil
}

// 处理接收到的消息并解析消息体
func (r *RabbitMQ) ListenForData() {
	msgs, err := r.ConsumeData()
	if err != nil {
		log.Fatalf("Error consuming messages: %v", err)
	}

	for msg := range msgs {
		var data Data
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			log.Printf("Error unmarshalling data: %v", err)
			continue
		}

		// 根据 DataType 解析不同类型的消息
		switch data.DataType {
		case reflect.TypeOf(TextMessage{}).Name():
			var textMsg TextMessage
			if err := json.Unmarshal(data.Body, &textMsg); err != nil {
				log.Fatalf("Error unmarshalling TextMessage: %v", err)
			}
			r.handleTextMessage(textMsg)
		case reflect.TypeOf(ImageMessage{}).Name():
			var imageMsg ImageMessage
			if err := json.Unmarshal(data.Body, &imageMsg); err != nil {
				log.Fatalf("Error unmarshalling ImageMessage: %v", err)
			}
			r.handleImageMessage(imageMsg)
		default:
			r.handleDefaultMessage(string(data.Body))
		}
	}
}

// 处理文本消息
func (r *RabbitMQ) handleTextMessage(textMsg TextMessage) {
	log.Printf("Received text message: Content=%s", textMsg.Content)
}

// 处理图像消息
func (r *RabbitMQ) handleImageMessage(imageMsg ImageMessage) {
	log.Printf("Received image message: URL=%s, AltText=%s", imageMsg.ImageURL, imageMsg.AltText)
}

// 处理默认消息
func (r *RabbitMQ) handleDefaultMessage(body string) {
	log.Printf("Received image message: body=%s", body)
}
