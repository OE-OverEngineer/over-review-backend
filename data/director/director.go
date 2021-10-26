package director

import "gorm.io/gorm"

type Director struct {
	gorm.Model
	FirstName string `gorm:"column:firstName"`
	LastName  string `gorm:"column:lastName"`
}
