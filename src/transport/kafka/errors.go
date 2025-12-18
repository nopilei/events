package kafka

import "errors"

var (
	ErrUnknownEventType = errors.New("cannot send to Kafka: Unknown Event")
	ErrMarshalError     = errors.New("cannot marshal while sending to Kafka")
	ErrProducerNotInitialized    = errors.New("producer was not initialized")
	ErrMessageSendTimeout   = errors.New("timeout sending message to Kafka")
)
