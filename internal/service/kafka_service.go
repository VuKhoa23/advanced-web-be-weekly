package service

type KafkaService interface {
	SendMessage(topic, value string) (string, error)
	Close() error
}
