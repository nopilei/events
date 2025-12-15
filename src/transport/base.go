package transport

type Broker interface{
	Send (data []byte) error
}