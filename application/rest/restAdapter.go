package rest

import (
	"github.com/OE-OverEngineer/over-review-backend/configuration"
	"github.com/OE-OverEngineer/over-review-backend/data/database"
	userStore "github.com/OE-OverEngineer/over-review-backend/data/user"
	"github.com/OE-OverEngineer/over-review-backend/domain/user"
	"github.com/gin-gonic/gin"
)

func Run(configuration *configuration.Config) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	_, err := database.Connection(configuration.Database)
	if err != nil {
		panic(err)
	}

	userRepository := userStore.NewUserRepositoryMock()
	userService := user.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)
	// userService := user.NewUserService(userRepository)
	// userService.GetAllUser()
	Routes(r, userHandler)
	r.Run(":" + configuration.Port)
}
