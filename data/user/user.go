package user

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/OE-OverEngineer/over-review-backend/data/role"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `gorm:"column:firstName"`
	LastName    string `gorm:"column:lastName"`
	Role        role.Role
	DisplayName string    `gorm:"column:displayName"`
	Avatar      string    `gorm:"column:avatar"`
	BirthDate   time.Time `gorm:"column:birthDate"`
	Email       string    `gorm:"column:email;uniqe"`
	Password    string    `gorm:"column:password"`
	Gender      Gender    `gorm:"column:gender"`
	Banned      bool      `gorm:"column:banned;default:false"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetUser(id int) (*User, error)
	GetByToken() (*User, error)
	DeleteUser(id int) error
}

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

func (t Gender) Value() (driver.Value, error) {
	switch t {
	case Male, Female:
		return string(t), nil
	}
	return nil, errors.New("invalid product type value")
}

func (t *Gender) Scan(value interface{}) error {
	var pt Gender
	if value == nil {
		*t = ""
		return nil
	}
	st, ok := value.(string)
	if !ok {
		return errors.New("invalid data for product type")
	}
	pt = Gender(st)
	switch pt {
	case Male, Female:
		*t = pt
		return nil
	}
	return errors.New("invalid product type value : ")
}
