package report

import (
	"github.com/OE-OverEngineer/over-review-backend/data/user"
	"gorm.io/gorm"
)

type Report struct {
	// ID        int
	gorm.Model
	UserID   int
	User     user.User
	ByUserID int
	ByUser   user.User
	Message  string
}
