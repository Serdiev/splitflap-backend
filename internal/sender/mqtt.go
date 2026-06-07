package sender

import (
	"os"
	config "splitflap-backend/configs"
	"splitflap-backend/internal/logger"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var cfg = config.New()

type MQTTSender struct {
	client    MQTT.Client
	topic     string
	initiated bool
}

func (m *MQTTSender) SendMessage(text string) error {
	if !m.initiated {
		return nil
	}

	token := m.client.Publish(m.topic, 0, false, text)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}

	return nil
}

// NewMQTTSender creates a new instance of MQTTSender.
func NewMQTTSender(topic string) *MQTTSender {
	if !cfg.MQTT.Enabled {
		logger.Info().Msg("MQTT is disabled, skipping sender creation")
		return &MQTTSender{
			initiated: false,
		}
	}

	opts := MQTT.NewClientOptions()
	opts.AddBroker(cfg.MQTT.BrokerUrl)
	opts.SetClientID("lcd_sender")
	opts.SetUsername(cfg.MQTT.Username)
	opts.SetPassword(cfg.MQTT.Password)

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Error().Err(token.Error()).Msg("Failed to connect to MQTT broker")
		return &MQTTSender{
			initiated: false,
		}
	}

	return &MQTTSender{
		client:    client,
		topic:     topic,
		initiated: true,
	}
}

func setupReadHandler(client MQTT.Client, topic string) {
	messageHandler := func(client MQTT.Client, msg MQTT.Message) {
		logger.Info().Msgf("Received message: %s\n", msg.Payload())
	}

	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		logger.Info().Msgf("Failed to subscribe to MQTT topic: %s", token.Error())
		os.Exit(1)
	}
}
