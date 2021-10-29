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
	FirstName   string
	LastName    string
	DisplayName string
	Avatar      string
	BirthDate   time.Time
	Email       string
	Password    string
	Banned      bool
	RoleId      int
	Role        role.Role
	Gender      Gender
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
