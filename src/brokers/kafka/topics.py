from pydantic import BaseModel

from src.event_schemas.order.events import OrderCreated
from src.event_schemas.user.events import UserCreated, UserUpdated

topics: dict[type[BaseModel], str] = {
    UserCreated: 'users',
    UserUpdated: 'users',
    OrderCreated: 'orders',
}