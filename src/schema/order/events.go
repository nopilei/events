// # from pydantic import BaseModel

// # class OrderCreated(BaseModel):
// #     id: int
// #     user_id: int
// #     product_id: int
// #     amount: int

// # registry: dict[str, type[BaseModel]] = {
// #     'order_created': OrderCreated,
// # }

package order

type OrderCreated struct {
	Id        int `json:"id" validate:"required"`
	UserId    int `json:"user_id" validate:"required"`
	ProductId int `json:"product_id" validate:"required"`
	Amount    int `json:"amount" validate:"required"`
}
func (event *OrderCreated) EventType () string{
	return "order_created"
}

