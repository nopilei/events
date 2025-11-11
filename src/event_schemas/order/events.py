from pydantic import BaseModel


class OrderCreated(BaseModel):
    id: int
    user_id: int
    product_id: int
    amount: int


registry: dict[str, type[BaseModel]] = {
    'order_created': OrderCreated,
}
