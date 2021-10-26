package database

import (
	"fmt"

	"github.com/OE-OverEngineer/over-review-backend/configuration"
	// "github.com/OE-OverEngineer/over-review-backend/data/actor"
	// "github.com/OE-OverEngineer/over-review-backend/data/comment"
	// "github.com/OE-OverEngineer/over-review-backend/data/director"
	// "github.com/OE-OverEngineer/over-review-backend/data/movie"
	// "github.com/OE-OverEngineer/over-review-backend/data/report"
	// "github.com/OE-OverEngineer/over-review-backend/data/review"
	"github.com/OE-OverEngineer/over-review-backend/data/role"
	"github.com/OE-OverEngineer/over-review-backend/data/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(dbConfig *configuration.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Bangkok",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	print("123")
	if err != nil {
		return nil, err
	}
	// db.AutoMigrate(user.User{}, actor.Actor{}, comment.Comment{}, director.Director{}, movie.Movie{}, report.Report{}, review.Review{}, role.Role{})
	db.AutoMigrate(role.Role{}, user.User{})
	return db, err
}
