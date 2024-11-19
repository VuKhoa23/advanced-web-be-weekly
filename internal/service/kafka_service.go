package service

import (
	"context"
)

type KafkaService interface {
    // Start begins consuming messages for the specified topics.
    Start(ctx context.Context, topics []string)
    // Close shuts down the Kafka consumer.
    Close() error
}
