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
}

type KafkaMessageHandler func(ctx context.Context, message []byte) error

func NewKafkaService(brokers []string, handler KafkaMessageHandler) service.KafkaService {
	groupID := constants.KEY

	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		fmt.Println("err ", err)
		return nil
	}

	return &KafkaService{consumerGroup: consumerGroup, handler: handler}
}

func (kc *KafkaService) Start(ctx context.Context, topics []string) {
	go func() {
		for {
			err := kc.consumerGroup.Consume(ctx, topics, &consumerHandler{handler: kc.handler})
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
}

func (h *consumerHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *consumerHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (h *consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		err := h.handler(context.Background(), message.Value)
		if err != nil {
			log.Printf("Failed to process Kafka message: %v", err)
		}
		session.MarkMessage(message, "")
	}
	return nil
}

