package user

type UserResponse struct {
	ID        uint    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type UserService interface {
	GetAllUser() ([]UserResponse, error)
	GetUser(int) (*UserResponse, error)
}
