package director

import "gorm.io/gorm"

type Director struct {
	gorm.Model
	FirstName string
	LastName  string
}
