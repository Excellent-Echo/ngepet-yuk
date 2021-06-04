package routes

import (
	"ngepet-yuk/handler"
	"ngepet-yuk/userDetail"

	"github.com/gin-gonic/gin"
)

var (
	userDetailRepository = userDetail.NewRepository(DB)
	userDetailService    = userDetail.NewService(userDetailRepository)
	userDetailHandler    = handler.NewUserDetailHandler(userDetailService, authService)
)

func UserDetailRoute(r *gin.Engine) {
	r.GET("/alluserdetails", handler.Middleware(userService, authService), userDetailHandler.ShowAllUserDetail)
	r.GET("/userdetails", handler.Middleware(userService, authService), userDetailHandler.GetUserDetailByUserIDHandler)
	r.POST("/userdetails", handler.Middleware(userService, authService), userDetailHandler.SaveNewUserDetailHandler)
	r.PUT("/userdetails/:userdetail_id", handler.Middleware(userService, authService), userDetailHandler.UpdateUserDetailByIDHandler)
}
