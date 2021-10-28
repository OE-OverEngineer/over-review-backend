package movie

import (
	"time"

	"github.com/OE-OverEngineer/over-review-backend/data/actor"
	"github.com/OE-OverEngineer/over-review-backend/data/review"

	// "github.com/OE-OverEngineer/over-review-backend/data/review"
	"github.com/OE-OverEngineer/over-review-backend/data/user"
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	DirectorID  int
	Title       string
	Descrption  string
	StartDate   time.Time
	BannerImage string
	TrailerLink string
	RequestByID int
	RequestBy   user.User
	Approved    bool
	Actors      []*actor.Actor `gorm:"many2many:actorsMovie"`
	Reviews     []review.Review
}
