package kafkaextend

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// KafkaConsumer 用于封装 Kafka 消费者
type KafkaConsumer struct {
	reader *kafka.Reader
}

// NewKafkaConsumer 创建并返回一个 KafkaConsumer 实例
func NewKafkaConsumer(broker string, topic string, groupID string) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{broker},
			Topic:   topic,
			GroupID: groupID,
		}),
	}
}

// ConsumeMessages 持续接收消息
func (kc *KafkaConsumer) ConsumeMessages() {
	for {
		// 从Kafka读取消息
		message, err := kc.reader.ReadMessage(nil)
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		fmt.Printf("Received message: Key=%s, Value=%s\n", string(message.Key), string(message.Value))
	}
}

// Close 关闭消费者连接
func (kc *KafkaConsumer) Close() {
	kc.reader.Close()
}
func TestConsumer() {
	// 配置 Kafka 信息
	broker := "localhost:9092" // Kafka地址
	topic := "example-topic"   // Kafka 主题
	groupID := "example-group" // 消费者组ID

	// 创建消费者实例
	consumer := NewKafkaConsumer(broker, topic, groupID)
	defer consumer.Close()

	// 开始消费消息
	consumer.ConsumeMessages()
}
