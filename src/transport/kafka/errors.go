package kafka

import "errors"

var (
	UnknownEventTypeError = errors.New("Cannot send to Kafka: Unknown Event")
	MarshalError = errors.New("Cannot marshal while sending to Kafka") 
)