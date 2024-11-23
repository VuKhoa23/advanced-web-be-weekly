package serviceimplement

import (
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/VuKhoa23/advanced-web-be/internal/constants"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
)

type KafkaService struct {
	client sarama.SyncProducer
}

func NewKafkaService(brokers []string) service.KafkaService {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return &KafkaService{
            client: nil,
        }
	}

	return &KafkaService{
		client: producer,
	}
}

func (p *KafkaService) SendMessage(topic, key, value string) (string, error) {
	correlationID := constants.CORRELATION_ID
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
		Headers: []sarama.RecordHeader{
            {
                Key:   []byte("correlation_id"),
                Value: []byte(correlationID),
            },
        },
	}

	_, _, err := p.client.SendMessage(msg)
	if err != nil {
		fmt.Println("Failed to send message: ", err)
		return "", err
	}

	// Wait for the reply
    reply := waitForReply(correlationID) // Wait for the reply with the same correlation ID
    return reply, nil
}

func waitForReply(correlationID string) string {
	replyChan := make(chan string)

	go func() {
		consumer, err := sarama.NewConsumer([]string{constants.BROKER}, nil)
        if err != nil {
            fmt.Println("Failed to create Kafka consumer:", err)
            return
        }
        defer consumer.Close()

		// Consume the reply topic and filter by correlation_id
        partitionConsumer, err := consumer.ConsumePartition(constants.REPLY_TOPIC, 0, sarama.OffsetNewest)
        if err != nil {
            fmt.Println("Failed to start Kafka partition consumer:", err)
            return
        }
        defer partitionConsumer.Close()

		for message := range partitionConsumer.Messages(){
			if correlationID == extractCorrelationID(message.Headers) {
                replyChan <- string(message.Value)
                break
            }
		}
	} ()

	// Wait for the reply
    select {
    case reply := <-replyChan:
        return reply
    case <-time.After(time.Second * 10):
        return "Reply timeout"
    }
}

func extractCorrelationID(headers []*sarama.RecordHeader) string {
    for _, header := range headers {
        if string(header.Key) == "correlation_id" {
            return string(header.Value)
        }
    }
    return ""
}

func (p *KafkaService) Close() error {
	return p.client.Close()
}
