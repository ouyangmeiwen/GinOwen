package kafkaextend

import (
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// KafkaProducer 用于封装 Kafka 生产者
type KafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer 创建并返回一个 KafkaProducer 实例
func NewKafkaProducer(broker string, topic string) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

// SendMessage 发送消息到 Kafka
func (kp *KafkaProducer) SendMessage(key string, value string) error {
	message := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	}

	err := kp.writer.WriteMessages(nil, message)
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	fmt.Println("Message sent successfully:", value)
	return nil
}

// Close 关闭生产者连接
func (kp *KafkaProducer) Close() {
	kp.writer.Close()
}
func TestPublisher() {
	// 配置 Kafka 信息
	broker := "localhost:9092" // Kafka地址
	topic := "example-topic"   // Kafka 主题

	// 创建生产者实例
	producer := NewKafkaProducer(broker, topic)
	defer producer.Close()

	// 发送5条消息
	for i := 0; i < 5; i++ {
		message := fmt.Sprintf("Message %d", i+1)
		err := producer.SendMessage(fmt.Sprintf("Key-%d", i+1), message)
		if err != nil {
			log.Fatal("Failed to send message:", err)
		}
		time.Sleep(1 * time.Second) // 等待1秒再发送下一个消息
	}
}
