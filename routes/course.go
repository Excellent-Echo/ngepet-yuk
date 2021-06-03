package routes

import (
	"ngepet-yuk/course"
	"ngepet-yuk/handler"

	"github.com/gin-gonic/gin"
)

var (
	courseRepository = course.NewRepository(DB)
	courseService    = course.NewService(courseRepository)
	courseHandler    = handler.NewCourseHandler(courseService)
)

func CoursesRoute(r *gin.Engine) {
	r.GET("/courses", courseHandler.ShowCourses)
	//r.GET("")
	r.GET("/courses/category/:category_id", courseHandler.ShowCourseFilterByCategory)
	r.GET("/courses/mastery/:mastery_id", courseHandler.ShowCourseFilterByMastery)
	r.GET("/courses/subcription/:sub_type", courseHandler.ShowCourseFilterBySubcription)
	r.POST("/courses", courseHandler.CreateCourseHandler)
	r.PUT("/courses/:course_id", courseHandler.UpdateCourseByIDHandler)
	r.DELETE("/courses/:course_id", courseHandler.DeleteCourseByIDHandler)
}
