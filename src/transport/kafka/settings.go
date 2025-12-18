package kafka

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var Topics = map[string]string{
	"user_created":  "users",
	"user_updated":  "users",
	"order_created": "orders",
}

type KafkaSettings struct {
	Brokers []string
	BatchTimeout time.Duration
	BatchSize int
	RequiredAcks  kafka.RequiredAcks
}

func LoadKafkaSettings() (KafkaSettings, error) {
	brokers := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if brokers == "" {
		return KafkaSettings{}, errors.New("KAFKA_BOOTSTRAP_SERVERS is not set")
	}

	return KafkaSettings{
		Brokers: strings.Split(brokers, ","),
		BatchTimeout: 100 * time.Millisecond,
		BatchSize: 1000,
		RequiredAcks: kafka.RequireAll,
	}, nil
}