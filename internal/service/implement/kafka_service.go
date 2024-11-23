package serviceimplement

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/VuKhoa23/advanced-web-be/internal/constants"
	"github.com/VuKhoa23/advanced-web-be/internal/service"
)

type KafkaService struct {
	consumerGroup   sarama.ConsumerGroup
	handler KafkaMessageHandler
	client sarama.SyncProducer	// must have client to send reply back
}

type KafkaMessageHandler func(ctx context.Context, message []byte) error

func NewKafkaService(brokers []string, handler KafkaMessageHandler) service.KafkaService {
	groupID := constants.KEY

	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		fmt.Println("err ", err)
		return nil
	}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return &KafkaService{
            client: nil, 
        }
	}

	return &KafkaService{consumerGroup: consumerGroup, handler: handler, client: producer}
}

// Start consuming messages from the given topics
func (kc *KafkaService) Start(ctx context.Context, topics []string) {
	go func() {
		for {
			err := kc.consumerGroup.Consume(ctx, topics, &consumerHandler{
				handler: kc.handler,
				client: kc.client,
			})
			if err != nil {
				log.Printf("Error consuming Kafka messages: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
}

func (k *KafkaService) Close() error {
	return k.consumerGroup.Close()
}

type consumerHandler struct {
	handler KafkaMessageHandler
	client  sarama.SyncProducer
}

func (h *consumerHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *consumerHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (h *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		correlationID := extractCorrelationID(message.Headers)
		// Validate the correlation ID
		if correlationID == "" || correlationID != constants.CORRELATION_ID {
			log.Printf("Invalid or missing correlation ID: %s. Skipping message.", correlationID)
			session.MarkMessage(message, "")
			continue // Skip processing this message
		}

		replyValue := fmt.Sprintf("Reply to request with correlation ID = %s", correlationID)
		err := h.handler(context.Background(), message.Value)
		if err != nil {
			log.Printf("Failed to process Kafka message: %v", err)
		}
		// Send the reply back to the producer with the same correlation_id
        replyMessage := &sarama.ProducerMessage{
            Topic: constants.REPLY_TOPIC,
            Key:   sarama.StringEncoder(constants.KEY),
            Value: sarama.StringEncoder(replyValue),
            Headers: []sarama.RecordHeader{
                {
                    Key:   []byte("correlation_id"),
                    Value: []byte(correlationID),
                },
            },
        }
		// Send the reply
        _, _, err = h.client.SendMessage(replyMessage)
        if err != nil {
            log.Printf("Failed to send reply: %v", err)
        }

        session.MarkMessage(message, "")
	}
	return nil
}

func extractCorrelationID(headers []*sarama.RecordHeader) string {
    for _, header := range headers {
        if string(header.Key) == "correlation_id" {
            return string(header.Value)
        }
    }
    return ""
}

