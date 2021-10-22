package rest

import (
	userStore "github.com/OE-OverEngineer/over-review-backend/data/user"
	"github.com/OE-OverEngineer/over-review-backend/domain/user"
	"github.com/gin-gonic/gin"
)

func Run(port string) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userRepository := userStore.NewUserRepositoryMock()
	userService := user.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)
	// userService := user.NewUserService(userRepository)
	// userService.GetAllUser()
	Routes(r, userHandler)
	r.Run(":" + port)
}
