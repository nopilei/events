package kafka

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/nopilei/events/src/schema"
	kafkago "github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafkago.Writer
}

var defaultProducer *KafkaProducer

func Send(ctx context.Context, data *schema.BaseEventData) error {
	if defaultProducer == nil {
		return ErrProducerNotInitialized
	}

	eventType := data.EventType
	topic, ok := Topics[eventType]
	if !ok {
		return ErrUnknownEventType
	}

	rawData, err := json.Marshal(data)
	if err != nil {
		return errors.Join(ErrMarshalError, err)
	}
	kafkaMessage := kafkago.Message{
		Topic: topic,
		Value: rawData,
	}

	return defaultProducer.writer.WriteMessages(ctx, kafkaMessage)
}

func InitProducer() error {
    if defaultProducer != nil{
        return nil
    }
	settings, err := LoadKafkaSettings()
	if err != nil {
		return err
	}
	defaultProducer = &KafkaProducer{
		writer: &kafkago.Writer{
			Addr:         kafkago.TCP(settings.Brokers...),
			BatchTimeout: settings.BatchTimeout,
            BatchSize: settings.BatchSize,
            RequiredAcks: settings.RequiredAcks,
            AllowAutoTopicCreation: true,

		},
	}
	return nil
}
