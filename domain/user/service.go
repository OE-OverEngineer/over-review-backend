package user

import (
	"github.com/OE-OverEngineer/over-review-backend/data/user"
)

// Service struct handles user bussiness logic tasks.
type service struct {
	repository user.UserRepository
}

func NewUserService(repository user.UserRepository) *service {
	return &service{repository: repository}
}
