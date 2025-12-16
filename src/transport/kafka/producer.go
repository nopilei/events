// from aiokafka import AIOKafkaProducer

// from src.brokers.base import AsyncProducer
// from src.brokers.exceptions import BrokerError
// from src.brokers.kafka.topics import topics
// from src.event_schemas.base import BaseEventData
// from src.event_schemas.registry import events_registry

// class KafkaProducer(AsyncProducer):
//     def __init__(self, producer: AIOKafkaProducer):
//         self.producer = producer

//     async def send(self, event: BaseEventData) -> None:
//         topic = self.get_topic_name(event)
//         data = event.model_dump_json().encode()
//         try:
//             await self.producer.send(topic=topic, value=data)
//         except Exception as exc:
//             raise BrokerError from exc

// def get_topic_name(self, event: BaseEventData) -> str:
//
//	event_type = events_registry[event.event_type]
//	return topics.get(event_type, event.event_type)
package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/nopilei/events/src/schema"
	kafkago "github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer      *kafkago.Writer
}

func (producer *KafkaProducer) Send(ctx context.Context, data schema.Event, timeout time.Duration) error{
    eventType := data.EventType()
    topic, ok := Topics[eventType]
    if !ok{
        return UnknownEventTypeError
    }

    rawData, err := json.Marshal(data)
    if err != nil{
        return errors.Join(MarshalError, err) 
    }
    kafkaMessage := kafkago.Message{
        Topic: topic,
        Value: rawData,
    }

    ctx, cancel := context.WithTimeout(ctx, timeout)
    defer cancel()
    return producer.writer.WriteMessages(ctx, kafkaMessage)
}

func New() (*KafkaProducer, error)  {
    settings, err := LoadKafkaSettings()
    if err != nil{
        return nil, err
    }
    return &KafkaProducer{
        writer: &kafkago.Writer{
            Addr: kafkago.TCP(settings.Brokers...),
            BatchTimeout: settings.BatchTimeout,
        },
    }, nil
}