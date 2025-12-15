
package kafka
// from pydantic import BaseModel

// from src.event_schemas.order.events import OrderCreated
// from src.event_schemas.user.events import UserCreated, UserUpdated

// topics: dict[type[BaseModel], str] = {
//     UserCreated: 'users',
//     UserUpdated: 'users',
//     OrderCreated: 'orders',
// }

var Topics = map[string]string{
	"user_created":  "users",
	"user_updated":  "users",
	"order_created": "orders",
}
