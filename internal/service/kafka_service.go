package service

import (
	"context"
)

type KafkaService interface {
    Start(ctx context.Context, topics []string)
    Close() error
}
