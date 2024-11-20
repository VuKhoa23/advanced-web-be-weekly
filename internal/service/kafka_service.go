package service

type KafkaService interface {
	SendMessage(topic, key, value string) (string, error)
	Close() error
}
