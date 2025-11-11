from pydantic import BaseModel


class UserCreated(BaseModel):
    id: int
    name: str


class UserUpdated(BaseModel):
    id: int
    name: str


registry: dict[str, type[BaseModel]] = {
    'user_created': UserCreated,
    'user_updated': UserUpdated,
}
