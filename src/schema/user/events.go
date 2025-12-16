// from pydantic import BaseModel

// class UserCreated(BaseModel):
//     id: int
//     name: str

// class UserUpdated(BaseModel):
//     id: int
//     name: str

// registry: dict[str, type[BaseModel]] = {
//     'user_created': UserCreated,
//     'user_updated': UserUpdated,
// }

package user

type UserCreated struct {
	Id   int `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
func (event *UserCreated) EventType () string{
	return "user_created"
}

type UserUpdated struct {
	Id   int `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
func (event *UserUpdated) EventType () string{
	return "user_updated"
}