package main

import (
	"ngepet-yuk/auth"
	"ngepet-yuk/config"
	"ngepet-yuk/handler"
	"ngepet-yuk/subtype"
	"ngepet-yuk/user"
	"ngepet-yuk/userDetail"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB = config.Connection()
	authService          = auth.NewService()

	userRepository = user.NewRepository(DB)
	userService    = user.NewService(userRepository)
	userHandler    = handler.NewUserHandler(userService, authService)

	userDetailRepository = userDetail.NewRepository(DB)
	userDetailService    = userDetail.NewService(userDetailRepository)
	userDetailHandler    = handler.NewUserDetailHandler(userDetailService, authService)

	subtypeRepository = subtype.NewRepository(DB)
	subtypeService    = subtype.NewService(subtypeRepository)
	subtypeHandler    = handler.NewSubTypeHandler(subtypeService)
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

	r.GET("/sub-types", subtypeHandler.ShowAllSubType)
	r.POST("/sub-types", subtypeHandler.CreateSubtypeHandler)
	r.PUT("/sub-types/:subtype_id", subtypeHandler.UpdateSubtypeByIDHandler)

	r.Run(":8080")
}
