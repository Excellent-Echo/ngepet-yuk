package handler

import (
	"net/http"
	"ngepet-yuk/course"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"

	"github.com/gin-gonic/gin"
)

type courseHandler struct {
	courseService course.Service
}

func NewCourseHandler(courseService course.Service) *courseHandler {
	return &courseHandler{courseService}
}

func (h *courseHandler) ShowCourses(c *gin.Context) {
	courses, err := h.courseService.GetAllCourse()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("success get all course", 200, "status OK", courses)
	c.JSON(200, response)
}

func (h *courseHandler) CreateCourseHandler(c *gin.Context) {
	var inputCourse entity.CoursesInput

	if err := c.ShouldBindJSON(&inputCourse); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newCourse, err := h.courseService.SaveNewCourse(inputCourse)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create course", 200, "Status OK", newCourse)
	c.JSON(200, response)
}

func (h *courseHandler) UpdateCourseByIDHandler(c *gin.Context) {
	id := c.Params.ByName("course_id")

	var updateCourseInput entity.UpdateCoursesInput

	if err := c.ShouldBindJSON(&updateCourseInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	course, err := h.courseService.UpdateCourseByID(id, updateCourseInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update course by ID", 200, "success", course)
	c.JSON(200, response)
}

func (h *courseHandler) DeleteCourseByIDHandler(c *gin.Context) {
	id := c.Params.ByName("course_id")

	course, err := h.courseService.DeleteCourseByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete course", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success delete course by ID", http.StatusOK, "success", course)
	c.JSON(http.StatusOK, response)
}

func (h *courseHandler) ShowCourseFilterByCategory(c *gin.Context) {
	id := c.Params.ByName("category_id")

	courses, err := h.courseService.FilterCoursesByCategory(id)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("success get all course filtered by category", 200, "status OK", courses)
	c.JSON(200, response)
}

func (h *courseHandler) ShowCourseFilterByMastery(c *gin.Context) {
	id := c.Params.ByName("mastery_id")

	courses, err := h.courseService.FilterCoursesByMastery(id)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("success get all course filtered by mastery", 200, "status OK", courses)
	c.JSON(200, response)
}

func (h *courseHandler) ShowCourseFilterBySubcription(c *gin.Context) {
	id := c.Params.ByName("sub_type")

	courses, err := h.courseService.FilterCoursesBySubcription(id)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("success get all course filtered by subcription", 200, "status OK", courses)
	c.JSON(200, response)
}

func (h *courseHandler) ShowCourseFilterByCategoryAndMasteryAndSubcription(c *gin.Context) {
	cid := c.Query("category_id")
	mid := c.Query("mastery_id")
	sid := c.Query("sub_type")

	courses, err := h.courseService.FilterCoursesByCategoryAndMasteryAndSubcription(cid, mid, sid)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("success get all course", 200, "status OK", courses)
	c.JSON(200, response)
}
