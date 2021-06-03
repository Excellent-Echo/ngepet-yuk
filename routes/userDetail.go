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

	r.GET("/userdetails/:user_id", userDetailHandler.GetUserDetailByUserIDHandler)
	r.POST("/userdetails", userDetailHandler.SaveNewUserDetailHandler)
	r.PUT("/userdetails/:userdetail_id", userDetailHandler.UpdateUserDetailByIDHandler)

}
