package actor

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	FirstName string `gorm:"column:firstName"`
	LastName  string `gorm:"column:lastName"`
}
