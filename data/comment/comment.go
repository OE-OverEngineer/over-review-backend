package comment

import (
	"github.com/OE-OverEngineer/over-review-backend/data/review"
	"github.com/OE-OverEngineer/over-review-backend/data/user"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	User    user.User
	Review  review.Review
	Comment string `gorm:"column:comment"`
}
