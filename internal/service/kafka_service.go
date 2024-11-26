package service

type KafkaService interface {
	SendMessage(topic, value string) error
	Close() error
}
