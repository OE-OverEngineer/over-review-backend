package review

import (
	"github.com/OE-OverEngineer/over-review-backend/data/movie"
	"github.com/OE-OverEngineer/over-review-backend/data/user"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	User    user.User
	Movie   movie.Movie
	Message string `gorm:"column:message"`
	Score   uint8  `gorm:"column:score"`
}
