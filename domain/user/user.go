package user

import (
	"github.com/OE-OverEngineer/over-review-backend/data/user"

	"github.com/gofrs/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

type UserService interface {
	CreateUser(*user.User) (UserResponse, error)
	GetUser(int) (UserResponse, error)
}
