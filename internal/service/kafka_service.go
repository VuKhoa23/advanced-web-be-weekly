package service

type KafkaService interface {
	SendMessage(topic, key, value string) error
	Close() error
}
