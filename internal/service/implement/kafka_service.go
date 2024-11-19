package serviceimplement

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
)

type KafkaService struct {
	client sarama.SyncProducer
}

// NewProducer creates a new producer instance
func NewKafkaService(brokers []string) service.KafkaService {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return &KafkaService{
            client: nil,  // Return nil client if creation fails
        }
	}

	return &KafkaService{
		client: producer,
	}
}

// SendMessage sends a message to the Kafka topic with the given key and value.
func (p *KafkaService) SendMessage(topic, key, value string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}

	_, _, err := p.client.SendMessage(msg)
	if err != nil {
		fmt.Println("Failed to send message: ", err)
		return err
	}
	fmt.Println("Message sent successfully")
	return nil
}

// Close closes the Kafka producer client
func (p *KafkaService) Close() error {
	return p.client.Close()
}
