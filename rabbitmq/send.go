package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// 发送不同类型的消息
func (r *RabbitMQ) SendData(data Data) error {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	// 发布消息到队列
	err = r.Channel.Publish(
		"",           // 默认交换机
		r.Queue.Name, // 队列名称
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        dataJSON,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Sent message: %s", dataJSON)
	return nil
}

// 发送不同类型的消息
func (r *RabbitMQ) PublishData(dataType int, body interface{}) error {
	// 将消息体结构体序列化为 JSON 字符串
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal body: %w", err)
	}

	// 创建 Data 结构体并序列化为 JSON
	data := Data{
		DataType: dataType,
		Body:     string(bodyJSON),
	}
	return r.SendData(data)
}
