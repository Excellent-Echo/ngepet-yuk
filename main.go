package main

import (
	"ngepet-yuk/auth"
	"ngepet-yuk/category"
	"ngepet-yuk/config"
	"ngepet-yuk/course"
	"ngepet-yuk/handler"
	"ngepet-yuk/mastery"
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

	categoryRepository = category.NewRepository(DB)
	categoryService    = category.NewService(categoryRepository)
	categoryHandler    = handler.NewCategoryHandler(categoryService)

	masteryRepository = mastery.NewRepository(DB)
	masteryService    = mastery.NewService(masteryRepository)
	masteryHandler    = handler.NewMasteryHandler(masteryService)

	courseRepository = course.NewRepository(DB)
	courseService    = course.NewService(courseRepository)
	courseHandler    = handler.NewCourseHandler(courseService)
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

	r.GET("/categories", categoryHandler.ShowCategories)
	r.POST("/categories", categoryHandler.CreateCategoryHandler)
	r.PUT("/categories/:category_id", categoryHandler.UpdateCategoryByIDHandler)

	r.GET("/masteries", masteryHandler.ShowMasteries)
	r.POST("/masteries", masteryHandler.CreateMasteryHandler)
	r.PUT("/masteries/:mastery_id", masteryHandler.UpdateMasteryByIDHandler)

	r.GET("/courses", courseHandler.ShowCourses)
	//r.GET("")
	r.GET("/courses/category/:category_id", courseHandler.ShowCourseFilterByCategory)
	r.GET("/courses/mastery/:mastery_id", courseHandler.ShowCourseFilterByMastery)
	r.GET("/courses/subcription/:sub_type", courseHandler.ShowCourseFilterBySubcription)
	r.POST("/courses", courseHandler.CreateCourseHandler)
	r.PUT("/courses/:course_id", courseHandler.UpdateCourseByIDHandler)
	r.DELETE("/courses/:course_id", courseHandler.DeleteCourseByIDHandler)

	r.Run(":8080")
}
