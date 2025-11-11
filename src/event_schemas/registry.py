from pydantic import BaseModel

from src.event_schemas.user.events import registry as user_registry
from src.event_schemas.order.events import registry as order_registry

events_registry: dict[str, type[BaseModel]] = (
        user_registry
        | order_registry
)
