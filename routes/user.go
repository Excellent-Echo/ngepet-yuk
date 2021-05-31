package routes

import (
	"todoAPIGolang/auth"
	"todoAPIGolang/config"
	"todoAPIGolang/handler"
	"todoAPIGolang/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewRepository(DB)
	userService             = user.NewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {

	r.GET("/users", userHandler.ShowAllUser)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.GET("/users/:user_id", userHandler.GetUserByIDHandler)
	r.DELETE("/users/:user_id", userHandler.DeleteUserByIDHandler)
	r.PUT("/users/:user_id", userHandler.UpdateUserByIDHandler)

}
