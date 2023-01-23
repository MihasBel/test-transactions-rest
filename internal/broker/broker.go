package broker

import (
	"github.com/MihasBel/test-transactions-rest/internal/app"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog"
)

// Broker represents kafka broker with logger and config
type Broker struct {
	b   *kafka.Producer
	cfg app.Configuration
	l   zerolog.Logger
}

// NewBroker returns new instance of Broker
func NewBroker(cfg app.Configuration, l zerolog.Logger) *Broker {
	b, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaURL,
		"client.id":         "PLAINTEXT",
		"acks":              "all"})
	if err != nil {
		l.Error().Err(err)
	}
	return &Broker{
		b:   b,
		cfg: cfg,
		l:   l,
	}
}
