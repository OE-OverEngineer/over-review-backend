package comment

import (
	"github.com/OE-OverEngineer/over-review-backend/data/user"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	MovieID  int
	UserID   int
	User     user.User
	ReviewID int
	Comment  string
}
