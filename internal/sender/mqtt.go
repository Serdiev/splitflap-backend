package sender

import (
	"fmt"
	"os"
	config "splitflap-backend/configs"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var cfg = config.New()

type MQTTSender struct {
	client    MQTT.Client
	topic     string
	initiated bool
}

// SendMessage sends the given text as a message using MQTT.
func (m *MQTTSender) SendMessage(text string) error {
	if !m.isInitiated() {
		return nil
	}

	token := m.client.Publish("topic", 0, false, text)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (m *MQTTSender) isInitiated() bool {
	return m.initiated
}

// NewMQTTSender creates a new instance of MQTTSender.
func NewMQTTSender(topic string) *MQTTSender {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(cfg.MQTT.BrokerUrl)
	opts.SetClientID("some_sender")

	// Create MQTT client
	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Failed to connect to MQTT broker:", token.Error())
		return &MQTTSender{
			initiated: false,
		}
	}
	defer client.Disconnect(250)

	return &MQTTSender{
		client:    client,
		topic:     topic,
		initiated: true,
	}
}

func setupReadHandler(client MQTT.Client, topic string) {
	messageHandler := func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Received message: %s\n", msg.Payload())
	}

	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		fmt.Println("Failed to subscribe to MQTT topic:", token.Error())
		os.Exit(1)
	}
}
