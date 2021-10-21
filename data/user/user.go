package user

type User struct {
	ID        int    `gorm:"column:id;primary_key"`
	FirstName string `gorm:"column:firstName"`
	LastName  string `gorm:"column:lastName"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetUser(int) (*User, error)
}
