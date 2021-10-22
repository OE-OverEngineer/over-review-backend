package user

import (
	"errors"

	"github.com/OE-OverEngineer/over-review-backend/data/user"
)

// Service struct handles user bussiness logic tasks.
type userService struct {
	userRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) UserService {
	return userService{userRepository: userRepository}
}

func (s userService) GetAllUser() ([]UserResponse, error) {
	users, err := s.userRepository.GetAll()
	if err != nil {
		return nil, err
	}

	userResponses := []UserResponse{}
	for _, user := range users {
		userResponse := UserResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (s userService) GetUser(id int) (*UserResponse, error) {
	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	userResponse := &UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.LastName,
	}
	return userResponse, nil
}
