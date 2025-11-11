from typing import Any

from pydantic import BaseModel, RootModel


class EventBody(RootModel[dict[str, Any]]):
    pass


class BaseEventData(BaseModel):
    event_type: str
    data: EventBody
