package mqttextend

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Instance *MQTTClient

// MQTTClient 是一个封装的结构体，包含客户端和连接选项
type MQTTClient struct {
	client  mqtt.Client
	options *mqtt.ClientOptions
}

// NewMQTTClient 创建一个新的 MQTTClient 实例
func NewMQTTClient(broker string, clientID string, username string, password string) *MQTTClient {
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID(clientID)
	options.SetUsername(username)
	options.SetPassword(password)
	options.SetDefaultPublishHandler(defaultMessageHandler)
	options.SetConnectRetry(true)
	options.SetKeepAlive(60 * time.Second)
	options.SetPingTimeout(1 * time.Second)

	client := mqtt.NewClient(options)

	return &MQTTClient{
		client:  client,
		options: options,
	}
}

// Connect 连接到 MQTT Broker
func (m *MQTTClient) Connect() error {
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	fmt.Println("MQTT connected successfully")
	return nil
}

// Publish 发送消息到指定的主题
func (m *MQTTClient) Publish(topic string, qos byte, retained bool, message string) error {
	token := m.client.Publish(topic, qos, retained, message)
	token.Wait()
	return token.Error()
}

// Subscribe 订阅指定主题并处理消息
func (m *MQTTClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) error {
	token := m.client.Subscribe(topic, qos, callback)
	token.Wait()
	return token.Error()
}

// Disconnect 断开连接
func (m *MQTTClient) Disconnect() {
	m.client.Disconnect(250)
	fmt.Println("MQTT disconnected")
}

// 默认消息处理函数
func defaultMessageHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
}

func RegisterMQTT() {
	// 创建 MQTT 客户端
	Instance = NewMQTTClient(
		"tcp://broker.emqx.io:1883", // 替换为你的 MQTT Broker 地址
		"testClientID",
		"",
		"",
	)

	// 连接到 MQTT Broker
	err := Instance.Connect()
	if err != nil {
		fmt.Printf("Failed to connect to MQTT broker: %v\n", err)
		return
	}
	defer Instance.Disconnect() // 放main 函数

	// 订阅主题
	err = Instance.Subscribe("test/topic", 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})
	if err != nil {
		fmt.Printf("Failed to subscribe to topic: %v\n", err)
		return
	}

	// 发布消息
	err = Instance.Publish("test/topic", 1, false, "Hello, MQTT!")
	if err != nil {
		fmt.Printf("Failed to publish message: %v\n", err)
		return
	}

	// 等待消息到来
	//select {} //不需要服务自动
}
