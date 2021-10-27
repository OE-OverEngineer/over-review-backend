package report

import (
	"github.com/OE-OverEngineer/over-review-backend/data/user"
	"gorm.io/gorm"
)

type Report struct {
	// ID        int    `gorm:"column:id;primary_key"`
	gorm.Model
	User    user.User
	ByUser  user.User
	Message string `gorm:"column:message"`
}
