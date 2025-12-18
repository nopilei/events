package schema

import (
	"github.com/nopilei/events/src/schema/order"
	"github.com/nopilei/events/src/schema/user"
)

type Event interface {
	EventType() string
}

var EventsRegistry = map[string]func() Event{
	"user_created":  func() Event { return &user.UserCreated{} },
	"user_updated":  func() Event { return &user.UserUpdated{} },
	"order_created": func() Event { return &order.OrderCreated{} },
}
