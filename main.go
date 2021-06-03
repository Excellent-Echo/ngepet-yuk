package main

import (
	"ngepet-yuk/auth"
	"ngepet-yuk/config"
	"ngepet-yuk/handler"
	"ngepet-yuk/user"
	"ngepet-yuk/userDetail"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB                   *gorm.DB = config.Connection()
	userRepository                = user.NewRepository(DB)
	userService                   = user.NewService(userRepository)
	userHandler                   = handler.NewUserHandler(userService, authService)
	userDetailRepository          = userDetail.NewRepository(DB)
	userDetailService             = userDetail.NewService(userDetailRepository)
	userDetailHandler             = handler.NewUserDetailHandler(userDetailService, authService)
	authService                   = auth.NewService()
)

func main() {
	r := gin.Default()
	r.GET("/users", userHandler.ShowAllUser)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.GET("/users/:user_id", userHandler.GetUserByIDHandler)
	r.DELETE("/users/:user_id", userHandler.DeleteUserByIDHandler)
	r.PUT("/users/:user_id", userHandler.UpdateUserByIDHandler)

	r.GET("/userdetails/:user_id", userDetailHandler.GetUserDetailByUserIDHandler)
	r.POST("/userdetails", userDetailHandler.SaveNewUserDetailHandler)
	r.PUT("/userdetails/:userdetail_id", userDetailHandler.UpdateUserDetailByIDHandler)

	r.Run(":8080")
}
