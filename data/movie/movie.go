package movie

import (
	"time"

	"github.com/OE-OverEngineer/over-review-backend/data/actor"
	"github.com/OE-OverEngineer/over-review-backend/data/director"

	// "github.com/OE-OverEngineer/over-review-backend/data/review"
	"github.com/OE-OverEngineer/over-review-backend/data/user"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Director    director.Director
	Title       string    `gorm:"column:title;type:varchar(64)"`
	Descrption  string    `gorm:"column:description"`
	StartDate   time.Time `gorm:"column:startDate"`
	BannerImage string    `gorm:"column:bannerImage"`
	TrailerLink string    `gorm:"column:trailerLink"`
	RequestBy   user.User
	Approved    bool `gorm:"column:approved"`
	Actors      []actor.Actor
	// Reviews     []review.Review
}
