package main

import (
	"ngepet-yuk/config"
	"ngepet-yuk/handler"
	"ngepet-yuk/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	userHandler             = handler.NewUserHandler(userService)
)

func main() {
	r := gin.Default()
	r.GET("/users", userHandler.ShowAllUser)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.GET("/users/:user_id", userHandler.GetUserByIDHandler)

	r.Run(":8080")
}
