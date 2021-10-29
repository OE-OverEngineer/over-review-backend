package review

import (
	// "github.com/OE-OverEngineer/over-review-backend/data/movie"
	"github.com/OE-OverEngineer/over-review-backend/data/user"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID  int
	User    user.User
	MovieID int
	Message string
	Score   uint8
}
