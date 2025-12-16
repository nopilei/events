package schema

import (
	"github.com/nopilei/events/src/schema/order"
	"github.com/nopilei/events/src/schema/user"
)

// from pydantic import BaseModel

// from src.event_schemas.user.events import registry as user_registry
// from src.event_schemas.order.events import registry as order_registry

// events_registry: dict[str, type[BaseModel]] = (
//
//	user_registry
//	| order_registry
//
// )
type Event interface{
	EventType() string
}

var EventsRegistry = map[string]func() Event{
	"user_created":  func() Event { return &user.UserCreated{} },
	"user_updated":  func() Event { return &user.UserUpdated{} },
	"order_created": func() Event { return &order.OrderCreated{} },
}
